package smsonline

import (
	"net/http"
	"testing"
)

func createClient() *SmsOnline {
	return NewSmsOnlineClientCustom("test", "secret", "", baseURL, http.DefaultClient)
}

func TestSmsOnline_SendSms(t *testing.T) {
	client := createClient()
	res, err := client.SendSms("test", "89119876543", "test", "", 0, false)

	if err != nil {
		t.Fatal(err)
	}

	// auth error
	if res.Code != CodeAuthenticationError {
		t.Fatal(res.Message)
	}
}

func TestSmsOnline_SendSms2(t *testing.T) {
	client := createClient()
	res, err := client.SendSms("test", "79119876543", "test", "UCS-2", 1000, true)

	if err != nil {
		t.Fatal(err)
	}

	// auth error
	if res.Code != CodeAuthenticationError {
		t.Fatal(res.Message)
	}
}

func TestSmsOnline_SendSimpleSms(t *testing.T) {
	client := createClient()
	res, err := client.SendSimpleSms("test", "89119876543", "test", "")

	if err != nil {
		t.Fatal(err)
	}

	// auth error
	if res.Code != CodeAuthenticationError {
		t.Fatal(res.Message)
	}
}
