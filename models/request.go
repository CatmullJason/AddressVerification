package models

import (
	"encoding/json"
)

type Request struct {
	Body    []byte
	Method  string
	Uri     string
	Headers []Header
}

type Header struct {
	Key   string
	Value string
}

func (req *Request) AddHeader(key, value string) {
	header := Header{key, value}
	req.Headers = append(req.Headers, header)
}

func (req *Request) AddBody(body interface{}) {
	b, err := json.Marshal(body)
	if err == nil {
		req.Body = b
	}
}

func (req *Request) SetMethod(method string) {
	req.Method = method
}

func (req *Request) SetUri(uri string) {
	req.Uri = uri
}
