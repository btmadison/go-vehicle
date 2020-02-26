package main

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/btmadison/btmadison/go-vehicle/pkg/crud"
	"github.com/btmadison/btmadison/go-vehicle/pkg/data/dynamo"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context) (Response, error) {
	repo := dynamo.NewRepository()
	svc := crud.NewService(repo)

	vehicles, err := svc.ReadAll()
	if err != nil {
		return Response{StatusCode: 404}, err
	}

	vehiclesJSON, _ := json.Marshal(vehicles)
	var buf bytes.Buffer

	body, err := json.Marshal(map[string]interface{}{
		"message": "Got Vehicles",
		"data":    vehiclesJSON,
	})
	if err != nil {
		return Response{StatusCode: 404}, err
	}
	json.HTMLEscape(&buf, body)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type":      "application/json",
			"X-BMad-Func-Reply": "list-handler",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
