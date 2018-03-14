package functions

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func NewRequest(body []byte, uri, apiKey, method string) ([]byte, error) {

	if method == "POST" || method == "GET" {
	} else {
		return nil, errors.New("Invalid http method")
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, uri, bytes.NewBuffer(body))
	req.Header.Add("Authorization", "Basic "+apiKey)
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
