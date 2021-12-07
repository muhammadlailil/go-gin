package log

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)
type todosHttp struct {
	UsersId int `json:"userId"`
	ID  int `json:"id"`
	Title  string `json:"title"`
	Comple  bool `json:"completed"`
}
func curl()  {
	url := "https://jsonplaceholder.typicode.com/todos"
	method := "GET"

	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	var result[]todosHttp
	json.Unmarshal(body, &result)

	for _, value := range result {
		fmt.Println(value.Title)
	}
}
