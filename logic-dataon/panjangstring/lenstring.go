package main

import (
	"fmt"
)

func main() {
	var s string

	fmt.Print("Masukkan string: ")
	fmt.Scan(&s)

	panjang := len(s)

	fmt.Println("Panjang dari string:", panjang)
}
