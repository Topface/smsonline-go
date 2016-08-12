package smsonline

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	// baseURL - default smsonline endpoint
	baseURL = "https://bulk.sms-online.com/"
)

// SmsOnline is sms online client
type SmsOnline struct {
	username   string
	secret     string
	charset    string
	baseURL    string
	httpClient *http.Client
}

// NewSmsOnlineClient initializes new SmsOnline client
func NewSmsOnlineClient(username, secret, charset string, httpClient *http.Client) *SmsOnline {
	if charset == "" {
		charset = defaultCharset
	}

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &SmsOnline{
		username:   username,
		secret:     secret,
		charset:    charset,
		baseURL:    baseURL,
		httpClient: httpClient,
	}
}

// NewSmsOnlineClientCustom initializes new SmsOnline client with custom BaseUrl
func NewSmsOnlineClientCustom(username, secret, charset, url string, httpClient *http.Client) *SmsOnline {
	client := NewSmsOnlineClient(username, secret, charset, httpClient)
	client.baseURL = url
	return client
}

// SendSimpleSms send simple sms
func (c *SmsOnline) SendSimpleSms(from, to, text, charset string) (*SmsResponse, error) {
	return c.SendSms(from, to, text, charset, 0, false)
}

// SendSms send sms with some additional options such as delay, ack, binary
func (c *SmsOnline) SendSms(from, to, text, charset string, delay int, ack bool) (*SmsResponse, error) {
	message := makeSms(from, text, to)
	message.setAck(ack)
	message.setDelay(delay)

	if charset == "" {
		charset = c.charset
	}
	message.setCharset(charset)

	response, err := c.send(message)
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
func (c *SmsOnline) send(m message) (*http.Response, error) {
	messageData := m.getMessageData(c.username, c.secret).Encode()
	req, err := http.NewRequest("POST", baseURL, bytes.NewBufferString(messageData))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(messageData)))

	return c.httpClient.Do(req)
}
