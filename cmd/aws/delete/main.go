package main

import (
	"context"
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

	vin := request.PathParameters["vin"]

	err := svc.Delete(vin)
	if err != nil {
		fmt.Println(err)
		return response{StatusCode: 404}, err
	}

	return response{StatusCode: 200}, nil
}

func main() {
	lambda.Start(handler)
}
