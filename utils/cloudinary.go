package utils

import (
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api"
)

// CloudinaryUpload uploads file to Cloudinary
func CloudinaryUpload(file multipart.File, filename string) (string, error) {
	cloudinary.UploadSettings.FileName = filename
	cloudinary.UploadSettings.Folder = "mygram_photos/"

	resp, err := cloudinary.Upload.Upload(file, api.UploadParams{})
	if err != nil {
		return "", err
	}

	return resp.SecureURL, nil
}
