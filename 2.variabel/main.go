package main

import "fmt"

func main() {

	// 1. Deklarasi Variabel Beserta Tipe Data
	var firstName string = "john"

	var lastName string
	lastName = "wick"

	fmt.Printf("halo %s %s!\n", firstName, lastName)

	// 2. Deklarasi Variabel Menggunakan Keyword var
	/* penulian variabel
	   var <nama-variabel> <tipe-data>
	   var <nama-variabel> <tipe-data> = <nilai>
	*/

	/*
	   Perhatikan bagian "halo %s %s!\n", karakter %s di situ akan diganti dengan data string yang berada di parameter ke-2, ke-3, dan seterusnya.

	   Contoh lain, ketiga baris kode berikut ini akan menghasilkan output yang sama, meskipun cara penulisannya berbeda.

	   fmt.Printf("halo john wick!\n")
	   fmt.Printf("halo %s %s!\n", firstName, lastName)
	   fmt.Println("halo", firstName, lastName + "!")
	*/

	// 3. Deklarasi Variabel Tanpa Tipe Data

	// var firstName string = "john"
	// lastName := "wick"

	// fmt.Printf("halo %s %s!\n", firstName, lastName)

	// menggunakan var, tanpa tipe data, menggunakan perantara "="
	// var firstName = "john"

	// tanpa var, tanpa tipe data, menggunakan perantara ":="
	// lastName := "wick"

	// 4. Deklarasi Multi Variabel
	/*
	   var first, second, third string
	   first, second, third = "satu", "dua", "tiga"
	*/

	// var fourth, fifth, sixth string = "empat", "lima", "enam"

	// seventh, eight, ninth := "tujuh", "delapan", "sembilan"

	// one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"

	// 5. Variabel Underscore _
	/*
	   Underscore (_) adalah reserved variable yang bisa dimanfaatkan untuk menampung nilai yang tidak dipakai. Bisa dibilang variabel ini merupakan keranjang sampah.

	   _ = "belajar Golang"
	   _ = "Golang itu mudah"
	   name, _ := "john", "wick"
	*/

	// 6.  Deklarasi Variabel Menggunakan Keyword new
	// Fungsi new() digunakan untuk membuat variabel pointer dengan tipe data tertentu. Nilai data default-nya akan menyesuaikan tipe datanya.
	/*
	   name := new(string)

	   fmt.Println(name)   // 0x20818a220
	   fmt.Println(*name)  // ""
	*/

	/*
	   Variabel name menampung data bertipe pointer string. Jika ditampilkan yang muncul bukanlah nilainya melainkan alamat memori nilai tersebut (dalam bentuk notasi heksadesimal). Untuk menampilkan nilai aslinya, variabel tersebut perlu di-dereference terlebih dahulu, caranya dengan menuliskan tanda asterisk (*) sebelum nama variabel.
	*/

	// 7. Deklarasi Variabel Menggunakan Keyword make
	// Fungsi make() ini hanya bisa digunakan untuk pembuatan beberapa jenis variabel saja, yaitu:

	// channel
	// slice
	// map
}
