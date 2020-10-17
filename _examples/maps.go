package main

import (
	"fmt"
	go_googleapis "github.com/BRUHItsABunny/go-googleapis"
)

func main() {

	// Fill in this stuff, whether you get your own or reverse engineer some is up to you
	apiKey := ""

	// Initialize client
	mapsClient := go_googleapis.GetMapsClient(apiKey)

	// Example coords
	nyLat := "40.730610"
	nyLng := "-73.935242"

	abLat := "42.933334"
	abLng := "-76.566666"

	nearby, token, err := mapsClient.Nearby(nyLat, nyLng, "restaurants", 10000)
	if err == nil {
		fmt.Println("Next page token:", token)
		for i, e := range nearby {
			fmt.Println("Result", i+1)
			fmt.Println(e.Name)
			fmt.Println(e.Vicinity)
			if e.BusinessStatus != nil {
				fmt.Println(*e.BusinessStatus)
			}
		}
	} else {
		fmt.Println("Error occurred:", err)
	}

	directions, err := mapsClient.Directions(nyLat, nyLng, abLat, abLng)
	if err == nil {
		for i, e := range directions {
			fmt.Println("Route", i+1)
			fmt.Println(e.Summary)
			for _, ee := range e.Legs {
				fmt.Println("From:", ee.StartAddress, "\nTo:", ee.EndAddress)
				for ii, eee := range ee.Steps {
					fmt.Println("Step", ii+1)
					fmt.Println(eee.HTMLInstructions)
					fmt.Println("Distance:", eee.Distance.Text, "\nShould take about", eee.Duration.Text)
				}
			}
		}
	} else {
		fmt.Println("Error occurred:", err)
	}
}
