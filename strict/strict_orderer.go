package strict

import (
	"fmt"
	"math"

	item "github.com/rohan-av/bigorder/item"
)

type StrictOrderer struct {
	Items         []*item.Item
	OutgoingComps chan [2]string
	IncomingComps chan [2]string
	Progress      [2]int // current, total (estimated)
}

func NewStrictOrderer(items []*item.Item, incomingComps, outgoingComps chan [2]string) *StrictOrderer {
	progress := [2]int{}
	progress[1] = getEstimatedLeft(1, len(items))
	orderer := StrictOrderer{
		Items:         items,
		OutgoingComps: outgoingComps,
		IncomingComps: incomingComps,
		Progress:      progress,
	}
	return &orderer
}

func getEstimatedLeft(start, end int) int {
	var res float64 = 0
	for i := start; i <= end; i++ {
		res = res + math.Log2(float64(i)+0.001)
	}
	return int(math.Floor(res))
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

func (s *StrictOrderer) GetNextComparison() [2]string {
	return <-s.OutgoingComps
}

func (s *StrictOrderer) SendComparison(higher, lower string) {
	s.IncomingComps <- [2]string{higher, lower}
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
		s.Progress[0] = getEstimatedLeft(1, i)
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
	if item1 == item2 {
		return idx1
	}
	s.OutgoingComps <- [2]string{item1, item2}
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
