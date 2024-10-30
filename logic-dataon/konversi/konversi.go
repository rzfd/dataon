package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var angka string

	fmt.Print("Masukkan angka: ")
	fmt.Scan(&angka)

	angka = strings.ReplaceAll(angka, ".", "")

	length := len(angka)

	for i, digit := range angka {
		digitInt, _ := strconv.Atoi(string(digit))

		placeValue := digitInt * intPow(10, length-i-1)

		if placeValue != 0 {
			fmt.Println(placeValue)
		}
	}
}

func intPow(base, exp int) int {
	result := 1
	for exp > 0 {
		result *= base
		exp--
	}
	return result
}
