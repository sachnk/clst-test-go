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

func TestEntityRegTMarginSimulationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Entities.RegTMarginSimulations.New(
		context.TODO(),
		"x",
		clsttest.EntityRegTMarginSimulationNewParams{
			Name:           clsttest.F("string"),
			IgnoreExisting: clsttest.F(true),
			Prices: clsttest.F([]clsttest.EntityRegTMarginSimulationNewParamsPrice{{
				Symbol:       clsttest.F("AAPL"),
				SymbolFormat: clsttest.F(clsttest.EntityRegTMarginSimulationNewParamsPricesSymbolFormatCms),
				Price:        clsttest.F("x"),
			}, {
				Symbol:       clsttest.F("AAPL"),
				SymbolFormat: clsttest.F(clsttest.EntityRegTMarginSimulationNewParamsPricesSymbolFormatCms),
				Price:        clsttest.F("x"),
			}, {
				Symbol:       clsttest.F("AAPL"),
				SymbolFormat: clsttest.F(clsttest.EntityRegTMarginSimulationNewParamsPricesSymbolFormatCms),
				Price:        clsttest.F("x"),
			}}),
			Trades: clsttest.F([]clsttest.EntityRegTMarginSimulationNewParamsTrade{{
				Symbol:       clsttest.F("AAPL"),
				SymbolFormat: clsttest.F(clsttest.EntityRegTMarginSimulationNewParamsTradesSymbolFormatCms),
				Side:         clsttest.F(clsttest.EntityRegTMarginSimulationNewParamsTradesSideBuy),
				Quantity:     clsttest.F("x"),
				Price:        clsttest.F("x"),
			}, {
				Symbol:       clsttest.F("AAPL"),
				SymbolFormat: clsttest.F(clsttest.EntityRegTMarginSimulationNewParamsTradesSymbolFormatCms),
				Side:         clsttest.F(clsttest.EntityRegTMarginSimulationNewParamsTradesSideBuy),
				Quantity:     clsttest.F("x"),
				Price:        clsttest.F("x"),
			}, {
				Symbol:       clsttest.F("AAPL"),
				SymbolFormat: clsttest.F(clsttest.EntityRegTMarginSimulationNewParamsTradesSymbolFormatCms),
				Side:         clsttest.F(clsttest.EntityRegTMarginSimulationNewParamsTradesSideBuy),
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

func TestEntityRegTMarginSimulationGet(t *testing.T) {
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
	_, err := client.Entities.RegTMarginSimulations.Get(
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
