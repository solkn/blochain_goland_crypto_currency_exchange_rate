package models

type CreateCryptoFiatInput struct {
	Coin string `json:"coin" binding:"required"`
	USD  float64 `json:"usd" binding:"required"`
	JPY  float64 `json:"jpy" binding:"required"`
	EUR  float64 `json:"eur" binding:"required"`
}

type UpdateCryptoFiatInput struct {
	Coin string `json:"coin" binding:"required"`
	USD  float64 `json:"usd" binding:"required"`
	JPY  float64 `json:"jpy" binding:"required"`
	EUR  float64 `json:"eur" binding:"required"`
}