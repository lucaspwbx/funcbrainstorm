package main

import (
	"encoding/json"
	"errors"
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

	getResponseFunc = func(r *http.Request) (PushResponse, error) {
		response, err := makeRequest(r)
		if err != nil {
			return nil, err
		}
		return response, nil
	}
)

type PushRequest struct {
	Method   string
	Endpoint string
	Params   Params
}

type Params map[string]string

type PushResponse map[string]interface{}

func makeRequest(r *http.Request) (PushResponse, error) {
	httpClient := &http.Client{}
	resp, err := httpClient.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response PushResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func NewPushRequest(method, endpoint string, params Params) *PushRequest {
	return &PushRequest{Method: method, Endpoint: endpoint, Params: params}
}

func ParseResponse(req *http.Request) (map[string]interface{}, error) {
	resp, err := getResponseFunc(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func GetContacts() (*http.Request, error) {
	//return newRequestFunc(NewPushRequest("GET", "/contacts", nil))
	//return newRequestFunc(NewPushRequest("GET", "http://www.google.com", nil))
	return newRequestFunc(NewPushRequest("GET", "https://api.pushbullet.com/v2/contacts", nil))
}

func CreateContact(params Params) (*http.Request, error) {
	request := NewPushRequest("POST", "/contacts", params)
	return newRequestFunc(request)
}

func UpdateContact(params Params) (*http.Request, error) {
	endpoint, err := ParseParamsId(params)
	if err != nil {
		return nil, err
	}
	request := NewPushRequest("POST", endpoint, params)
	return newRequestFunc(request)
}

func DeleteContact(params Params) (*http.Request, error) {
	endpoint, err := ParseParamsId(params)
	if err != nil {
		return nil, err
	}
	request := NewPushRequest("DELETE", endpoint, params)
	return newRequestFunc(request)
}

func ParseParamsId(params Params) (string, error) {
	if id, ok := params["iden"]; ok {
		return fmt.Sprintf("/contacts/%s", id), nil
	}
	return "", errors.New("No id")
}

func main() {
	req, _ := GetContacts()
	fmt.Println(req)
	req2, _ := CreateContact(Params{"teste": "bla"})
	fmt.Println(req2)
	req3, _ := UpdateContact(Params{"iden": "123"})
	fmt.Println(req3)
	req4, _ := DeleteContact(Params{"iden": "123"})
	fmt.Println(req4)
	fmt.Println(req)
	bla, err := ParseResponse(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bla)
}
