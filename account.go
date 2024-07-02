// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clsttest

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/sachnk/clst-test-go/internal/apijson"
	"github.com/sachnk/clst-test-go/internal/param"
	"github.com/sachnk/clst-test-go/internal/requestconfig"
	"github.com/sachnk/clst-test-go/option"
	"github.com/sachnk/clst-test-go/shared"
)

// AccountService contains methods and other services that help with interacting
// with the clearstreet API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountService] method instead.
type AccountService struct {
	Options      []option.RequestOption
	Orders       *AccountOrderService
	Trades       *AccountTradeService
	Positions    *AccountPositionService
	LocateOrders *AccountLocateOrderService
	EasyBorrows  *AccountEasyBorrowService
}

// NewAccountService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAccountService(opts ...option.RequestOption) (r *AccountService) {
	r = &AccountService{}
	r.Options = opts
	r.Orders = NewAccountOrderService(opts...)
	r.Trades = NewAccountTradeService(opts...)
	r.Positions = NewAccountPositionService(opts...)
	r.LocateOrders = NewAccountLocateOrderService(opts...)
	r.EasyBorrows = NewAccountEasyBorrowService(opts...)
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
func (r *AccountService) GetPNLSummary(ctx context.Context, accountID string, opts ...option.RequestOption) (res *PNLSummaryForAccount, err error) {
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

type LocateOrder struct {
	// Account ID for the account.
	AccountID string `json:"account_id,required"`
	// Unique locate ID assigned by us.
	LocateOrderID string `json:"locate_order_id,required"`
	// Unique MPID assigned by us.
	Mpid string `json:"mpid,required"`
	// The timestamp indicating when the locate order was requested.
	RequestedAt int64 `json:"requested_at,required"`
	// String representation of quantity.
	RequestedQuantity string `json:"requested_quantity,required"`
	// The status of the locate order.
	Status LocateOrderStatus `json:"status,required"`
	Symbol string            `json:"symbol,required"`
	// The timestamp indicating when the locate order was last updated.
	UpdatedAt int64 `json:"updated_at,required"`
	// The rate charged if the instrument is held overnight.
	BorrowRate string `json:"borrow_rate"`
	// Comment from the desk.
	DeskComment string `json:"desk_comment"`
	// The timestamp indicating when the locate-order will expire.
	ExpiresAt int64 `json:"expires_at"`
	// The locate ID, available once the locate order has been offered
	LocateID string `json:"locate_id"`
	// The timestamp indicating when the locate-order was located.
	LocatedAt int64 `json:"located_at"`
	// The quantity that has been located.
	LocatedQuantity string `json:"located_quantity"`
	// The reference ID provided by you.
	ReferenceID string `json:"reference_id"`
	// The total cost of the locate.
	TotalCost string `json:"total_cost"`
	// Comment from the trader.
	TraderComment string          `json:"trader_comment"`
	JSON          locateOrderJSON `json:"-"`
}

// locateOrderJSON contains the JSON metadata for the struct [LocateOrder]
type locateOrderJSON struct {
	AccountID         apijson.Field
	LocateOrderID     apijson.Field
	Mpid              apijson.Field
	RequestedAt       apijson.Field
	RequestedQuantity apijson.Field
	Status            apijson.Field
	Symbol            apijson.Field
	UpdatedAt         apijson.Field
	BorrowRate        apijson.Field
	DeskComment       apijson.Field
	ExpiresAt         apijson.Field
	LocateID          apijson.Field
	LocatedAt         apijson.Field
	LocatedQuantity   apijson.Field
	ReferenceID       apijson.Field
	TotalCost         apijson.Field
	TraderComment     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *LocateOrder) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r locateOrderJSON) RawJSON() string {
	return r.raw
}

// The status of the locate order.
type LocateOrderStatus string

