package main

import "fmt"

var (
	requestFunc = func(method, endpoint string, params Params) string {
		debug := fmt.Sprintf("Doing %s request to: %s\n", method, endpoint)
		debug += fmt.Sprintf("Params:")
		debug += "--------\n"
		for k, v := range params {
			debug += fmt.Sprintf("Key: %s, Value: %s\n", k, v)
		}
		return debug
	}
)

type Params map[string]string

func GetContacts(params Params) string {
	return requestFunc("GET", "/contacts", params)
}

func CreateContact(params Params) string {
	return requestFunc("POST", "/contacts", params)
}

func UpdateContact(params Params) string {
	contact_iden := params["iden"]
	endpoint := fmt.Sprintf("/contacts/%s", contact_iden)
	return requestFunc("POST", endpoint, params)
}

func DeleteContact(params Params) string {
	contact_iden := params["iden"]
	endpoint := fmt.Sprintf("/contacts/%s", contact_iden)
	return requestFunc("DELETE", endpoint, params)
}

func main() {
	fmt.Println(GetContacts(map[string]string{"title": "NotaUm"}))
	fmt.Println(CreateContact(map[string]string{"name": "Lucas", "email": "foo@bar.com"}))
	fmt.Println(UpdateContact(map[string]string{"iden": "1234", "name": "Lucas"}))
	fmt.Println(DeleteContact(map[string]string{"iden": "1234"}))
}
