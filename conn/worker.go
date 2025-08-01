package conn

import (
	"github.com/hossainabid/go-ims/config"
	"github.com/hossainabid/go-ims/worker"
)

var workerPool *worker.Pool

func ConnectWorker() {
	workerPool = worker.NewPool(config.App().NumberOfWorkers, 2*config.App().NumberOfWorkers)
	workerPool.Start()
}

func WorkerPool() *worker.Pool {
	return workerPool
}
