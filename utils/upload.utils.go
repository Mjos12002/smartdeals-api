package utils

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// A utility to handle file uploads
func HandleFileUpload(context *gin.Context, destination string) (string, error) {

	// Declare the multiform
	multiPForm, error := context.MultipartForm()

	if error != nil {
		return error.Error(), error
	}

	// Get the files uploaded
	files := multiPForm.File["logo_url"]
	fileURL := []string{}

	for _, file := range files {

		logoFileExt := filepath.Ext(file.Filename)
		newFilename := uuid.New().String() + logoFileExt

		logoFileDest := filepath.Join(fmt.Sprintf("./resources/%s", destination), newFilename)
		fmt.Printf("Destination - %s\n", logoFileDest)
		if err := context.SaveUploadedFile(file, logoFileDest); err != nil {
			return err.Error(), err
		}
		fileURL = append(fileURL, logoFileDest)

	}
	return strings.Join(fileURL, ", "), nil
}
