package persistance

import "github.com/tidwall/buntdb"

// Filekey defines the key in the datastore for files
const Filekey = "file:"

func (p *Persistance) indexFile() {
	p.DB.CreateIndex("file", "file:*", buntdb.IndexInt)
}

// StoreFile stores a file and its path in the Datastore
func (p *Persistance) StoreFile(fileID, path string) error {
	err := p.store(Filekey+fileID, path)
	if err != nil {
		return err
	}

	return nil
}

// ReadFilePath reads the Path of a given file in the datastore
func (p *Persistance) ReadFilePath(fileID string) (string, error) {
	path, err := p.read(Filekey + fileID)
	if err != nil {
		return "", err
	}

	return path, nil
}

func (p *Persistance) ReadAllFiles() (*map[string]string, error) {
	collection, err := p.iterateStringString("file")

	return collection, err
}
