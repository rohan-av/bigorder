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
		{Name: "ewbaite"},
		{Name: "okhuman"},
		{Name: "black"},
		{Name: "pacific"},
	}
	orderer := strict.StrictOrderer{Items: arr}
	orderer.Sort()
	fmt.Println(orderer.GetItems()[7].GetName())
	fmt.Println(orderer.GetItems()[6].GetName())
	fmt.Println(orderer.GetItems()[5].GetName())
}
