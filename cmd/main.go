package main

import (
	"context"
	"github.com/badfan/inno-taxi-user-service/app/api"
	"github.com/badfan/inno-taxi-user-service/app/api/v1"
	"github.com/badfan/inno-taxi-user-service/app/handlers"
	"github.com/badfan/inno-taxi-user-service/app/resources"
	"github.com/badfan/inno-taxi-user-service/app/services"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(handler http.Handler, port string) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func InitLogger() *zap.SugaredLogger {
	logger, _ := zap.NewProduction()
	sugarLogger := logger.Sugar()
	return sugarLogger
}

func InitRouter(handler *handlers.Handler) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	apiV1 := v1.NewApiV1(handler)
	apiGroup := api.NewApiGroup(handler, apiV1)
	apiGroup.InitRouterGroups(router)

	return router
}

func main() {
	logger := InitLogger()
	defer logger.Sync()

	resource, err := resources.NewResource(logger)
	if err != nil {
		logger.Fatalf("db error: %s", err.Error())
	}
	defer resource.Db.Close()

	service := services.NewService(resource, logger)
	handler := handlers.NewHandler(service, logger)

	router := InitRouter(handler)

	server := new(Server)
	if err := server.Run(router, "8080"); err != nil {
		logger.Fatalf("error occured while running http server: %s", err.Error())
	}
}
