package main

import (
	"fmt"
)

func main() {
	// //ini namanya komentar
	// hallo := "ini tipe datanya string"
	// fmt.Println(hallo)

	// var iniSlice []int32 = []int32{1, 2, 2, 3, 4, 5, 6, 7, 8, 4, 11}
	// //kalau di python :
	// //iniSlice = [1,2,312,41324234,234234,3234]
	// fmt.Println("isi dari slice : ", iniSlice)
	// fmt.Println("alamat dari variable iniSlice : ", &iniSlice[0])
	// fmt.Println("alamat dari variable hallo : ", &hallo)

	// oke := "fakhri"
	// oke2 := "op2"
	// // || = melambangan or
	// // && = melambangkan and
	// if oke == "fakhri" && oke2 == "op" {
	// 	fmt.Println("oke memiliki nilai true")
	// } else {
	// 	fmt.Println("oke memiliki nilai false")
	// }

	// for i := 0; i < len(iniSlice); i++ {
	// 	if i == 7 {
	// 		break
	// 	}
	// 	if i%2 == 0 {
	// 		continue
	// 	} else {
	// 		fmt.Println("i ke-", i)
	// 	}
	// }

	// fungsiCetak()
	// // fungsiPenjumlahan(1, 2)
	// terserah := fungsiPenjumlahan(1, 2)
	// fmt.Println(terserah)

	// terserah2 := module1.FungsiPerkalian(4, 5)
	// fmt.Println(terserah2)

	// ovis := module1.Mahasiswa{
	// 	Nama:  "ovi sanjaya",
	// 	Nim:   "192499999999",
	// 	Prodi: "teknik kuli",
	// }

	// fmt.Println(ovis.Nama)
}

func fungsiCetak() {
	fmt.Println("ini dicetak di dalam fungsi")
}

func fungsiPenjumlahan(a int, b int) int {
	return a + b
}
