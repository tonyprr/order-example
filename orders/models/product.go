package orders

type Product struct {
	id       int
	name     string
	category Category
}

func (prod *Product) SetId(id int) {
	prod.id = id
}

func (prod Product) GetId() int {
	return prod.id
}

func (prod *Product) SetCategory(category Category) {
	prod.category = category
}

func (prod Product) GetCategory() Category {
	return prod.category
}

func (prod *Product) SetName(name string) {
	prod.name = name
}

func (prod Product) GetName() string {
	return prod.name
}

func NewProduct(id int, name string, category Category) *Product {
	return &Product{
		id:       id,
		name:     name,
		category: category,
	}
}
