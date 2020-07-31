package main

import "fmt"

func main() {
	var angka int = 30
	if angka%2 == 1 {
		fmt.Println(angka, " = ganjil")
	} else {
		fmt.Println(angka, " = genap")
	}
}
