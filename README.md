# zaif

zaif is exchange market for bitcoin and monacoin.
you can also trade if you have api-key of zaif.

# install
```Go
go get github.com/yanakend/zaif
# install depend on package
go get github.com/google/go-querystring/query
```

# usage
```Go
package main

import (
	"github.com/yanakend/zaif-api"
	"log"
)

func main() {
	zaifPrivateAPI := zaif_api.NewPrivateAPI(
		"Input key here",
		"Input secret key here",
	)
	price := zaif_api.PublicAPI.LastPrice("btc_jpy")
	log.Println(price)
	history := zaifPrivateAPI.DepositHistory(zaif_api.BodyDepositHistory{Currency: "jpy", Order: "ASC"})
	log.Println(history)
}
```

# Public API

#### PublicAPI.LastPrice
#### PublicAPI.Ticker
#### PublicAPI.Trades
#### PublicAPI.Depth
#### PublicAPI.Withdraw
#### PublicAPI.DepositHistory
#### PublicAPI.WithdrawHistory

# Private API

#### PrivateAPI.GetInfo
#### PrivateAPI.ActiveOrders
#### PrivateAPI.Trade
#### PrivateAPI.Cancel
#### PrivateAPI.Withdraw
#### PrivateAPI.DepositHistory
#### PrivateAPI.WithdrawHistory


