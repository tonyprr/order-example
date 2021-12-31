package orders

import (
	pojos "order-example/orders/models"
)

type ProductGateway interface {
	GetProducts() []*pojos.Product
	AddProduct(p *pojos.Product) (bool, error)
}

type ProductInMemory struct {
	products []*pojos.Product
}

func (pm ProductInMemory) GetProducts() []*pojos.Product {
	return pm.products
}

func (pm *ProductInMemory) AddProduct(p *pojos.Product) (bool, error) {
	pm.products = append(pm.products, p)

	return true, nil
}

func NewProductGateway() ProductGateway {
	return &ProductInMemory{}
}
