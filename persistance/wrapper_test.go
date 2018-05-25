package persistance

import "testing"

func TestDeleteNotExists(t *testing.T) {
	p, err := InitPersistance(TestingDBPath)
	if err != nil {
		t.Error(err)
	}

	err = p.delete("NOTEXISTS")

	if err == nil {
		t.Error("ERROR, this deletion shouldn't be succesful")
	}
}

func TestStoreDuplicatedKey(t *testing.T) {
	p, err := InitPersistance(TestingDBPath)
	if err != nil {
		t.Error(err)
	}

	p.store("DUP", "DUP")

	err = p.store("DUP", "DUP")

	if err == nil {
		t.Error("ERROR, this store op. shoudn't be succesful, there's a duplicate in the datastore")
	}
}
