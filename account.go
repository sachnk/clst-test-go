// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clsttest

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/clst-test-go/internal/apijson"
	"github.com/stainless-sdks/clst-test-go/internal/param"
	"github.com/stainless-sdks/clst-test-go/internal/requestconfig"
	"github.com/stainless-sdks/clst-test-go/option"
	"github.com/stainless-sdks/clst-test-go/shared"
)

// AccountService contains methods and other services that help with interacting
// with the clearstreet API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountService] method instead.
type AccountService struct {
	Options []option.RequestOption
}

// NewAccountService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAccountService(opts ...option.RequestOption) (r *AccountService) {
	r = &AccountService{}
	r.Options = opts
	return
}

// Get an account by its ID.
func (r *AccountService) Get(ctx context.Context, accountID string, opts ...option.RequestOption) (res *Account, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all available accounts.
func (r *AccountService) List(ctx context.Context, opts ...option.RequestOption) (res *AccountListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Creates multiple orders in a single request, up to 1000. Note that a successful
// call to this endpoint does not necessarily mean your orders have been accepted,
// e.g. a downstream venue might reject your order. You should therefore utilize
// our WebSocket APIs to listen for changes in order lifecycle events.
//
// The response will contain an array of objects, indicating whether your order was
// submitted. If the order was submitted, the `order_id` field will be populated
// with the order ID assigned to this order. If the order was rejected, the
// `reason` field will be populated with the reason for rejection. The data array
// returned in the response object is guaranteed to be ordered in the same order as
// the orders you provided in the request. Again, note that even if your order was
// submitted, that doesn't mean it was _accepted_, and may still be rejected by
// downstream venues.
func (r *AccountService) NewOrdersInBulk(ctx context.Context, accountID string, body AccountNewOrdersInBulkParams, opts ...option.RequestOption) (res *AccountNewOrdersInBulkResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/bulk-orders", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List PNL details for a given account.
func (r *AccountService) GetPNLDetails(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountGetPNLDetailsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pnl-details", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get PNL summary for a given account.
func (r *AccountService) GetPNLSummary(ctx context.Context, accountID string, opts ...option.RequestOption) (res *shared.PNLSummary, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pnl-summary", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Account struct {
	// Account ID for the account.
	AccountID string `json:"account_id,required"`
	// Entity ID for the legal entity.
	EntityID string      `json:"entity_id,required"`
	Name     string      `json:"name,required"`
	JSON     accountJSON `json:"-"`
}

// accountJSON contains the JSON metadata for the struct [Account]
type accountJSON struct {
	AccountID   apijson.Field
	EntityID    apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Account) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountJSON) RawJSON() string {
	return r.raw
}

type PNLSummary struct {
	// Profit and loss from intraday trading activities.
	DayPNL float64 `json:"day_pnl,required"`
	// Entity ID for the legal entity.
	EntityID string `json:"entity_id,required"`
	// Net value of instruments held in the portfolio.
	Equity float64 `json:"equity,required"`
	// Absolute market value of long and short market values.
	GrossMarketValue float64 `json:"gross_market_value,required"`
	// Market value of securities positioned long.
	LongMarketValue float64 `json:"long_market_value,required"`
	// Market value net of long and short market values.
	NetMarketValue float64 `json:"net_market_value,required"`
	// `total_pnl + total_fees`
	NetPNL float64 `json:"net_pnl,required"`
	// Profit and loss from previous trading date.
	OvernightPNL float64 `json:"overnight_pnl,required"`
	// Profit and loss realized from position closing trading activity
	RealizedPNL float64 `json:"realized_pnl,required"`
	// Market value of securities positioned short.
	ShortMarketValue float64 `json:"short_market_value,required"`
	// Net value of instruments held in the portfolio at the start of a trading day.
	SodEquity float64 `json:"sod_equity,required"`
	// Absolute market value at the start of a trading day.
	SodGrossMarketValue float64 `json:"sod_gross_market_value,required"`
	// Market value of securities positioned long at the start of a trading day.
	SodLongMarketValue float64 `json:"sod_long_market_value,required"`
	// Market value of securities positioned short at the start of a trading day.
	SodShortMarketValue float64 `json:"sod_short_market_value,required"`
	// Milliseconds since epoch.
	Timestamp int64 `json:"timestamp,required"`
	// Total fees incurred from trading activities.
	TotalFees float64 `json:"total_fees,required"`
	// `realized_pnl + unrealized_pnl`
	TotalPNL float64 `json:"total_pnl,required"`
	// Profit and loss from market changes.
	UnrealizedPNL float64        `json:"unrealized_pnl,required"`
	JSON          pnlSummaryJSON `json:"-"`
}

// pnlSummaryJSON contains the JSON metadata for the struct [shared.PNLSummary]
type pnlSummaryJSON struct {
	DayPNL              apijson.Field
	EntityID            apijson.Field
	Equity              apijson.Field
	GrossMarketValue    apijson.Field
	LongMarketValue     apijson.Field
	NetMarketValue      apijson.Field
	NetPNL              apijson.Field
	OvernightPNL        apijson.Field
	RealizedPNL         apijson.Field
	ShortMarketValue    apijson.Field
	SodEquity           apijson.Field
	SodGrossMarketValue apijson.Field
	SodLongMarketValue  apijson.Field
	SodShortMarketValue apijson.Field
	Timestamp           apijson.Field
	TotalFees           apijson.Field
	TotalPNL            apijson.Field
	UnrealizedPNL       apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *shared.PNLSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pnlSummaryJSON) RawJSON() string {
	return r.raw
}

type AccountListResponse struct {
	Data []Account               `json:"data"`
	JSON accountListResponseJSON `json:"-"`
}

// accountListResponseJSON contains the JSON metadata for the struct
// [AccountListResponse]
type accountListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountNewOrdersInBulkResponse struct {
	// Array indicating whether each respective order was submitted or not. This array
	// is guaranteed to be sorted in the same order as the orders you provided in your
	// request.
	Data []AccountNewOrdersInBulkResponseData `json:"data,required"`
	// Total number of orders rejected
	Rejected int64 `json:"rejected,required"`
	// Total number of orders submitted
	Submitted int64                              `json:"submitted,required"`
	JSON      accountNewOrdersInBulkResponseJSON `json:"-"`
}

// accountNewOrdersInBulkResponseJSON contains the JSON metadata for the struct
// [AccountNewOrdersInBulkResponse]
type accountNewOrdersInBulkResponseJSON struct {
	Data        apijson.Field
	Rejected    apijson.Field
	Submitted   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountNewOrdersInBulkResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountNewOrdersInBulkResponseJSON) RawJSON() string {
	return r.raw
}

type AccountNewOrdersInBulkResponseData struct {
	// True if the order was submitted successfully, false otherwise.
	Submitted bool `json:"submitted,required"`
	// If the order was submitted, the order ID assigned to this order. Empty if the
	// order was rejected.
	OrderID string `json:"order_id"`
	// If the order rejected, the reason for rejection. Empty if the order was
	// accepted.
	Reason string                                 `json:"reason"`
	JSON   accountNewOrdersInBulkResponseDataJSON `json:"-"`
}

// accountNewOrdersInBulkResponseDataJSON contains the JSON metadata for the struct
// [AccountNewOrdersInBulkResponseData]
type accountNewOrdersInBulkResponseDataJSON struct {
	Submitted   apijson.Field
	OrderID     apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountNewOrdersInBulkResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountNewOrdersInBulkResponseDataJSON) RawJSON() string {
	return r.raw
}

type AccountGetPNLDetailsResponse struct {
	Data []AccountGetPNLDetailsResponseData `json:"data,required"`
	JSON accountGetPNLDetailsResponseJSON   `json:"-"`
}

// accountGetPNLDetailsResponseJSON contains the JSON metadata for the struct
// [AccountGetPNLDetailsResponse]
type accountGetPNLDetailsResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGetPNLDetailsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGetPNLDetailsResponseJSON) RawJSON() string {
	return r.raw
}

type AccountGetPNLDetailsResponseData struct {
	// Account ID for the account.
	AccountID string `json:"account_id,required"`
	// The asset class of the symbol.
	AssetClass AccountGetPNLDetailsResponseDataAssetClass `json:"asset_class,required"`
	// Quantity of a given instrument bought.
	BoughtQuantity string `json:"bought_quantity,required"`
	// Total buys of a given instrument.
	Buys int64 `json:"buys,required"`
	// Profit and loss from intraday trading activities.
	DayPNL float64 `json:"day_pnl,required"`
	// Name of the legal entity.
	EntityID string `json:"entity_id,required"`
	// Absolute market value of long and short market values.
	GrossMarketValue float64 `json:"gross_market_value,required"`
	// Market value net of long and short market values.
	NetMarketValue float64 `json:"net_market_value,required"`
	// Profit and loss from previous trading date.
	OvernightPNL float64 `json:"overnight_pnl,required"`
	// Price used in this pnl calculation.
	Price float64 `json:"price,required"`
	// String representation of quantity.
	Quantity string `json:"quantity,required"`
	// Profit and loss realized from position closing trading activity.
	RealizedPNL float64 `json:"realized_pnl,required"`
	// Total sells of a given instrument.
	Sells int64 `json:"sells,required"`
	// Market value of a given instrument a the start of a trading day.
	SodMarketValue float64 `json:"sod_market_value,required"`
	// Price at the start of a trading day.
	SodPrice float64 `json:"sod_price,required"`
	// Quantity of a given instrument at the start of a trading day.
	SodQuantity string `json:"sod_quantity,required"`
	// Quantity of a given instrument sold.
	SoldQuantity string `json:"sold_quantity,required"`
	Symbol       string `json:"symbol,required"`
	// Description of the symbol.
	SymbolDescription string `json:"symbol_description,required"`
	// Milliseconds since epoch.
	Timestamp int64 `json:"timestamp,required"`
	// Total fees incurred from trading activities.
	TotalFees float64 `json:"total_fees,required"`
	// `realized_pnl + unrealized_pnl`
	TotalPNL float64 `json:"total_pnl,required"`
	// The underlying instrument.
	Underlier string `json:"underlier,required"`
	// Profit and loss from market changes.
	UnrealizedPNL float64                              `json:"unrealized_pnl,required"`
	JSON          accountGetPNLDetailsResponseDataJSON `json:"-"`
}

// accountGetPNLDetailsResponseDataJSON contains the JSON metadata for the struct
// [AccountGetPNLDetailsResponseData]
type accountGetPNLDetailsResponseDataJSON struct {
	AccountID         apijson.Field
	AssetClass        apijson.Field
	BoughtQuantity    apijson.Field
	Buys              apijson.Field
	DayPNL            apijson.Field
	EntityID          apijson.Field
	GrossMarketValue  apijson.Field
	NetMarketValue    apijson.Field
	OvernightPNL      apijson.Field
	Price             apijson.Field
	Quantity          apijson.Field
	RealizedPNL       apijson.Field
	Sells             apijson.Field
	SodMarketValue    apijson.Field
	SodPrice          apijson.Field
	SodQuantity       apijson.Field
	SoldQuantity      apijson.Field
	Symbol            apijson.Field
	SymbolDescription apijson.Field
	Timestamp         apijson.Field
	TotalFees         apijson.Field
	TotalPNL          apijson.Field
	Underlier         apijson.Field
	UnrealizedPNL     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountGetPNLDetailsResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGetPNLDetailsResponseDataJSON) RawJSON() string {
	return r.raw
}

// The asset class of the symbol.
type AccountGetPNLDetailsResponseDataAssetClass string

const (
	AccountGetPNLDetailsResponseDataAssetClassOther  AccountGetPNLDetailsResponseDataAssetClass = "other"
	AccountGetPNLDetailsResponseDataAssetClassEquity AccountGetPNLDetailsResponseDataAssetClass = "equity"
	AccountGetPNLDetailsResponseDataAssetClassOption AccountGetPNLDetailsResponseDataAssetClass = "option"
	AccountGetPNLDetailsResponseDataAssetClassDebt   AccountGetPNLDetailsResponseDataAssetClass = "debt"
)

func (r AccountGetPNLDetailsResponseDataAssetClass) IsKnown() bool {
	switch r {
	case AccountGetPNLDetailsResponseDataAssetClassOther, AccountGetPNLDetailsResponseDataAssetClassEquity, AccountGetPNLDetailsResponseDataAssetClassOption, AccountGetPNLDetailsResponseDataAssetClassDebt:
		return true
	}
	return false
}

type AccountNewOrdersInBulkParams struct {
	// An array of orders to create.
	Orders param.Field[[]AccountNewOrdersInBulkParamsOrder] `json:"orders,required"`
}

func (r AccountNewOrdersInBulkParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountNewOrdersInBulkParamsOrder struct {
	// The type of order, can be one of the following:
	//
	//   - `limit`: A limit order will execute at-or-better than the limit price you
	//     specify
	//   - `market`: An order that will execute at the prevailing market prices
	OrderType param.Field[AccountNewOrdersInBulkParamsOrdersOrderType] `json:"order_type,required"`
	// The maximum quantity to be executed.
	Quantity param.Field[string] `json:"quantity,required"`
	// Buy, sell, sell-short indicator.
	Side param.Field[AccountNewOrdersInBulkParamsOrdersSide] `json:"side,required"`
	// Strategy type used for execution, can be one of below. Note, we use sensible
	// defaults for strategy parameters at the moment. In future, we will provide a way
	// to provide specify these parameters.
	//
	// - `sor`: Smart order router
	// - `dark`: Dark pool
	// - `ap`: Arrival price
	// - `pov`: Percentage of volume
	// - `twap`: Time weighted average price
	// - `vwap`: Volume weighted average price
	//
	// For more information on these strategies, please refer to our
	// [documentation](https://docs.clearstreet.io/studio/docs/execution-strategies).
	StrategyType param.Field[AccountNewOrdersInBulkParamsOrdersStrategyType] `json:"strategy_type,required"`
	// The symbol this order is for. See `symbol_format` for supported symbol formats.
	Symbol param.Field[string] `json:"symbol,required"`
	// The lifecycle enforcement of this order.
	//
	//   - `day`: The order will exist for the duration of the current trading session
	//   - `ioc`: The order will immediately be executed or cancelled
	//   - `day-plus`: The order will exist only for the duration the current trading
	//     session plus extended hours, if applicable
	//   - `at-open`: The order will exist only for the opening auction of the next
	//     session
	//   - `at-close`: The order will exist only for the closing auction of the current
	//     session
	TimeInForce param.Field[AccountNewOrdersInBulkParamsOrdersTimeInForce] `json:"time_in_force,required"`
	// If you're short-selling and using an away broker for a locate, provide the
	// broker name here.
	LocateBroker param.Field[string] `json:"locate_broker"`
	// The price to execute at-or-better.
	Price param.Field[string] `json:"price"`
	// An ID that you provide.
	ReferenceID param.Field[string] `json:"reference_id"`
	// Denotes the format of the provided `symbol` field.
	SymbolFormat param.Field[AccountNewOrdersInBulkParamsOrdersSymbolFormat] `json:"symbol_format"`
}

func (r AccountNewOrdersInBulkParamsOrder) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of order, can be one of the following:
//
//   - `limit`: A limit order will execute at-or-better than the limit price you
//     specify
//   - `market`: An order that will execute at the prevailing market prices
type AccountNewOrdersInBulkParamsOrdersOrderType string

const (
	AccountNewOrdersInBulkParamsOrdersOrderTypeLimit  AccountNewOrdersInBulkParamsOrdersOrderType = "limit"
	AccountNewOrdersInBulkParamsOrdersOrderTypeMarket AccountNewOrdersInBulkParamsOrdersOrderType = "market"
)

func (r AccountNewOrdersInBulkParamsOrdersOrderType) IsKnown() bool {
	switch r {
	case AccountNewOrdersInBulkParamsOrdersOrderTypeLimit, AccountNewOrdersInBulkParamsOrdersOrderTypeMarket:
		return true
	}
	return false
}

// Buy, sell, sell-short indicator.
type AccountNewOrdersInBulkParamsOrdersSide string

const (
	AccountNewOrdersInBulkParamsOrdersSideBuy       AccountNewOrdersInBulkParamsOrdersSide = "buy"
	AccountNewOrdersInBulkParamsOrdersSideSell      AccountNewOrdersInBulkParamsOrdersSide = "sell"
	AccountNewOrdersInBulkParamsOrdersSideSellShort AccountNewOrdersInBulkParamsOrdersSide = "sell-short"
)

func (r AccountNewOrdersInBulkParamsOrdersSide) IsKnown() bool {
	switch r {
	case AccountNewOrdersInBulkParamsOrdersSideBuy, AccountNewOrdersInBulkParamsOrdersSideSell, AccountNewOrdersInBulkParamsOrdersSideSellShort:
		return true
	}
	return false
}

// Strategy type used for execution, can be one of below. Note, we use sensible
// defaults for strategy parameters at the moment. In future, we will provide a way
// to provide specify these parameters.
//
// - `sor`: Smart order router
// - `dark`: Dark pool
// - `ap`: Arrival price
// - `pov`: Percentage of volume
// - `twap`: Time weighted average price
// - `vwap`: Volume weighted average price
//
// For more information on these strategies, please refer to our
// [documentation](https://docs.clearstreet.io/studio/docs/execution-strategies).
type AccountNewOrdersInBulkParamsOrdersStrategyType string

const (
	AccountNewOrdersInBulkParamsOrdersStrategyTypeSor  AccountNewOrdersInBulkParamsOrdersStrategyType = "sor"
	AccountNewOrdersInBulkParamsOrdersStrategyTypeDark AccountNewOrdersInBulkParamsOrdersStrategyType = "dark"
	AccountNewOrdersInBulkParamsOrdersStrategyTypeAp   AccountNewOrdersInBulkParamsOrdersStrategyType = "ap"
	AccountNewOrdersInBulkParamsOrdersStrategyTypePov  AccountNewOrdersInBulkParamsOrdersStrategyType = "pov"
	AccountNewOrdersInBulkParamsOrdersStrategyTypeTwap AccountNewOrdersInBulkParamsOrdersStrategyType = "twap"
	AccountNewOrdersInBulkParamsOrdersStrategyTypeVwap AccountNewOrdersInBulkParamsOrdersStrategyType = "vwap"
)

func (r AccountNewOrdersInBulkParamsOrdersStrategyType) IsKnown() bool {
	switch r {
	case AccountNewOrdersInBulkParamsOrdersStrategyTypeSor, AccountNewOrdersInBulkParamsOrdersStrategyTypeDark, AccountNewOrdersInBulkParamsOrdersStrategyTypeAp, AccountNewOrdersInBulkParamsOrdersStrategyTypePov, AccountNewOrdersInBulkParamsOrdersStrategyTypeTwap, AccountNewOrdersInBulkParamsOrdersStrategyTypeVwap:
		return true
	}
	return false
}

// The lifecycle enforcement of this order.
//
//   - `day`: The order will exist for the duration of the current trading session
//   - `ioc`: The order will immediately be executed or cancelled
//   - `day-plus`: The order will exist only for the duration the current trading
//     session plus extended hours, if applicable
//   - `at-open`: The order will exist only for the opening auction of the next
//     session
//   - `at-close`: The order will exist only for the closing auction of the current
//     session
type AccountNewOrdersInBulkParamsOrdersTimeInForce string

const (
	AccountNewOrdersInBulkParamsOrdersTimeInForceDay     AccountNewOrdersInBulkParamsOrdersTimeInForce = "day"
	AccountNewOrdersInBulkParamsOrdersTimeInForceIoc     AccountNewOrdersInBulkParamsOrdersTimeInForce = "ioc"
	AccountNewOrdersInBulkParamsOrdersTimeInForceDayPlus AccountNewOrdersInBulkParamsOrdersTimeInForce = "day-plus"
	AccountNewOrdersInBulkParamsOrdersTimeInForceAtOpen  AccountNewOrdersInBulkParamsOrdersTimeInForce = "at-open"
	AccountNewOrdersInBulkParamsOrdersTimeInForceAtClose AccountNewOrdersInBulkParamsOrdersTimeInForce = "at-close"
)

func (r AccountNewOrdersInBulkParamsOrdersTimeInForce) IsKnown() bool {
	switch r {
	case AccountNewOrdersInBulkParamsOrdersTimeInForceDay, AccountNewOrdersInBulkParamsOrdersTimeInForceIoc, AccountNewOrdersInBulkParamsOrdersTimeInForceDayPlus, AccountNewOrdersInBulkParamsOrdersTimeInForceAtOpen, AccountNewOrdersInBulkParamsOrdersTimeInForceAtClose:
		return true
	}
	return false
}

// Denotes the format of the provided `symbol` field.
type AccountNewOrdersInBulkParamsOrdersSymbolFormat string

const (
	AccountNewOrdersInBulkParamsOrdersSymbolFormatCms AccountNewOrdersInBulkParamsOrdersSymbolFormat = "cms"
	AccountNewOrdersInBulkParamsOrdersSymbolFormatOsi AccountNewOrdersInBulkParamsOrdersSymbolFormat = "osi"
)

func (r AccountNewOrdersInBulkParamsOrdersSymbolFormat) IsKnown() bool {
	switch r {
	case AccountNewOrdersInBulkParamsOrdersSymbolFormatCms, AccountNewOrdersInBulkParamsOrdersSymbolFormatOsi:
		return true
	}
	return false
}
