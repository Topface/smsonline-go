package smsonline

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"net/url"
	"strconv"
)

const (
	// type of report settings
	typeReport   = 0
	typeNoReport = 1

	// max message delay
	maxDelay       = 86400
	defaultCharset = "UTF-8"
)

type message struct {
	from       string
	to         string
	text       string
	delay      int
	reportType int
	charset    string
}

// make sms message
func makeSms(from, text string, to string) message {
	return message{
		from:       from,
		text:       text,
		to:         to,
		reportType: typeNoReport,
	}
}

// set charset
func (m *message) setCharset(charset string) {
	if charset != "" {
		m.charset = charset
	}
}

// set delay
func (m *message) setDelay(delay int) {
	if delay >= 0 && delay <= maxDelay {
		m.delay = delay
	} else {
		m.delay = maxDelay
	}
}

// set acknowledgment
func (m *message) setAck(ack bool) {
	if ack {
		m.reportType = typeReport
	} else {
		m.reportType = typeNoReport
	}
}

// make secret sign
func (m *message) getSign(user, secret string) string {
	var signBuffer bytes.Buffer
	signBuffer.WriteString(user)
	signBuffer.WriteString(m.from)
	signBuffer.WriteString(m.to)
	signBuffer.WriteString(m.text)
	signBuffer.WriteString(secret)

	return fmt.Sprintf("%x", md5.Sum(signBuffer.Bytes()))
}

// get message data
func (m *message) getMessageData(user, secret string) url.Values {
	formValues := url.Values{}

	formValues.Set("charset", m.charset)
	formValues.Set("delay", strconv.Itoa(m.delay))
	formValues.Set("txt", m.text)
	formValues.Set("dlr", strconv.Itoa(m.reportType))
	formValues.Set("phone", m.to)
	formValues.Set("user", user)
	formValues.Set("from", m.from)
	formValues.Set("sign", m.getSign(user, secret))

	return formValues
}
