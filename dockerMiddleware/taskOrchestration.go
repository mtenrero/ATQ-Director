package dockerMiddleware

import (
	"errors"
	"time"

	"github.com/mtenrero/ATQ-Director/persistance"

	dockerTypes "docker.io/go-docker/api/types"

	"github.com/mtenrero/ATQ-Director/app"
	"github.com/mtenrero/ATQ-Director/types"
)

// GlobalTimeoutSeconds specify the predefined timeout for the orchestration processes
const GlobalTimeoutSeconds = 60

// TaskMasterWorker initializes a new Master/Worker Task Type
func TaskMasterWorker(task *app.TaskPayload, persistance *persistance.Persistance) (*app.AtqTaskFull, error) {
	var worker *app.AtqService
	var master *app.AtqService
	//var peers *[]network.PeerInfo
	var err error

	// Register Task in the Datastore
	err = persistance.StoreTask(task.Name, task.Worker.Fileid, *task.Delay)
	if err != nil {
		return nil, err
	}

	worker, err = InitService(Worker, task.Name, task.Worker, nil)
	if err != nil {
		return nil, err
	}

	// Wait until service upstart
	errWaiting := VIPSWaiter(*worker.ID, *task.Worker.Replicas, GlobalTimeoutSeconds)
	if err != nil {
		return nil, errWaiting
	}

	// Register Worker in the Datastore
	err = persistance.StoreTaskWorker(task.Name, *worker.ID, task.Worker.Alias)

	err = WorkerHealthchecks(*worker.ID, *task.Worker.Replicas, GlobalTimeoutSeconds)
	if err != nil {
		return nil, err
	}

	// Obtain Service Attached Network
	workerNetworkID, err := ServiceAttachedNetworkID(*worker.ID)
	if err != nil {
		return nil, err
	}

	master, err = InitService(Master, task.Name, task.Master, workerNetworkID)
	if err != nil {
		return nil, err
	}

	// Wait until service upstart
	errWaiting = VIPSWaiter(*worker.ID, *task.Worker.Replicas+*task.Master.Replicas, GlobalTimeoutSeconds)
	if err != nil {
		return nil, errWaiting
	}

	// Register Master in the Datastore
	err = persistance.StoreTaskMaster(task.Name, *master.ID, task.Master.Alias)

	status := "started"

	taskFull := app.AtqTaskFull{
		ID:     &task.Name,
		Delay:  task.Delay,
		Status: &status,
		Master: master,
		Worker: worker,
	}

	// TO-DO To be implemented
	return &taskFull, err
}

// InitService initializes the service
func InitService(serviceType Service, globalAlias string, service *app.ServicePayload, networkID *string) (*app.AtqService, error) {

	var workerBaseAlias = globalAlias + "_" + service.Alias + serviceType.Name()
	var volumeBindPath = service.Fileid
	var serviceCreateResponse *dockerTypes.ServiceCreateResponse
	var err error

	serviceImage := types.ServiceImage{
		ImageName: service.Image,
		TTY:       *service.Tty,
	}

	switch serviceType {
	case Master:
		if networkID == nil {
			return nil, errors.New("networkID must be specified when creating Master Services")
		}
		serviceCreateResponse, err = ComposeService(&serviceImage, globalAlias, workerBaseAlias, volumeBindPath, replicatedService(*service.Replicas), networkID)
	case Worker:
		serviceCreateResponse, err = ComposeService(&serviceImage, globalAlias, workerBaseAlias, volumeBindPath, replicatedService(*service.Replicas), nil)
	}

	if err != nil {
		return nil, err
	}

	serviceResponse := app.AtqService{
		Alias:  &service.Alias,
		ID:     &serviceCreateResponse.ID,
		FileID: service.Fileid,
	}

	return &serviceResponse, nil
}

// WorkerHealthchecks runs the given healthchecks in a Service to ensure all worker containers are ready for use
func WorkerHealthchecks(containerID string, replicas int, timeoutSeconds int) error {

	err := VIPSWaiter(containerID, replicas, timeoutSeconds)
	if err != nil {
		return err
	}
	return nil
}

// VIPSWaiter await for a given amount of VIPS specified in the parameter.
// Timeout in seconds
func VIPSWaiter(serviceID string, replicas int, timeout int) error {

	var vipsExpected = replicas

	timeoutchan := time.After(time.Duration(timeout) * time.Second)
	tick := time.Tick(500 * time.Millisecond)

	// Obtain Service Attached Network
	networkID, err := ServiceAttachedNetworkID(serviceID)
	if err != nil {
		return err
	}

	for {
		select {
		case <-timeoutchan:
			return errors.New("Timed out waiting for all VIPS. Containers are not ready")
		case <-tick:
			vips, err := NetworkVIPs(*networkID)
			if err != nil {
				return err
			}
			vipsAmount := len(*vips)

			if vipsAmount == vipsExpected {
				return nil
			}
		}
	}
}
