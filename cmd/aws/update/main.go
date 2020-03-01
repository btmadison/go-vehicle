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

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (response, error) {
	repo := dynamo.NewRepository()
	svc := crud.NewService(repo)
	vin := request.PathParameters["vin"]
	body := []byte(request.Body)
	var v crud.Vehicle

	err := json.Unmarshal(body, &v)
	if err != nil {
		fmt.Println(err)
		return response{StatusCode: 404}, err
	}

	_, err = svc.Update(vin, v)
	if err != nil {
		fmt.Println(err)
		return response{StatusCode: 404}, err
	}

	return response{StatusCode: 200}, nil
}

func main() {
	lambda.Start(handler)
}
