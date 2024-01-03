package main

import (
	"fmt"
)

func findShort(long string) string {
	//api stuff
	return
}

func main() {
	var input string
	fmt.Print("Insert your long URL: ")
	fmt.Scan(&input)

	store := make(map[string]string)

	if len(store) == 0 {
		store[input] = findShort(input)
	} else {
		for k, v := range store {
			if k == input {
				break
			}
		}
	}

}
