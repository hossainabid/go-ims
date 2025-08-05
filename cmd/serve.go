package cmd

import (
	"github.com/hossainabid/go-ims/config"
	"github.com/hossainabid/go-ims/conn"
	"github.com/hossainabid/go-ims/controllers"
	"github.com/hossainabid/go-ims/middlewares"
	asynq_repo "github.com/hossainabid/go-ims/repositories/asynq"
	db_repo "github.com/hossainabid/go-ims/repositories/db"
	"github.com/hossainabid/go-ims/routes"
	"github.com/hossainabid/go-ims/server"
	"github.com/hossainabid/go-ims/services"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use: "serve",
	Run: serve,
}

func serve(cmd *cobra.Command, args []string) {
	// clients
	dbClient := conn.Db()
	redisClient := conn.Redis()
	asynqClient := conn.Asynq()
	asynqInspector := conn.AsynqInspector()

	// repositories
	dbRepo := db_repo.NewRepository(dbClient)
	asynqRepo := asynq_repo.NewRepository(config.Asynq(), asynqClient, asynqInspector)

	// services
	redisSvc := services.NewRedisService(redisClient)
	productSvc := services.NewProductServiceImpl(dbRepo)
	userSvc := services.NewUserServiceImpl(redisSvc, dbRepo)
	stockHistorySvc := services.NewStockHistoryServiceImpl(dbRepo)
	tokenSvc := services.NewTokenServiceImpl(redisSvc)
	authSvc := services.NewAuthServiceImpl(userSvc, tokenSvc)
	asynqSvc := services.NewAsynqService(config.Asynq(), asynqRepo)

	// controllers
	productCtrl := controllers.NewProductController(productSvc)
	userCtrl := controllers.NewUserController(userSvc)
	stockHistoryCtrl := controllers.NewStockHistoryController(stockHistorySvc, asynqSvc)
	authCtrl := controllers.NewAuthController(authSvc)

	// middlewares
	authMiddleware := middlewares.NewAuthMiddleware(authSvc, userSvc)

	// Server
	var echo_ = echo.New()
	var Routes = routes.New(echo_, productCtrl, userCtrl, stockHistoryCtrl, authCtrl, authMiddleware)
	var Server = server.New(echo_)

	// Spooling
	Routes.Init()

	// Stopping running workers
	Server.Start()
}
