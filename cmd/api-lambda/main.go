package main

import (
	"shrinkIt/internal/api"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	// Start the Lambda function
	lambda.Start(api.ShrinkItApiRequestHandler)
}
