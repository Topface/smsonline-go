package smsonline

import (
	"testing"
)

func createClient() *SmsOnline {
	return NewSmsOnlineClientCustom("test", "secret", "", baseURL)
}

func TestSmsOnline_SendSms(t *testing.T) {
	client := createClient()
	res, err := client.SendSms("test", "89119876543", "test", "", 0, false, false)

	if err != nil {
		t.Fatal(err)
	}

	if res.Code != CodeSyntaxDataError {
		t.Fatal(res.Message)
	}
}

func TestSmsOnline_SendSms2(t *testing.T) {
	client := createClient()
	res, err := client.SendSms("test", "89119876543", "test", "UCS2", 1000, true, true)

	if err != nil {
		t.Fatal(err)
	}

	if res.Code != CodeSyntaxDataError {
		t.Fatal(res.Message)
	}
}

func TestSmsOnline_SendSimpleSms(t *testing.T) {
	client := createClient()
	res, err := client.SendSimpleSms("test", "89119876543", "test", "")

	if err != nil {
		t.Fatal(err)
	}

	if res.Code != CodeSyntaxDataError {
		t.Fatal(res.Message)
	}
}
