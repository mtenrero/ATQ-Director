package main

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"

	"strconv"

	"github.com/goadesign/goa"
	"github.com/mholt/archiver"
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
	os.MkdirAll(c.Persistance.GlusterPath+"/files", 0755)

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

		fileName := part.FileName()

		// Ensure file extension
		isNotZip := ensureZip(fileName)
		if isNotZip != nil {
			errr := isNotZip.Error()
			atqNotZip := app.AtqDatabindUploadError{
				Error: &errr,
			}

			return ctx.TheFileDoesnTHaveAnAcceptedCompressionError(&atqNotZip)
		}

		// Open file for later usage
		file, fileErr := os.OpenFile(c.Persistance.GlusterPath+"/files/"+fileName, os.O_WRONLY|os.O_CREATE, 0666)
		if fileErr != nil {
			errr := fileErr.Error()
			atqUploadError := app.AtqDatabindUploadError{
				Error: &errr,
			}
			return ctx.UploadErrorError(&atqUploadError)
		}

		// Ensure file is closed if error occurred
		defer file.Close()

		// Copy mulitpart readed file content into FileSystem
		io.Copy(file, part)

		file.Close()

		timestampUID := time.Now().Unix()
		timestampUIDString := strconv.Itoa(int(timestampUID))

		// Unzip File
		directory, err := unzip(fileName, timestampUIDString)
		if err != nil {
			errr := err.Error()
			atqError := app.AtqDatabindUploadError{
				Error: &errr,
			}

			ctx.UploadErrorError(&atqError)
		}

		// Delete Original file
		deleteFile(fileName)

		// Save File to datastore
		fullPath, _ := filepath.Abs(c.Persistance.GlusterPath + "/files/" + directory)
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

func ensureZip(fileName string) error {

	var extension = filepath.Ext(fileName)

	if extension != ".zip" {
		return errors.New("The uploaded file is not a ZIP file")
	}

	return nil
}

func unzip(fileName, fileID string) (string, error) {
	input := "./files/" + fileName
	output := "./files/" + fileID
	err := archiver.Zip.Open(input, output)

	return output, err
}

func deleteFile(fileName string) error {
	file := "./files/" + fileName
	err := os.Remove(file)
	return err
}
