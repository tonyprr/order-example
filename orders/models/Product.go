package orders

type Product struct {
	id       int
	category Category
	name     string
}

func (prod *Product) setId(id int) {
	prod.id = id
}

func (prod Product) getId() int {
	return prod.id
}
