package tracker

import (
	"net/http"
	"testing"
)

// Test cases for BuildSwapTransaction
func TestGetRate(t *testing.T) {
	testCases := []struct {
		name          string
		fromMint      string
		toMint        string
		fromAmount    float64
		slipPage      float64
		errorExpected bool
	}{
		{
			name:          "ValidInput",
			fromMint:      "So11111111111111111111111111111111111111112",
			toMint:        "4k3Dyjzvzp8eMZWUXbBCjEvwSkkk59S5iCNLY3QrkX6R",
			fromAmount:    12.5,
			slipPage:      2.0,
			errorExpected: false,
		},
		{
			name:          "ValidInput",
			fromMint:      "So11111111111111111111111111111111111111112",
			toMint:        "AAAAAA",
			fromAmount:    12.5,
			slipPage:      2.0,
			errorExpected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			client := http.Client{}
			swapper := NewSolanaTracker(&client)
			_, err := swapper.GetRate(tc.fromMint, tc.toMint, tc.fromAmount, tc.slipPage)
			if tc.errorExpected && err == nil {
				t.Errorf("expected an error but got none")
			}
			if !tc.errorExpected && err != nil {
				t.Errorf("did not expect any error but got %v", err)
			}

		})
	}
}

// Test cases for BuildSwapTransaction
func TestGetSwapTransaction(t *testing.T) {
	testCases := []struct {
		name          string
		fromMint      string
		toMint        string
		fromAmount    float64
		slipPage      float64
		payerAddress  string
		errorExpected bool
	}{
		{
			name:          "ValidInput",
			fromMint:      "So11111111111111111111111111111111111111112",
			toMint:        "4k3Dyjzvzp8eMZWUXbBCjEvwSkkk59S5iCNLY3QrkX6R",
			fromAmount:    12.5,
			slipPage:      2.0,
			payerAddress:  "44ByXN2NgCeHJdsmbUvKrjcdRQJLJDuY8DqnqWMcwy5U",
			errorExpected: false,
		},
		{
			name:          "ValidInput",
			fromMint:      "So11111111111111111111111111111111111111112",
			toMint:        "4k3Dyjzvzp8eMZWUXbBCjEvwSkkk59S5iCNLY3QrkX6R",
			fromAmount:    12.5,
			slipPage:      2.0,
			payerAddress:  "this does not exist",
			errorExpected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			client := http.Client{}
			swapper := NewSolanaTracker(&client)
			_, err := swapper.GetSwapTransaction(tc.fromMint, tc.toMint, tc.fromAmount, tc.slipPage, tc.payerAddress)
			if tc.errorExpected && err == nil {
				t.Errorf("expected an error but got none")
			}
			if !tc.errorExpected && err != nil {
				t.Errorf("did not expect any error but got %v", err)
			}

		})
	}
}
