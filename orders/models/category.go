package orders

type Category struct {
	id   int
	name string
}

func (cat *Category) SetId(id int) {
	cat.id = id
}

func (cat Category) GetId() int {
	return cat.id
}

func (cat *Category) SetName(name string) {
	cat.name = name
}

func (cat Category) GetName() string {
	return cat.name
}
