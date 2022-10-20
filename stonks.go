package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	today := time.Now()
	tomorrow := today.AddDate(0, 0, 1)

	min := 93
	max := 100
	answers := []string{
		"green",
		"red",
	}
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("Tomorrow:", tomorrow.Month(), tomorrow.Day(), tomorrow.Year(), "will be a", answers[rand.Intn(len(answers))], "day")
	fmt.Println("Accuracy: ", rand.Intn(max-min)+min, "%")
}
