package persistance

import (
	"log"
	"os"

	"github.com/tidwall/buntdb"
)

const dbPath = "/storage/atq.db"

// InitPersistance initializes and loads the K/V datastore
func InitPersistance(path string, glusterPath string) (*Persistance, error) {
	// Attempt to create storage directory ignoring errors
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll("/storage", os.ModePerm)
	}

	db, err := buntdb.Open(glusterPath + path)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	persistance := Persistance{
		DB:          db,
		GlusterPath: glusterPath,
	}

	// Initialize indexes
	persistance.indexFile()

	return &persistance, nil
}

// ClosePersistance closes the Persistance Datastore
func (persistance *Persistance) ClosePersistance() {
	persistance.DB.Close()
}
