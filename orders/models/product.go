package orders

type Product struct {
	id       int
	name     string
	category Category
}

func (prod *Product) setId(id int) {
	prod.id = id
}

func (prod Product) getId() int {
	return prod.id
}

func (prod *Product) setCategory(category Category) {
	prod.category = category
}

func (prod Product) getCategory() Category {
	return prod.category
}

func (prod *Product) setName(name string) {
	prod.name = name
}

func (prod Product) getName() string {
	return prod.name
}

func NewProduct(id int, name string, category Category) *Product {
	return &Product{
		id:       id,
		name:     name,
		category: category,
	}
}
