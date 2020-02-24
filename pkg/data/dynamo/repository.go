package dynamo

import (
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

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
	dyn := newDynSession()

	params := &dynamodb.ScanInput{
		TableName: aws.String(m.tableName),
	}

	result, _ := dyn.Scan(params)

	return mapResultsToVehicles(result.Items)
}

func newDynSession() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	return dynamodb.New(sess)
}

func mapResultsToVehicles(items []map[string]*dynamodb.AttributeValue) []crud.Vehicle {
	vehicles := []crud.Vehicle{}
	for _, item := range items {
		v := Vehicle{}
		err := dynamodbattribute.UnmarshalMap(item, &v)
		if err != nil {
			fmt.Println(err)
		} else if v.Sk == "metadata" {
			cv := v.ToCrudVehicle()
			vehicles = append(vehicles, cv)
		}
	}
	return vehicles
}

// GetOneByID returns vehicle with given VIN number
func (m *Repository) GetOneByID(vin string) (crud.Vehicle, error) {
	dyn := newDynSession()

	result, err := dyn.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(m.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"pk": {
				S: aws.String(vin),
			},
			"sk": {
				S: aws.String(m.metaKey),
			},
		},
	})

	if err != nil {
		return crud.Vehicle{}, err
	}

	v := Vehicle{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &v)
	if err != nil {
		return crud.Vehicle{}, errors.New(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	if v.Pk == "" {
		return crud.Vehicle{}, errors.New("vin not found: " + vin)
	}

	return v.ToCrudVehicle(), nil
}

// Upsert will Insert or Update existing Vehicle based on globally unique VIN#
func (m *Repository) Upsert(v crud.Vehicle) {
}

// Delete vehicle from in memory inventory
func (m *Repository) Delete(vin string) {
}
