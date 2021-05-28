package main

import (
	"fmt"

	"github.com/rohan-av/bigorder/item"
	"github.com/rohan-av/bigorder/strict"
)

func main() {
	arr := []*item.Item{
		{Name: "pinkerton"},
		{Name: "blue"},
		{Name: "green"},
		{Name: "white"},
		{Name: "hurley"},
		{Name: "ewbaite"},
		{Name: "okhuman"},
		{Name: "black"},
		{Name: "pacific"},
	}

	orderer, err := strict.NewStrictOrderer(arr)
	if err != nil {
		fmt.Printf("error: cannot create orderer: %v", err)
		return
	}

	fmt.Println(orderer.GetProgress())
	done := make(chan bool)

	go orderer.Sort()
	go func() {
		// endless for loop to handle human comparisons
		for {
			if items, ok := orderer.GetNextComparison(); ok {
				fmt.Printf("Which is better? %v or %v?\n", items[0].GetName(), items[1].GetName())
				var userChoice string
				fmt.Scanln(&userChoice)
				if userChoice == items[0].GetName() {
					orderer.SendComparison(items[0], items[1])
				} else {
					orderer.SendComparison(items[1], items[0])
				}
				fmt.Println(orderer.GetProgress())
			} else {
				fmt.Println("channel closed")
				done <- true
				return
			}
		}
	}()

	// do other thing

	<-done
	fmt.Printf("1st: %v\n", orderer.GetItems()[8].GetName())
	fmt.Printf("2nd: %v\n", orderer.GetItems()[7].GetName())
	fmt.Printf("3rd: %v\n", orderer.GetItems()[6].GetName())
	fmt.Println("finished")
}
