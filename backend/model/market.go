package model

import (
	"io/ioutil"
	"net/http"

	"fmt"
	"github.com/bitly/go-simplejson"
	"time"
	"explorer/model/cache"
)

const (
	CoinMarketCapCacheKey     = "coinMarketCapInfo"
	CoinMarketCapIOSOfBTCTUrl = "https://api.coinmarketcap.com/v2/ticker/2405?convert=BTC"
	CoinMarketCapIOSOfETHTUrl = "https://api.coinmarketcap.com/v2/ticker/2405?convert=ETH"
)

type MarketInfo struct {
	Price            string  `json:"price"`
	Volume24h        int64   `json:"volume_24h"`
	PercentChange24h float64 `json:"percent_change_24h"`
	MarketCap        int64   `json:"market_cap"`
	BtcPrice         string  `json:"btc_price"`
	EthPrice         string  `json:"eth_price"`
	LastUpdate       string  `json:"last_update"`
}

func GetMarketInfo() (*MarketInfo, error) {
	// get cache first.
	if marketInfoInterface, ok := cache.GlobalCache.Get(CoinMarketCapCacheKey); ok {
		if marketInfo, ok := marketInfoInterface.(*MarketInfo); ok {
			return marketInfo, nil
		}
	}

	resp, err := http.Get(CoinMarketCapIOSOfBTCTUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	respObj, err := simplejson.NewJson(body)
	if err != nil {
		return nil, err
	}

	var ethPrice float64 = 0.000067
	if ethResp, err := http.Get(CoinMarketCapIOSOfETHTUrl); err == nil {
		defer ethResp.Body.Close()

		body, err := ioutil.ReadAll(ethResp.Body)
		if err != nil {
			return nil, err
		}

		if ethRespObj, err := simplejson.NewJson(body); err == nil {
			if price, err := ethRespObj.GetPath("data", "quotes", "ETH", "price").Float64(); err == nil && price > 0 {
				ethPrice = price
			}
		}
	}

	price, _ := respObj.GetPath("data", "quotes", "USD", "price").Float64()

	volume24h, _ := respObj.GetPath("data", "quotes", "USD", "volume_24h").Float64()

	percentChange24h, _ := respObj.GetPath("data", "quotes", "USD", "percent_change_24h").Float64()

	marketCap, _ := respObj.GetPath("data", "quotes", "USD", "market_cap").Float64()

	btcPrice, _ := respObj.GetPath("data", "quotes", "BTC", "price").Float64()

	lastUpdate, _ := respObj.GetPath("data", "last_updated").Float64()

	mi := &MarketInfo{
		fmt.Sprintf("%.3f", price),
		int64(volume24h),
		percentChange24h,
		int64(marketCap),
		fmt.Sprintf("%.10f", btcPrice),
		fmt.Sprintf("%.6f", ethPrice),
		modifyIntToTimeStr(int64(lastUpdate)),
	}

	// set to cache
	cache.GlobalCache.Set(CoinMarketCapCacheKey, mi, time.Minute)

	return mi, nil
}
