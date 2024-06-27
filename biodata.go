package main

import (
	"fmt"
	"os"
	"strconv"
)

// Friend struct to hold the friend's data
type Friend struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

// Function to get friend data by ID
func getFriendByID(id int) (Friend, bool) {
	friends := map[int]Friend{
		1: {"Budi", "Jakarta", "Programmer", "Ingin belajar bahasa Go untuk backend"},
		2: {"Ani", "Bandung", "Desainer", "Tertarik dengan efisiensi Go"},
		3: {"Siti", "Surabaya", "Data Analyst", "Ingin mengolah data dengan Go"},
		// Tambahkan data teman lainnya di sini
	}

	friend, exists := friends[id]
	return friend, exists
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run biodata.go <absen>")
		return
	}

	absen, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Absen harus berupa angka")
		return
	}

	friend, exists := getFriendByID(absen)
	if !exists {
		fmt.Println("Data teman dengan absen", absen, "tidak ditemukan")
		return
	}

	fmt.Println("Nama: ", friend.Nama)
	fmt.Println("Alamat: ", friend.Alamat)
	fmt.Println("Pekerjaan: ", friend.Pekerjaan)
	fmt.Println("Alasan memilih kelas Golang: ", friend.Alasan)
}
