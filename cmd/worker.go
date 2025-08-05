package cmd

import (
	asynq_ "github.com/hibiken/asynq"
	"github.com/hossainabid/go-ims/config"
	"github.com/hossainabid/go-ims/conn"
	"github.com/hossainabid/go-ims/controllers"
	asynq_repo "github.com/hossainabid/go-ims/repositories/asynq"
	db_repo "github.com/hossainabid/go-ims/repositories/db"
	mail_repo "github.com/hossainabid/go-ims/repositories/mail"
	"github.com/hossainabid/go-ims/services"
	"github.com/hossainabid/go-ims/types"
	"github.com/hossainabid/go-ims/worker"
	"github.com/spf13/cobra"
)

var workerCmd = &cobra.Command{
	Use: "worker",
	Run: runWorker,
}

func runWorker(cmd *cobra.Command, args []string) {
	// clients
	dbClient := conn.Db()
	emailClient := conn.EmailClient()
	asynqClient := conn.Asynq()
	asynqInspector := conn.AsynqInspector()
	workerPool := conn.WorkerPool()

	// repositories
	dbRepo := db_repo.NewRepository(dbClient)
	asynqRepo := asynq_repo.NewRepository(config.Asynq(), asynqClient, asynqInspector)
	mailRepo := mail_repo.NewRepository(emailClient, config.Email())

	// services
	_ = services.NewAsynqService(config.Asynq(), asynqRepo)
	productSvc := services.NewProductServiceImpl(dbRepo)
	mailSvc := services.NewMailService(dbRepo, dbRepo, mailRepo, workerPool)

	// controllers
	asynqCtrl := controllers.NewAsynqController(productSvc, mailSvc)

	mux := asynq_.NewServeMux()

	mux.HandleFunc(types.AsynqTaskTypeStockSync.String(), asynqCtrl.ProcessStockSyncTask)
	// Start the Asynq worker
	worker.StartAsynqWorker(mux)

}
