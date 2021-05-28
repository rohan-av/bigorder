package main

type Album struct {
	Name string
}

func (a *Album) GetName() string {
	return a.Name
}
