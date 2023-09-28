package main

import (
	"challenge_3/helpers"
	"fmt"
	"os"
	"strconv"
)

/**
* Ini bukan cara terbaik, kayaknya malah ribet wkw.
* pengen explore saja sih Mba
*/

func main() {
	arg := os.Args[1]                   // only check the first param
	numberArg, err := strconv.Atoi(arg) // check is it a number

	var result helpers.Employee

	if err == nil { // arg is a number
		if helpers.ShortbyID(numberArg, &result) {
			result.Show()
		} else {
			fmt.Printf("Tidak ada Employee dengan ID %d\n", numberArg)
		}
	} else { // arg is a string
		if helpers.Shortbyname(arg, &result) {
			result.Show()
		} else {
			fmt.Printf("Tidak ada Employee dengan nama %s\n", arg)
		}
	}
}
