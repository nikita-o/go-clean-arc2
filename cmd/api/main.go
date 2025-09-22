package main

import (
	"clean-arc-2/internal/handlers/http/controllers"
	"clean-arc-2/internal/handlers/http/routes"
	"clean-arc-2/internal/infrastructure/config"
	"clean-arc-2/internal/infrastructure/database"
	"clean-arc-2/internal/infrastructure/logger"
	"clean-arc-2/internal/repositories"
	use_case "clean-arc-2/internal/use-case"
	"github.com/gin-gonic/gin"
)

func main() {
	log := logger.NewLogger()
	defer log.Sync()

	env := config.NewConfig()
	pg := database.NewPostgresDatabase(env.DbUrl)
	rep := repositories.NewRepositories(pg)
	useCases := use_case.NewUseCases(*rep)
	c := controllers.NewControllers(*useCases)

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())
	//r.Use(middleware.ZapLoggerMiddleware(log))

	//timeout := time.Duration(env.ContextTimeout) * time.Second

	routes.NewApiRouter(r, *c)

	err := r.Run(env.ServerAddress)
	if err != nil {
		log.Error(err.Error())
		return
	}
}
