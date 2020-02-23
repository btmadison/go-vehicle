package dynamo

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"

	"github.com/btmadison/btmadison/go-vehicle/pkg/crud"
)

type Repository struct {
	tableName string
	metaKey   string
}

func NewRepository() *Repository {
	store := new(Repository)
	store.tableName = "govehicles"
	store.metaKey = "metadata"
	return store
}

// GetAll gets all vehicles
func (m *Repository) GetAllVehicles() []crud.Vehicle {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := dynamodb.New(sess)
	filt := expression.Name("sk").Equal(expression.Value("metadata"))

	expr, err := expression.NewBuilder().WithFilter(filt).Build()
	if err != nil {
		fmt.Println("Got error building expression:")
		fmt.Println(err.Error())
		os.Exit(1)
	}
	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		TableName:                 aws.String(m.tableName),
	}

	result, err := svc.Scan(params)
	fmt.Println(result)

	vehicles := []crud.Vehicle{}

	for _, item := range result.Items {
		v := Vehicle{}
		err = dynamodbattribute.UnmarshalMap(item, &v)
		fmt.Println(v)
		if err != nil {
			fmt.Println(err)
		} else if v.sk == "metadata" {
			v := crud.Vehicle{
				Vin:        v.pk,
				Make:       v.make,
				Model:      v.model,
				Year:       v.year,
				Dealership: v.dealership,
			}

			vehicles = append(vehicles, v)
			fmt.Println("appended")
		} else {
			fmt.Println(v)
		}
	}
	fmt.Println(len(vehicles))
	return vehicles
}

// GetOneByID returns vehicle with given VIN number
func (m *Repository) GetOneByID(vin string) (crud.Vehicle, error) {
	return crud.Vehicle{
		Vin:        "5e4ef1ad40e8ab28f6e75138",
		Make:       "Acura",
		Model:      "Farger",
		Year:       1969,
		Dealership: "Rent-A-Wreck",
	}, nil
}

// Upsert will Insert or Update existing Vehicle based on globally unique VIN#
func (m *Repository) Upsert(v crud.Vehicle) {
}

// Delete vehicle from in memory inventory
func (m *Repository) Delete(vin string) {
}
