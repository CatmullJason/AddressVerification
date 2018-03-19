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
	client    *http.Client
	transport *http.Transport = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		Dial: (&net.Dialer{
			Timeout:   0, // no timeout
			KeepAlive: 0,
		}).Dial,
		TLSHandshakeTimeout: 8 * time.Second,
	}
)

/**************************************************
* SendRequest : Receives http request method, uri
* headers, and body to make an API request and return
* the response.
***************************************************/
func SendRequest(request models.Request) ([]byte, error) {

	client = &http.Client{Transport: transport}
	req, err := http.NewRequest(request.Method, request.Uri, bytes.NewBuffer(request.Body))

	//Headers
	for _, header := range request.Headers {
		req.Header.Add(header.Key, header.Value)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
