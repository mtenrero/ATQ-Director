package dockerMiddleware

import (
	"github.com/mtenrero/ATQ-Director/app"
)

// TaskMasterWorker initializes a new Master/Worker Task Type
func TaskMasterWorker(task *app.TaskPayload) {
	InitWorkerServices(task.Worker)
}

// InitWorkerServices initializes the Worker Service and attach them to a random generated Network
func InitWorkerServices(worker *app.ServicePayload) (*[]app.ServicePayload, error) {

	return nil, nil
}

// InitMasterService initializes the Master Service and attach it to the given Network name
func InitMasterService(master *app.ServicePayload) (*app.ServicePayload, error) {

	return nil, nil
}

// WorkerHealthchecks runs the given healthchecks in a Service to ensure all worker containers are ready for use
func WorkerHealthchecks() error {

	return nil
}
