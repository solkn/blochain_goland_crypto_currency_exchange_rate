package models

type CryptoFiatCurrency struct{
	Id uint `gorm:"primary_key;auto_increment" json:"id"`
	Coin string `json:"coin"`
	USD float64 `json:"usd"`
	JPY float64 `json:"jpy"`
	EUR float64 `json:"eur"`
}