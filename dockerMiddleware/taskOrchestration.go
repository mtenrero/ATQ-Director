package dockerMiddleware

import (
	"errors"
	"time"

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
	var volumeBindPath = worker.Fileid

	serviceImage := types.ServiceImage{
		ImageName: worker.Image,
		TTY:       *worker.Tty,
	}

	serviceCreateResponse, err := ComposeService(&serviceImage, globalAlias, workerBaseAlias, volumeBindPath, replicatedService(*worker.Replicas))

	if err != nil {
		return nil, err
	}

	taskResponse := app.AtqTask{
		ID: &serviceCreateResponse.ID,
	}

	errWaiting := VIPSWaiter(*taskResponse.ID, *worker.Replicas, 60)
	if err != nil {
		return nil, errWaiting
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

// VIPSWaiter await for a given amount of VIPS specified in the parameter.
// Timeout in seconds
func VIPSWaiter(serviceID string, replicas int, timeout int) error {

	var vipsExpected = replicas - 1

	timeoutchan := time.After(time.Duration(timeout) * time.Second)
	tick := time.Tick(500 * time.Millisecond)

	for {
		select {
		case <-timeoutchan:
			return errors.New("Timed out waiting for all VIPS. Containers are not ready")
		case <-tick:
			peers, err := NetworkContainerPeers(serviceID)
			if err != nil {
				return err
			}
			vipsAmount := len(*peers)

			if vipsAmount == vipsExpected {
				return nil
			}
		}
	}
}
