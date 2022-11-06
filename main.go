package main

import "fmt"

type car struct {
	model string
	year  int
}

type Product struct {
	car      []car
	name     string
	price    float64
	Quantity int
}

func (p Product) showProduct() {
	fmt.Println("product name:", p.name)
	fmt.Println("product price", p.price)
	fmt.Println("product Quantity", p.price)
	fmt.Println("product car", p.car)
	fmt.Println("----------------")
}

func (p Product) showProducStatus() {
	if p.Quantity > 0 {
		fmt.Println("product is available")
	} else {
		fmt.Println("product is available")
	}
}

type store struct {
	Products          []Product
	Numberofproducts  int
	totalSoldProducts int
	totalRevenue      float64
}

func (s *store) addProduct(p Product) {
	s.Products = append(s.Products, p)
	s.Numberofproducts++
}

func (s *store) sellProduct(name string, Quantity int) {
	for i, p := range s.Products {
		if p.name == name {
			if p.Quantity >= Quantity {
				s.Products[i].Quantity -= Quantity
				s.totalSoldProducts += Quantity
				s.totalRevenue += p.price * float64(Quantity)
			} else {
				fmt.Println("Not Enough Quantity")
			}
		}
	}
}

func (s *store) showNumOfProducts() {
	fmt.Println("Total products:", s.Numberofproducts)
	fmt.Println("-------------")
}

func (s *store) showProducts() {
	for _, p := range s.Products {
		p.showProduct()
	}

}

func (s *store) showProductStatus() {
	for _, p := range s.Products {
		p.showProducStatus()
		fmt.Println("--------------")
	}
}

func (s *store) showTotalSoldProducts() {
	fmt.Println("Total sold products:", s.totalSoldProducts)
}

func (s *store) showTotalRevenue() {
	fmt.Println("Total revenue:", s.totalRevenue)
}

func main() {
	car1 := car{"Lexus", 2022}
	car2 := car{"Audi", 2021}
	car3 := car{"Honda", 2019}

	product1 := Product{[]car{car1, car2}, "car", 2000, 1000}
	product2 := Product{[]car{car3}, "car", 4000, 1500}

	store := store{}
	store.addProduct(product1)
	store.addProduct(product2)

	store.showNumOfProducts()

	store.showProducts()

	store.showProductStatus()

	store.sellProduct("car", 4)

	store.showTotalSoldProducts()

	store.showTotalRevenue()

}
