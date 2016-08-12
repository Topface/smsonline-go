## Go client for SMS Online provider

[![Build Status](https://travis-ci.org/Topface/smsonline-go.svg?branch=master)](https://travis-ci.org/Topface/smsonline-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/Topface/smsonline-go)](https://goreportcard.com/report/github.com/Topface/smsonline-go)

Library for sending sms through https://sms-online.com provider

[smsonline-bulk](http://smsonline-bulk.readthedocs.io/en/latest/bulk_sms.html)

##Usages

```go
func main() {
    client := smsonline.NewSmsOnlineClient("user", "secret", "UTF-8")
    response, err := client.SendSimpleSms("from", "to", "text", "charset")
    
    if err != nil {
        log.Println(err)
    }
    
    if response.Code != smsonline.CodeOk {
        log.Println(response.Message)
    }
}
```

* `charset` - charset, default 'UTF-8'
* `from` - sender name
* `to` - receiver phone number
* `text` - sms message text
* `delay` - message sending delay
* `ack` - acknowledgment: 0 - reporting, 1 - no report (default)

##Authors
 [xjewer](github.com/xjewer)

##Licence

[MIT](/LICENSE)
