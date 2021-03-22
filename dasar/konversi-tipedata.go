package main

import "fmt"

func main() {
	//variabel awal int32
	var nilai32 int32 = 129
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)

	//jika nilai melebihi batas tampungan maka ia akan balik lagi ke nilai dasar tipe data
	fmt.Println(nilai8)
}
