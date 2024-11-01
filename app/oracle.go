package app

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"slices"
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/server/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/skip-mev/connect/v2/cmd/constants/marketmaps"
	oracleconfig "github.com/skip-mev/connect/v2/oracle/config"
	connecttypes "github.com/skip-mev/connect/v2/pkg/types"
	oracleclient "github.com/skip-mev/connect/v2/service/clients/oracle"
	servicemetrics "github.com/skip-mev/connect/v2/service/metrics"
	servicetypes "github.com/skip-mev/connect/v2/service/servers/oracle/types"

	prediction "github.com/artela-network/artela-rollkit/contracts/generated"
)

// initializeOracle initializes the oracle client and metrics.
func (app *App) initializeOracle(appOpts types.AppOptions) (oracleclient.OracleClient, servicemetrics.Metrics, error) {
	// Read general config from app-opts, and construct oracle service.
	cfg, err := oracleconfig.ReadConfigFromAppOpts(appOpts)
	if err != nil {
		return nil, nil, err
	}

	app.Logger().Info("Reading Oracle config", "OracleAddress", cfg.OracleAddress)
	// If app level instrumentation is enabled, then wrap the oracle service with a metrics client
	// to get metrics on the oracle service (for ABCI++). This will allow the instrumentation to track
	// latency in VerifyVoteExtension requests and more.
	oracleMetrics, err := servicemetrics.NewMetricsFromConfig(cfg, app.ChainID())
	if err != nil {
		app.Logger().Error("Failed to create Oracle metrics", "error", err)
		return nil, nil, err
	}

	// Create the oracle service.
	oracleClient, err := oracleclient.NewPriceDaemonClientFromConfig(
		cfg,
		app.Logger().With("client", "oracle"),
		oracleMetrics,
	)
	if err != nil {
		app.Logger().Error("Failed to create Oracle client", "error", err)
		return nil, nil, err
	}

	// Connect to the oracle service (default timeout of 5 seconds).
	go func() {
		app.Logger().Info("attempting to start oracle client...", "address", cfg.OracleAddress)
		if err := oracleClient.Start(context.Background()); err != nil {
			app.Logger().Error("failed to start oracle client", "err", err)
			panic(err)
		}
	}()

	return oracleClient, oracleMetrics, nil
}

