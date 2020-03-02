package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/btmadison/go-vehicle/pkg/crud"
	"github.com/btmadison/go-vehicle/pkg/data/dynamo"
)

type response events.APIGatewayProxyResponse

func Handler(ctx context.Context) (response, error) {
	repo := dynamo.NewRepository()
	svc := crud.NewService(repo)

	vehicles, err := svc.ReadAll()
	if err != nil {
		return response{StatusCode: 404}, err
	}

	vehiclesJSON, _ := json.Marshal(vehicles)
	return response{Body: string(vehiclesJSON), StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
