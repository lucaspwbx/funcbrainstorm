package main

import (
	"fmt"
	"net/http"
)

var (
	logFunc = func(r *PushRequest) string {
		debug := fmt.Sprintf("Doing %s request to: %s\n", r.Method, r.Endpoint)
		debug += fmt.Sprintf("Params:")
		debug += "--------\n"
		for k, v := range r.Params {
			debug += fmt.Sprintf("Key: %s, Value: %s\n", k, v)
		}
		return debug
	}

	newRequestFunc = func(r *PushRequest) (*http.Request, error) {
		log := logFunc(r)
		fmt.Println(log)
		req, err := http.NewRequest(r.Method, r.Endpoint, nil)
		req.SetBasicAuth("user", "password")
		if err != nil {
			return nil, err
		}
		return req, nil
	}
)

type PushRequest struct {
	Method   string
	Endpoint string
	Params   Params
}

type Params map[string]string

func NewPushRequest(method, endpoint string, params Params) *PushRequest {
	return &PushRequest{Method: method, Endpoint: endpoint, Params: params}
}

func NewHttpRequest() (*http.Request, error) {
	return newRequestFunc(NewPushRequest("GET", "/contacts", nil))
}

func GetContacts(params Params) string {
	//return logFunc("GET", "/contacts", params)
	return "not implemented"
}

func CreateContact(params Params) string {
	//return logFunc("POST", "/contacts", params)
	return "not implemented"
}

func UpdateContact(params Params) string {
	//contact_iden := params["iden"]
	//endpoint := fmt.Sprintf("/contacts/%s", contact_iden)
	//return logFunc("POST", endpoint, params)
	return "not implemented"
}

func DeleteContact(params Params) string {
	//contact_iden := params["iden"]
	//endpoint := fmt.Sprintf("/contacts/%s", contact_iden)
	//return logFunc("DELETE", endpoint, params)
	return "not implemented"
}

func main() {
	req, _ := NewHttpRequest()
	fmt.Println(req)
}
