package uploader

import (
	"context"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadImage(
	cloudName string,
	apiKey string,
	apiSecret string,
	filePath string,
) (string, error) {

	cld, err := cloudinary.NewFromParams(
		cloudName,
		apiKey,
		apiSecret,
	)

	if err != nil {
		return "", err
	}

	resp, err := cld.Upload.Upload(
		context.Background(),
		filePath,
		uploader.UploadParams{},
	)

	if err != nil {
		return "", err
	}

	return resp.SecureURL, nil
}