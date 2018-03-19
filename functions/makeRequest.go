package functions

import (
	"AddressVerification/models"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func SendRequest(request models.Request) ([]byte, error) {
	fmt.Println(string(request.Body))
	client := &http.Client{}
	req, err := http.NewRequest(request.Method, request.Uri, bytes.NewBuffer(request.Body))

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
