package http

import (
	"clean-arc-2/internal/handlers/http/controllers"
	"clean-arc-2/internal/handlers/http/middlewares"
	"clean-arc-2/internal/handlers/http/routes"
	"clean-arc-2/internal/use-case"
	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	controllers *controllers.Controllers
	middlewares *middlewares.Middlewares
	Server      *gin.Engine
}

type DataInitHttpServer struct {
	UseCases *use_case.UseCases
}

func NewHttpServer(data DataInitHttpServer) *HttpServer {
	server := HttpServer{}
	server.controllers = controllers.NewControllers(controllers.DataInitControllers{
		UseCases: data.UseCases,
	})
	server.middlewares = middlewares.NewMiddlewares(middlewares.DataInitMiddlewares{})

	gin.SetMode(gin.ReleaseMode)
	server.Server = gin.New()
	server.Server.Use(gin.Recovery())
	//r.Use(middleware.ZapLoggerMiddleware(log))

	routes.NewApiRouter(server.Server, server.controllers, server.middlewares)

	return &server
}
