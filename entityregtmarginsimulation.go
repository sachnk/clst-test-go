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
)

// EntityRegTMarginSimulationService contains methods and other services that help
// with interacting with the clearstreet API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEntityRegTMarginSimulationService] method instead.
type EntityRegTMarginSimulationService struct {
	Options []option.RequestOption
}

// NewEntityRegTMarginSimulationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewEntityRegTMarginSimulationService(opts ...option.RequestOption) (r *EntityRegTMarginSimulationService) {
	r = &EntityRegTMarginSimulationService{}
	r.Options = opts
	return
}

// Simulate Reg-T margin calculation for a given hypothetical set of prices and/or
// trades. This is useful for understanding the impact of price fluctuations or
// trades on margin requirements. Once a simulation is created, it remains
// available for 48-hours, after which it will automatically be deleted.
//
// Simulations created through the API are visible in the Studio UI under the Risk
// & Margin section, after enabling the "Risk Simulations" toggle.
func (r *EntityRegTMarginSimulationService) New(ctx context.Context, entityID string, body EntityRegTMarginSimulationNewParams, opts ...option.RequestOption) (res *EntityRegTMarginSimulationNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if entityID == "" {
		err = errors.New("missing required entity_id parameter")
		return
	}
	path := fmt.Sprintf("entities/%s/regt-margin-simulations", entityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a Reg-T margin simluation that was previously created. Note, simulations are
// automatically deleted after 48-hours.
func (r *EntityRegTMarginSimulationService) Get(ctx context.Context, entityID string, simulationID SimulationIDParam, opts ...option.RequestOption) (res *RegTMarginSimulation, err error) {
	opts = append(r.Options[:], opts...)
	if entityID == "" {
		err = errors.New("missing required entity_id parameter")
		return
	}
	if simulationID == "" {
		err = errors.New("missing required simulation_id parameter")
		return
	}
	path := fmt.Sprintf("entities/%s/regt-margin-simulations/%s", entityID, simulationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type RegTMarginSimulation struct {
	// The margin calculation after applying simulated trades.
	After RegTMargin `json:"after,required"`
	// The margin calculation before applying simulated trades.
	Before RegTMargin `json:"before,required"`
	// Timestamp of when this simulation was created.
	CreatedAt int64 `json:"created_at,required"`
	// Name of this simulation that you provided when creating it.
	Name string `json:"name,required"`
	// Unique ID for a simulation.
	SimulationID SimulationID             `json:"simulation_id,required" format:"uuid"`
	JSON         regTMarginSimulationJSON `json:"-"`
}

// regTMarginSimulationJSON contains the JSON metadata for the struct
// [RegTMarginSimulation]
type regTMarginSimulationJSON struct {
	After        apijson.Field
	Before       apijson.Field
	CreatedAt    apijson.Field
	Name         apijson.Field
	SimulationID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RegTMarginSimulation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regTMarginSimulationJSON) RawJSON() string {
	return r.raw
}

type SimulationID = string

type SimulationIDParam = string

type EntityRegTMarginSimulationNewResponse struct {
	// Unique ID for a simulation.
	SimulationID SimulationID                              `json:"simulation_id,required" format:"uuid"`
	JSON         entityRegTMarginSimulationNewResponseJSON `json:"-"`
}

// entityRegTMarginSimulationNewResponseJSON contains the JSON metadata for the
// struct [EntityRegTMarginSimulationNewResponse]
type entityRegTMarginSimulationNewResponseJSON struct {
	SimulationID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EntityRegTMarginSimulationNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityRegTMarginSimulationNewResponseJSON) RawJSON() string {
	return r.raw
}

type EntityRegTMarginSimulationNewParams struct {
	// A name for this simulation for reference.
	Name param.Field[string] `json:"name,required"`
	// If true, the simulation will ignore any existing positions and balances in the
	// account. Set to true if you want to simulate from a clean slate, i.e. an empty
	// account.
	IgnoreExisting param.Field[bool] `json:"ignore_existing"`
	// List of prices to use in the simulation, i.e. fair-market-values you specify for
	// each symbol. If this is not provided, current market prices will be used, if
	// they are available.
	Prices param.Field[[]EntityRegTMarginSimulationNewParamsPrice] `json:"prices"`
	// List of hypothetical trades to include in the simulation, if any.
	Trades param.Field[[]EntityRegTMarginSimulationNewParamsTrade] `json:"trades"`
}

func (r EntityRegTMarginSimulationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EntityRegTMarginSimulationNewParamsPrice struct {
	// The price to use in the simulation.
	Price param.Field[string] `json:"price,required"`
	// The symbol for the instrument.
	Symbol param.Field[string] `json:"symbol,required"`
	// Denotes the format of the provided `symbol` field.
	SymbolFormat param.Field[EntityRegTMarginSimulationNewParamsPricesSymbolFormat] `json:"symbol_format"`
}

func (r EntityRegTMarginSimulationNewParamsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Denotes the format of the provided `symbol` field.
type EntityRegTMarginSimulationNewParamsPricesSymbolFormat string

const (
	EntityRegTMarginSimulationNewParamsPricesSymbolFormatCms EntityRegTMarginSimulationNewParamsPricesSymbolFormat = "cms"
	EntityRegTMarginSimulationNewParamsPricesSymbolFormatOsi EntityRegTMarginSimulationNewParamsPricesSymbolFormat = "osi"
)

func (r EntityRegTMarginSimulationNewParamsPricesSymbolFormat) IsKnown() bool {
	switch r {
	case EntityRegTMarginSimulationNewParamsPricesSymbolFormatCms, EntityRegTMarginSimulationNewParamsPricesSymbolFormatOsi:
		return true
	}
	return false
}

type EntityRegTMarginSimulationNewParamsTrade struct {
	// The price of the simulated trade.
	Price param.Field[string] `json:"price,required"`
	// The quantity of the simulated trade.
	Quantity param.Field[string] `json:"quantity,required"`
	// The side of the simulated trade.
	Side param.Field[EntityRegTMarginSimulationNewParamsTradesSide] `json:"side,required"`
	// The symbol for the instrument.
	Symbol param.Field[string] `json:"symbol,required"`
	// Denotes the format of the provided `symbol` field.
	SymbolFormat param.Field[EntityRegTMarginSimulationNewParamsTradesSymbolFormat] `json:"symbol_format"`
}

func (r EntityRegTMarginSimulationNewParamsTrade) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The side of the simulated trade.
type EntityRegTMarginSimulationNewParamsTradesSide string

const (
	EntityRegTMarginSimulationNewParamsTradesSideBuy  EntityRegTMarginSimulationNewParamsTradesSide = "buy"
	EntityRegTMarginSimulationNewParamsTradesSideSell EntityRegTMarginSimulationNewParamsTradesSide = "sell"
)

func (r EntityRegTMarginSimulationNewParamsTradesSide) IsKnown() bool {
	switch r {
	case EntityRegTMarginSimulationNewParamsTradesSideBuy, EntityRegTMarginSimulationNewParamsTradesSideSell:
		return true
	}
	return false
}

// Denotes the format of the provided `symbol` field.
type EntityRegTMarginSimulationNewParamsTradesSymbolFormat string

const (
	EntityRegTMarginSimulationNewParamsTradesSymbolFormatCms EntityRegTMarginSimulationNewParamsTradesSymbolFormat = "cms"
	EntityRegTMarginSimulationNewParamsTradesSymbolFormatOsi EntityRegTMarginSimulationNewParamsTradesSymbolFormat = "osi"
)

func (r EntityRegTMarginSimulationNewParamsTradesSymbolFormat) IsKnown() bool {
	switch r {
	case EntityRegTMarginSimulationNewParamsTradesSymbolFormatCms, EntityRegTMarginSimulationNewParamsTradesSymbolFormatOsi:
		return true
	}
	return false
}