const (
	LocateOrderStatusPending   LocateOrderStatus = "pending"
	LocateOrderStatusOffered   LocateOrderStatus = "offered"
	LocateOrderStatusFilled    LocateOrderStatus = "filled"
	LocateOrderStatusRejected  LocateOrderStatus = "rejected"
	LocateOrderStatusDeclined  LocateOrderStatus = "declined"
	LocateOrderStatusExpired   LocateOrderStatus = "expired"
	LocateOrderStatusCancelled LocateOrderStatus = "cancelled"
)

func (r LocateOrderStatus) IsKnown() bool {
	switch r {
	case LocateOrderStatusPending, LocateOrderStatusOffered, LocateOrderStatusFilled, LocateOrderStatusRejected, LocateOrderStatusDeclined, LocateOrderStatusExpired, LocateOrderStatusCancelled:
		return true
	}
	return false
}

type Order struct {
	// Account ID for the account.
	AccountID string `json:"account_id,required"`
	// When the order was created in milliseconds since epoch.
	CreatedAt int64 `json:"created_at,required"`
	// The quantity that has been filled.
	FilledQuantity string `json:"filled_quantity,required"`
	// An internally generated unique ID for this order.
	OrderID string `json:"order_id,required"`
	// The type of order, can be one of the following:
	//
	//   - `limit`: A limit order will execute at-or-better than the limit price you
	//     specify
	//   - `market`: An order that will execute at the prevailing market prices
	OrderType OrderOrderType `json:"order_type,required"`
	// The requested quantity on this order.
	Quantity string `json:"quantity,required"`
	// Buy, sell, sell-short indicator.
	Side OrderSide `json:"side,required"`
	// Simplified order state, which is inferred from `OrderStatus`. Makes it easier to
	// determine whether an order can be executed against.
	//
	//   - `open`: Order _can_ potentially be executed against.
	//   - `rejected`: Order _cannot_ be executed against because it was rejected. This
	//     is a terminal state.
	//   - `closed`: Order _cannot_ be executed against. This is a terminal state.
	State OrderState `json:"state,required"`
	// Granular order status using
	// [standard values come FIX tag 39](https://www.fixtrading.org/online-specification/order-state-changes).
	Status OrderStatus `json:"status,required"`
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
	StrategyType OrderStrategyType `json:"strategy_type,required"`
	Symbol       string            `json:"symbol,required"`
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
	TimeInForce OrderTimeInForce `json:"time_in_force,required"`
	// When the order was updated in milliseconds since epoch.
	UpdatedAt int64 `json:"updated_at,required"`
	// A monotonically increasing number indicating the version of this order. A higher
	// number indicates a more recent version of the order.
	Version int64 `json:"version,required"`
	// Calculated average price of all fills on this order.
	AveragePrice float64 `json:"average_price"`
	// The last reason why this order was updated
	OrderUpdateReason OrderOrderUpdateReason `json:"order_update_reason"`
	// The requsted limit price on this order.
	Price string `json:"price"`
	// The ID you provided when creating this order.
	ReferenceID string `json:"reference_id"`
	// Free form text typically contains reasons for a reject.
	Text string    `json:"text"`
	JSON orderJSON `json:"-"`
}

// orderJSON contains the JSON metadata for the struct [Order]
type orderJSON struct {
	AccountID         apijson.Field
	CreatedAt         apijson.Field
	FilledQuantity    apijson.Field
	OrderID           apijson.Field
	OrderType         apijson.Field
	Quantity          apijson.Field
	Side              apijson.Field
	State             apijson.Field
	Status            apijson.Field
	StrategyType      apijson.Field
	Symbol            apijson.Field
	TimeInForce       apijson.Field
	UpdatedAt         apijson.Field
	Version           apijson.Field
	AveragePrice      apijson.Field
	OrderUpdateReason apijson.Field
	Price             apijson.Field
	ReferenceID       apijson.Field
	Text              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *Order) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orderJSON) RawJSON() string {
	return r.raw
}

// The type of order, can be one of the following:
//
//   - `limit`: A limit order will execute at-or-better than the limit price you
//     specify
//   - `market`: An order that will execute at the prevailing market prices
type OrderOrderType string

