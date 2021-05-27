package strict

import (
	"fmt"

	item "github.com/rohan-av/bigorder/item"
)

type StrictOrderer struct {
	Items         []*item.Item
	OutgoingComps chan [2]string
	IncomingComps chan [2]string
}

func (s *StrictOrderer) GetItems() []*item.Item {
	return s.Items
}

func (s *StrictOrderer) GetSortedList() []*item.Item {
	if _, ok := <-s.IncomingComps; !ok {
		return s.GetItems()
	} else {
		return nil // sorting is not over yet
	}
}

func (s *StrictOrderer) GetNextComparison() [2]string {
	return <-s.OutgoingComps
}

func (s *StrictOrderer) SendComparison(higher, lower string) {
	go func() {
		s.IncomingComps <- [2]string{higher, lower}
	}()
}

// for debugging purposes
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

func (s *StrictOrderer) Sort() {
	// splits array into sorted and unsorted regions
	for i := 1; i < s.Len(); i++ {
		fmt.Printf("i = %v\n", i)
		s.binarySearch(0, i, i)
		s.PrintItems()
	}
	close(s.IncomingComps)
	close(s.OutgoingComps)
}

func (s *StrictOrderer) setItem(idx int, item *item.Item) {
	s.GetItems()[idx] = item
}

func (s *StrictOrderer) swap(idx1, idx2 int) {
	tmp := s.GetItems()[idx1]
	s.setItem(idx1, s.GetItems()[idx2])
	s.setItem(idx2, tmp)
}

func (s *StrictOrderer) insertItem(insertIdx, itemIdx int) {
	for i := itemIdx; i > insertIdx; i-- {
		s.swap(i, i-1)
	}
}

func (s *StrictOrderer) compare(idx1, idx2 int) int {
	items := s.GetItems()
	item1 := items[idx1].GetName()
	item2 := items[idx2].GetName()
	s.OutgoingComps <- [2]string{item1, item2}
	comparison := <-s.IncomingComps
	if comparison[0] == item1 {
		return idx1
	} else {
		return idx2
	}
}

func (s *StrictOrderer) binarySearch(start, end, item int) {
	median := (start + end) / 2
	if end < start {
		s.insertItem(start, item)
		return
	}
	if s.compare(median, item) == median {
		s.binarySearch(start, median-1, item)
	} else {
		s.binarySearch(median+1, end, item)
	}
}
