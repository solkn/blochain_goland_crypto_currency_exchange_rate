package controllers

import (
	"fmt"
	"net/http"

	"project/db"
	"project/models"

	"github.com/gin-gonic/gin"
	"github.com/lucazulian/cryptocomparego"
	"github.com/lucazulian/cryptocomparego/context"
)

func GetCryptoCurrencies(context *gin.Context) {

	var currency []models.CryptoFiatCurrency

	db.DB.Find(&currency)

	context.JSON(http.StatusOK, gin.H{"data": currency})

}
func CreateCryptoCurrency(ct *gin.Context) {

	ctx := context.TODO()

	client := cryptocomparego.NewClient(nil)
	coinList, _, err1 := client.Coin.List(ctx)

	// fmt.Print("coin,", coinList)

	if err1 != nil {
		fmt.Printf("Something bad happened: %s\n", err1)
		// return err1
		fmt.Print("coin,", coinList)

	}

	price, _, err2 := client.Price.List(ctx, &cryptocomparego.PriceRequest{Fsym: "BTC", Tsyms: []string{"USD", "JPY", "EUR"}})

	fmt.Println("price,", price[0].Value)

	var price_rslt = price[0]

	fmt.Println("name", price_rslt.Name)
	fmt.Println("value", price_rslt.Value)
	fmt.Println("value")

	if err2 != nil {
		fmt.Print("mm", price_rslt)
	}

	var input models.CreateCryptoFiatInput
	if err := ct.ShouldBindJSON(&input); err != nil {
		ct.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	currency := models.CryptoFiatCurrency{
		Coin: input.Coin,
		USD:  input.USD,
		JPY:  input.JPY,
		EUR:  input.EUR,
	}

	fmt.Println("currency", currency)

	var a = db.DB.Create(&currency)

	fmt.Print("currency", a)

	ct.JSON(http.StatusOK, gin.H{"data": currency})

}

func GetCryptoToAllCurrencyExchangeRates(context *gin.Context) {
	var currency models.CryptoFiatCurrency

	if err := db.DB.Where("id = ?", context.Param("id")).First(&currency).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Curency not found!"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": currency})
}
func GetSpecifiedCryptoCurrencyExchangeRates(context *gin.Context) {
	var currency models.CryptoFiatCurrency

	if err := db.DB.Where("id = ?", context.Param("id")).First(&currency).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Curency not found!"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": currency})
}
func UpdateCryptoCurrencyExchangeRates(context *gin.Context) {
	var currency models.CryptoFiatCurrency
	if err := db.DB.Where("id = ?", context.Param("id")).First(&currency).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "currency not found!"})
		return
	}

	var input models.UpdateCryptoFiatInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Model(&currency).Updates(input)

	context.JSON(http.StatusOK, gin.H{"data": currency})
}