// fetchAndStoreOracleData fetches the latest data and updates both chain state and Ethereum contract
func (app *App) fetchAndStoreOracleData(ctx sdk.Context) error {
	// Ensure Oracle client and Ethereum components are initialized
	if app.oracleClient == nil {
		return fmt.Errorf("oracle client not initialized")
	}
	if app.ethClient == nil {
		// Connect to Ethereum
		var err error
		app.ethClient, err = ethclient.Dial("http://localhost:8545")
		if err != nil {
			return fmt.Errorf("failed to connect to ethereum: %w", err)
		}
	}
	if app.predictionMarket == nil {
		return fmt.Errorf("prediction market contract not initialized")
	}

	// Create request context with timeout
	reqCtx, cancel := context.WithTimeout(ctx.Context(), time.Second*5)
	defer cancel()

	// Fetch latest prices from Oracle client
	oracleResp, err := app.oracleClient.Prices(ctx.WithContext(reqCtx), &servicetypes.QueryPricesRequest{})
	if err != nil {
		app.Logger().Error(
			"failed to retrieve oracle prices",
			"ctx_err", reqCtx.Err(),
			"err", err,
		)
		return err
	}

	if oracleResp == nil {
		return fmt.Errorf("oracle returned nil prices")
	}

	app.Logger().Info("Retrieved oracle prices", "count", len(oracleResp.Prices), "prices", oracleResp.Prices)

	// Convert Oracle response to price map
	//prices, err := ConvertOraclePrices(oracleResp)
	//if err != nil {
	//	app.Logger().Error("failed to convert oracle prices", "err", err)
	//	return err
	//}

	//app.Logger().Info("Converted oracle prices", "count", len(prices), "prices", prices)

	// Get markets that need updating
	//market, err := app.MarketMapKeeper.GetMarket(ctx, "WILL_BERNIE_SANDERS_WIN_THE_2024_US_PRESIDENTIAL_ELECTION?YES/USD")
	//if err != nil {
	//	app.Logger().Error("failed to get market", "err", err)
	//	return err
	//}

	// Prepare Ethereum transaction auth
	auth, err := createEthereumAuth(app.ethClient, app.contractConfig.DeployerPrivateKey)
	if err != nil {
		app.Logger().Error("failed to create ethereum auth", "err", err)
		return err
	}

	m := Market{
		ID: 1,
		CurrencyPair: connecttypes.CurrencyPair{
			Base:  "WILL_BERNIE_SANDERS_WIN_THE_2024_US_PRESIDENTIAL_ELECTION?YES/USD",
			Quote: "USD",
		},
		IsActive: true,
	}

	// Get corresponding currency pair for this market
	cp := m.CurrencyPair
	price, exists := oracleResp.Prices[cp.Base]

	app.Logger().Debug("retrieved price", "market_id", m.ID, "currency_pair", cp, "price", price)

	if !exists {
		app.Logger().Debug("no price for market", "market_id", m.ID, "currency_pair", cp)
	}

	//if price.Sign() < 0 {
	//	app.Logger().Error("price is negative", "market_id", m.ID, "currency_pair", cp)
	//}

	// Create and store QuotePrice in chain state
	//quotePrice := oracletypes.QuotePrice{
	//	Price:          math.NewIntFromBigInt(price),
	//	BlockTimestamp: ctx.BlockHeader().Time,
	//	BlockHeight:    uint64(ctx.BlockHeight()),
	//}

	//// Store in chain state
	//if err := app.OracleKeeper.SetPriceForCurrencyPair(ctx, cp, quotePrice); err != nil {
	//	app.Logger().Error(
	//		"failed to set price in chain state",
	//		"market_id", m.ID,
	//		"currency_pair", cp,
	//		"err", err,
	//	)
	//	return err
	//}

	// Convert price to odds format for the prediction market
	odds := convertPriceToOdds(price)

	details, err := app.predictionMarket.GetMarketDetails(
		nil, //auth,
		big.NewInt(m.ID),
	)
	if err != nil {
		app.Logger().Error(
			"failed to get market details",
			"market_id", m.ID,
			"err", err,
		)
		return err
	}

	app.Logger().Info(
		"market details",
		"market_id", m.ID,
		"description", details.Description,
		"current odds", details.CurrentOdds,
	)

	tx, err := app.predictionMarket.UpdateOracleData(
		auth,
		big.NewInt(m.ID),
		odds,
		big.NewInt(ctx.BlockHeader().Time.Unix()),
	)
	if err != nil {
		app.Logger().Error(
			"failed to update ethereum contract",
			"market_id", m.ID,
			"err", err,
		)
		return err
	}

	// Wait for transaction confirmation
	receipt, err := app.ethClient.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		app.Logger().Error(
			"failed to get transaction receipt",
			"market_id", m.ID,
			"tx_hash", tx.Hash().String(),
			"err", err,
		)
		return err
	}

	if receipt.Status == 0 {
		app.Logger().Error(
			"ethereum transaction failed",
			"market_id", m.ID,
			"tx_hash", tx.Hash().String(),
		)
		return fmt.Errorf("ethereum transaction failed")
	}

	app.Logger().Info(
		"successfully updated market data",
		"market_id", m.ID,
		"currency_pair", cp.String(),
		"odd", odds.String(),
		"tx_hash", tx.Hash().String(),
	)

	return nil
}

// Helper function to convert price to odds format
func convertPriceToOdds(price string) *big.Int {
	// Convert price to big.Int
	priceBigInt, ok := new(big.Int).SetString(price, 10)
	if !ok {
		return big.NewInt(0)
	}

	return priceBigInt
}

// Market struct - adjust based on your needs
type Market struct {
	ID           int64
	CurrencyPair connecttypes.CurrencyPair
	IsActive     bool
	// Add other relevant fields
}

type ContractConfig struct {
	DeployerPrivateKey string         `json:"deployer_private_key"`
	OracleAddress      common.Address `json:"oracle_address"`
}

