package persistance

import (
	"github.com/tidwall/buntdb"
)

// store Stores a String in the given Key
func (p *Persistance) store(key, value string) error {
	err := p.DB.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(key, value, nil)
		return err
	})

	return err
}

func (p *Persistance) read(key string) (string, error) {
	return "", nil
}
