package utils

import (
	"aws-sns-sesevent-processor/model"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
)

func parseRecord(record events.SNSEntity) (sesEvent model.SesEvent, err error) {
	err = json.Unmarshal([]byte(record.Message), &sesEvent)
	return
}

func ParseEvent(events events.SNSEvent) (sesEvent model.SesEvent, err error) {
	if len(events.Records) > 0 {
		return parseRecord(events.Records[0].SNS)
	}
	return
}
