package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/researchhub/pdf2html/pdf2html-lambda/internal"
)

func main() {
	lambda.Start(internal.HandleRequest)
}
