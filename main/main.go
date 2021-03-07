package main

import (
	"fmt"

	"github.com/rohan-av/bigorder/item"
	"github.com/rohan-av/bigorder/strict"
)

func main() {
	i := item.Item{Name: "hello world"}
	arr := []*item.Item{&i}
	orderer := strict.StrictOrderer{Items: arr}
	fmt.Println(orderer.GetItems()[0].GetName())
}
