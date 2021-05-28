package item

type Item struct {
	Name string
	// additional properties can include Tier/Rating etc.
}

func (i *Item) GetName() string {
	return i.Name
}
