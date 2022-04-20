package main

import (
	"fmt"
	"log"

	"github.com/grokify/go-adaptivecards/demo"
)

func main() {
	err := demo.WriteJSONFiles(
		"example_card_bullets_min.json",
		"example_card_bullets.json", "", "  ",
		demo.ExampleCardBullets(""), 0644)
	if err != nil {
		log.Fatal(err)
	}

	err = demo.WriteJSONFiles(
		"example_card_bullets-multi_min.json",
		"example_card_bullets-mult.json", "", "  ",
		demo.ExampleCardBulletsMulti(), 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DONE")
}
