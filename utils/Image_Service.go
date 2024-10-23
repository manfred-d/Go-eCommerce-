package utils

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

// SaveImage saves the uploaded image and returns its URL
func SaveImage(file multipart.File, header *multipart.FileHeader) (string, error) {
	// Check file type
	if !IsImageFile(header.Filename) {
		return "", errors.New("invalid file type")
	}
	// Generate a unique filename
	ext := filepath.Ext(header.Filename)
	newFilename := uuid.New().String() + ext

	// Ensure the upload directory exists
	uploadDir := "./uploads"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return "", err
	}

	// Create the file
	dst, err := os.Create(filepath.Join(uploadDir, newFilename))
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// Copy the uploaded file to the destination file
	if _, err = io.Copy(dst, file); err != nil {
		return "", err
	}

	// Return the relative URL of the image
	return "/uploads/" + newFilename, nil
}

// DeleteImage removes an image file
func DeleteImage(imageURL string) error {
	if imageURL == "" {
		return nil
	}
	filename := strings.TrimPrefix(imageURL, "/uploads/")
	err := os.Remove(filepath.Join("./uploads", filename))

	if err != nil && !os.IsNotExist(err) {
		return err
	}
	return nil
}

// IsImageFile checks if the file has an allowed image extension
func IsImageFile(filename string) bool {
	allowedExt := map[string]bool{
		".jpg": true, ".jpeg": true, ".png": true, ".gif": true,
	}
	ext := strings.ToLower(filepath.Ext(filename))
	return allowedExt[ext]
}
