# solana-tracker-golang

This is a golang api for the solana tracker API.

Which is hosted here https://docs.solanatracker.io/swap-api/swap and I'm not affiliated with them or their services.

## Usage Example

The easiest way is to look at the main.go method, which provides the following example

```go

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
	tx, err := swap.GetSwapTransaction("So11111111111111111111111111111111111111112", "4k3Dyjzvzp8eMZWUXbBCjEvwSkkk59S5iCNLY3QrkX6R", 10, 15, "44ByXN2NgCeHJdsmbUvKrjcdRQJLJDuY8DqnqWMcwy5U")

	if err != nil {
		log.Fatal(err)
	}

	//do something with these information now
	log.Println(rate)
	log.Println(tx)
}
```

### Getting the lib into your project

```bash
go get github.com/codingsandmore/solana-tracker-golang
```

### Need coffee!!!

we would love to receive your support and would appreciate it, if you send us some SOL for coffee, if you consider this
useful :)

```
D79Tjjs6GpsUL2Pxd7PNsWuBaRJs5Qdt1XNorXRD5Azd
```

### Collaboration

Feel free to reach out to me, I'm always interested in some fun collaborations.

### Disclaimer

all tools provided, are for learning purposes only, without any form or kind of warranty. Please use at your own risk.