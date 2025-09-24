package main

import (
	"clean-arc-2/internal/handlers/http"
	"clean-arc-2/internal/handlers/http/controllers"
	"clean-arc-2/internal/infrastructure/config"
	"clean-arc-2/internal/infrastructure/database"
	"clean-arc-2/internal/infrastructure/logger"
	"clean-arc-2/internal/repositories"
	"clean-arc-2/internal/use-case"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"time"
)

type AppApi struct {
	Config   *config.Config
	Logger   *zap.Logger
	Timeout  time.Duration
	Database *pgxpool.Pool
	//Redis        *redis.Client

	Repositories *repositories.Repositories
	UseCases     *use_case.UseCases
	Controllers  *controllers.Controllers

	HttpServer *http.HttpServer
}

func main() {
	app := AppApi{}

	app.Config = config.NewConfig()
	app.Logger = logger.NewLogger()
	defer app.Logger.Sync()
	app.Timeout = time.Duration(app.Config.ContextTimeout) * time.Second

	// Инициализация сервисов приложения
	app.Database = database.NewPostgresDatabase(app.Config.DbUrl)

	app.Repositories = repositories.NewRepositories(repositories.DataInitRepositories{
		Database: app.Database,
	})

	app.UseCases = use_case.NewUseCases(use_case.DataInitUseCases{
		Repositories: app.Repositories,
	})

	app.HttpServer = http.NewHttpServer(http.DataInitHttpServer{
		UseCases: app.UseCases,
	})
	
	err := app.HttpServer.Server.Run(app.Config.ServerAddress)
	if err != nil {
		app.Logger.Error(err.Error())
		return
	}
}
