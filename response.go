package smsonline

import (
	"encoding/xml"
)

const (
	// CodeOk successful request
	CodeOk = 0
	// CodeSyntaxDataError syntax data error
	CodeSyntaxDataError = -1
	// CodeAuthenticationError authentication error
	CodeAuthenticationError = -2
	// CodeReject reject
	CodeReject = -3
	// CodeSystemError system error
	CodeSystemError = -4
	// CodeLimitReached limit reached
	CodeLimitReached = -5
)

// SmsResponse is a sms online response data
type SmsResponse struct {
	Code       int               `xml:"code"`
	Message    string            `xml:"tech_message"`
	MessageIds map[string]string `xml:"msg_id"`
}

// make sms online response
func makeResponse(data []byte) (response *SmsResponse, err error) {
	response = new(SmsResponse)
	err = xml.Unmarshal(data, response)
	return
}
