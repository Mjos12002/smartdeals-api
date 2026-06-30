package utils

import (
	"mime/multipart"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// A utility to handle file uploads
func HandleFileUpload(logoFile multipart.File, logoFileHeader *multipart.FileHeader, context *gin.Context) error {

	defer logoFile.Close()

	logoFileExt := filepath.Ext(logoFileHeader.Filename)

	newFilename := uuid.New().String() + logoFileExt

	logoFileDest := filepath.Join("./resources/products/", newFilename)

	if err := context.SaveUploadedFile(logoFileHeader, logoFileDest); err != nil {
		return err
	}

	return nil
}
