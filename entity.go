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

// EntityService contains methods and other services that help with interacting
// with the clearstreet API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEntityService] method instead.
type EntityService struct {
	Options []option.RequestOption
}

// NewEntityService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEntityService(opts ...option.RequestOption) (r *EntityService) {
	r = &EntityService{}
	r.Options = opts
	return
}

// Get an entity by its ID.
func (r *EntityService) Get(ctx context.Context, entityID string, opts ...option.RequestOption) (res *Entity, err error) {
	opts = append(r.Options[:], opts...)
	if entityID == "" {
		err = errors.New("missing required entity_id parameter")
		return
	}
	path := fmt.Sprintf("entities/%s", entityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all available entities.
func (r *EntityService) List(ctx context.Context, opts ...option.RequestOption) (res *EntityListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "entities"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Simulate Reg-T margin calculation for a given hypothetical set of prices and/or
// trades. This is useful for understanding the impact of price fluctuations or
// trades on margin requirements. Once a simulation is created, it remains
// available for 48-hours, after which it will automatically be deleted.
//
// Simulations created through the API are visible in the Studio UI under the Risk
// & Margin section, after enabling the "Risk Simulations" toggle.
func (r *EntityService) NewRegTMarginSimulation(ctx context.Context, entityID string, body EntityNewRegTMarginSimulationParams, opts ...option.RequestOption) (res *EntityNewRegTMarginSimulationResponse, err error) {
	opts = append(r.Options[:], opts...)
	if entityID == "" {
		err = errors.New("missing required entity_id parameter")
		return
	}
	path := fmt.Sprintf("entities/%s/regt-margin-simulations", entityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get PNL summary for all accounts in an entity.
func (r *EntityService) GetPNLSummary(ctx context.Context, entityID string, opts ...option.RequestOption) (res *shared.PNLSummary, err error) {
	opts = append(r.Options[:], opts...)
	if entityID == "" {
		err = errors.New("missing required entity_id parameter")
		return
	}
	path := fmt.Sprintf("entities/%s/pnl-summary", entityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get latest portfolio margin calculation for the given entity
func (r *EntityService) GetPortfolioMargin(ctx context.Context, entityID string, opts ...option.RequestOption) (res *PortfolioMargin, err error) {
	opts = append(r.Options[:], opts...)
	if entityID == "" {
		err = errors.New("missing required entity_id parameter")
		return
	}
	path := fmt.Sprintf("entities/%s/portfolio-margin", entityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get the latest Reg-T margin calculation for the given entity
func (r *EntityService) GetRegTMargin(ctx context.Context, entityID string, opts ...option.RequestOption) (res *RegTMargin, err error) {
	opts = append(r.Options[:], opts...)
	if entityID == "" {
		err = errors.New("missing required entity_id parameter")
		return
	}
	path := fmt.Sprintf("entities/%s/regt-margin", entityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get a Reg-T margin simluation that was previously created. Note, simulations are
// automatically deleted after 48-hours.
func (r *EntityService) GetRegTMarginSimulation(ctx context.Context, entityID string, simulationID string, opts ...option.RequestOption) (res *RegTMarginSimulation, err error) {
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

type Entity struct {
	ClientCode string `json:"client_code,required"`
	// Entity ID for the legal entity.
	EntityID  string     `json:"entity_id,required"`
	LegalName string     `json:"legal_name"`
	JSON      entityJSON `json:"-"`
}

// entityJSON contains the JSON metadata for the struct [Entity]
type entityJSON struct {
	ClientCode  apijson.Field
	EntityID    apijson.Field
	LegalName   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Entity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityJSON) RawJSON() string {
	return r.raw
}

type PortfolioMargin struct {
	// Sum of add-on margin requirements. Formula:
	// `liquidity_add_on + concentration_add_on + discretionary_requirement`
	AddOnRequirement float64 `json:"add_on_requirement"`
	// The percentage add-on margin requirements in terms of total house requirement.
	// Formula: `add_on_requirement / house_requirement`
	AddOnRequirementPercent float64 `json:"add_on_requirement_percent"`
	// A component margin requirement that captures risk based on gross exposure to
	// total equity.
	ConcentrationAddOn float64 `json:"concentration_add_on"`
	// The percentage concentration add-on margin requirements in terms of total house
	// requirement. Formula: `concentration_add_on / house_requirement`
	ConcentrationAddOnPercent float64 `json:"concentration_add_on_percent"`
	// A component margin requirement that captures miscellaneous risk factors.
	DiscretionaryRequirement float64 `json:"discretionary_requirement"`
	// The percentage discretionary margin requirements in terms of total house
	// requirement Formula: `discretionary_requirement / house_requirement`
	DiscretionaryRequirementPercent float64 `json:"discretionary_requirement_percent"`
	// The maring amount by taking the difference between total equity and the
	// effective requirement. A negative number reflects an effective margin deficit.
	EffeciveExcess float64 `json:"effecive_excess"`
	// The enforced margin requirement in effect.
	EffectiveRequirement float64 `json:"effective_requirement"`
	// Portfolio margin groups
	Groups []PortfolioMarginGroup `json:"groups"`
	// The margin amount by taking the difference between total equity and the house
	// requirement. A negative number reflects a house margin deficit.
	HouseExcess float64 `json:"house_excess"`
	// Margin requirements based on Clear Street's house margin methodology.
	HouseRequirement float64 `json:"house_requirement"`
	// A component margin requirement that captures risk based on liquidity, Market
	// Cap, and Average Daily Volume factors.
	LiquidityAddOn float64 `json:"liquidity_add_on"`
	// The percentage liquidity add-on margin requirements in terms of total house
	// requirement. Formula: `liquidity_add_on / house_requirement`
	LiquidityAddOnPercent float64 `json:"liquidity_add_on_percent"`
	// Sum of market values across all positions.
	NetMarketValue float64 `json:"net_market_value"`
	// A component margin requirement that captures risk for security instruments that
	// are not margin eligible.
	NonMarginableRequirement float64 `json:"non_marginable_requirement"`
	// The percentage non-marginable requirement in terms of total house requirement
	// Formula: `non_marginable_requirement / house_requirement`
	NonMarginableRequirementPercent float64 `json:"non_marginable_requirement_percent"`
	// A component margin requirement that captures base-case risk under house margin
	// methodology.
	RiskBasedRequirement float64 `json:"risk_based_requirement"`
	// The percentage risk_base margin requirement in terms of total house requirement
	// Formula: `risk_based_requirement / house_requirement`
	RiskBasedRequirementPercent float64 `json:"risk_based_requirement_percent"`
	// Timestamp of when this margin was calculated.
	Timestamp int64 `json:"timestamp"`
	// A component margin requirement that captures risk based on vega.
	VegaRequirement float64 `json:"vega_requirement"`
	// Unique identifier for this margin calculation.
	Version string              `json:"version"`
	JSON    portfolioMarginJSON `json:"-"`
}

// portfolioMarginJSON contains the JSON metadata for the struct [PortfolioMargin]
type portfolioMarginJSON struct {
	AddOnRequirement                apijson.Field
	AddOnRequirementPercent         apijson.Field
	ConcentrationAddOn              apijson.Field
	ConcentrationAddOnPercent       apijson.Field
	DiscretionaryRequirement        apijson.Field
	DiscretionaryRequirementPercent apijson.Field
	EffeciveExcess                  apijson.Field
	EffectiveRequirement            apijson.Field
	Groups                          apijson.Field
	HouseExcess                     apijson.Field
	HouseRequirement                apijson.Field
	LiquidityAddOn                  apijson.Field
	LiquidityAddOnPercent           apijson.Field
	NetMarketValue                  apijson.Field
	NonMarginableRequirement        apijson.Field
	NonMarginableRequirementPercent apijson.Field
	RiskBasedRequirement            apijson.Field
	RiskBasedRequirementPercent     apijson.Field
	Timestamp                       apijson.Field
	VegaRequirement                 apijson.Field
	Version                         apijson.Field
	raw                             string
	ExtraFields                     map[string]apijson.Field
}

func (r *PortfolioMargin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r portfolioMarginJSON) RawJSON() string {
	return r.raw
}

type PortfolioMarginGroup struct {
	// The enforced margin requirement in effect for the group.
	EffectiveRequirement float64 `json:"effective_requirement,required"`
	// The percentage effective margin requirement in terms of the group market value.
	// Formula: `(effective_requirement / net_market_value)`
	MarginPercent float64 `json:"margin_percent,required"`
	// The percentage effective margin requirement in terms of the total effective
	// requirement. Formula: `(effective_requirement / sum(effective_requirement))`
	MarginPercentContribution float64 `json:"margin_percent_contribution,required"`
	// The aggregated market value of all instruments for the group.
	MarketValue float64 `json:"market_value,required"`
	// The percentage market value of the group in terms of the total net_market_value
	// of all positions. Formula: `(market_value / net_market_value)`
	MarketValuePercent float64 `json:"market_value_percent,required"`
	// A list of securities that comprise this group.
	Members []PortfolioMarginGroupsMember `json:"members,required"`
	// Unique name of the group, typically the symbol of the underlier.
	Name string `json:"name,required"`
	// A component margin requirement that captures risk for the group based on gross
	// exposure to total equity
	ConcentrationRequirement float64 `json:"concentration_requirement"`
	// A component margin requirement that captures miscellaneous risk factors for the
	// group.
	DiscretionaryRequirement float64 `json:"discretionary_requirement"`
	// A component margin requirement that captures risk for the group based on
	// liquidity, Market Cap, and Average Daily Volume factors.
	LiquidityRequirement float64 `json:"liquidity_requirement"`
	// A component margin requirement that captures risk for the group that are not
	// margin eligible.
	NonMarginableRequirement float64 `json:"non_marginable_requirement"`
	// Margin requirements based on OCC TIMS regulatory margin methodology
	RegulatoryRequirement float64 `json:"regulatory_requirement"`
	// A component margin requirement that captures base-case risk for the group under
	// house margin methodology
	RiskBasedRequirement float64 `json:"risk_based_requirement"`
	// Maps shock scenarios to their resulting pnl.
	Shocks map[string]float64 `json:"shocks"`
	// Margin requirements based on value-at-risk over any 5-day period in a 2 year
	// historic lookback
	VarRequirement float64                  `json:"var_requirement"`
	JSON           portfolioMarginGroupJSON `json:"-"`
}

// portfolioMarginGroupJSON contains the JSON metadata for the struct
// [PortfolioMarginGroup]
type portfolioMarginGroupJSON struct {
	EffectiveRequirement      apijson.Field
	MarginPercent             apijson.Field
	MarginPercentContribution apijson.Field
	MarketValue               apijson.Field
	MarketValuePercent        apijson.Field
	Members                   apijson.Field
	Name                      apijson.Field
	ConcentrationRequirement  apijson.Field
	DiscretionaryRequirement  apijson.Field
	LiquidityRequirement      apijson.Field
	NonMarginableRequirement  apijson.Field
	RegulatoryRequirement     apijson.Field
	RiskBasedRequirement      apijson.Field
	Shocks                    apijson.Field
	VarRequirement            apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *PortfolioMarginGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r portfolioMarginGroupJSON) RawJSON() string {
	return r.raw
}

type PortfolioMarginGroupsMember struct {
	// The asset class of the symbol.
	AssetClass PortfolioMarginGroupsMembersAssetClass `json:"asset_class"`
	// Market value of the instrument.
	MarketValue float64 `json:"market_value"`
	// The percentage market value of the instrument in terms of the total
	// `net_market_value` of all positions held. Formula:
	// `market_value / net_market_value`
	MarketValuePercent float64 `json:"market_value_percent"`
	// The quantity held for this instrument.
	Quantity string `json:"quantity"`
	// Maps shock scenarios to their resulting pnl.
	Shocks map[string]float64 `json:"shocks"`
	// The symbol for the instrument.
	Symbol string                          `json:"symbol"`
	JSON   portfolioMarginGroupsMemberJSON `json:"-"`
}

// portfolioMarginGroupsMemberJSON contains the JSON metadata for the struct
// [PortfolioMarginGroupsMember]
type portfolioMarginGroupsMemberJSON struct {
	AssetClass         apijson.Field
	MarketValue        apijson.Field
	MarketValuePercent apijson.Field
	Quantity           apijson.Field
	Shocks             apijson.Field
	Symbol             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PortfolioMarginGroupsMember) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r portfolioMarginGroupsMemberJSON) RawJSON() string {
	return r.raw
}

// The asset class of the symbol.
type PortfolioMarginGroupsMembersAssetClass string

const (
	PortfolioMarginGroupsMembersAssetClassOther  PortfolioMarginGroupsMembersAssetClass = "other"
	PortfolioMarginGroupsMembersAssetClassEquity PortfolioMarginGroupsMembersAssetClass = "equity"
	PortfolioMarginGroupsMembersAssetClassOption PortfolioMarginGroupsMembersAssetClass = "option"
	PortfolioMarginGroupsMembersAssetClassDebt   PortfolioMarginGroupsMembersAssetClass = "debt"
)

func (r PortfolioMarginGroupsMembersAssetClass) IsKnown() bool {
	switch r {
	case PortfolioMarginGroupsMembersAssetClassOther, PortfolioMarginGroupsMembersAssetClassEquity, PortfolioMarginGroupsMembersAssetClassOption, PortfolioMarginGroupsMembersAssetClassDebt:
		return true
	}
	return false
}

type RegTMargin struct {
	// The remaining amount of start_of_day_buying_power that captures any day-trading
	// activity.
	DayTradeBuyingPower float64 `json:"day_trade_buying_power,required"`
	// The enforced margin requirement in effect.
	EffectiveRequirement float64 `json:"effective_requirement,required"`
	// Margin requirements based on regulatory rules.
	ExchangeRequirement float64 `json:"exchange_requirement,required"`
	// Reg-T margin groups
	Groups []RegTMarginGroup `json:"groups,required"`
	// The margin amount by taking the difference between total equity and the house
	// requirement. A negative number reflects a house margin deficit.
	HouseExcess float64 `json:"house_excess,required"`
	// Margin requirements based on Clear Street's house margin methodology.
	HouseRequirement float64 `json:"house_requirement,required"`
	// Market value net of long and short market values.
	NetMarketValue float64 `json:"net_market_value,required"`
	// The limit, or "up-to" amount, of securities value that can be purchased and held
	// overnight.
	OvernightBuyingPower float64 `json:"overnight_buying_power,required"`
	// Special Memorandum Account (SMA). The regulatory line of credit amount for
	// margin trading based on market value, trading activity, and available cash.
	Sma float64 `json:"sma,required"`
	// The limit, or "up-to" amount, of securities value that can be day-traded for a
	// given trading day.
	SodBuyingPower float64 `json:"sod_buying_power,required"`
	// Timestamp of when this margin was calculated.
	Timestamp int64 `json:"timestamp,required"`
	// Unique identifier for this margin calculation.
	Version string `json:"version,required"`
	// The maring amount by taking the difference between total equity and the
	// effective requirement. A negative number reflects an effective margin deficit.
	EffectiveExcess float64 `json:"effective_excess"`
	// The margin amount by taking the difference between total equity and the exchange
	// requirement. A negative number reflects an regulatory margin deficit.
	ExchangeExcess float64        `json:"exchange_excess"`
	JSON           regTMarginJSON `json:"-"`
}

// regTMarginJSON contains the JSON metadata for the struct [RegTMargin]
type regTMarginJSON struct {
	DayTradeBuyingPower  apijson.Field
	EffectiveRequirement apijson.Field
	ExchangeRequirement  apijson.Field
	Groups               apijson.Field
	HouseExcess          apijson.Field
	HouseRequirement     apijson.Field
	NetMarketValue       apijson.Field
	OvernightBuyingPower apijson.Field
	Sma                  apijson.Field
	SodBuyingPower       apijson.Field
	Timestamp            apijson.Field
	Version              apijson.Field
	EffectiveExcess      apijson.Field
	ExchangeExcess       apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *RegTMargin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regTMarginJSON) RawJSON() string {
	return r.raw
}

type RegTMarginGroup struct {
	// The enforced margin requirement in effect for the symbol group.
	EffectiveRequirement float64 `json:"effective_requirement,required"`
	// Margin requirements based on regulatory rules for the symbol group.
	ExchangeRequirement float64 `json:"exchange_requirement,required"`
	// Margin requirements based on Clear Street's house margin methodology for the
	// symbol group.
	HouseRequirement float64 `json:"house_requirement,required"`
	// The percentage effective margin requirement in terms of the symbol group market
	// value. Formula: `(effective_requirement / net_market_value)`
	MarginPercent float64 `json:"margin_percent,required"`
	// The percentage effective margin requirement in terms of the total effective
	// requirement. Formula: `(effective_requirement / sum(effective_requirement))`
	MarginPercentContribution float64 `json:"margin_percent_contribution,required"`
	// The aggregated market value of all instruments for the symbol group.
	MarketValue float64 `json:"market_value,required"`
	// The percentage market value of the symbol group in terms of the total
	// net_market_value of all positions. Formula: `(market_value / net_market_value)`
	MarketValuePercent float64 `json:"market_value_percent,required"`
	// A list of securities that comprise this group.
	Members []RegTMarginGroupsMember `json:"members,required"`
	// Unique name of the group, typically the symbol of the underlier.
	Name string              `json:"name,required"`
	JSON regTMarginGroupJSON `json:"-"`
}

// regTMarginGroupJSON contains the JSON metadata for the struct [RegTMarginGroup]
type regTMarginGroupJSON struct {
	EffectiveRequirement      apijson.Field
	ExchangeRequirement       apijson.Field
	HouseRequirement          apijson.Field
	MarginPercent             apijson.Field
	MarginPercentContribution apijson.Field
	MarketValue               apijson.Field
	MarketValuePercent        apijson.Field
	Members                   apijson.Field
	Name                      apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *RegTMarginGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regTMarginGroupJSON) RawJSON() string {
	return r.raw
}

type RegTMarginGroupsMember struct {
	// The asset class of the symbol.
	AssetClass RegTMarginGroupsMembersAssetClass `json:"asset_class,required"`
	// Market value of the instrument.
	MarketValue float64 `json:"market_value,required"`
	// The percentage market value of the instrument in terms of the total
	// `net_market_value` of all positions held. Formula:
	// `market_value / net_market_value`
	MarketValuePercent float64 `json:"market_value_percent,required"`
	// The quantity held for this instrument.
	Quantity string `json:"quantity,required"`
	// The symbol for the instrument.
	Symbol string                     `json:"symbol,required"`
	JSON   regTMarginGroupsMemberJSON `json:"-"`
}

// regTMarginGroupsMemberJSON contains the JSON metadata for the struct
// [RegTMarginGroupsMember]
type regTMarginGroupsMemberJSON struct {
	AssetClass         apijson.Field
	MarketValue        apijson.Field
	MarketValuePercent apijson.Field
	Quantity           apijson.Field
	Symbol             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RegTMarginGroupsMember) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regTMarginGroupsMemberJSON) RawJSON() string {
	return r.raw
}

// The asset class of the symbol.
type RegTMarginGroupsMembersAssetClass string

const (
	RegTMarginGroupsMembersAssetClassOther  RegTMarginGroupsMembersAssetClass = "other"
	RegTMarginGroupsMembersAssetClassEquity RegTMarginGroupsMembersAssetClass = "equity"
	RegTMarginGroupsMembersAssetClassOption RegTMarginGroupsMembersAssetClass = "option"
	RegTMarginGroupsMembersAssetClassDebt   RegTMarginGroupsMembersAssetClass = "debt"
)

func (r RegTMarginGroupsMembersAssetClass) IsKnown() bool {
	switch r {
	case RegTMarginGroupsMembersAssetClassOther, RegTMarginGroupsMembersAssetClassEquity, RegTMarginGroupsMembersAssetClassOption, RegTMarginGroupsMembersAssetClassDebt:
		return true
	}
	return false
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
	SimulationID string                   `json:"simulation_id,required" format:"uuid"`
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

type EntityListResponse struct {
	Data []Entity               `json:"data"`
	JSON entityListResponseJSON `json:"-"`
}

// entityListResponseJSON contains the JSON metadata for the struct
// [EntityListResponse]
type entityListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntityListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityListResponseJSON) RawJSON() string {
	return r.raw
}

type EntityNewRegTMarginSimulationResponse struct {
	// Unique ID for a simulation.
	SimulationID string                                    `json:"simulation_id,required" format:"uuid"`
	JSON         entityNewRegTMarginSimulationResponseJSON `json:"-"`
}

// entityNewRegTMarginSimulationResponseJSON contains the JSON metadata for the
// struct [EntityNewRegTMarginSimulationResponse]
type entityNewRegTMarginSimulationResponseJSON struct {
	SimulationID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EntityNewRegTMarginSimulationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityNewRegTMarginSimulationResponseJSON) RawJSON() string {
	return r.raw
}

type EntityNewRegTMarginSimulationParams struct {
	// A name for this simulation for reference.
	Name param.Field[string] `json:"name,required"`
	// If true, the simulation will ignore any existing positions and balances in the
	// account. Set to true if you want to simulate from a clean slate, i.e. an empty
	// account.
	IgnoreExisting param.Field[bool] `json:"ignore_existing"`
	// List of prices to use in the simulation, i.e. fair-market-values you specify for
	// each symbol. If this is not provided, current market prices will be used, if
	// they are available.
	Prices param.Field[[]EntityNewRegTMarginSimulationParamsPrice] `json:"prices"`
	// List of hypothetical trades to include in the simulation, if any.
	Trades param.Field[[]EntityNewRegTMarginSimulationParamsTrade] `json:"trades"`
}

func (r EntityNewRegTMarginSimulationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EntityNewRegTMarginSimulationParamsPrice struct {
	// The price to use in the simulation.
	Price param.Field[string] `json:"price,required"`
	// The symbol for the instrument.
	Symbol param.Field[string] `json:"symbol,required"`
	// Denotes the format of the provided `symbol` field.
	SymbolFormat param.Field[EntityNewRegTMarginSimulationParamsPricesSymbolFormat] `json:"symbol_format"`
}

func (r EntityNewRegTMarginSimulationParamsPrice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Denotes the format of the provided `symbol` field.
type EntityNewRegTMarginSimulationParamsPricesSymbolFormat string

const (
	EntityNewRegTMarginSimulationParamsPricesSymbolFormatCms EntityNewRegTMarginSimulationParamsPricesSymbolFormat = "cms"
	EntityNewRegTMarginSimulationParamsPricesSymbolFormatOsi EntityNewRegTMarginSimulationParamsPricesSymbolFormat = "osi"
)

func (r EntityNewRegTMarginSimulationParamsPricesSymbolFormat) IsKnown() bool {
	switch r {
	case EntityNewRegTMarginSimulationParamsPricesSymbolFormatCms, EntityNewRegTMarginSimulationParamsPricesSymbolFormatOsi:
		return true
	}
	return false
}

type EntityNewRegTMarginSimulationParamsTrade struct {
	// The price of the simulated trade.
	Price param.Field[string] `json:"price,required"`
	// The quantity of the simulated trade.
	Quantity param.Field[string] `json:"quantity,required"`
	// The side of the simulated trade.
	Side param.Field[EntityNewRegTMarginSimulationParamsTradesSide] `json:"side,required"`
	// The symbol for the instrument.
	Symbol param.Field[string] `json:"symbol,required"`
	// Denotes the format of the provided `symbol` field.
	SymbolFormat param.Field[EntityNewRegTMarginSimulationParamsTradesSymbolFormat] `json:"symbol_format"`
}

func (r EntityNewRegTMarginSimulationParamsTrade) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The side of the simulated trade.
type EntityNewRegTMarginSimulationParamsTradesSide string

const (
	EntityNewRegTMarginSimulationParamsTradesSideBuy  EntityNewRegTMarginSimulationParamsTradesSide = "buy"
	EntityNewRegTMarginSimulationParamsTradesSideSell EntityNewRegTMarginSimulationParamsTradesSide = "sell"
)

func (r EntityNewRegTMarginSimulationParamsTradesSide) IsKnown() bool {
	switch r {
	case EntityNewRegTMarginSimulationParamsTradesSideBuy, EntityNewRegTMarginSimulationParamsTradesSideSell:
		return true
	}
	return false
}

// Denotes the format of the provided `symbol` field.
type EntityNewRegTMarginSimulationParamsTradesSymbolFormat string

const (
	EntityNewRegTMarginSimulationParamsTradesSymbolFormatCms EntityNewRegTMarginSimulationParamsTradesSymbolFormat = "cms"
	EntityNewRegTMarginSimulationParamsTradesSymbolFormatOsi EntityNewRegTMarginSimulationParamsTradesSymbolFormat = "osi"
)

func (r EntityNewRegTMarginSimulationParamsTradesSymbolFormat) IsKnown() bool {
	switch r {
	case EntityNewRegTMarginSimulationParamsTradesSymbolFormatCms, EntityNewRegTMarginSimulationParamsTradesSymbolFormatOsi:
		return true
	}
	return false
}
