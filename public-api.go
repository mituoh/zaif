package zaif

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const publicEndPointURL = "https://api.zaif.jp/api/1"

// TmpPublicAPI ...
type TmpPublicAPI struct{}

// PublicAPI ...
var PublicAPI = TmpPublicAPI{}

// LastPrice 終値
type LastPrice struct {
	LastPrice float64 `json:"last_price"` // 終値
}

// Ticker ティッカー
type Ticker struct {
	Last   float64 `json:"last"`   // 終値
	High   float64 `json:"high"`   // 過去24時間の高値
	Low    float64 `json:"low"`    // 過去24時間の安値
	Vmap   float64 `json:"vmap"`   // 過去24時間の加重平均
	Volume float64 `json:"volume"` // 過去24時間の出来高
	Bid    float64 `json:"bid"`    // 買気配値
	Ask    float64 `json:"ask"`    // 売気配値
}

// Trade 取引情報
type Trade struct {
	Date         int     `json:"date"`
	Price        float64 `json:"price"`
	Amount       float64 `json:"amount"`
	Tid          int     `json:"tid"`
	CurrencyPair string  `json:"currency_pair"`
	TradeType    string  `json:"trade_type"`
}

// Ask 価格, 数量
type Ask []float64

// Depth 板
type Depth struct {
	Asks []Ask `json:"asks"`
}

// LastPrice 終値取得
func (api TmpPublicAPI) LastPrice(currencyPair string) (dat LastPrice, err error) {
	res, err := http.Get(publicEndPointURL + "/last_price/" + currencyPair)
	if err != nil {
		return LastPrice{}, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return LastPrice{}, err
	}
	json.Unmarshal(body, &dat)
	return dat, nil
}

// Ticker ティッカー取得
func (api TmpPublicAPI) Ticker(currencyPair string) (dat Ticker, err error) {
	res, err := http.Get(publicEndPointURL + "/ticker/" + currencyPair)
	if err != nil {
		return Ticker{}, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Ticker{}, err
	}
	json.Unmarshal(body, &dat)
	return dat, nil
}

// Trades 全ての取引履歴取得
func (api TmpPublicAPI) Trades(currencyPair string) (dat []Trade, err error) {
	res, err := http.Get(publicEndPointURL + "/trades/" + currencyPair)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(body, &dat)
	return dat, nil
}

// Depth 板情報取得
func (api TmpPublicAPI) Depth(currencyPair string) (dat Depth, err error) {
	res, err := http.Get(publicEndPointURL + "/depth/" + currencyPair)
	if err != nil {
		return Depth{}, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Depth{}, err
	}
	json.Unmarshal(body, &dat)
	return dat, nil
}
