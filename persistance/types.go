package persistance

import (
	"github.com/tidwall/buntdb"
)

// Persistance contains all relevant data to the Framework datastore
type Persistance struct {
	DB *buntdb.DB
}
