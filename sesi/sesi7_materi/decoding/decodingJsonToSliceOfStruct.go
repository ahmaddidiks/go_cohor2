package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

func decodingJsonToSliceOfStruct() {
	var jsonString = `[
			{
				"full_name": "Kim Jennie",
				"email": "jennie@mail.com",
				"age": 23
			},
			{
				"full_name": "Kim Jisoo",
				"email": "jisoo@mail.com",
				"age": 23
			}
		]
	`

	var result []Employee

	err := json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i, v := range result {
		fmt.Printf("Index %d: %+v\n", i+1, v)
	}
}
