package smsonline // import "github.com/xjewer/smsonline"

import (
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	baseURL = "https://bulk.sms-online.com"
)

// SmsOnline is sms online client
type SmsOnline struct {
	username string
	secret   string
	charset  string
	baseURL  string
}

// NewSmsOnlineClient initializes new SmsOnline client
func NewSmsOnlineClient(username, secret, charset string) *SmsOnline {
	if charset == "" {
		charset = defaultCharset
	}
	return &SmsOnline{
		username: username,
		secret:   secret,
		charset:  charset,
		baseURL:  baseURL,
	}
}

// NewSmsOnlineClientCustom initializes new SmsOnline client with custom BaseUrl
func NewSmsOnlineClientCustom(username, secret, charset, url string) *SmsOnline {
	client := NewSmsOnlineClient(username, secret, charset)
	client.baseURL = url
	return client
}

// SendSimpleSms send simple sms
func (client *SmsOnline) SendSimpleSms(from, to, text, charset string) (*SmsResponse, error) {
	return client.SendSms(from, to, text, charset, 0, false, false)
}

// SendSms send sms with some additional options such as delay, ack, binary
func (client *SmsOnline) SendSms(from, to, text, charset string, delay int, ack, binary bool) (*SmsResponse, error) {
	message := makeSms(from, text, to)
	message.setAck(ack)
	message.setDelay(delay)
	message.setBinaryType(binary)

	if charset == "" {
		charset = client.charset
	}
	message.setCharset(charset)

	response, err := client.send(message)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	smsResponse, err := makeResponse(responseBody)
	return smsResponse, err
}

// send data to sms online
func (client *SmsOnline) send(m message) (*http.Response, error) {
	messageData := m.getMessageData(client.username, client.secret).Encode()
	req, err := http.NewRequest("POST", client.baseURL, strings.NewReader(messageData))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/text")

	return http.DefaultClient.Do(req)
}
