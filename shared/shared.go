// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"github.com/sachnk/clst-test-go/internal/apijson"
)

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

// pnlSummaryJSON contains the JSON metadata for the struct [PNLSummary]
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

func (r *PNLSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pnlSummaryJSON) RawJSON() string {
	return r.raw
}
