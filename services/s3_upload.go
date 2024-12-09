package services

import (
	"ametory-crud/config"
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
)

type s3Uploader struct{}

func (s3Uploader) UploadFile(c *gin.Context) (string, error) {
	flipped := c.PostForm("flipped")
	file, err := c.FormFile("file")

	if err != nil {
		return "", err
	}
	fileContent, err := file.Open()
	if err != nil {
		return "", err
	}
	defer fileContent.Close()

	fileBytes, err := io.ReadAll(fileContent)
	if err != nil {
		return "", err
	}

	if flipped == "1" {
		img, _, err := image.Decode(bytes.NewReader(fileBytes))
		if err != nil {
			return "", fmt.Errorf("error decoding image: %v", err)
		}
		flippedImg := flipImageHorizontally(img)
		buf := new(bytes.Buffer)
		ext := filepath.Ext(file.Filename)
		if ext == ".png" {
			png.Encode(buf, flippedImg)
			fileBytes = buf.Bytes()
		} else {
			jpeg.Encode(buf, flippedImg, &jpeg.Options{Quality: 90})
			fileBytes = buf.Bytes()
		}
	}

	var folder string = "images"
	if config.App.S3.Folder != "" {
		folder = config.App.S3.Folder
	}

	var key string = folder + "/" + file.Filename

	return key, uploadFileToS3(fileBytes, key)
}

func uploadFileToS3(fileBytes []byte, key string) error {
	// Use the S3Configuration to get the AccessKey and SecretKey
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(config.App.S3.Region),
		Credentials: credentials.NewStaticCredentials(config.App.S3.AccessKey, config.App.S3.SecretKey, ""),
	})

	if err != nil {
		return fmt.Errorf("failed to create session: %v", err)
	}

	uploader := s3.New(sess)
	reader := bytes.NewReader(fileBytes)

	_, err = uploader.PutObject(&s3.PutObjectInput{
		Body: reader,
		Bucket: aws.String(
			config.App.S3.Bucket,
		),
		Key: aws.String(key),
		ACL: aws.String("public-read"),
	})
	if err != nil {
		return fmt.Errorf("failed to upload file: %v", err)
	}

	return nil
}