const (
	OrderOrderTypeLimit  OrderOrderType = "limit"
	OrderOrderTypeMarket OrderOrderType = "market"
)

func (r OrderOrderType) IsKnown() bool {
	switch r {
	case OrderOrderTypeLimit, OrderOrderTypeMarket:
		return true
	}
	return false
}

// Buy, sell, sell-short indicator.
type OrderSide string

const (
	OrderSideBuy       OrderSide = "buy"
	OrderSideSell      OrderSide = "sell"
	OrderSideSellShort OrderSide = "sell-short"
)

func (r OrderSide) IsKnown() bool {
	switch r {
	case OrderSideBuy, OrderSideSell, OrderSideSellShort:
		return true
	}
	return false
}

// Simplified order state, which is inferred from `OrderStatus`. Makes it easier to
// determine whether an order can be executed against.
//
//   - `open`: Order _can_ potentially be executed against.
//   - `rejected`: Order _cannot_ be executed against because it was rejected. This
//     is a terminal state.
//   - `closed`: Order _cannot_ be executed against. This is a terminal state.
type OrderState string

const (
	OrderStateOpen     OrderState = "open"
	OrderStateRejected OrderState = "rejected"
	OrderStateClosed   OrderState = "closed"
)

func (r OrderState) IsKnown() bool {
	switch r {
	case OrderStateOpen, OrderStateRejected, OrderStateClosed:
		return true
	}
	return false
}

// Granular order status using
// [standard values come FIX tag 39](https://www.fixtrading.org/online-specification/order-state-changes).
type OrderStatus string

const (
	OrderStatusNew                OrderStatus = "new"
	OrderStatusPartiallyFilled    OrderStatus = "partially-filled"
	OrderStatusFilled             OrderStatus = "filled"
	OrderStatusCanceled           OrderStatus = "canceled"
	OrderStatusReplaced           OrderStatus = "replaced"
	OrderStatusPendingCancel      OrderStatus = "pending-cancel"
	OrderStatusStopped            OrderStatus = "stopped"
	OrderStatusRejected           OrderStatus = "rejected"
	OrderStatusSuspended          OrderStatus = "suspended"
	OrderStatusPendingNew         OrderStatus = "pending-new"
	OrderStatusCalculated         OrderStatus = "calculated"
	OrderStatusExpired            OrderStatus = "expired"
	OrderStatusAcceptedForBidding OrderStatus = "accepted-for-bidding"
	OrderStatusPendingReplace     OrderStatus = "pending-replace"
	OrderStatusDoneForDay         OrderStatus = "done-for-day"
)

