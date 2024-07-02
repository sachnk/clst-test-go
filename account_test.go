// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package clsttest_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/sachnk/clst-test-go"
	"github.com/sachnk/clst-test-go/internal/testutil"
	"github.com/sachnk/clst-test-go/option"
)

func TestAccountGet(t *testing.T) {
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
	_, err := client.Accounts.Get(context.TODO(), "x")
	if err != nil {
		var apierr *clsttest.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountList(t *testing.T) {
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
	_, err := client.Accounts.List(context.TODO())
	if err != nil {
		var apierr *clsttest.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountNewOrdersInBulk(t *testing.T) {
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
	_, err := client.Accounts.NewOrdersInBulk(
		context.TODO(),
		"x",
		clsttest.AccountNewOrdersInBulkParams{
			Orders: clsttest.F([]clsttest.AccountNewOrdersInBulkParamsOrder{{
				ReferenceID:  clsttest.F("my-order-id-123"),
				OrderType:    clsttest.F(clsttest.AccountNewOrdersInBulkParamsOrdersOrderTypeLimit),
				Side:         clsttest.F(clsttest.AccountNewOrdersInBulkParamsOrdersSideBuy),
				Quantity:     clsttest.F("x"),
				Price:        clsttest.F("x"),
				TimeInForce:  clsttest.F(clsttest.AccountNewOrdersInBulkParamsOrdersTimeInForceDay),
				StrategyType: clsttest.F(clsttest.AccountNewOrdersInBulkParamsOrdersStrategyTypeSor),
				LocateBroker: clsttest.F("x"),
				Symbol:       clsttest.F("AAPL"),
				SymbolFormat: clsttest.F(clsttest.AccountNewOrdersInBulkParamsOrdersSymbolFormatCms),
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

func TestAccountGetPNLDetails(t *testing.T) {
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
	_, err := client.Accounts.GetPNLDetails(context.TODO(), "x")
	if err != nil {
		var apierr *clsttest.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountGetPNLSummary(t *testing.T) {
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
	_, err := client.Accounts.GetPNLSummary(context.TODO(), "x")
	if err != nil {
		var apierr *clsttest.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
