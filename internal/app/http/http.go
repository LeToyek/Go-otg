package http

import (
	// golang package
	"context"
	"fmt"
	"go-otg/internal/app/http/server"
	"go-otg/internal/repository/db"
	"go-otg/internal/repository/redis"
	"go-otg/internal/server/http"
	"os"

	// external package
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// NewApplication new application.
func NewApplication() {

	psqlInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	sqlDB, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		fmt.Print(err.Error())
		fmt.Println("EMPTY DB CONNECTION STRING")
		// log.Fatal(err)
	}

	client := redis.NewRedisClient(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PASSWORD"))

	repositoryRedis := redis.NewRedis(client)
	repositoryDB := db.New(sqlDB)
	// fmt.Print("Error Here")
	fmt.Print(repositoryDB.GetUserByID(context.Background(), "asd"))
	// fmt.Print(" <<- error")
	resources := server.NewResources(repositoryDB, repositoryRedis)
	services := server.NewServices(resources)
	usecases := server.NewUsecases(services)
	handlers := server.NewHandlers(usecases)
	httpServer := http.NewHTTPServer(handlers)

	app := gin.Default()
	httpServer.RegisterHandler(app)
	app.Run(os.Getenv("PORT"))
}
