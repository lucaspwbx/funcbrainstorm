package main

import (
	"fmt"
	"net/http"
)

var (
	logFunc = func(method, endpoint string, params Params) string {
		debug := fmt.Sprintf("Doing %s request to: %s\n", method, endpoint)
		debug += fmt.Sprintf("Params:")
		debug += "--------\n"
		for k, v := range params {
			debug += fmt.Sprintf("Key: %s, Value: %s\n", k, v)
		}
		return debug
	}

	newRequestFunc = func(method, endpoint string, params Params) (*http.Request, error) {
		log := logFunc(method, endpoint, params)
		fmt.Println(log)
		req, err := http.NewRequest(method, endpoint, nil)
		req.SetBasicAuth("user", "password")
		if err != nil {
			return nil, err
		}
		return req, nil
	}
)

type Params map[string]string

func NewRequest() (*http.Request, error) {
	return newRequestFunc("GET", "/contacts", nil)
}

func GetContacts(params Params) string {
	return logFunc("GET", "/contacts", params)
}

func CreateContact(params Params) string {
	return logFunc("POST", "/contacts", params)
}

func UpdateContact(params Params) string {
	contact_iden := params["iden"]
	endpoint := fmt.Sprintf("/contacts/%s", contact_iden)
	return logFunc("POST", endpoint, params)
}

func DeleteContact(params Params) string {
	contact_iden := params["iden"]
	endpoint := fmt.Sprintf("/contacts/%s", contact_iden)
	return logFunc("DELETE", endpoint, params)
}

func main() {
	//fmt.Println(GetContacts(Params{"title": "NotaUm"}))
	//fmt.Println(CreateContact(Params{"name": "Lucas", "email": "foo@bar.com"}))
	//fmt.Println(UpdateContact(Params{"iden": "1234", "name": "Lucas"}))
	//fmt.Println(DeleteContact(Params{"iden": "1234"}))
	req, _ := NewRequest()
	fmt.Println(req)
}
