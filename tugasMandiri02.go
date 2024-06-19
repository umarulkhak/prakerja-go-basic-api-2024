package main

import (
	"fmt"
)

func main() {
	// Jumlah baris piramida
	n := 10

	for i := 0; i < n; i++ {
		// Cetak spasi
		for j := 0; j < n-i-1; j++ {
			fmt.Print(" ")
		}
		// Cetak bintang
		for k := 0; k < 2*i+1; k++ {
			fmt.Print("*")
		}
		// Pindah ke baris berikutnya
		fmt.Println()
	}
}
