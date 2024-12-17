// 1. Konstanta
// Konstanta adalah jenis variabel yang nilainya tidak bisa diubah setelah dideklarasikan. Inisialisasi nilai konstanta hanya dilakukan sekali saja di awal, setelah itu variabel tidak bisa diubah nilainya.

// package main

// import "fmt"

// func main() {
// 	// const firstName string = "john"
// 	// fmt.Print("halo ", firstName, "!\n")

// 	const lastName = "wick"
// 	fmt.Print("nice to meet you ", lastName, "!\n")
// }


// 2. Deklarasi Multi Konstanta
package main

import "fmt"

func main() {
	const (
		square          = "kotak"
		isToday bool    = true
		numeric uint8   = 1
		floatNum        = 2.2
	)
}

/*
square, dideklarasikan dengan metode type inference dengan tipe data string dan nilainya "kotak"
isToday, dideklarasikan dengan metode manifest typing dengan tipe data bool dan nilainya true
numeric, dideklarasikan dengan metode manifest typing dengan tipe data uint8 dan nilainya 1
floatNum, dideklarasikan dengan metode type inference dengan tipe data float dan nilainya 2.2
*/

// Contoh deklarasi konstanta dengan tipe data dan nilai yang sama:

const (
    a = "konstanta"
    b
)

/*
Ketika tipe data dan nilai tidak dituliskan dalam deklarasi konstanta, maka tipe data dan nilai yang dipergunakan adalah sama seperti konstanta yang dideklarasikan diatasnya.

a dideklarasikan dengan metode type inference dengan tipe data string dan nilainya "konstanta"
b dideklarasikan dengan metode type inference dengan tipe data string dan nilainya "konstanta"
*/

const (
    today string = "senin"
    sekarang
    isToday2 = true
)

/*
today dideklarasikan dengan metode manifest typing dengan tipe data string dan nilainya "senin"
sekarang dideklarasikan dengan metode manifest typing dengan tipe data string dan nilainya "senin"
isToday2 dideklarasikan dengan metode type inference dengan tipe data bool dan nilainya true
*/

// Berikut contoh deklrasi multiple konstanta dalam satu baris:

const satu, dua = 1, 2
const three, four string = "tiga", "empat"

/*
satu, dideklarasikan dengan metode type inference dengan tipe data int dan nilainya 1
dua, dideklarasikan dengan metode type inference dengan tipe data int dan nilainya 2
three, dideklarasikan dengan metode manifest typing dengan tipe data string dan nilainya "tiga"
four, dideklarasikan dengan metode manifest typing dengan tipe data string dan nilainya "empat"
*/
