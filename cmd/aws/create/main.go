package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/btmadison/go-vehicle/pkg/crud"
	"github.com/btmadison/go-vehicle/pkg/data/dynamo"
)

type response events.APIGatewayProxyResponse

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (response, error) {
	repo := dynamo.NewRepository()
	svc := crud.NewService(repo)

	body := []byte(request.Body)
	var v crud.Vehicle

	err := json.Unmarshal(body, &v)
	if err != nil {
		fmt.Println(err)
		return response{StatusCode: 404}, err
	}
	svc.Create(v)

	return response{StatusCode: 200}, nil
}

func main() {
	lambda.Start(handler)
}
