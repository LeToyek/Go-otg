package http

import (
	// golang package
	"fmt"
	"go-otg/internal/app/http/server"
	"go-otg/internal/repository/db"
	"go-otg/internal/repository/redis"
	"go-otg/internal/server/http"
	"go-otg/internal/utils"

	// external package
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

// NewApplication new application.
func NewApplication() {
	sqlDB, err := sqlx.Connect("pq", "")
	if err != nil {
		fmt.Println("EMPTY DB CONNECTION STRING")
		// log.Fatal(err)
	}
	utils.StartENV()
	client := redis.NewRedisClient(viper.GetString("REDIS_HOST"), viper.GetString("REDIS_PASSWORD"))

	repositoryRedis := redis.NewRedis(client)
	repositoryDB := db.New(sqlDB)
	resources := server.NewResources(repositoryDB, repositoryRedis)
	services := server.NewServices(resources)
	usecases := server.NewUsecases(services)
	handlers := server.NewHandlers(usecases)
	httpServer := http.NewHTTPServer(handlers)

	app := gin.Default()
	httpServer.RegisterHandler(app)
	app.Run(viper.GetString("PORT"))
}
