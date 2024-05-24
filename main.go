package main

import (
	"github.com/codingsandmore/solana-tracker-golang/tracker"
	"log"
	"net/http"
)

func main() {

	//get a new tracker
	swap := tracker.NewSolanaTracker(&http.Client{})

	//get a quote

	rate, err := swap.GetRate("So11111111111111111111111111111111111111112", "4k3Dyjzvzp8eMZWUXbBCjEvwSkkk59S5iCNLY3QrkX6R", 10, 15)

	//get the swap transaction
	tx, err := swap.GetSwapTransaction("So11111111111111111111111111111111111111112", "4k3Dyjzvzp8eMZWUXbBCjEvwSkkk59S5iCNLY3QrkX6R", 10, 15, "44ByXN2NgCeHJdsmbUvKrjcdRQJLJDuY8DqnqWMcwy5U", 0.015)

	if err != nil {
		log.Fatal(err)
	}

	//do something with these information now
	log.Println(rate)
	log.Println(tx)
}
