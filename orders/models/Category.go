package orders

type Category struct {
	id   int
	name string
}

func (cat *Category) setId(id int) {
	cat.id = id
}

func (cat Category) getId() int {
	return cat.id
}

func (cat *Category) setName(name string) {
	cat.name = name
}

func (cat Category) getName() string {
	return cat.name
}
