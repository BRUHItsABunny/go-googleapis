package main

import (
	"fmt"
	go_googleapis "github.com/BRUHItsABunny/go-googleapis"
)

func main() {

	// Initialize client
	translateClient := go_googleapis.GetTranslateClient()

	// Do TTS, success will return filename of where mp3 is stored
	fileName, err := translateClient.TTS("The bunny is happy\nThe chicken is also happy", "en")
	if err == nil {
		// No error, enjoy mp3
		fmt.Println("Open " + fileName + " to listen to the file")
	} else {
		fmt.Println("Error occurred", err)
	}

	// Translate sentences to Dutch
	result, err := translateClient.Translate("The bunny is happy\nThe chicken is also happy", "auto", "nl")
	if err == nil {
		fmt.Println("Translated from " + result.SrcLang + " to " + result.DstLang)
		fmt.Println(result.Origin + "\n")
		fmt.Println(result.Translate + "\n")
	} else {
		fmt.Println("Error occurred", err)
	}
}
