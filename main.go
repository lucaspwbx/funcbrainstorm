package main

import (
	"bytes"
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
		//	req.SetBasicAuth("user", "password")
		req.SetBasicAuth("99el4gWxWuoh1of9vEGQvKZSHZwveuOh", "")
		if err != nil {
			return nil, err
		}
		return req, nil
	}

	createFunc = func(r *PushRequest) (*http.Request, error) {
		data, _ := json.Marshal(r.Params)
		req, err := http.NewRequest(r.Method, r.Endpoint, bytes.NewBuffer(data))
		//	req.SetBasicAuth("user", "password")
		if err != nil {
			return nil, err
		}
		req.SetBasicAuth("99el4gWxWuoh1of9vEGQvKZSHZwveuOh", "")
		req.Header.Set("Content-Type", "application/json")
		return req, nil
	}
)

type PushRequest struct {
	Method   string
	Endpoint string
	Params   Params
}

type Params map[string]string

func ExecuteRequest(r *http.Request) (*http.Response, error) {
	httpClient := &http.Client{}
	resp, err := httpClient.Do(r)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type Me struct {
	Iden     string `json:"iden"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

type Contact struct {
	Iden   string `json:"iden"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Active bool   `json:"active"`
}

type ContactsCollection struct {
	Contacts []Contact `json:contacts"`
}

func (c *ContactsCollection) Get() {
	fmt.Println("getting collection")
	req, err := newRequestFunc(NewPushRequest("GET", "https://api.pushbullet.com/v2/contacts", nil))
	if err != nil {
		fmt.Println(err)
	}
	resp, err := ExecuteRequest(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = c.ParseResponse(resp)
	if err != nil {
		fmt.Println(err)
	}
}

func (c *Contact) Create(params Params) {
	request := NewPushRequest("POST", "https://api.pushbullet.com/v2/contacts", params)
	fmt.Println(request)
	req, err := createFunc(request)
	fmt.Println(req)
	if err != nil {
		fmt.Println(err)
	}
	resp, err := ExecuteRequest(req)
	if err != nil {
		fmt.Println(err)
	}
	if resp.StatusCode == http.StatusOK {
		fmt.Println("Contact added")
	} else {
		fmt.Printf("Error during request, status code: %d", resp.StatusCode)
	}
}

func (c *Contact) Update(params Params) {
	var endpoint string
	if id, ok := params["iden"]; ok {
		endpoint = fmt.Sprintf("https://api.pushbullet.com/v2/contacts/%s", id)
		delete(params, "iden")
	} else {
		//TODO
		//return error
	}
	request := NewPushRequest("POST", endpoint, params)
	req, err := createFunc(request)
	fmt.Println(req)
	if err != nil {
		fmt.Println(err)
	}
	resp, err := ExecuteRequest(req)
	if err != nil {
		fmt.Println(err)
	}
	if resp.StatusCode == http.StatusOK {
		fmt.Println("Contact updated")
	} else {
		fmt.Printf("Error during request, status code: %d", resp.StatusCode)
	}
}

func (c *Contact) Delete(params Params) {
	var endpoint string
	if id, ok := params["iden"]; ok {
		endpoint = fmt.Sprintf("https://api.pushbullet.com/v2/contacts/%s", id)
		delete(params, "iden")
	} else {
		//TODO
		//return error
	}
	request := NewPushRequest("DELETE", endpoint, nil)
	req, err := newRequestFunc(request)
	fmt.Println(req)
	if err != nil {
		fmt.Println(err)
	}
	resp, err := ExecuteRequest(req)
	if err != nil {
		fmt.Println(err)
	}
	if resp.StatusCode == http.StatusOK {
		fmt.Println("Contact deleted")
	} else {
		fmt.Printf("Error during request, status code: %d", resp.StatusCode)
	}
}

func (c *Contact) ParseResponse(res *http.Response) error {
	defer res.Body.Close()
	err := json.NewDecoder(res.Body).Decode(c)
	if err != nil {
		return err
	}
	return nil
}

func (c *ContactsCollection) ParseResponse(res *http.Response) error {
	defer res.Body.Close()
	err := json.NewDecoder(res.Body).Decode(c)
	if err != nil {
		return err
	}
	return nil
}

func (m *Me) Get() {
	req, err := newRequestFunc(NewPushRequest("GET", "https://api.pushbullet.com/v2/users/me", nil))
	if err != nil {
		fmt.Println(err)
	}
	resp, err := ExecuteRequest(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = m.ParseResponse(resp)
	if err != nil {
		fmt.Println(err)
	}
}

func (m *Me) ParseResponse(res *http.Response) error {
	defer res.Body.Close()
	err := json.NewDecoder(res.Body).Decode(m)
	if err != nil {
		return err
	}
	return nil
}

func NewPushRequest(method, endpoint string, params Params) *PushRequest {
	return &PushRequest{Method: method, Endpoint: endpoint, Params: params}
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
	//req, _ := GetContacts()
	//req, _ := GetMe()
	//fmt.Println(req)
	//req2, _ := CreateContact(Params{"teste": "bla"})
	//fmt.Println(req2)
	//req3, _ := UpdateContact(Params{"iden": "123"})
	//fmt.Println(req3)
	//req4, _ := DeleteContact(Params{"iden": "123"})
	//fmt.Println(req4)
	//fmt.Println(req)
	//resp, err := ExecuteRequest(req)
	//if err != nil {
	//fmt.Println(err)
	//return
	//}

	coll := &ContactsCollection{}
	coll.Get()
	fmt.Println(coll)

	me := &Me{}
	me.Get()
	fmt.Println(me)

	//contact := &Contact{}
	//fmt.Println(contact)
	//contact.Create(Params{"name": "foo", "email": "baz@bar.com"})
	//contact.Update(Params{"iden": "ujDigsFMxWesjAdP9ajVsq", "name": "fooupdated"})
	//contact.Delete(Params{"iden": "ujDigsFMxWesjAwktvpfqu"})
}
