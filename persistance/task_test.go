package persistance

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStoreTask(t *testing.T) {
	p, err := InitPersistance(TestingDBPath)
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
	p, err := InitPersistance(TestingDBPath)
	if err != nil {
		t.Error(err)
	}

	tmp := "/tmp"
	err = p.StoreTask("TASK", &tmp, 3)
	if err != nil {
		t.Error(err)
	}

	err = p.StoreTaskMaster("TASK", "MASTER")
	if err != nil {
		t.Error(err)
	}

	err = p.StoreTaskWorker("TASK", "WORKER")
	if err != nil {
		t.Error(err)
	}

	task, err := p.ReadTask("TASK")
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 3, *task.Delay)
	assert.Equal(t, "MASTER", task.Master.Alias)
	assert.Equal(t, "WORKER", task.Worker.Alias)
	assert.Equal(t, "/tmp", *task.Master.Fileid)
	assert.Equal(t, "/tmp", *task.Worker.Fileid)

	defer teardown(p, true)
}

func teardown(p *Persistance, close bool) {
	p.delete("task:TASK:fileID")
	p.delete("task:TASK:delay")
	p.delete("task:TASK:master")
	p.delete("task:TASK:worker")
	p.delete("task:TASK:status")

	if close {
		p.ClosePersistance()
	}
}
