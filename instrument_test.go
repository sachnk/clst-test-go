// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clsttest_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/clst-test-go"
	"github.com/stainless-sdks/clst-test-go/internal/testutil"
	"github.com/stainless-sdks/clst-test-go/option"
)

func TestInstrumentGetWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := clsttest.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Instruments.Get(
		context.TODO(),
		"AAPL",
		clsttest.InstrumentGetParams{
			SymbolFormat: clsttest.F(clsttest.InstrumentGetParamsSymbolFormatCms),
		},
	)
	if err != nil {
		var apierr *clsttest.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
