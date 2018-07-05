package persistance

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStoreTask(t *testing.T) {
	p, err := InitPersistance(TestingDBPath, ".")
	if err != nil {
		t.Error(err)
	}

	tmp := "/tmp"
	err = p.StoreTask("TASK", &tmp, 6)
	if err != nil {
		t.Error(err)
	}

	value, err := p.read("task:" + "TASK" + TaskDelay)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "6", value, "Delay Mismatch!")

	value, err = p.read("task:" + "TASK" + TaskFileId)

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "/tmp", value, "FileID Mismatch!")

	defer teardown(p, true)
}

func TestWholeTask(t *testing.T) {
	p, err := InitPersistance(TestingDBPath, ".")
	if err != nil {
		t.Error(err)
	}

	tmp := "/tmp"
	err = p.StoreTask("TASK", &tmp, 3)
	if err != nil {
		t.Error(err)
	}

	err = p.StoreTaskMaster("TASK", "MASTER", "MASTERALIAS")
	if err != nil {
		t.Error(err)
	}

	err = p.StoreTaskWorker("TASK", "WORKER", "WORKERALIAS")
	if err != nil {
		t.Error(err)
	}

	task, err := p.ReadTask("TASK")
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 3, *task.Delay)
	assert.Equal(t, "MASTER", *task.Master.ID)
	assert.Equal(t, "WORKER", *task.Worker.ID)
	assert.Equal(t, "/tmp", *task.Master.FileID)
	assert.Equal(t, "/tmp", *task.Worker.FileID)
	assert.Equal(t, "MASTERALIAS", *task.Master.Alias)
	assert.Equal(t, "WORKERALIAS", *task.Worker.Alias)

	defer teardown(p, true)
}

func teardown(p *Persistance, close bool) {
	p.delete("task:TASK" + TaskFileId)
	p.delete("task:TASK" + TaskDelay)
	p.delete("task:TASK" + TaskMaster)
	p.delete("task:TASK" + TaskWorker)
	p.delete("task:TASK" + TaskStatus)
	p.delete("task:TASK" + TaskAliasMaster)
	p.delete("task:TASK" + TaskAliasWorker)

	if close {
		p.ClosePersistance()
	}
}
