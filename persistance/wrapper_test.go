package persistance

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteNotExists(t *testing.T) {
	p, err := InitPersistance(TestingDBPath, ".")
	if err != nil {
		t.Error(err)
	}

	err = p.delete("NOTEXISTS")

	if err == nil {
		t.Error("ERROR, this deletion shouldn't be succesful")
	}
}

func TestStoreDuplicatedKey(t *testing.T) {
	p, err := InitPersistance(TestingDBPath, ".")
	if err != nil {
		t.Error(err)
	}

	p.store("DUP", "DUP")

	err = p.store("DUP", "DUP")

	if err == nil {
		t.Error("ERROR, this store op. shoudn't be succesful, there's a duplicate in the datastore")
	}
}

func TestIterateStringString(t *testing.T) {
	p, err := InitPersistance(TestingDBPath, ".")
	if err != nil {
		t.Error(err)
	}

	p.StoreTask("task1", nil, 3)
	p.StoreTask("task2", nil, 0)

	_, err = p.iterateStringString("taskalias")

	assert.NoError(t, err, "Errlr iterating index")
}
