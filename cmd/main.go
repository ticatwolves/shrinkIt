package main

import (
	"shrinkIt/internal/handlers"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	// Start the Lambda function
	lambda.Start(handlers.ShrinkItApiRequestHandler)
}
