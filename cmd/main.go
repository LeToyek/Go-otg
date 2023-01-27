package main

import (
	"go-otg/internal/app/http"
	"log"

	"github.com/spf13/viper"
)

// main function
func main() {
	err := StartENV()
	if err != nil {
		log.Fatal(err.Error())
	}
	if 1+1 == 2 {
		http.NewApplication()
	}
}
func StartENV() error {
	viper.SetConfigType("yaml")
	viper.SetConfigName("app.config")
	viper.AddConfigPath("../config")
	viper.AddConfigPath("./config")

	return viper.ReadInConfig()
}
