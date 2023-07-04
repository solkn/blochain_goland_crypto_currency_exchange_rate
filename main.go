package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"project/controllers"
	"project/db"

	"github.com/lucazulian/cryptocomparego"
	"github.com/lucazulian/cryptocomparego/context"
)

func main() {

	ctx := context.TODO()

	client := cryptocomparego.NewClient(nil)
	coinList, _, err1 := client.Coin.List(ctx)

	// fmt.Print("coin,", coinList)

	if err1 != nil {
		fmt.Printf("Error: %s\n", err1)
		// return err1
		fmt.Print("coin,", coinList)

	}

	price, _, err2 := client.Price.List(ctx, &cryptocomparego.PriceRequest{Fsym: "USDT", Tsyms: []string{"USD", "JPY", "EUR"}})

	fmt.Println("price,", price)

	if err2 != nil {
		fmt.Println("error,", err2)
	}

	router := gin.Default()

	router.SetTrustedProxies(nil)

	db.ConnectDatabase()

	router.GET("/rates", controllers.GetCryptoCurrencies)
	router.POST("/rates", controllers.CreateCryptoCurrency)
	router.GET("/rates/:id", controllers.GetCryptoToAllCurrencyExchangeRates)
	router.PATCH("/rates/:id", controllers.UpdateCryptoCurrencyExchangeRates)

	err := router.Run("127.0.0.1:8000")
	if err != nil {
		return
	}
}
