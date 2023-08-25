package main

import "fmt"

type name struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}
type desc struct {
	Description string `json:"description"`
}

func main() {
	// a := name{"foo", "bandarlampung"}
	// b := desc{"Description"}

	// c := struct {
	// 	name
	// 	desc
	// }{a, b}

	price := []int{1000, 2000}

	fmt.Println(price[0])
}
