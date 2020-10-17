package main

import (
	"fmt"
	go_googleapis "github.com/BRUHItsABunny/go-googleapis"
	"os"
)

func main() {

	// Fill in this stuff, whether you get your own or reverse engineer some is up to you
	apiKey := ""
	androidPackage := ""
	androidCert := ""

	// Initialize client
	visionClient := go_googleapis.GetVisionClient(apiKey, androidPackage, androidCert)

	// Read image to file
	f, err := os.Open("face.jpg")
	imageJob := go_googleapis.ImageJob{Reader: f}

	if err == nil {
		result, err := visionClient.ImageProperties(&imageJob)
		if err == nil {
			fmt.Println("Dominant colors:")
			for _, e := range result.DominantColors.Colors {
				fmt.Println(e.Color)
			}
		} else {
			fmt.Println("Error during API call", err)
		}
	} else {
		fmt.Println("Error reading file", err)
	}
}
