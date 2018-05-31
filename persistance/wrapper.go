package persistance

import (
	"errors"

	"github.com/tidwall/buntdb"
)

// store Stores a String in the given Key
func (p *Persistance) store(key, value string) error {
	_, err := p.read(key)
	if err == nil {
		return errors.New("The given KEY already exists in the Datastore!")
	}

	err = p.DB.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(key, value, nil)
		return err
	})

	return err
}

// read Reads from datastore
func (p *Persistance) read(key string) (string, error) {
	var value string
	err := p.DB.View(func(tx *buntdb.Tx) error {
		val, err := tx.Get(key)
		if err != nil {
			return err
		}
		value = val
		return nil
	})

	if err != nil {
		return "", err
	}

	return value, nil
}

func (p *Persistance) delete(key string) error {
	_, err := p.read(key)
	if err != nil {
		return errors.New("The given KEY doesn't exists in the Datastore!")
	}

	err = p.DB.Update(func(tx *buntdb.Tx) error {
		_, err := tx.Delete(key)
		return err
	})

	return err
}

// iterate returns the contents of a previos defined index
func (p *Persistance) iterateStringString(index string) (*map[string]string, error) {
	var collection map[string]string
	collection = make(map[string]string)

	err := p.DB.View(func(tx *buntdb.Tx) error {
		tx.Ascend(index, func(dbkey, dbval string) bool {
			collection[dbkey] = dbval
			return true
		})
		return nil
	})

	return &collection, err
}
