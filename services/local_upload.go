package services

import (
	"ametory-crud/config"
	"ametory-crud/utils"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type Uploader interface {
	UploadFile(c *gin.Context) (string, error)
}

func UploadFile(c *gin.Context) (string, error) {
	uploader := getUploader()
	return uploader.UploadFile(c)
}

func getUploader() Uploader {
	switch config.App.Server.StorageProvider {

	case "google":
		return &googleUploader{}
	case "s3":
		return &s3Uploader{}
	default:
		return &localUploader{}
	}
}

type localUploader struct{}

func (localUploader) UploadFile(c *gin.Context) (string, error) {
	file, err := c.FormFile("file")
	if err != nil {
		return err.Error(), err
	}

	flipped := c.PostForm("flipped")

	assetsFolder := "./assets/images/"
	if _, err := os.Stat(assetsFolder); os.IsNotExist(err) {
		os.Mkdir(assetsFolder, os.ModePerm)
	}

	// Generate a unique filename
	filename := fmt.Sprintf("%v%s", utils.GetCurrentTimestamp(), path.Ext(file.Filename))

	// Save the uploaded file to the assets folder
	destination := filepath.Join(assetsFolder, filename)
	if err := c.SaveUploadedFile(file, destination); err != nil {
		return err.Error(), err
	}

	if flipped == "1" {

		srcFile, err := os.Open(destination)
		if err != nil {
			return err.Error(), err
		}
		defer srcFile.Close()

		// Detect the file type
		ext := filepath.Ext(file.Filename)
		var img image.Image

		switch ext {

		case ".png":
			img, err = png.Decode(srcFile)
			if err != nil {
				return err.Error(), err
			}
		default:
			img, err = jpeg.Decode(srcFile)
			if err != nil {
				return err.Error(), err
			}
		}

		flippedImg := flipImageHorizontally(img)

		newDest := filepath.Join(assetsFolder, "flipped_"+filename)

		// Save the flipped image
		outFile, err := os.Create(newDest)
		if err != nil {
			return err.Error(), err
		}
		defer outFile.Close()

		switch ext {
		case ".png":
			err = png.Encode(outFile, flippedImg)
		default:
			err = jpeg.Encode(outFile, flippedImg, nil)
		}
		if err != nil {
			return err.Error(), err
		}
		destination = newDest
	}

	return destination, nil
}

func flipImageHorizontally(img image.Image) image.Image {
	bounds := img.Bounds()
	flipped := image.NewRGBA(bounds)
	for x := 0; x < bounds.Dx(); x++ {
		for y := 0; y < bounds.Dy(); y++ {
			flipped.Set(bounds.Dx()-x-1, y, img.At(x, y))
		}
	}
	return flipped
}
