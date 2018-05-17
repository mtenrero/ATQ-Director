package dockerMiddleware

import (
	"github.com/mtenrero/ATQ-Director/app"
	"github.com/mtenrero/ATQ-Director/types"
)

// TaskMasterWorker initializes a new Master/Worker Task Type
func TaskMasterWorker(task *app.TaskPayload) (*app.AtqTask, error) {
	worker, err := InitWorkerService(task.Name, task.Worker)

	return worker, err
}

// InitWorkerService initializes the Worker Service and attach them to a random generated Network
func InitWorkerService(globalAlias string, worker *app.ServicePayload) (*app.AtqTask, error) {

	var workerBaseAlias = globalAlias + "_" + worker.Alias

	serviceImage := types.ServiceImage{
		ImageName: worker.Image,
		TTY:       *worker.Tty,
	}

	serviceCreateResponse, err := ComposeService(&serviceImage, globalAlias, workerBaseAlias, "", replicatedService(*worker.Replicas))

	if err != nil {
		return nil, err
	}

	taskResponse := app.AtqTask{
		ID: &serviceCreateResponse.ID,
	}

	return &taskResponse, nil
}

// InitMasterService initializes the Master Service and attach it to the given Network name
func InitMasterService(master *app.ServicePayload) (*app.ServicePayload, error) {

	return nil, nil
}

// WorkerHealthchecks runs the given healthchecks in a Service to ensure all worker containers are ready for use
func WorkerHealthchecks() error {

	return nil
}
