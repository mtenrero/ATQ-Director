package persistance

import (
	"log"
	"os"

	"github.com/tidwall/buntdb"
)

// InitPersistance initializes and loads the K/V datastore
func InitPersistance(path string, glusterPath string, discoveryHost string) (*Persistance, error) {
	// Attempt to create storage directory ignoring errors
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll("./storage", os.ModePerm)
	}

	db, err := buntdb.Open(glusterPath + path)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	persistance := Persistance{
		DB:            db,
		GlusterPath:   glusterPath,
		DiscoveryHost: discoveryHost,
	}

	// Initialize indexes
	persistance.indexFile()
	persistance.IndexTask()

	return &persistance, nil
}

// ClosePersistance closes the Persistance Datastore
func (persistance *Persistance) ClosePersistance() {
	persistance.DB.Close()
}
