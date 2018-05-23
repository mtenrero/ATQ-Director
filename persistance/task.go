package persistance

import (
	"github.com/tidwall/buntdb"
)

const TaskFileId = ":fileId"
const TaskDelay = ":delay"
const TaskMaster = ":master"
const TaskWorker = ":worker"

// IndexTask set the Index in the Task K/V storage
func (p *Persistance) IndexTask() {
	p.DB.CreateIndex("taskalias", "task:*:alias", buntdb.IndexString)
	p.DB.CreateIndex("taskid", "task:*:", buntdb.IndexUint)
}

// StoreTask stores a Task in the Datastore
func (p *Persistance) StoreTask(alias string, fileID *string, delay int) error {
	var baseKey = "task:" + alias + "-" + string(UID())

	if fileID != nil {
		var fileIdKey = baseKey + TaskFileId
		p.store(fileIdKey, string(UID()))
	}

	var delayKey = baseKey + TaskDelay
	return p.store(delayKey, string(delay))
}

// StoreTaskMaster persists the Master service ID inside taskId properties
func (p *Persistance) StoreTaskMaster(taskId, serviceID string) error {
	var baseKey = "task:" + taskId

	var master = baseKey + TaskMaster
	err := p.store(master, serviceID)
	return err
}

// StoreTaskWorker persists the Master service ID inside taskId properties
func (p *Persistance) StoreTaskWorker(taskId, serviceID string) error {
	var baseKey = "task:" + taskId

	var worker = baseKey + TaskWorker
	err := p.store(worker, serviceID)
	return err
}
