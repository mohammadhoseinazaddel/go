package main

import (
	"TopLearn/WebApp/datalayer"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	person := datalayer.Person{
		Name:   "Ali",
		Family: "Rezaee",
	}
	var data bytes.Buffer
	json.NewEncoder(&data).Encode(person)

	res, err := http.Post("http://localhost:8383/api/person/add", "application/json", &data)
	if err != nil || res.StatusCode != 200 {
		fmt.Println(res.Status, err)
	}

}
