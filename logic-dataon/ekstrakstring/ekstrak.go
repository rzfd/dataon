package main

import (
	"fmt"
	"strings"
)

func main() {
	var email string

	fmt.Print("Masukkan email: ")
	fmt.Scan(&email)

	parts := strings.Split(email, "@")

	if len(parts) != 2 {
		fmt.Println("Format email tidak valid")
		return
	}

	e := parts[1]

	fmt.Println("Domain:", e)
}
