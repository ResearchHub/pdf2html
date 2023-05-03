package internal

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func getS3Client() (*s3.S3, error) {
	config := aws.NewConfig().
		WithRegion(GetAwsRegion()).
		WithEndpoint(GetAwsS3Endpoint()).
		WithS3ForcePathStyle(true)

	sess, err := session.NewSession()
	if err != nil {
		return nil, err
	}

	return s3.New(sess, config), nil
}
