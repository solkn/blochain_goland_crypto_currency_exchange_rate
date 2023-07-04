package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"project/models"
)

var DB *gorm.DB

func createTables(dbConn *gorm.DB) []error {

	dbConn.DropTableIfExists(&models.CryptoFiatCurrency{}).GetErrors()

	errs := dbConn.CreateTable(&models.CryptoFiatCurrency{}).GetErrors()

	if len(errs) > 0 {
		return errs
	}
	return nil
}

func ConnectDatabase() {


	dbconn, err := gorm.Open("postgres",
		"postgres://postgres:postgres@localhost/crypto?sslmode=disable")


	if err != nil {
		panic(err)
	}

	// createTables(dbconn)

	DB = dbconn

}
