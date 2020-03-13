package tests

import (
	"aws-sns-sesevent-processor/model"
	"aws-sns-sesevent-processor/utils"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"testing"
)

func Test_ParseEvent(t *testing.T) {
	t.Run("ParseEvent_OK", func(t *testing.T) {
		content, err := ioutil.ReadFile("sendevent.json")
		if err != nil {
			log.Fatal(err)
		}
		var snsEvent events.SNSEvent
		err = json.Unmarshal(content, &snsEvent)
		var sesEvent model.SesEvent
		sesEvent, err = utils.ParseEvent(snsEvent)
		if err != nil {
			log.Fatal(err)
		}
		assert.Equal(t, model.Send.Name(), sesEvent.EventType)
	})
}
