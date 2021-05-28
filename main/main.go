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
	incoming := make(chan [2]string)
	outgoing := make(chan [2]string)
	orderer := strict.NewStrictOrderer(arr, outgoing, incoming)
	//orderer := strict.StrictOrderer{arr, outgoing, incoming, [2]int{}}
	fmt.Println(orderer.GetProgress())
	go orderer.Sort()
	for {
		select {
		case items, ok := <-incoming:
			if ok {
				fmt.Printf("Which is better? %v or %v?\n", items[0], items[1])
				var userChoice string
				fmt.Scanln(&userChoice)
				if userChoice == items[0] {
					outgoing <- items
				} else {
					outgoing <- [2]string{items[1], items[0]}
				}
				fmt.Println(orderer.GetProgress())
			} else {
				fmt.Println("channel closed")
				fmt.Printf("1st: %v\n", orderer.GetItems()[8].GetName())
				fmt.Printf("2nd: %v\n", orderer.GetItems()[7].GetName())
				fmt.Printf("3rd: %v\n", orderer.GetItems()[6].GetName())
				return
			}
		default:
			// fmt.Println("no value ready")
		}
	}
}
