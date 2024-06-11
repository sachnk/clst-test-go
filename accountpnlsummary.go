// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clsttest

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/clst-test-go/internal/requestconfig"
	"github.com/stainless-sdks/clst-test-go/option"
)

// AccountPNLSummaryService contains methods and other services that help with
// interacting with the clearstreet API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountPNLSummaryService] method instead.
type AccountPNLSummaryService struct {
	Options []option.RequestOption
}

// NewAccountPNLSummaryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountPNLSummaryService(opts ...option.RequestOption) (r *AccountPNLSummaryService) {
	r = &AccountPNLSummaryService{}
	r.Options = opts
	return
}

// Get PNL summary for a given account.
func (r *AccountPNLSummaryService) Get(ctx context.Context, accountID string, opts ...option.RequestOption) (res *PNLSummary, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pnl-summary", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
