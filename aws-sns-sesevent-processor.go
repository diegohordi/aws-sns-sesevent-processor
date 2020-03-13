package main

import (
	"aws-sns-sesevent-processor/services"
	"aws-sns-sesevent-processor/utils"
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"os"
)
import "github.com/aws/aws-lambda-go/events"

func HandleRequest(context context.Context, event events.SNSEvent) {
	sesEvent, err := utils.ParseEvent(event)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	err = services.Process(event.Records[0].SNS.Timestamp.String(), sesEvent)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}

func main() {
	lambda.Start(HandleRequest)
}
