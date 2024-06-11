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

func TestEntityRegtMarginSimulationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Entities.RegtMarginSimulations.New(
		context.TODO(),
		"x",
		clsttest.EntityRegtMarginSimulationNewParams{
			Name:           clsttest.F("string"),
			IgnoreExisting: clsttest.F(true),
			Prices: clsttest.F([]clsttest.EntityRegtMarginSimulationNewParamsPrice{{
				Symbol:       clsttest.F("AAPL"),
				SymbolFormat: clsttest.F(clsttest.EntityRegtMarginSimulationNewParamsPricesSymbolFormatCms),
				Price:        clsttest.F("x"),
			}, {
				Symbol:       clsttest.F("AAPL"),
				SymbolFormat: clsttest.F(clsttest.EntityRegtMarginSimulationNewParamsPricesSymbolFormatCms),
				Price:        clsttest.F("x"),
			}, {
				Symbol:       clsttest.F("AAPL"),
				SymbolFormat: clsttest.F(clsttest.EntityRegtMarginSimulationNewParamsPricesSymbolFormatCms),
				Price:        clsttest.F("x"),
			}}),
			Trades: clsttest.F([]clsttest.EntityRegtMarginSimulationNewParamsTrade{{
				Symbol:       clsttest.F("AAPL"),
				SymbolFormat: clsttest.F(clsttest.EntityRegtMarginSimulationNewParamsTradesSymbolFormatCms),
				Side:         clsttest.F(clsttest.EntityRegtMarginSimulationNewParamsTradesSideBuy),
				Quantity:     clsttest.F("x"),
				Price:        clsttest.F("x"),
			}, {
				Symbol:       clsttest.F("AAPL"),
				SymbolFormat: clsttest.F(clsttest.EntityRegtMarginSimulationNewParamsTradesSymbolFormatCms),
				Side:         clsttest.F(clsttest.EntityRegtMarginSimulationNewParamsTradesSideBuy),
				Quantity:     clsttest.F("x"),
				Price:        clsttest.F("x"),
			}, {
				Symbol:       clsttest.F("AAPL"),
				SymbolFormat: clsttest.F(clsttest.EntityRegtMarginSimulationNewParamsTradesSymbolFormatCms),
				Side:         clsttest.F(clsttest.EntityRegtMarginSimulationNewParamsTradesSideBuy),
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

func TestEntityRegtMarginSimulationGet(t *testing.T) {
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
	_, err := client.Entities.RegtMarginSimulations.Get(
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
