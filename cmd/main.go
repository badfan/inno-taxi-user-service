package main

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/badfan/inno-taxi-user-service/app/rpc"

	"github.com/badfan/inno-taxi-user-service/app/services/order"

	"github.com/badfan/inno-taxi-user-service/app/services/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/badfan/inno-taxi-user-service/app"

	"github.com/badfan/inno-taxi-user-service/app/api"
	v1 "github.com/badfan/inno-taxi-user-service/app/api/v1"
	"github.com/badfan/inno-taxi-user-service/app/handlers"
	"github.com/badfan/inno-taxi-user-service/app/resources"
	"github.com/badfan/inno-taxi-user-service/app/services/auth"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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

func InitGRPCClient(apiConfig *app.APIConfig, logger *zap.SugaredLogger) *grpc.ClientConn {
	var options []grpc.DialOption
	options = append(options, grpc.WithTransportCredentials(insecure.NewCredentials()))

	orderClientConn, err := grpc.Dial("localhost:"+apiConfig.RPCOrderPort, options...)
	if err != nil {
		logger.Fatalf("error occurred while connecting to order GRPC server: %s", err.Error())
	}

	return orderClientConn
}

func InitGRPCServer(rpcService *rpc.RPCService, apiConfig *app.APIConfig, logger *zap.SugaredLogger) {
	listener, err := net.Listen("tcp", "localhost:"+apiConfig.RPCUserPort)
	if err != nil {
		logger.Fatalf("failed to up user GRPC server: %s", err.Error())
	}

	var options []grpc.ServerOption
	rpcServer := grpc.NewServer(options...)
	rpc.RegisterUserServiceServer(rpcServer, rpcService.UserServiceServer)
	rpcServer.Serve(listener)
}

func main() {
	logger := InitLogger()
	defer logger.Sync()

	apiConfig, err := app.NewAPIConfig()
	if err != nil {
		logger.Fatalf("error occurred while preparing apiconfig: %s", err.Error())
	}
	dbConfig, err := app.NewDBConfig()
	if err != nil {
		logger.Fatalf("error occurred while preparing dbconfig: %s", err.Error())
	}

	resource, err := resources.NewResource(dbConfig, logger)
	if err != nil {
		logger.Fatalf("error occurred while creating new resource: %s", err.Error())
	}
	defer resource.Db.Close()

	orderClientConn := InitGRPCClient(apiConfig, logger)
	defer orderClientConn.Close()

	rpcService := rpc.NewRPCService()
	InitGRPCServer(rpcService, apiConfig, logger)

	orderService := order.NewOrderService(rpcService, orderClientConn)
	authService := auth.NewAuthenticationService(resource, logger)
	userService := user.NewUserService(resource, orderService, apiConfig, logger)
	handler := handlers.NewHandler(authService, userService, logger)

	router := InitRouter(handler)

	server := new(Server)
	if err := server.Run(router, apiConfig.APIPort); err != nil {
		logger.Fatalf("error occurred while running http server: %s", err.Error())
	}
}
