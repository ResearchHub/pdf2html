package internal

import "os"

func GetAwsRegion() string {
	return os.Getenv("AWS_REGION")
}

func GetAwsS3Endpoint() string {
	return os.Getenv("AWS_S3_ENDPOINT")
}
