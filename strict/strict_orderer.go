package strict

import (
	item "github.com/rohan-av/bigorder/item"
)

type StrictOrderer struct {
	Items []*item.Item
}

func (s *StrictOrderer) GetItems() []*item.Item {
	return s.Items
}
