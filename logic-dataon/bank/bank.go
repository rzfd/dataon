package main

import (
	"fmt"
)

func main() {
	var totalUang int
	fmt.Print("Masukkan jumlah uang yang akan diambil: Rp ")
	fmt.Scan(&totalUang)

	lembar := []int{100000, 50000, 20000, 5000}
	koin := []int{100, 50}

	jumlahPecahan := make(map[int]int)

	for _, p := range lembar {
		if totalUang >= p {
			jumlahPecahan[p] = totalUang / p
			totalUang = totalUang % p
		}
	}

	for _, p := range koin {
		if totalUang >= p {
			jumlahPecahan[p] = totalUang / p
			totalUang = totalUang % p
		}
	}

	fmt.Println("Jumlah pecahan yang dibutuhkan:")

	fmt.Println("Uang Kertas:")
	for _, p := range lembar {
		if jumlahPecahan[p] > 0 {
			fmt.Printf("Pecahan Rp %d: %d lembar\n", p, jumlahPecahan[p])
		}
	}

	fmt.Println("Koin:")
	for _, p := range koin {
		if jumlahPecahan[p] > 0 {
			fmt.Printf("Pecahan Rp %d: %d koin\n", p, jumlahPecahan[p])
		}
	}
	if totalUang > 0 {
		fmt.Printf("Sisa uang yang tidak dapat dipecah: Rp %d\n", totalUang)
	}
}
