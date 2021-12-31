package main

import (
	"fmt"
	gateway "order-example/orders/gateway"
	models "order-example/orders/models"
)

func main() {
	cate1 := models.Category{}
	cate1.SetId(1)
	cate1.SetName("Comida")

	cate2 := models.Category{}
	cate2.SetId(2)
	cate2.SetName("Limpieza")

	prod1 := models.NewProduct(1, "Pollo a la Brasa", cate1)
	prod2 := models.NewProduct(2, "Pizza", cate1)
	prod3 := models.NewProduct(3, "Clorox", cate2)

	prod_gway := gateway.NewProductGateway()
	prod_gway.AddProduct(prod1)
	prod_gway.AddProduct(prod2)
	prod_gway.AddProduct(prod3)

	prod3.SetName("Clorox modif...")

	fmt.Println(prod_gway.GetProducts())

	for _, prod := range prod_gway.GetProducts() {
		fmt.Println(*prod)
	}
}
