package binance

import (
	"context"
	"encoding/json"
)

// ExchangeInfoService exchange info service
type ExchangeInfoService struct {
	c *Client
}

// Do send request
func (s *ExchangeInfoService) Do(ctx context.Context, opts ...RequestOption) (res *ExchangeInfo, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/api/v1/exchangeInfo",
		secType:  secTypeNone,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(ExchangeInfo)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// ExchangeInfo exchange info
type ExchangeInfo struct {
	Timezone        string        `json:"timezone"`
	ServerTime      int64         `json:"serverTime"`
	RateLimits      []RateLimit   `json:"rateLimits"`
	ExchangeFilters []interface{} `json:"exchangeFilters"`
	Symbols         []Symbol      `json:"symbols"`
}

// RateLimit struct
type RateLimit struct {
	RateLimitType string `json:"rateLimitType"`
	Interval      string `json:"interval"`
	Limit         int64  `json:"limit"`
}

// Symbol market symbol
type Symbol struct {
	Symbol             string                   `json:"symbol"`
	Status             string                   `json:"status"`
	BaseAsset          string                   `json:"baseAsset"`
	BaseAssetPrecision int                      `json:"baseAssetPrecision"`
	QuoteAsset         string                   `json:"quoteAsset"`
	QuotePrecision     int                      `json:"quotePrecision"`
	OrderTypes         []string                 `json:"orderTypes"`
	IcebergAllowed     bool                     `json:"icebergAllowed"`
	Filters            []map[string]interface{} `json:"filters"`
}

// AssetDetailService asset detail service
type AssetDetailService struct {
	c *Client
}

// Do send request
func (ads *AssetDetailService) Do(ctx context.Context, opts ...RequestOption) (res *AssetDetail, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/wapi/v3/assetDetail.html",
		secType:  secTypeNone,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(AssetDetail)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

//AssetDetail asset detail
type AssetDetail struct {
	AssetDetail map[string]DetailOfAsset `json:"assetDetail"`
}

//DetailOfAsset return detail of an asset
type DetailOfAsset struct {
	MinWithdrawAmount float64 `json:"minWithdrawAmount"`
	DepositStatus     bool    `json:"depositStatus"`
	WithdrawFee       float64 `json:"withdrawFee"`
	WithdrawStatus    bool    `json:"withdrawStatus"`
	DepositTip        string  `json:"depositTip"`
}