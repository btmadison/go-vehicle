package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type response events.APIGatewayProxyResponse

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (response, error) {
	// repo := dynamo.NewRepository()
	// svc := crud.NewService(repo)

	vehicle := request.Body
	fmt.Printf(vehicle)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return response{StatusCode: 404}, err
	// }

	return response{StatusCode: 200}, nil
}

func main() {
	lambda.Start(handler)
}
