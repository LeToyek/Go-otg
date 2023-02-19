package http

import (
	// golang package

	"fmt"
	"go-otg/internal/app/http/server"
	"go-otg/internal/infrastructure"
	hash "go-otg/internal/infrastructure/hash"
	"go-otg/internal/infrastructure/uuid"
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

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost",
		"5432",
		"postgres",
		"handoko",
		"go-otg")
	if os.Getenv("ENV") == "PRODUCTION" {
		psqlInfo = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_OPEN_PORT"),
			os.Getenv("DB_USERNAME"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"))
	}

	sqlDB, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		fmt.Print(err.Error())
		fmt.Println("EMPTY DB CONNECTION STRING")
	}

	client := redis.NewRedisClient(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PASSWORD"))

	repositoryRedis := redis.NewRedis(client)
	repositoryDB := db.New(sqlDB)
	infra := infrastructure.New(hash.NewHasher(), uuid.NewUUIDService())
	resources := server.NewResources(repositoryDB, repositoryRedis)
	services := server.NewServices(infra, resources)
	usecases := server.NewUsecases(services)
	handlers := server.NewHandlers(usecases)
	httpServer := http.NewHTTPServer(handlers)

	app := gin.Default()
	httpServer.RegisterHandler(app)
	app.Run(os.Getenv("PORT"))
}
