# Entities

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#Entity">Entity</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#PNLSummary">PNLSummary</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#PortfolioMargin">PortfolioMargin</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#RegtMargin">RegtMargin</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#EntityListResponse">EntityListResponse</a>

Methods:

- <code title="get /entities/{entity_id}">client.Entities.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#EntityService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, entityID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#Entity">Entity</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /entities">client.Entities.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#EntityService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#EntityListResponse">EntityListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /entities/{entity_id}/pnl-summary">client.Entities.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#EntityService.GetPNLSummary">GetPNLSummary</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, entityID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#PNLSummary">PNLSummary</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /entities/{entity_id}/portfolio-margin">client.Entities.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#EntityService.GetPortfolioMargin">GetPortfolioMargin</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, entityID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#PortfolioMargin">PortfolioMargin</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /entities/{entity_id}/regt-margin">client.Entities.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#EntityService.GetRegtMargin">GetRegtMargin</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, entityID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#RegtMargin">RegtMargin</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## RegtMarginSimulations

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#SimulationIDParam">SimulationIDParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#RegtMarginSimulation">RegtMarginSimulation</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#SimulationID">SimulationID</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#EntityRegtMarginSimulationNewResponse">EntityRegtMarginSimulationNewResponse</a>

Methods:

- <code title="post /entities/{entity_id}/regt-margin-simulations">client.Entities.RegtMarginSimulations.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#EntityRegtMarginSimulationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, entityID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#EntityRegtMarginSimulationNewParams">EntityRegtMarginSimulationNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#EntityRegtMarginSimulationNewResponse">EntityRegtMarginSimulationNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /entities/{entity_id}/regt-margin-simulations/{simulation_id}">client.Entities.RegtMarginSimulations.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#EntityRegtMarginSimulationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, entityID <a href="https://pkg.go.dev/builtin#string">string</a>, simulationID <a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#SimulationIDParam">SimulationIDParam</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#RegtMarginSimulation">RegtMarginSimulation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Accounts

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#Account">Account</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#LocateOrder">LocateOrder</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#Order">Order</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#PNLSummary">PNLSummary</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#Position">Position</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#Trade">Trade</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#AccountListResponse">AccountListResponse</a>

Methods:

- <code title="get /accounts/{account_id}">client.Accounts.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#AccountService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#Account">Account</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts">client.Accounts.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#AccountService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#AccountListResponse">AccountListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## BulkOrders

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#AccountBulkOrderNewResponse">AccountBulkOrderNewResponse</a>

Methods:

- <code title="post /accounts/{account_id}/bulk-orders">client.Accounts.BulkOrders.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#AccountBulkOrderService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#AccountBulkOrderNewParams">AccountBulkOrderNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#AccountBulkOrderNewResponse">AccountBulkOrderNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Orders

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#AccountOrderNewResponse">AccountOrderNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#AccountOrderGetResponse">AccountOrderGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#AccountOrderListResponse">AccountOrderListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#AccountOrderDeleteResponse">AccountOrderDeleteResponse</a>

Methods:

- <code title="post /accounts/{account_id}/orders">client.Accounts.Orders.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#AccountOrderService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#AccountOrderNewParams">AccountOrderNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#AccountOrderNewResponse">AccountOrderNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_id}/orders/{order_id}">client.Accounts.Orders.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#AccountOrderService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, orderID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#AccountOrderGetResponse">AccountOrderGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_id}/orders">client.Accounts.Orders.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#AccountOrderService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#AccountOrderListParams">AccountOrderListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#AccountOrderListResponse">AccountOrderListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /accounts/{account_id}/orders">client.Accounts.Orders.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#AccountOrderService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#AccountOrderDeleteParams">AccountOrderDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#AccountOrderDeleteResponse">AccountOrderDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /accounts/{account_id}/orders/{order_id}">client.Accounts.Orders.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#AccountOrderService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, orderID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Trades

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#AccountTradeListResponse">AccountTradeListResponse</a>

Methods:

- <code title="get /accounts/{account_id}/trades/{trade_id}">client.Accounts.Trades.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#AccountTradeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, tradeID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#Trade">Trade</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_id}/trades">client.Accounts.Trades.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#AccountTradeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#AccountTradeListParams">AccountTradeListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#AccountTradeListResponse">AccountTradeListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Positions

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#AccountPositionListResponse">AccountPositionListResponse</a>

Methods:

- <code title="get /accounts/{account_id}/positions/{symbol}">client.Accounts.Positions.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#AccountPositionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, symbol <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#Position">Position</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_id}/positions">client.Accounts.Positions.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#AccountPositionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#AccountPositionListParams">AccountPositionListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#AccountPositionListResponse">AccountPositionListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## LocateOrders

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#AccountLocateOrderListResponse">AccountLocateOrderListResponse</a>

Methods:

- <code title="post /accounts/{account_id}/locate-orders">client.Accounts.LocateOrders.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#AccountLocateOrderService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#AccountLocateOrderNewParams">AccountLocateOrderNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#LocateOrder">LocateOrder</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{account_id}/locate-orders/{locate_order_id}">client.Accounts.LocateOrders.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#AccountLocateOrderService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, locateOrderID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#LocateOrder">LocateOrder</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /accounts/{account_id}/locate-orders/{locate_order_id}">client.Accounts.LocateOrders.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#AccountLocateOrderService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, locateOrderID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#AccountLocateOrderUpdateParams">AccountLocateOrderUpdateParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /accounts/{account_id}/locate-orders">client.Accounts.LocateOrders.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#AccountLocateOrderService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#AccountLocateOrderListResponse">AccountLocateOrderListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## EasyBorrows

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#AccountEasyBorrowListResponse">AccountEasyBorrowListResponse</a>

Methods:

- <code title="get /accounts/{account_id}/easy-borrows">client.Accounts.EasyBorrows.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#AccountEasyBorrowService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#AccountEasyBorrowListResponse">AccountEasyBorrowListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## PNLSummary

Methods:

- <code title="get /accounts/{account_id}/pnl-summary">client.Accounts.PNLSummary.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#AccountPNLSummaryService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#PNLSummary">PNLSummary</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## PNLDetails

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#AccountPNLDetailListResponse">AccountPNLDetailListResponse</a>

Methods:

- <code title="get /accounts/{account_id}/pnl-details">client.Accounts.PNLDetails.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#AccountPNLDetailService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#AccountPNLDetailListResponse">AccountPNLDetailListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Instruments

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#Instrument">Instrument</a>

Methods:

- <code title="get /instruments/{symbol}">client.Instruments.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#InstrumentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, symbol <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#InstrumentGetParams">InstrumentGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go">clsttest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/clst-test-go#Instrument">Instrument</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
