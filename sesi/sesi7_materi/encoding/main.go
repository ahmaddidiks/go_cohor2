package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	FullName string `json:"name"`
	Age      int
}

func main() {
	var object = []User{{"Rose", 22}, {"Lisa", 20}}

	var jsonData, err = json.Marshal(object)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var jsonString = string(jsonData)
	fmt.Println(jsonString)
}
