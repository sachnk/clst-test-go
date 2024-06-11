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

func TestAccountBulkOrderNew(t *testing.T) {
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
	_, err := client.Accounts.BulkOrders.New(
		context.TODO(),
		"x",
		clsttest.AccountBulkOrderNewParams{
			Orders: clsttest.F([]clsttest.AccountBulkOrderNewParamsOrder{{
				ReferenceID:  clsttest.F("my-order-id-123"),
				OrderType:    clsttest.F(clsttest.AccountBulkOrderNewParamsOrdersOrderTypeLimit),
				Side:         clsttest.F(clsttest.AccountBulkOrderNewParamsOrdersSideBuy),
				Quantity:     clsttest.F("x"),
				Price:        clsttest.F("x"),
				TimeInForce:  clsttest.F(clsttest.AccountBulkOrderNewParamsOrdersTimeInForceDay),
				StrategyType: clsttest.F(clsttest.AccountBulkOrderNewParamsOrdersStrategyTypeSor),
				LocateBroker: clsttest.F("x"),
				Symbol:       clsttest.F("AAPL"),
				SymbolFormat: clsttest.F(clsttest.AccountBulkOrderNewParamsOrdersSymbolFormatCms),
			}}),
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
