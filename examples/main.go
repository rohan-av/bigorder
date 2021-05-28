package main

import (
	"fmt"

	"github.com/rohan-av/bigorder/item"
	"github.com/rohan-av/bigorder/strict"
)

func main() {
	arr := []item.Item{
		&Album{Name: "pinkerton"},
		&Album{Name: "blue"},
		&Album{Name: "green"},
		&Album{Name: "white"},
		&Album{Name: "hurley"},
		&Album{Name: "ewbaite"},
		&Album{Name: "okhuman"},
		&Album{Name: "black"},
		&Album{Name: "pacific"},
	}

	orderer, err := strict.NewStrictOrderer(arr)
	if err != nil {
		fmt.Printf("error: cannot create orderer: %v", err)
		return
	}

	fmt.Println(orderer.GetProgress())
	go orderer.Sort()

	for {
		if items, ok := orderer.GetNextComparison(); ok {
			fmt.Printf("Which is better? %v or %v?\n", items[0], items[1])
			var userChoice string
			fmt.Scanln(&userChoice)
			if userChoice == items[0] {
				orderer.SendComparison(items[0], items[1])
			} else {
				orderer.SendComparison(items[1], items[0])
			}
			fmt.Println(orderer.GetProgress())
		} else {
			fmt.Println("channel closed")
			fmt.Printf("1st: %v\n", orderer.GetItems()[8].GetName())
			fmt.Printf("2nd: %v\n", orderer.GetItems()[7].GetName())
			fmt.Printf("3rd: %v\n", orderer.GetItems()[6].GetName())
			return
		}
	}
}
