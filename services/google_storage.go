package services

import (
	"ametory-crud/config"
	"bytes"
	"context"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"path/filepath"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

var FirebaseAPP = &firebase.App{}

func InitFirebaseApp() (*firebase.App, error) {
	opt := option.WithCredentialsFile(config.App.Google.FirebaseCredentialFile)

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}

	FirebaseAPP = app

	return app, nil
}

type googleUploader struct{}

func (googleUploader) UploadFile(c *gin.Context) (string, error) {
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
	if config.App.Google.FirebaseFolderFile != "" {
		folder = config.App.Google.FirebaseFolderFile
	}
	return uploadFileToFirebaseStorage(fileBytes, folder, file.Filename)
}

func uploadFileToFirebaseStorage(file []byte, folder string, fileName string) (string, error) {
	// Use the initialization token we have to create a authenticated client
	client, err := FirebaseAPP.Storage(context.Background())
	if err != nil {
		return "", fmt.Errorf("error getting Storage client: %v", err)
	}

	var objString = folder + "/" + fileName
	fmt.Println("BUCKET", config.App.Google.FirebaseStorageBucket)
	fmt.Println("objString", objString)
	bucket, err := client.Bucket(config.App.Google.FirebaseStorageBucket)
	if err != nil {
		return "", fmt.Errorf("error getting bucket: %v", err)
	}

	wc := bucket.Object(objString).NewWriter(context.Background())
	if _, err = wc.Write(file); err != nil {
		return "", fmt.Errorf("error writing object to bucket: %v", err)
	}

	if err := wc.Close(); err != nil {
		return "", fmt.Errorf("error closing writer: %v", err)
	}

	fmt.Println("ERROR", makePublic(objString))

	publicURL := fmt.Sprintf("https://storage.googleapis.com/%s/%s", config.App.Google.FirebaseStorageBucket, objString)
	fmt.Printf("Public URL: %s\n", publicURL)

	return objString, nil

}

func makePublic(object string) error {

	client, err := FirebaseAPP.Storage(context.Background())
	if err != nil {
		return fmt.Errorf("storage.NewClient: %w", err)
	}

	bucket, err := client.Bucket(config.App.Google.FirebaseStorageBucket)
	if err != nil {
		return fmt.Errorf("error getting bucket: %v", err)
	}

	acl := bucket.Object(object).ACL()
	if err := acl.Set(context.Background(), storage.AllUsers, storage.RoleReader); err != nil {
		return fmt.Errorf("ACLHandle.Set: %w", err)
	}
	return nil
}
