package persistance

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStoreAndReadFile(t *testing.T) {
	p, err := InitPersistance(TestingDBPath, ".", "localhost")
	if err != nil {
		t.Error(err)
	}

	p.StoreFile("file", "filepath")

	filepath, err := p.ReadFilePath("file")

	assert.NoError(t, err, "Error reading FilePath, the error may be in storing file")
	assert.Equal(t, "filepath", filepath)
}

func TestReadFileNotPresent(t *testing.T) {
	p, err := InitPersistance(TestingDBPath, ".", "localhost")
	if err != nil {
		t.Error(err)
	}

	_, err = p.ReadFilePath("notPresentFile")

	assert.Error(t, err, "Error not fired readinf FileID not present")
}

func TestReadAllFiles(t *testing.T) {
	p, err := InitPersistance(TestingDBPath, ".", "localhost")
	if err != nil {
		t.Error(err)
	}

	p.StoreFile("file1", "filepath1")
	p.StoreFile("file2", "filepath2")

	files, err := p.ReadAllFiles()

	filesMap := *files

	assert.NoError(t, err, "Error fired reading All Files using index")

	assert.Equal(t, "filepath1", filesMap["file:file1"])
	assert.Equal(t, "filepath2", filesMap["file:file2"])
}
