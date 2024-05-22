package tracker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type SolanaTracker struct {
	client         *http.Client
	retryOn500     bool
	retryOn500Wait time.Duration
}

// NewSolanaTracker creates a new instance of SolanaTracker with the specified client.
// It sets the retryOn500 flag to true and the retryOn500Wait duration to 500 milliseconds.
func NewSolanaTracker(client *http.Client) *SolanaTracker {
	return &SolanaTracker{
		client:         client,
		retryOn500:     true,
		retryOn500Wait: 500 * time.Millisecond,
	}
}

// GetRate retrieves the exchange rate between two tokens.
// It takes the source token mint, destination token mint,
// token amount, and slippage percentage as parameters.
// The function makes an HTTP GET request to the Solana Tracker API
// to fetch the exchange rate.
// If the request fails with a status code of 500 and the retryOn500 flag is set to true,
// the function will sleep for the specified duration and retry the request.
// If the status code is not 200, an error is returned with the details.
func (s *SolanaTracker) GetRate(fromMint string, toMint string, amount float64, slipPage float64) (*RateResponse, error) {
	url := fmt.Sprintf("https://swap-api.solanatracker.io/rate?from=%s&to=%s&amount=%f&slippage=%f", fromMint, toMint, amount, slipPage)
	response, err := s.client.Get(url)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode == 500 && s.retryOn500 {
		time.Sleep(s.retryOn500Wait)
		return s.GetRate(fromMint, toMint, amount, slipPage)
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("status code %d. Details: %s", response.StatusCode, string(body))
	}
	var result RateResponse

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetSwapTransaction retrieves information about a swap transaction between two tokens.
// It takes the source token mint, destination token mint, source token amount, slippage percentage,
// and payer address as parameters. It returns a SwapResponse containing the transaction information
// or an error if the request fails.
// The function makes an HTTP GET request to the Solana Tracker API to fetch the swap transaction details.
// If the request fails with a status code of 500 and the retryOn500 flag is set to true, the function will
// sleep for the specified duration and retry the request.
// If the status code is not 200, an error is returned with the
func (s *SolanaTracker) GetSwapTransaction(fromMint string, toMint string, fromAmount float64, slipPage float64, payerAddress string) (*SwapResponse, error) {

	url := fmt.Sprintf("https://swap-api.solanatracker.io/swap?from=%s&to=%s&fromAmount=%f&slippage=%f&payer=%s", fromMint, toMint, fromAmount, slipPage, payerAddress)

	response, err := s.client.Get(url)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode == 500 && s.retryOn500 {
		time.Sleep(s.retryOn500Wait)
		return s.GetSwapTransaction(fromMint, toMint, fromAmount, slipPage, payerAddress)
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("status code %d. Details: %s", response.StatusCode, string(body))
	}
	var result SwapResponse

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}
	return &result, nil
}
