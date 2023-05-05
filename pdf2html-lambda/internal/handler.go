package internal

import (
	"context"
	"errors"
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Location struct {
	BucketName string `json:"bucket_name"`
	ObjectKey  string `json:"object_key"`
}

type Event struct {
	S3Input S3Location `json:"s3_input"`
	S3Ouput S3Location `json:"s3_output"`
}

func HandleRequest(ctx context.Context, event Event) error {
	if event.S3Input.BucketName == "" {
		return errors.New("input bucket name is required")
	}
	if event.S3Input.ObjectKey == "" {
		return errors.New("input object key is required")
	}
	if event.S3Ouput.BucketName == "" {
		return errors.New("output bucket name is required")
	}
	if event.S3Ouput.ObjectKey == "" {
		return errors.New("output object key is required")
	}

	s3Client, err := getS3Client()
	if err != nil {
		return err
	}

	res, err := s3Client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(event.S3Input.BucketName),
		Key:    aws.String(event.S3Input.ObjectKey),
	})
	if err != nil {
		return err
	}

	// ensure response gets closed
	defer res.Body.Close()

	// create a temp file to copy the uploaded file to
	tmpFile, err := os.CreateTemp("", "*.pdf")
	if err != nil {
		return err
	}

	// remove the temp file when we are done
	tmpFilePath := tmpFile.Name()
	defer os.Remove(tmpFilePath)

	// copy the file content to the temp file
	n, err := io.Copy(tmpFile, res.Body)
	if err != nil {
		return err
	}

	// flush the temp file to disk
	log.Printf("wrote %d bytes to file %s\n", n, tmpFilePath)
	err = tmpFile.Sync()
	if err != nil {
		return err
	}

	// output the html to a temp file too
	outFilePath := replaceExtension(tmpFilePath, ".html")
	defer os.Remove(outFilePath)

	// ensure we are dealing with PDF content
	file, err := os.Open(tmpFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	err = validateFileType(file, "application/pdf")
	if err != nil {
		return err
	}

	// generate the html file
	log.Println("converting:", tmpFilePath, outFilePath)
	err = pdf2html(tmpFilePath, outFilePath)
	if err != nil {
		return err
	}

	// open the generated file for reading
	htmlFile, err := os.Open(outFilePath)
	if err != nil {
		return err
	}

	// stream the generated output to s3
	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(event.S3Ouput.BucketName),
		Key:         aws.String(event.S3Ouput.ObjectKey),
		ContentType: aws.String("text/html"),
		Body:        htmlFile,
	})
	if err != nil {
		return err
	}

	return nil
}
