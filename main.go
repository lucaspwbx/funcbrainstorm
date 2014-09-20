package main

import "fmt"

var (
	funcao = func(method, endpoint string, params Params) string {
		teste := fmt.Sprintf("Doing %s request to: %s", method, endpoint)
		teste += "\n"
		teste += fmt.Sprintf("Params:")
		teste += "--------\n"
		for k, v := range params {
			teste += fmt.Sprintf("Key: %s, Value: %s", k, v)
			teste += "\n"
		}
		return teste
	}
)

type Params map[string]string

func GetContacts(params Params) string {
	//fn := func(method, endpoint string, params Params) string {
	//teste := fmt.Sprintf("Doing %s request to: %s", method, endpoint)
	//teste += "\n"
	//teste += fmt.Sprintf("Params:")
	//teste += "--------\n"
	//for k, v := range params {
	//teste += fmt.Sprintf("Key: %s, Value: %s", k, v)
	//}
	//return teste
	//}
	//return fn("GET", "/contacts", params)
	return funcao("GET", "/contacts", params)
}

func CreateContact(params Params) string {
	return funcao("POST", "/contacts", params)
}

func UpdateContact(params Params) string {
	//fn := func(method, endpoint string, params Params) string {
	//teste := fmt.Sprintf("Doing %s request to: %s", method, endpoint)
	//teste += "\n"
	//teste += fmt.Sprintf("Params:")
	//teste += "--------\n"
	//for k, v := range params {
	//teste += fmt.Sprintf("Key: %s, Value: %s", k, v)
	//teste += "\n"
	//}
	//	return teste
	//}
	//return fn("POST", "/contacts", params)
	iden := params["iden"]
	endpoint := fmt.Sprintf("/contacts/%s", iden)
	return funcao("POST", endpoint, params)
}

func main() {
	fmt.Println(GetContacts(map[string]string{"title": "NotaUm"}))
	fmt.Println(CreateContact(map[string]string{"name": "Lucas", "email": "foo@bar.com"}))
	fmt.Println(UpdateContact(map[string]string{"iden": "1234", "name": "Lucas"}))
}
