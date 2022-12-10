package main

import (
	go_gcp_misc "go-gcp-misc"
)

func main() {
	cat := go_gcp_misc.Cat1{Id: 3, Name: "Murzik", Age: 5, Color: "white"}
	go_gcp_misc.CreateCat1("Cat1", &cat)
}
