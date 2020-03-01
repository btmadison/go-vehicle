package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/btmadison/btmadison/go-vehicle/pkg/crud"
	"github.com/btmadison/btmadison/go-vehicle/pkg/data/dynamo"
)

type response events.APIGatewayProxyResponse

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (response, error) {
	repo := dynamo.NewRepository()
	svc := crud.NewService(repo)

	vin := request.PathParameters["vin"]

	vehicle, err := svc.ReadOneByID(vin)
	if err != nil {
		fmt.Println(err)
		return response{StatusCode: 404}, err
	}

	vehicleJSON, _ := json.Marshal(vehicle)
	return response{Body: string(vehicleJSON), StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
