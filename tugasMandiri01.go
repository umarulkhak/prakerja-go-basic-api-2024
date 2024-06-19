package main

import "fmt"

func main() {
    // Deklarasi variabel
    var name string = "John Doe"
    var age int = 10
    var isStudent bool = true

    // Print variabel dengan go verb yang sesuai
    fmt.Printf("%s\n", name)    // string
    fmt.Printf("%d\n", age)     // integer
    fmt.Printf("%t\n", isStudent) // boolean

    // Print variabel dengan go verb untuk mengecek tipe data
    fmt.Printf("%T\n", name)    // string
    fmt.Printf("%T\n", age)     // int
    fmt.Printf("%T\n", isStudent) // bool
}
