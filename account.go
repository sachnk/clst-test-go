// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clsttest

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/clst-test-go/internal/apijson"
	"github.com/stainless-sdks/clst-test-go/internal/requestconfig"
	"github.com/stainless-sdks/clst-test-go/option"
)

// AccountService contains methods and other services that help with interacting
// with the clearstreet API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountService] method instead.
type AccountService struct {
	Options      []option.RequestOption
	BulkOrders   *AccountBulkOrderService
	Orders       *AccountOrderService
	Trades       *AccountTradeService
	Positions    *AccountPositionService
	LocateOrders *AccountLocateOrderService
	EasyBorrows  *AccountEasyBorrowService
	PnlSummary   *AccountPnlSummaryService
	PnlDetails   *AccountPnlDetailService
}

// NewAccountService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAccountService(opts ...option.RequestOption) (r *AccountService) {
	r = &AccountService{}
	r.Options = opts
	r.BulkOrders = NewAccountBulkOrderService(opts...)
	r.Orders = NewAccountOrderService(opts...)
	r.Trades = NewAccountTradeService(opts...)
	r.Positions = NewAccountPositionService(opts...)
	r.LocateOrders = NewAccountLocateOrderService(opts...)
	r.EasyBorrows = NewAccountEasyBorrowService(opts...)
	r.PnlSummary = NewAccountPnlSummaryService(opts...)
	r.PnlDetails = NewAccountPnlDetailService(opts...)
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

type PnlSummary struct {
	// Profit and loss from intraday trading activities.
	DayPnl float64 `json:"day_pnl,required"`
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
	NetPnl float64 `json:"net_pnl,required"`
	// Profit and loss from previous trading date.
	OvernightPnl float64 `json:"overnight_pnl,required"`
	// Profit and loss realized from position closing trading activity
	RealizedPnl float64 `json:"realized_pnl,required"`
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
	TotalPnl float64 `json:"total_pnl,required"`
	// Profit and loss from market changes.
	UnrealizedPnl float64        `json:"unrealized_pnl,required"`
	JSON          pnlSummaryJSON `json:"-"`
}

// pnlSummaryJSON contains the JSON metadata for the struct [PnlSummary]
type pnlSummaryJSON struct {
	DayPnl              apijson.Field
	EntityID            apijson.Field
	Equity              apijson.Field
	GrossMarketValue    apijson.Field
	LongMarketValue     apijson.Field
	NetMarketValue      apijson.Field
	NetPnl              apijson.Field
	OvernightPnl        apijson.Field
	RealizedPnl         apijson.Field
	ShortMarketValue    apijson.Field
	SodEquity           apijson.Field
	SodGrossMarketValue apijson.Field
	SodLongMarketValue  apijson.Field
	SodShortMarketValue apijson.Field
	Timestamp           apijson.Field
	TotalFees           apijson.Field
	TotalPnl            apijson.Field
	UnrealizedPnl       apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *PnlSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pnlSummaryJSON) RawJSON() string {
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
