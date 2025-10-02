package main

import (
	"github.com/ticatwolves/shrinkit/internal/handlers"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	// Start the Lambda function
	lambda.Start(handlers.ShrinkItApiRequestHandler)
}