func (r OrderStatus) IsKnown() bool {
	switch r {
	case OrderStatusNew, OrderStatusPartiallyFilled, OrderStatusFilled, OrderStatusCanceled, OrderStatusReplaced, OrderStatusPendingCancel, OrderStatusStopped, OrderStatusRejected, OrderStatusSuspended, OrderStatusPendingNew, OrderStatusCalculated, OrderStatusExpired, OrderStatusAcceptedForBidding, OrderStatusPendingReplace, OrderStatusDoneForDay:
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
type OrderStrategyType string

const (
	OrderStrategyTypeSor  OrderStrategyType = "sor"
	OrderStrategyTypeDark OrderStrategyType = "dark"
	OrderStrategyTypeAp   OrderStrategyType = "ap"
	OrderStrategyTypePov  OrderStrategyType = "pov"
	OrderStrategyTypeTwap OrderStrategyType = "twap"
	OrderStrategyTypeVwap OrderStrategyType = "vwap"
)

func (r OrderStrategyType) IsKnown() bool {
	switch r {
	case OrderStrategyTypeSor, OrderStrategyTypeDark, OrderStrategyTypeAp, OrderStrategyTypePov, OrderStrategyTypeTwap, OrderStrategyTypeVwap:
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
type OrderTimeInForce string

const (
	OrderTimeInForceDay     OrderTimeInForce = "day"
	OrderTimeInForceIoc     OrderTimeInForce = "ioc"
	OrderTimeInForceDayPlus OrderTimeInForce = "day-plus"
	OrderTimeInForceAtOpen  OrderTimeInForce = "at-open"
	OrderTimeInForceAtClose OrderTimeInForce = "at-close"
)

func (r OrderTimeInForce) IsKnown() bool {
	switch r {
	case OrderTimeInForceDay, OrderTimeInForceIoc, OrderTimeInForceDayPlus, OrderTimeInForceAtOpen, OrderTimeInForceAtClose:
		return true
	}
	return false
}

// The last reason why this order was updated
type OrderOrderUpdateReason string

const (
	OrderOrderUpdateReasonPlace           OrderOrderUpdateReason = "place"
	OrderOrderUpdateReasonModify          OrderOrderUpdateReason = "modify"
	OrderOrderUpdateReasonCancel          OrderOrderUpdateReason = "cancel"
	OrderOrderUpdateReasonExecutionReport OrderOrderUpdateReason = "execution-report"
	OrderOrderUpdateReasonCancelReject    OrderOrderUpdateReason = "cancel-reject"
)

func (r OrderOrderUpdateReason) IsKnown() bool {
	switch r {
	case OrderOrderUpdateReasonPlace, OrderOrderUpdateReasonModify, OrderOrderUpdateReasonCancel, OrderOrderUpdateReasonExecutionReport, OrderOrderUpdateReasonCancelReject:
		return true
	}
	return false
}

type PNLSummaryForAccount struct {
	AccountID string                   `json:"account_id,required"`
	JSON      pnlSummaryForAccountJSON `json:"-"`
	shared.PNLSummary
}

// pnlSummaryForAccountJSON contains the JSON metadata for the struct
// [PNLSummaryForAccount]
type pnlSummaryForAccountJSON struct {
	AccountID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PNLSummaryForAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pnlSummaryForAccountJSON) RawJSON() string {
	return r.raw
}

type Position struct {
	// Account ID for the account.
	AccountID string `json:"account_id"`
	// String representation of quantity.
	Quantity string       `json:"quantity"`
	Symbol   string       `json:"symbol"`
	JSON     positionJSON `json:"-"`
}

// positionJSON contains the JSON metadata for the struct [Position]
type positionJSON struct {
	AccountID   apijson.Field
	Quantity    apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Position) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r positionJSON) RawJSON() string {
	return r.raw
}

type Trade struct {
	// When this trade happened in milliseconds since epoch.
	CreatedAt int64 `json:"created_at,required"`
	// The order ID of the order this trade occurred on.
	OrderID string `json:"order_id,required"`
	// The traded price.
	Price string `json:"price,required"`
	// The amount that was traded.
	Quantity string `json:"quantity,required"`
	// The side this trade occurred on.
	Side TradeSide `json:"side,required"`
	// Unique trade ID assigned by us.
	TradeID string `json:"trade_id,required"`
	// Account ID for the account.
	AccountID string `json:"account_id"`
	// The symbol this trade was for.
	Symbol string    `json:"symbol"`
	JSON   tradeJSON `json:"-"`
}

// tradeJSON contains the JSON metadata for the struct [Trade]
type tradeJSON struct {
	CreatedAt   apijson.Field
	OrderID     apijson.Field
	Price       apijson.Field
	Quantity    apijson.Field
	Side        apijson.Field
	TradeID     apijson.Field
	AccountID   apijson.Field
	Symbol      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Trade) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tradeJSON) RawJSON() string {
	return r.raw
}

// The side this trade occurred on.
type TradeSide string

const (
	TradeSideBuy       TradeSide = "buy"
	TradeSideSell      TradeSide = "sell"
	TradeSideSellShort TradeSide = "sell-short"
)

func (r TradeSide) IsKnown() bool {
	switch r {
	case TradeSideBuy, TradeSideSell, TradeSideSellShort:
		return true
	}
	return false
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
