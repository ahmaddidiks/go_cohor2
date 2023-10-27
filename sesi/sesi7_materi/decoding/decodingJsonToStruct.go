package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

func decodingJsonToStruct() {
	var jsonString = `
		{
			"full_name": "Fitri Ayu",
			"email": "fitri@mail.com",
			"age": 23
		}
	`

	var result Student

	var err = json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("full_name: ", result.FullName)
	fmt.Println("email: ", result.Email)
	fmt.Println("age: ", result.Age)

}
