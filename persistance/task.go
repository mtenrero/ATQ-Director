package persistance

import (
	"strconv"

	"github.com/mtenrero/ATQ-Director/app"
	"github.com/tidwall/buntdb"
)

const TaskFileId = ":fileId"
const TaskDelay = ":delay"
const TaskMaster = ":master"
const TaskWorker = ":worker"
const TaskStatus = ":status"
const TaskAliasMaster = ":masteralias"
const TaskAliasWorker = ":workeralias"

// IndexTask set the Index in the Task K/V storage
func (p *Persistance) IndexTask() {
	p.DB.CreateIndex("taskalias", "task:*:alias", buntdb.IndexString)
	p.DB.CreateIndex("taskid", "task:*:", buntdb.IndexUint)
}

// StoreTask stores a Task in the Datastore
func (p *Persistance) StoreTask(alias string, fileID *string, delay int) error {
	var baseKey = "task:" + alias

	if fileID != nil {
		var fileIdKey = baseKey + TaskFileId
		p.store(fileIdKey, *fileID)
	}

	delayString := strconv.FormatInt(int64(delay), 10)
	var delayKey = baseKey + TaskDelay
	err := p.store(delayKey, delayString)
	if err != nil {
		return err
	}

	var statusKey = baseKey + TaskStatus
	return p.store(statusKey, "initializing")
}

// StoreTaskMaster persists the Master service ID inside taskId properties
func (p *Persistance) StoreTaskMaster(taskId, serviceID, alias string) error {
	var baseKey = "task:" + taskId

	var master = baseKey + TaskMaster
	err := p.store(master, serviceID)
	if err != nil {
		return err
	}

	var aliasKey = baseKey + TaskAliasMaster
	err = p.store(aliasKey, alias)

	return err
}

// StoreTaskWorker persists the Master service ID inside taskId properties
func (p *Persistance) StoreTaskWorker(taskId, serviceID, alias string) error {
	var baseKey = "task:" + taskId

	var worker = baseKey + TaskWorker
	err := p.store(worker, serviceID)
	if err != nil {
		return err
	}

	var aliasKey = baseKey + TaskAliasWorker
	err = p.store(aliasKey, alias)

	return err
}

func (p *Persistance) ReadTask(taskId string) (*app.AtqTaskFull, error) {
	var fileID string
	var fileIDerr error

	var err error

	var delay string
	var masterID string
	var workerID string
	var status string

	fileID, fileIDerr = p.read("task:" + taskId + TaskFileId)
	if fileIDerr != nil {
		fileID = "/tmp"
	}

	delay, err = p.read("task:" + taskId + TaskDelay)
	if err != nil {
		return nil, err
	}

	masterID, err = p.read("task:" + taskId + TaskMaster)
	if err != nil {
		return nil, err
	}

	workerID, err = p.read("task:" + taskId + TaskWorker)
	if err != nil {
		return nil, err
	}

	delayInt, err := strconv.Atoi(delay)
	if err != nil {
		return nil, err
	}

	status, err = p.read("task:" + taskId + TaskStatus)
	if err != nil {
		return nil, err
	}

	workerAlias, err := p.read("task:" + taskId + TaskAliasWorker)
	if err != nil {
		return nil, err
	}

	masterAlias, err := p.read("task:" + taskId + TaskAliasMaster)
	if err != nil {
		return nil, err
	}

	task := app.AtqTaskFull{
		Delay: &delayInt,
		ID:    &taskId,
		Master: &app.AtqService{
			FileID: &fileID,
			ID:     &masterID,
			Alias:  &masterAlias,
		},
		Status: &status,
		Worker: &app.AtqService{
			FileID: &fileID,
			ID:     &workerID,
			Alias:  &workerAlias,
		},
	}

	return &task, nil
}
