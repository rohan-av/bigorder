package strict

import (
	"fmt"

	item "github.com/rohan-av/bigorder/item"
)

type StrictOrderer struct {
	Items []*item.Item
}

func (s *StrictOrderer) GetItems() []*item.Item {
	return s.Items
}

func (s *StrictOrderer) SetItem(idx int, item *item.Item) {
	s.GetItems()[idx] = item
}

func (s *StrictOrderer) PrintItems() {
	fmt.Print("\n")
	for i := 0; i < s.Len(); i++ {
		fmt.Println(s.GetItems()[i].GetName())
	}
	fmt.Print("\n")
}

func (s *StrictOrderer) Len() int {
	return len(s.GetItems())
}

func (s *StrictOrderer) Compare(idx1, idx2 int) int {
	items := s.GetItems()
	name1 := items[idx1].GetName()
	name2 := items[idx2].GetName()
	fmt.Printf("Which is better? %v or %v?\n", name1, name2)
	var userChoice string
	fmt.Scanln(&userChoice)
	if userChoice == name1 {
		return idx1
	} else {
		return idx2
	}
}

func (s *StrictOrderer) Swap(idx1, idx2 int) {
	tmp := s.GetItems()[idx1]
	s.SetItem(idx1, s.GetItems()[idx2])
	s.SetItem(idx2, tmp)
}

func (s *StrictOrderer) InsertItem(insertIdx, itemIdx int) {
	// items := s.GetItems()
	// item := items[itemIdx]
	for i := itemIdx; i > insertIdx; i-- {

		s.Swap(i, i-1)
	}
}

func (s *StrictOrderer) BinarySearch(start, end, item int) {
	median := (start + end) / 2
	if end < start {
		s.InsertItem(start, item)
		return
	}
	if s.Compare(median, item) == median {
		s.BinarySearch(start, median-1, item)
	} else {
		s.BinarySearch(median+1, end, item)
	}
}

func (s *StrictOrderer) Sort() {
	// splits array into sorted and unsorted regions
	for i := 1; i < s.Len(); i++ {
		fmt.Printf("i = %v\n", i)
		s.BinarySearch(0, i, i)
		s.PrintItems()
	}
}
