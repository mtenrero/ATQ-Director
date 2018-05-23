package persistance

import (
	"log"

	"github.com/tidwall/buntdb"
)

// InitPersistance initializes and loads the K/V datastore
func InitPersistance() (*Persistance, error) {
	db, err := buntdb.Open("./storage/atq.db")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	persistance := Persistance{
		DB: db,
	}

	return &persistance, nil
}

// ClosePersistance closes the Persistance Datastore
func (persistance *Persistance) ClosePersistance() {
	persistance.DB.Close()
}
