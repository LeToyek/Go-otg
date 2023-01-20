package utils

import (
	"github.com/spf13/viper"
)

func StartENV() {
	viper.Set("PORT", "localhost:5000")
	viper.Set("REDIS_HOST", "localhost:6379")
	viper.Set("REDIS_PASSWORD", "")
}
