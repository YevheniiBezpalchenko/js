package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/get_products", GetProducts)
	http.ListenAndServe(":8080", nil)
}

type Product struct {
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

func GetProducts(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Add("Access-Control-Allow-Origin", "*")
	products := make([]Product, 0)
	var product Product
	for i := 1; i <= 15; i++ {
		product = Product{
			Name:  "Product #" + strconv.Itoa(i),
			Price: float32(rand.Int()%10000) / 100,
		}
		products = append(products, product)
	}
	res, _ := json.Marshal(products)
	fmt.Fprintln(resp, string(res))

}
