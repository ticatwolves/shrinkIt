package handlers

import (
	"encoding/json"
	"os"
	"shrinkIt/internal/db"
	schema "shrinkIt/internal/schemas"
	"shrinkIt/lib"

	"github.com/aws/aws-lambda-go/events"
)

func handlePost(request *events.APIGatewayV2HTTPRequest) (schema.ResponsePayload, error) {
	var requestPayload schema.RequestPayload
	err := lib.RequestParser([]byte(request.Body), &requestPayload)
	if err != nil {
		return schema.ResponsePayload{Status: 400, Message: string(err.Error())}, err
	}
	hashBytes := lib.GenerateHash(requestPayload.Url)
	var urlHash string = string(hashBytes)[:5]
	db.Insert(&urlHash, &requestPayload.Url)
	hashUrl := os.Getenv("API_ENDPOINT") + "?urlHash=" + urlHash
	return schema.ResponsePayload{Status: 200, Message: hashUrl}, nil
}

func handleGet(request *events.APIGatewayV2HTTPRequest, headers map[string]string) (schema.ResponsePayload, error) {
	hash := request.QueryStringParameters["urlHash"]
	if hash == "" {
		return schema.ResponsePayload{Status: 400, Message: "Invalid URL"}, nil
	}
	shrinkit := db.GetByHash(&hash)
	headers["Location"] = shrinkit.ActualUrl
	return schema.ResponsePayload{Status: 302, Message: ""}, nil
}

func ShrinkItApiRequestHandler(request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	DefaultHeaders := make(map[string]string)
	DefaultHeaders["Content-Type"] = "application/json"
	var response schema.ResponsePayload
	switch request.RequestContext.HTTP.Method {
	case "POST":
		response, _ = handlePost(&request)
	case "GET":
		response, _ = handleGet(&request, DefaultHeaders)
	}
	body, _ := json.Marshal(response)
	return events.APIGatewayV2HTTPResponse{
		StatusCode: response.Status,
		Body:       string(body),
		Headers:    DefaultHeaders,
	}, nil
}
