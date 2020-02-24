package dynamo

import (
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"github.com/btmadison/btmadison/go-vehicle/pkg/crud"
)

// Repository for dynamodb govehicle connection
type Repository struct {
	tableName string
	metaKey   string
	region    string
}

// NewRepository initializes a new dynamodb govehicle repository
func NewRepository() *Repository {
	store := new(Repository)
	store.tableName = "govehicles"
	store.metaKey = "metadata"
	store.region = os.Getenv("region")
	return store
}

// GetAllVehicles gets all vehicles
func (m *Repository) GetAllVehicles() []crud.Vehicle {
	dyn := newDynSession(m)

	params := &dynamodb.ScanInput{
		TableName: aws.String(m.tableName),
	}

	result, _ := dyn.Scan(params)

	return mapResultsToVehicles(result.Items)
}

// GetOneByID returns vehicle with given VIN number
func (m *Repository) GetOneByID(vin string) (crud.Vehicle, error) {
	dyn := newDynSession(m)

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
		return crud.Vehicle{}, fmt.Errorf("Failed to unmarshal Record, %v", err)
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

func newDynSession(m *Repository) *dynamodb.DynamoDB {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region: aws.String(m.region),
		},
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
