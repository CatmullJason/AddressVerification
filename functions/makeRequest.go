package functions

import (
	"AddressVerification/models"
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

var (
	apiKey    string
	client    *http.Client
	transport = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		Dial: (&net.Dialer{
			Timeout:   0, // no timeout
			KeepAlive: 0,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}
)

/**************************************************
* SendRequest : Receives http request method, uri
* headers, and body to make API request and return
* the response.
***************************************************/
func SendRequest(request models.Request) (address []byte, err error) {

	client = &http.Client{Transport: transport}
	req, err := http.NewRequest(request.Method, request.Uri, bytes.NewBuffer(request.Body))

	//Headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", apiKey)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func SetAPIKey(key string) {
	apiKey = key
}