func (app *App) DeployContracts() error {
	config := app.contractConfig

	app.Logger().Info("Deploying contracts", "config", config)

	// Connect to Ethereum
	ethClient, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		return fmt.Errorf("failed to connect to ethereum: %w", err)
	}

	// 1. Deploy BetToken
	betTokenAddress, err := app.deployBetToken(ethClient, config)
	if err != nil {
		return fmt.Errorf("failed to deploy bet token: %w", err)
	}

	// 2. Deploy PredictionMarket with the BetToken address
	predictionMarketAddress, predictionMarket, err := app.deployPredictionMarket(
		ethClient,
		config,
		betTokenAddress,
		config.OracleAddress,
	)
	if err != nil {
		return fmt.Errorf("failed to deploy prediction market: %w", err)
	}

	// 3. Approve PredictionMarket contract to spend tokens
	if err := app.approvePredictionMarket(
		ethClient,
		config,
		betTokenAddress,
		predictionMarketAddress,
	); err != nil {
		return fmt.Errorf("failed to approve prediction market: %w", err)
	}

	// Store addresses in app state or config
	app.betTokenAddress = betTokenAddress

	app.predictionMarketAddress = predictionMarketAddress
	app.predictionMarket = predictionMarket

	app.Logger().Info("Contracts deployed", "BetToken", betTokenAddress, "PredictionMarket", predictionMarketAddress)

	app.ethClient = ethClient

	return nil
}

func (app *App) deployBetToken(
	client *ethclient.Client,
	config ContractConfig,
) (common.Address, error) {
	// Create auth for deployment
	auth, err := createEthereumAuth(client, config.DeployerPrivateKey)
	if err != nil {
		return common.Address{}, err
	}
	// Deploy BetToken contract
	address, tx, _, err := prediction.DeployBetToken(
		auth,
		client,
		big.NewInt(1000000000000000000), // 1M
	)
	if err != nil {
		return common.Address{}, err
	}

	// Wait for deployment to complete
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return common.Address{}, err
	}

	app.Logger().Info("Deployed BetToken contract", "receipt", receipt)

	return address, nil
}

func (app *App) deployPredictionMarket(
	client *ethclient.Client,
	config ContractConfig,
	betTokenAddress common.Address,
	oracleAddress common.Address,
) (common.Address, *prediction.PredictionMarket, error) {
	// Create auth for deployment
	auth, err := createEthereumAuth(client, config.DeployerPrivateKey)
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("failed to create ethereum auth: %w", err)
	}
	// Deploy PredictionMarket contract
	address, tx, instance, err := prediction.DeployPredictionMarket(
		auth,
		client,
		betTokenAddress,
		oracleAddress,
	)
	if err != nil {
		return common.Address{}, nil, err
	}

	// Wait for deployment to complete
	_, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return common.Address{}, nil, err
	}

	// Create auth for deployment
	auth, err = createEthereumAuth(client, config.DeployerPrivateKey)
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("failed to create ethereum auth: %w", err)
	}

	// Set oracle address
	tx, err = instance.SetOracle(auth, oracleAddress)
	if err != nil {
		return common.Address{}, nil, err
	}

	app.Logger().Info("Setting oracle address for prediction market contract", "oracle", oracleAddress.String())
	app.Logger().Info("Setting oracle address for prediction market contract", "tx", tx)

	// Wait for oracle setup to complete
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return common.Address{}, nil, err
	}

	app.Logger().Info("Set oracle address for prediction market contract", "receipt", receipt)

	// Create auth for deployment
	auth, err = createEthereumAuth(client, config.DeployerPrivateKey)
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("failed to create ethereum auth: %w", err)
	}

	tx, err = instance.CreateMarket(auth, "WILL_BERNIE_SANDERS_WIN_THE_2024_US_PRESIDENTIAL_ELECTION?YES/USD")
	if err != nil {
		return common.Address{}, nil, err
	}

	app.Logger().Info("Creating market", "description", "WILL_BERNIE_SANDERS_WIN_THE_2024_US_PRESIDENTIAL_ELECTION?YES/USD", "tx", tx)

	// Wait for market creation to complete
	receipt, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return common.Address{}, nil, err
	}

	app.Logger().Info("Created market", "description", "WILL_BERNIE_SANDERS_WIN_THE_2024_US_PRESIDENTIAL_ELECTION?YES/USD", "receipt", receipt, "logs", receipt.Logs)

	return address, instance, nil
}

