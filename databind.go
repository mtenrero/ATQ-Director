package main

import (
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"

	"strconv"

	"github.com/goadesign/goa"
	"github.com/mtenrero/ATQ-Director/app"
	"github.com/mtenrero/ATQ-Director/persistance"
)

// DatabindController implements the databind resource.
type DatabindController struct {
	*goa.Controller
	*sync.Mutex
	*persistance.Persistance
}

// NewDatabindController creates a databind controller.
func NewDatabindController(service *goa.Service, persistance *persistance.Persistance) *DatabindController {
	return &DatabindController{
		Controller:  service.NewController("DatabindController"),
		Mutex:       &sync.Mutex{},
		Persistance: persistance,
	}
}

// List runs the list action.
func (c *DatabindController) List(ctx *app.ListDatabindContext) error {

	p := c.Persistance

	filesCollecion, err := p.ReadAllFiles()
	if err != nil {

	}

	res := parseDatabind(filesCollecion)
	return ctx.OK(res)
}

// Upload runs the upload action.
func (c *DatabindController) Upload(ctx *app.UploadDatabindContext) error {

	reader, err := ctx.MultipartReader()

	// Create files directory if doesn't exists
	os.MkdirAll("./files", 0755)

	// Reply with error message if errored
	if err != nil {
		errr := err.Error()
		atqUploadError := app.AtqDatabindUploadError{
			Error: &errr,
		}
		return ctx.UploadErrorError(&atqUploadError)
	}

	// Reply with error if multipart load not detected
	if reader == nil {
		errr := "The payload must be a Multipart request"
		atqUploadError := app.AtqDatabindUploadError{
			Error: &errr,
		}
		return ctx.UploadErrorError(&atqUploadError)
	}

	// Read Multipart file
	for {
		part, err := reader.NextPart()

		// End reading if EOF
		if err == io.EOF {
			break
		}

		// Reply with error if error detected in current part
		if err != nil {
			errr := "Failed to load part: " + err.Error()
			atqUploadError := app.AtqDatabindUploadError{
				Error: &errr,
			}
			return ctx.UploadErrorError(&atqUploadError)
		}

		// Open file for later usage
		fileName := part.FileName()
		file, fileErr := os.OpenFile("./files/"+fileName, os.O_WRONLY|os.O_CREATE, 0666)
		if fileErr != nil {
			errr := fileErr.Error()
			atqUploadError := app.AtqDatabindUploadError{
				Error: &errr,
			}
			return ctx.UploadErrorError(&atqUploadError)
		}

		defer file.Close()

		// Copy mulitpart readed file content into FileSystem
		io.Copy(file, part)

		timestampUID := time.Now().Unix()
		timestampUIDString := strconv.Itoa(int(timestampUID))

		// Save File to datastore
		fullPath, _ := filepath.Abs("./files/" + fileName)
		saveFileToDatastore(timestampUIDString, fullPath, c.Persistance)

		atqUpload := app.AtqDatabindUpload{
			ID: &timestampUIDString,
		}
		return ctx.OK(&atqUpload)
	}

	// Put your logic here

	res := &app.AtqDatabindUpload{}
	return ctx.OK(res)
	// DatabindController_Upload: end_implement
}

func saveFileToDatastore(fileID, path string, p *persistance.Persistance) error {

	return p.StoreFile(fileID, path)
}

func parseDatabind(collection *map[string]string) app.AtqDatabindUploadCollection {
	typeCollection := make(app.AtqDatabindUploadCollection, 0)

	for key := range *collection {
		fileType := app.AtqDatabindUpload{
			ID: &key,
		}
		typeCollection = append(typeCollection, &fileType)
	}

	return typeCollection
}
