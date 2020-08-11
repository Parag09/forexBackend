package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// creating session with dynamoDB
var db = dynamodb.New(session.New(), aws.NewConfig().WithRegion("ap-south-1"))

// function to put item from forex api to dynamo db
func putItem(forexWrite ForexData) error {
	input := &dynamodb.PutItemInput{
		TableName: aws.String("Forextable"),
		Item: map[string]*dynamodb.AttributeValue{
			"Curr": {
				S: aws.String(forexWrite.Curr),
			},
			"Rate": {
				S: aws.String(forexWrite.Rate),
			},
			"Timestamp": {
				S: aws.String(forexWrite.Timestamp),
			},
		},
	}

	_, err := db.PutItem(input)
	return err
}
