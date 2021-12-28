package orders

import (
	pojos "order-example/orders/models"
)

type ProductGateway interface {
	GetProducts() []*pojos.Product
}
