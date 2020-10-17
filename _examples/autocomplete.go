package main

import (
	"fmt"
	go_googleapis "github.com/BRUHItsABunny/go-googleapis"
)

func main() {

	// Initialize client
	autoCompleteClient := go_googleapis.GetAutoCompleteClient()

	// Get array of suggestions
	result, err := autoCompleteClient.Suggest("Te", "en") // Language doesn't work YET
	if err == nil {
		// Process list of suggestions
		fmt.Println("You were typing: Te...")
		for _, e := range result {
			fmt.Println(e)
		}
	} else {
		fmt.Println("Error occurred", err)
	}
}
