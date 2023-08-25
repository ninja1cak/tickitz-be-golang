package pkg

import (
	"context"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func Cloudinary(file interface{}) (string, error) {

	cd_url := os.Getenv("CLOUDINARY_URL")

	cld, err := cloudinary.NewFromURL(cd_url)
	if err != nil {
		return "", err
	}
	result, err := cld.Upload.Upload(context.Background(), file, uploader.UploadParams{})
	return result.URL, nil
}