func (app *App) approvePredictionMarket(
	client *ethclient.Client,
	config ContractConfig,
	betTokenAddress common.Address,
	predictionMarketAddress common.Address,
) error {
	// Load BetToken contract
	betToken, err := prediction.NewBetToken(betTokenAddress, client)
	if err != nil {
		return err
	}

	auth, err := createEthereumAuth(client, config.DeployerPrivateKey)
	if err != nil {
		return fmt.Errorf("failed to create ethereum auth: %w", err)
	}

	// Approve PredictionMarket to spend maximum amount
	tx, err := betToken.Approve(
		auth,
		predictionMarketAddress,
		new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(1)), // max uint256
	)
	if err != nil {
		return err
	}

	// Wait for approval to complete
	_, err = bind.WaitMined(context.Background(), client, tx)
	return err
}

// Contract deployment and setup
func (app *App) SetupPredictionMarket(ctx sdk.Context) error {
	// add core markets
	coreMarkets := marketmaps.PolymarketMarketMap
	markets := coreMarkets.Markets

	app.Logger().Info("Setting up markets", "Market count", len(markets))
	// sort keys so we can deterministically iterate over map items.
	keys := make([]string, 0, len(markets))
	for name := range markets {
		keys = append(keys, name)
	}
	slices.Sort(keys)

	for _, marketName := range keys {
		// create market
		market := markets[marketName]
		app.Logger().Info("Creating market", "market", marketName)
		err := app.MarketMapKeeper.CreateMarket(ctx, market)
		if err != nil {
			app.Logger().Error("Failed to create market", "market", marketName, "error", err)
			return err
		}

		// invoke hooks. this syncs the market to x/oracle.
		err = app.MarketMapKeeper.Hooks().AfterMarketCreated(ctx, market)
		if err != nil {
			app.Logger().Error("Failed to sync market with Oracle", "market", marketName, "error", err)
			return err
		}
	}

	return nil
}

// Helper to create Ethereum auth
func createEthereumAuth(client *ethclient.Client, privateKeyHex string) (*bind.TransactOpts, error) {
	if strings.HasPrefix(privateKeyHex, "0x") {
		privateKeyHex = privateKeyHex[2:]
	}
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, err
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}

	// Get the sender's address
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// Try both PendingNonceAt and NonceAt to ensure we get the correct nonce
	pendingNonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}

	latestNonce, err := client.NonceAt(context.Background(), fromAddress, nil) // nil for latest block
	if err != nil {
		return nil, err
	}

	// Use the maximum of pending and latest nonce to be safe
	nonce := pendingNonce
	if latestNonce > pendingNonce {
		nonce = latestNonce
	}

	// Log the nonces for debugging
	log.Printf("Address: %s, Latest Nonce: %d, Pending Nonce: %d, Using Nonce: %d",
		fromAddress.Hex(), latestNonce, pendingNonce, nonce)

	auth.Nonce = big.NewInt(int64(nonce))

	// Get gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	auth.GasPrice = gasPrice

	// Set gas limit
	auth.GasLimit = uint64(5000000)

	return auth, nil
}

// ConvertOraclePrices converts oracleResp.Prices to map[connecttypes.CurrencyPair]*big.Int
//func ConvertOraclePrices(oracleResp *servicetypes.QueryPricesResponse) (map[connecttypes.CurrencyPair]*big.Int, error) {
//	result := make(map[connecttypes.CurrencyPair]*big.Int)
//
//	// Iterate over the prices in the oracle response
//	for pairStr, priceStr := range oracleResp.Prices {
//		// Parse the string into a CurrencyPair
//		currencyPair, err := parseCurrencyPair(pairStr)
//		if err != nil {
//			return nil, fmt.Errorf("failed to parse currency pair %s: %w", pairStr, err)
//		}
//
//		// Convert the price from string to *big.Int
//		priceBigInt, ok := new(big.Int).SetString(priceStr, 10)
//		if !ok {
//			return nil, fmt.Errorf("failed to convert price %s to big.Int", priceStr)
//		}
//
//		// Store the result in the map
//		result[currencyPair] = priceBigInt
//	}
//
//	return result, nil
//}

// parseCurrencyPair is a helper function that converts a string representation of a currency pair to connecttypes.CurrencyPair
// More tests are needed
//func parseCurrencyPair(pairStr string) (connecttypes.CurrencyPair, error) {
//	parts := strings.Split(pairStr, "/")
//	if len(parts) != 2 {
//		return connecttypes.CurrencyPair{}, fmt.Errorf("invalid currency pair format: %s", pairStr)
//	}
//	// Construct the CurrencyPair object
//	return connecttypes.CurrencyPair{
//		Base:  parts[0],
//		Quote: parts[1],
//	}, nil
//}
