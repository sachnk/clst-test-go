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

func TestAccountOrderNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Orders.New(
		context.TODO(),
		"x",
		clsttest.AccountOrderNewParams{
			OrderType:    clsttest.F(clsttest.AccountOrderNewParamsOrderTypeLimit),
			Quantity:     clsttest.F("x"),
			Side:         clsttest.F(clsttest.AccountOrderNewParamsSideBuy),
			StrategyType: clsttest.F(clsttest.AccountOrderNewParamsStrategyTypeSor),
			Symbol:       clsttest.F("AAPL"),
			TimeInForce:  clsttest.F(clsttest.AccountOrderNewParamsTimeInForceDay),
			LocateBroker: clsttest.F("x"),
			Price:        clsttest.F("x"),
			ReferenceID:  clsttest.F("my-order-id-123"),
			SymbolFormat: clsttest.F(clsttest.AccountOrderNewParamsSymbolFormatCms),
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

func TestAccountOrderGet(t *testing.T) {
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
	_, err := client.Accounts.Orders.Get(
		context.TODO(),
		"x",
		"x",
	)
	if err != nil {
		var apierr *clsttest.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAccountOrderListWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Orders.List(
		context.TODO(),
		"x",
		clsttest.AccountOrderListParams{
			From:      clsttest.F(int64(1710613560668)),
			PageSize:  clsttest.F(int64(1)),
			PageToken: clsttest.F("string"),
			To:        clsttest.F(int64(1710613560668)),
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

func TestAccountOrderDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Accounts.Orders.Delete(
		context.TODO(),
		"x",
		clsttest.AccountOrderDeleteParams{
			Symbol:       clsttest.F("AAPL"),
			SymbolFormat: clsttest.F(clsttest.AccountOrderDeleteParamsSymbolFormatCms),
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

func TestAccountOrderCancel(t *testing.T) {
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
	err := client.Accounts.Orders.Cancel(
		context.TODO(),
		"x",
		"x",
	)
	if err != nil {
		var apierr *clsttest.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
