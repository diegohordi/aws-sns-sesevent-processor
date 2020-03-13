package services

import (
	"aws-sns-sesevent-processor/model"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)
import "github.com/aws/aws-sdk-go/aws/session"
import "github.com/aws/aws-sdk-go/service/dynamodb"

const (
	tableName = "sesevent"
)

func registerEvent(event model.SesEvent) (err error) {
	var session = session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	var db = dynamodb.New(session)
	item, err := dynamodbattribute.MarshalMap(event)
	if err != nil {
		fmt.Println("Error marshalling event item.")
		return
	}
	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(tableName),
	}
	_, err = db.PutItem(input)
	if err != nil {
		fmt.Println("Error registering item into the DynamoDB.")
		return
	}
	fmt.Println("Event '" + event.MessageId + "' (" + event.SnsPublishTime + " registered successfully.")
	return
}

func Process(snsPublishTime string, event model.SesEvent) (err error) {
	event.MessageId = event.Mail.MessageId
	event.SnsPublishTime = snsPublishTime
	err = registerEvent(event)
	return
}
