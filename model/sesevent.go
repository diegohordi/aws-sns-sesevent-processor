/**
* AWS reference
*
* https://docs.aws.amazon.com/ses/latest/DeveloperGuide/event-publishing-retrieving-sns-contents.html#event-publishing-retrieving-sns-contents-send-object
 */

package model

import "time"

type EventType int
type BounceType int

const (
	Delivery EventType = iota
	Send
	Reject
	Open
	Click
	Bounce
	Complaint
	Failure
)

func (eventType EventType) Name() string {
	return [...]string{"Delivery", "Send", "Reject", "Open", "Click", "Bounce", "Complaint", "Failure"}[eventType]
}

const (
	Undetermined BounceType = iota
	Permanent
	Transient
)

func (bounceType BounceType) Name() string {
	return [...]string{"Undetermined", "Permanent", "Transient"}[bounceType]
}

type Header struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Mail struct {
	Timestamp        string                 `json:"timestamp"`
	MessageId        string                 `json:"messageId"`
	Source           string                 `json:"source"`
	SourceArn        string                 `json:"sourceArn"`
	SendingAccountId string                 `json:"sendingAccountId"`
	Destination      []string               `json:"destination"`
	HeadersTruncated bool                   `json:"headersTruncated"`
	Headers          []Header               `json:"headers"`
	CommonHeaders    map[string]interface{} `json:"commonHeaders"`
	Tags             map[string][]string    `json:"tags"`
}

type SendEvent struct {
}

type BounceEvent struct {
	BounceType        string             `json:"bounceType"`
	BounceSubType     string             `json:"bounceSubType"`
	BouncedRecipients []RecipientAddress `json:"bouncedRecipients"`
	Timestamp         time.Time          `json:"timestamp"`
	FeedbackId        string             `json:"feedbackId"`
	ReportingMTA      string             `json:"reportingMTA"`
}

type ComplaintEvent struct {
	ComplainedRecipients  []RecipientAddress `json:"complainedRecipients"`
	Timestamp             time.Time          `json:"timestamp"`
	FeedbackId            string             `json:"feedbackId"`
	ComplaintSubType      string             `json:"complaintSubType,omitempty"`
	UserAgent             string             `json:"userAgent,omitempty"`
	ComplaintFeedbackType string             `json:"complaintFeedbackType,omitempty"`
	ArrivalDate           string             `json:"arrivalDate,omitempty"`
}

type DeliveryEvent struct {
	Timestamp            time.Time `json:"timestamp"`
	ProcessingTimeMillis int64     `json:"processingTimeMillis"`
	Recipients           []string  `json:"recipients"`
	SmtpResponse         string    `json:"smtpResponse"`
	ReportingMTA         string    `json:"reportingMTA"`
}

type RecipientAddress struct {
	EmailAddress   string `json:"emailAddress"`
	Action         string `json:"action,omitempty"`
	Status         string `json:"status,omitempty"`
	DiagnosticCode string `json:"diagnosticCode,omitempty"`
}

type RejectEvent struct {
	reason string `json:"reason"`
}

type OpenEvent struct {
	IpAddress string    `json:"ipAddress"`
	Timestamp time.Time `json:"timestamp"`
	UserAgent string    `json:"userAgent"`
}

type ClickEvent struct {
	IpAddress string              `json:"ipAddress"`
	Timestamp time.Time           `json:"timestamp"`
	UserAgent string              `json:"userAgent"`
	Link      string              `json:"link"`
	LinkTags  map[string][]string `json:"LinkTags"`
}

type FailureEvent struct {
	TemplateName string `json:"templateName"`
	ErrorMessage string `json:"errorMessage"`
}

type SesEvent struct {
	MessageId      string
	SnsPublishTime string
	EventType      string         `json:"eventType"`
	Mail           Mail           `json:"mail"`
	Bounce         BounceEvent    `json:"bounce,omitempty"`
	Complaint      ComplaintEvent `json:"complaint,omitempty"`
	Delivery       DeliveryEvent  `json:"delivery,omitempty"`
	Send           SendEvent      `json:"send,omitempty"`
	Reject         RejectEvent    `json:"reject,omitempty"`
	Open           OpenEvent      `json:"open,omitempty"`
	Click          ClickEvent     `json:"click,omitempty"`
	Failure        string         `json:"failure,omitempty"`
}
