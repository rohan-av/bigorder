package strict

import (
	"fmt"
	"math"

	item "github.com/rohan-av/bigorder/item"
)

type StrictOrderer struct {
	Items         []*item.Item
	OutgoingComps chan [2]*item.Item
	IncomingComps chan [2]*item.Item
	Progress      [2]int // current, total (estimated)
}

func NewStrictOrderer(items []*item.Item) (*StrictOrderer, error) {
	if len(items) == 0 || len(items) == 1 {
		return nil, fmt.Errorf("no sorting required")
	}

	progress := [2]int{0, getEstimatedLeft(1, len(items))}
	outgoingComps := make(chan [2]*item.Item)
	incomingComps := make(chan [2]*item.Item)
	orderer := StrictOrderer{
		Items:         items,
		OutgoingComps: outgoingComps,
		IncomingComps: incomingComps,
		Progress:      progress,
	}
	return &orderer, nil
}

func (s *StrictOrderer) GetItems() []*item.Item {
	return s.Items
}

func (s *StrictOrderer) GetProgress() [2]int {
	return s.Progress
}

func (s *StrictOrderer) GetSortedList() []*item.Item {
	if _, ok := <-s.IncomingComps; !ok {
		return s.GetItems()
	} else {
		return nil // sorting is not over yet
	}
}

func (s *StrictOrderer) GetNextComparison() ([2]*item.Item, bool) {
	comp, ok := <-s.OutgoingComps
	return comp, ok
}

func (s *StrictOrderer) SendComparison(higher, lower *item.Item) {
	s.IncomingComps <- [2]*item.Item{higher, lower}
}

func (s *StrictOrderer) Len() int {
	return len(s.GetItems())
}

func (s *StrictOrderer) Sort() {
	// splits array into sorted and unsorted regions
	for i := 1; i < s.Len(); i++ {
		fmt.Printf("i = %v\n", i)
		s.Progress[0] = getEstimatedLeft(1, i)
		s.binarySearch(0, i, i)
		s.printItems()
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
	item1 := items[idx1]
	item2 := items[idx2]
	if item1 == item2 {
		return idx1
	}
	s.OutgoingComps <- [2]*item.Item{item1, item2}
	comparison := <-s.IncomingComps
	s.Progress[0] = s.Progress[0] + 1
	if comparison[0] == item1 {
		return idx1
	} else {
		return idx2
	}
}

func (s *StrictOrderer) binarySearch(start, end, item int) {
	median := start + (end-start)/2
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

func getEstimatedLeft(start, end int) int {
	var res float64 = 0
	for i := start; i <= end; i++ {
		res = res + math.Log2(float64(i)+0.001)
	}
	return int(math.Floor(res))
}

// for debugging purposes
func (s *StrictOrderer) printItems() {
	fmt.Print("\n")
	for i := 0; i < s.Len(); i++ {
		fmt.Println(s.GetItems()[i].GetName())
	}
	fmt.Print("\n")
}
