package app

import (
	"context"
	"fmt"
	"math/big"
	"slices"
	"strings"
	"time"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/server/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/skip-mev/connect/v2/cmd/constants/marketmaps"
	oracleconfig "github.com/skip-mev/connect/v2/oracle/config"
	connecttypes "github.com/skip-mev/connect/v2/pkg/types"
	oracleclient "github.com/skip-mev/connect/v2/service/clients/oracle"
	servicemetrics "github.com/skip-mev/connect/v2/service/metrics"
	servicetypes "github.com/skip-mev/connect/v2/service/servers/oracle/types"
	oracletypes "github.com/skip-mev/connect/v2/x/oracle/types"
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

//func (app *App) initializeABCIExtensions(oracleClient oracleclient.OracleClient, oracleMetrics servicemetrics.Metrics) {
//	// Create the proposal handler that will be used to fill proposals with
//	// transactions and oracle data.
//	proposalHandler := proposals.NewProposalHandler(
//		app.Logger(),
//		baseapp.NoOpPrepareProposal(),
//		baseapp.NoOpProcessProposal(),
//		ve.NewDefaultValidateVoteExtensionsFn(app.StakingKeeper),
//		compression.NewCompressionVoteExtensionCodec(
//			compression.NewDefaultVoteExtensionCodec(),
//			compression.NewZLibCompressor(),
//		),
//		compression.NewCompressionExtendedCommitCodec(
//			compression.NewDefaultExtendedCommitCodec(),
//			compression.NewZStdCompressor(),
//		),
//		currencypair.NewDeltaCurrencyPairStrategy(app.OracleKeeper),
//		oracleMetrics,
//	)
//	app.SetPrepareProposal(proposalHandler.PrepareProposalHandler())
//	app.SetProcessProposal(proposalHandler.ProcessProposalHandler())
//
//	// Create the aggregation function that will be used to aggregate oracle data
//	// from each validator.
//	aggregatorFn := voteweighted.MedianFromContext(
//		app.Logger(),
//		app.StakingKeeper,
//		voteweighted.DefaultPowerThreshold,
//	)
//	veCodec := compression.NewCompressionVoteExtensionCodec(
//		compression.NewDefaultVoteExtensionCodec(),
//		compression.NewZLibCompressor(),
//	)
//	ecCodec := compression.NewCompressionExtendedCommitCodec(
//		compression.NewDefaultExtendedCommitCodec(),
//		compression.NewZStdCompressor(),
//	)
//
//	// Create the pre-finalize block hook that will be used to apply oracle data
//	// to the state before any transactions are executed (in finalize block).
//	oraclePreBlockHandler := oraclepreblock.NewOraclePreBlockHandler(
//		app.Logger(),
//		aggregatorFn,
//		app.OracleKeeper,
//		oracleMetrics,
//		currencypair.NewDeltaCurrencyPairStrategy(app.OracleKeeper), // IMPORTANT: always construct new currency pair strategy objects when functions require them as arguments.
//		veCodec,
//		ecCodec,
//	)
//
//	app.SetPreBlocker(oraclePreBlockHandler.WrappedPreBlocker(app.ModuleManager))
//
//	// Create the vote extensions handler that will be used to extend and verify
//	// vote extensions (i.e. oracle data).
//	voteExtensionsHandler := ve.NewVoteExtensionHandler(
//		app.Logger(),
//		oracleClient,
//		time.Second, // timeout
//		currencypair.NewDeltaCurrencyPairStrategy(app.OracleKeeper), // IMPORTANT: always construct new currency pair strategy objects when functions require them as arguments.
//		veCodec,
//		aggregator.NewOraclePriceApplier(
//			aggregator.NewDefaultVoteAggregator(
//				app.Logger(),
//				aggregatorFn,
//				// we need a separate price strategy here, so that we can optimistically apply the latest prices
//				// and extend our vote based on these prices
//				currencypair.NewDeltaCurrencyPairStrategy(app.OracleKeeper), // IMPORTANT: always construct new currency pair strategy objects when functions require them as arguments.
//			),
//			app.OracleKeeper,
//			veCodec,
//			ecCodec,
//			app.Logger(),
//		),
//		oracleMetrics,
//	)
//	app.SetExtendVoteHandler(voteExtensionsHandler.ExtendVoteHandler())
//	app.SetVerifyVoteExtensionHandler(voteExtensionsHandler.VerifyVoteExtensionHandler())
//}

// fetchAndStoreOracleData fetches the latest data from the Oracle client and stores it in the chain's state.
func (app *App) fetchAndStoreOracleData(ctx sdk.Context) error {
	// Ensure the Oracle client has been initialized
	if app.oracleClient == nil {
		return fmt.Errorf("oracle client not initialized")
	}

	// Create a request context with a 1-second timeout
	reqCtx, cancel := context.WithTimeout(ctx.Context(), time.Second*5)
	defer cancel()

	// Fetch the latest prices from the Oracle client
	oracleResp, err := app.oracleClient.Prices(ctx.WithContext(reqCtx), &servicetypes.QueryPricesRequest{})
	if err != nil {
		app.Logger().Error(
			"failed to retrieve oracle prices",
			"ctx_err", reqCtx.Err(),
			"err", err,
		)
		return err
	}

	// Handle a nil response from the Oracle client
	if oracleResp == nil {
		return fmt.Errorf("oracle returned nil prices")
	}

	// Convert Oracle response to a map of currency pairs to prices
	prices, err := ConvertOraclePrices(oracleResp)
	if err != nil {
		app.Logger().Error("failed to convert oracle prices", "err", err)
		return err
	}

	// Retrieve all currency pairs from the OracleKeeper
	currencyPairs := app.OracleKeeper.GetAllCurrencyPairs(ctx)

	// Iterate over currency pairs and process prices
	for _, cp := range currencyPairs {
		price, exists := prices[cp]

		// Log and skip if price is missing or nil
		if !exists || price == nil {
			app.Logger().Debug("no price for currency pair", "currency_pair", cp.String())
			continue
		}

		// Log and skip if price is negative
		if price.Sign() < 0 {
			app.Logger().Error("price is negative", "currency_pair", cp.String(), "price", price.String())
			continue
		}

		// Create a QuotePrice and store it
		quotePrice := oracletypes.QuotePrice{
			Price:          math.NewIntFromBigInt(price),
			BlockTimestamp: ctx.BlockHeader().Time,
			BlockHeight:    uint64(ctx.BlockHeight()),
		}

		// Store the price using the OracleKeeper, logging any errors
		if err := app.OracleKeeper.SetPriceForCurrencyPair(ctx, cp, quotePrice); err != nil {
			app.Logger().Error(
				"failed to set price for currency pair",
				"currency_pair", cp.String(),
				"quote_price", quotePrice.Price.String(),
				"err", err,
			)
			return err
		}

		// Log successful price storage
		app.Logger().Debug(
			"set price for currency pair",
			"currency_pair", cp.String(),
			"quote_price", quotePrice.Price.String(),
		)
	}

	return nil
}

func (app *App) setupMarkets(ctx sdk.Context) error {
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

// ConvertOraclePrices converts oracleResp.Prices to map[connecttypes.CurrencyPair]*big.Int
func ConvertOraclePrices(oracleResp *servicetypes.QueryPricesResponse) (map[connecttypes.CurrencyPair]*big.Int, error) {
	result := make(map[connecttypes.CurrencyPair]*big.Int)

	// Iterate over the prices in the oracle response
	for pairStr, priceStr := range oracleResp.Prices {
		// Parse the string into a CurrencyPair
		currencyPair, err := parseCurrencyPair(pairStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse currency pair %s: %w", pairStr, err)
		}

		// Convert the price from string to *big.Int
		priceBigInt, ok := new(big.Int).SetString(priceStr, 10)
		if !ok {
			return nil, fmt.Errorf("failed to convert price %s to big.Int", priceStr)
		}

		// Store the result in the map
		result[currencyPair] = priceBigInt
	}

	return result, nil
}

// parseCurrencyPair is a helper function that converts a string representation of a currency pair to connecttypes.CurrencyPair
// More tests are needed
func parseCurrencyPair(pairStr string) (connecttypes.CurrencyPair, error) {
	parts := strings.Split(pairStr, "/")
	if len(parts) != 2 {
		return connecttypes.CurrencyPair{}, fmt.Errorf("invalid currency pair format: %s", pairStr)
	}
	// Construct the CurrencyPair object
	return connecttypes.CurrencyPair{
		Base:  parts[0],
		Quote: parts[1],
	}, nil
}
