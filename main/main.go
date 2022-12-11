package main

import go_gcp_misc "go-gcp-misc"

func main() {
	//cat := go_gcp_misc.Cat1{Id: 3, Name: "Murzik", Age: 5, Color: "white"}
	//go_gcp_misc.CreateCat1("Cat1", &cat)

	//for a := 0; a <= 5; a++ {
	//	go_gcp_misc.CreateMessage()
	//}

	go_gcp_misc.Listen()
}
