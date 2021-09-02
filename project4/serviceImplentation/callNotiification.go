package serviceimplentation

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type CallNotification struct {
}

func (call *CallNotification) Notify() {
	// Let's set some initial default variables
	accountSid := "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	authToken := "YYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYY"
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Calls.json"

	// Build out the data for our message
	v := url.Values{}
	v.Set("To", "+155555555555")
	v.Set("From", "+15555551234")
	v.Set("Url", "[ngrokurl]/twiml")
	rb := *strings.NewReader(v.Encode())

	// Create Client
	client := &http.Client{}

	req, _ := http.NewRequest("POST", urlStr, &rb)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// make request
	resp, _ := client.Do(req)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		err := json.Unmarshal(bodyBytes, &data)
		if err == nil {
			fmt.Println(data["sid"])
		}
	} else {
		fmt.Println(resp.Status)
		// w.Write([]byte("Hello World!"))
	}

}
