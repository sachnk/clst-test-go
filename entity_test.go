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

func TestEntityGet(t *testing.T) {
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
	_, err := client.Entities.Get(context.TODO(), "x")
	if err != nil {
		var apierr *clsttest.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEntityList(t *testing.T) {
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
	_, err := client.Entities.List(context.TODO())
	if err != nil {
		var apierr *clsttest.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEntityNewRegTMarginSimulationWithOptionalParams(t *testing.T) {
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
	_, err := client.Entities.NewRegTMarginSimulation(
		context.TODO(),
		"x",
		clsttest.EntityNewRegTMarginSimulationParams{
			Name:           clsttest.F("string"),
			IgnoreExisting: clsttest.F(true),
			Prices: clsttest.F([]clsttest.EntityNewRegTMarginSimulationParamsPrice{{
				Symbol:       clsttest.F("AAPL"),
				SymbolFormat: clsttest.F(clsttest.EntityNewRegTMarginSimulationParamsPricesSymbolFormatCms),
				Price:        clsttest.F("x"),
			}, {
				Symbol:       clsttest.F("AAPL"),
				SymbolFormat: clsttest.F(clsttest.EntityNewRegTMarginSimulationParamsPricesSymbolFormatCms),
				Price:        clsttest.F("x"),
			}, {
				Symbol:       clsttest.F("AAPL"),
				SymbolFormat: clsttest.F(clsttest.EntityNewRegTMarginSimulationParamsPricesSymbolFormatCms),
				Price:        clsttest.F("x"),
			}}),
			Trades: clsttest.F([]clsttest.EntityNewRegTMarginSimulationParamsTrade{{
				Symbol:       clsttest.F("AAPL"),
				SymbolFormat: clsttest.F(clsttest.EntityNewRegTMarginSimulationParamsTradesSymbolFormatCms),
				Side:         clsttest.F(clsttest.EntityNewRegTMarginSimulationParamsTradesSideBuy),
				Quantity:     clsttest.F("x"),
				Price:        clsttest.F("x"),
			}, {
				Symbol:       clsttest.F("AAPL"),
				SymbolFormat: clsttest.F(clsttest.EntityNewRegTMarginSimulationParamsTradesSymbolFormatCms),
				Side:         clsttest.F(clsttest.EntityNewRegTMarginSimulationParamsTradesSideBuy),
				Quantity:     clsttest.F("x"),
				Price:        clsttest.F("x"),
			}, {
				Symbol:       clsttest.F("AAPL"),
				SymbolFormat: clsttest.F(clsttest.EntityNewRegTMarginSimulationParamsTradesSymbolFormatCms),
				Side:         clsttest.F(clsttest.EntityNewRegTMarginSimulationParamsTradesSideBuy),
				Quantity:     clsttest.F("x"),
				Price:        clsttest.F("x"),
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

func TestEntityGetPNLSummary(t *testing.T) {
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
	_, err := client.Entities.GetPNLSummary(context.TODO(), "x")
	if err != nil {
		var apierr *clsttest.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEntityGetPortfolioMargin(t *testing.T) {
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
	_, err := client.Entities.GetPortfolioMargin(context.TODO(), "x")
	if err != nil {
		var apierr *clsttest.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEntityGetRegTMargin(t *testing.T) {
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
	_, err := client.Entities.GetRegTMargin(context.TODO(), "x")
	if err != nil {
		var apierr *clsttest.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEntityGetRegTMarginSimulation(t *testing.T) {
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
	_, err := client.Entities.GetRegTMarginSimulation(
		context.TODO(),
		"x",
		"6460030d-8ed4-19d3-818e-e87b36e90005",
	)
	if err != nil {
		var apierr *clsttest.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
