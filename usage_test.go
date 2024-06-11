// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clsttest_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/clst-test-go"
	"github.com/stainless-sdks/clst-test-go/internal/testutil"
	"github.com/stainless-sdks/clst-test-go/option"
)

func TestUsage(t *testing.T) {
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
	instrument, err := client.Instruments.Get(
		context.TODO(),
		"REPLACE_ME",
		clsttest.InstrumentGetParams{},
	)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", instrument.AssetClass)
}
