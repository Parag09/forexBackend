package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// Declared a new DynamoDB instance
var db = dynamodb.New(session.New(), aws.NewConfig().WithRegion("ap-south-1"))

// function to get single item from dynamodb
func getItem(Curr string) (*ForexData, error) {
	// Prepare the input for the query.
	input := &dynamodb.GetItemInput{
		TableName: aws.String("Forextable"),
		Key: map[string]*dynamodb.AttributeValue{
			"Curr": {
				S: aws.String(Curr),
			},
		},
	}

	// Retrieving the item from DynamoDB. If no matching item is found, return nil.
	result, err := db.GetItem(input)
	if err != nil {
		return nil, err
	}
	if result.Item == nil {
		return nil, nil
	}

	// The result.Item object returned has the underlying type map[string]*AttributeValue.
	// UnmarshalMap helper to parse this straight into the fields of a struct. Note:
	forexitem := new(ForexData)
	err = dynamodbattribute.UnmarshalMap(result.Item, forexitem)
	if err != nil {
		return nil, err
	}

	return forexitem, nil
}

//function to get all the items from dynamodb
func getAllItem() ([]ForexData, error) {
	params := &dynamodb.ScanInput{
		TableName: aws.String("Forextable"),
	}
	result, err := db.Scan(params)
	if err != nil {
		return nil, err
	}
	if result.Items == nil {
		return nil, nil
	}

	forexItems := []ForexData{}
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &forexItems)
	if err != nil {
		fmt.Println(err)
	}

	return forexItems, nil

}
