package api

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

// HandleRequest is the business logic for processing the HTTP request
func ShrinkItApiRequestHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Request received:", request)

	// Custom business logic here

	// Simple response
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Request processed successfully",
		Headers:    map[string]string{"Content-Type": "text/plain"},
	}, nil
}
