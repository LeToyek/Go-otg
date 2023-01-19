package utils

import (
	"github.com/spf13/viper"
)

func StartENV() {
	viper.Set("PORT", "localhost:5000")
}
