// Perulangan
/*
Perulangan
Perulangan adalah proses mengulang-ulang eksekusi blok kode tanpa henti, selama kondisi yang dijadikan acuan terpenuhi. Biasanya disiapkan variabel untuk iterasi atau variabel penanda kapan perulangan akan diberhentikan.

Di Go keyword perulangan hanya for saja, tetapi meski demikian, kemampuannya merupakan gabungan for, foreach, dan while ibarat bahasa pemrograman lain.
*/

// 1. Perulangan Menggunakan Keyword for
/*
Ada beberapa cara standar menggunakan for. Cara pertama dengan memasukkan variabel counter perulangan beserta kondisinya setelah keyword. Perhatikan dan praktekan kode berikut.
*/

// package main

// import "fmt"

// func main() {
// 	for i := 0; i < 5; i++ {
// 		fmt.Println("Angka", i)
// 	}
// }

/*
Perulangan di atas hanya akan berjalan ketika variabel i bernilai di bawah 5, dengan ketentuan setiap kali perulangan, nilai variabel i akan di-iterasi atau ditambahkan 1 (i++ artinya ditambah satu, sama seperti i = i + 1). Karena i pada awalnya bernilai 0, maka perulangan akan berlangsung 5 kali, yaitu ketika i bernilai 0, 1, 2, 3, dan 4.
*/

// 2. Penggunaan Keyword for Dengan Argumen Hanya Kondisi
/*
Cara ke-2 adalah dengan menuliskan kondisi setelah keyword for (hanya kondisi). Deklarasi dan iterasi variabel counter tidak dituliskan setelah keyword, hanya kondisi perulangan saja. Konsepnya mirip seperti while milik bahasa pemrograman lain.

Kode berikut adalah contoh for dengan argumen hanya kondisi (seperti if), output yang dihasilkan sama seperti penerapan for cara pertama.
*/

// package main

// import "fmt"

// func main() {
// 	var i = 0

// 	for i < 5 {
// 		fmt.Println("Angka", i)
// 		i++
// 	}
// }

// 3. Penggunaan Keyword for Tanpa Argumen
/*
Cara ke-3 adalah for ditulis tanpa kondisi. Dengan ini akan dihasilkan perulangan tanpa henti (sama dengan for true). Pemberhentian perulangan dilakukan dengan menggunakan keyword break.
*/

// package main

// import "fmt"

// func main() {
// 	var i = 0

// 	for {
// 		fmt.Println("Angka", i)

// 		i++
// 		if i == 5 {
// 			break
// 		}
// 	}
// }

/*
Dalam perulangan tanpa henti di atas, variabel i yang nilai awalnya 0 di-inkrementasi. Ketika nilai i sudah mencapai 5, keyword break digunakan, dan perulangan akan berhenti.
*/

// 4. Penggunaan Keyword for - range
/*
Cara ke-4 adalah perulangan dengan menggunakan kombinasi keyword for dan range. Cara ini biasa digunakan untuk me-looping data gabungan (misalnya string, array, slice, map).
*/

// package main

// import "fmt"

// func main() {
// 	var xs = "123" // string

// 	for i, v := range xs {
// 		fmt.Println("Index=", i, "Value=", v)
// 	}

// 	var ys = [5]int{10, 20, 30, 40, 50} // array
// 	for _, v := range ys {
// 		fmt.Println("Value=", v)
// 	}

// 	var zs = ys[0:2] // slice
// 	for _, v := range zs {
// 		fmt.Println("Value=", v)
// 	}

// 	var kvs = map[byte]int{'a': 0, 'b': 1, 'c': 2} // map
// 	for k, v := range kvs {
// 		fmt.Println("Key=", k, "Value=", v)
// 	}

// 	// boleh juga baik k dan atau v nya diabaikan
// 	for range kvs {
// 		fmt.Println("Done")
// 	}

// 	// selain itu, bisa juga dengan cukup menentukan nilai numerik perulangan
// 	for i := range 5 {
// 		fmt.Print(i) // 01234
// 	}
// }

// 5. Penggunaan Keyword break & continue
/*
Keyword break digunakan untuk menghentikan secara paksa sebuah perulangan, sedangkan continue dipakai untuk memaksa maju ke perulangan berikutnya.

Berikut merupakan contoh penerapan continue dan break. Kedua keyword tersebut dimanfaatkan untuk menampilkan angka genap berurutan yang lebih besar dari 0 dan kurang dari atau sama dengan 8.
*/

// package main

// import "fmt"

// func main() {
// 	for i := 1; i <= 10; i++ {
// 		if i % 2 == 0 {
// 			continue
// 		}

// 		if i > 9 {
// 			break
// 		}

// 		fmt.Println("Angka ganjil:", i)
// 	}

// 	for i := 1; i <= 10; i++ {
// 		if i % 2 == 1 {
// 			continue
// 		}

// 		if i > 8 {
// 			break
// 		}

// 		fmt.Println("Angka genap:", i)
// 	}
// }

// 6.peluang bersarang
/*
Tak hanya seleksi kondisi yang bisa bersarang, perulangan juga bisa. Cara pengaplikasiannya kurang lebih sama, tinggal tulis blok statement perulangan di dalam perulangan.
*/

// package main

// import "fmt"

// func main() {

// 	// membuat bintang
// 	for i := 5; i >= 2; i-- {
// 		for j := i; j <= 5; j++ {
// 			fmt.Print("* ")
// 		}
// 		fmt.Println()
// 	}

// 	for i := 1; i <= 5; i++ {
// 		for j := i; j <= 5; j++ {
// 			fmt.Print("* ")
// 		}
// 		fmt.Println()
// 	}

// 	// for i := 0; i < 5; i++ {
// 	// 	for j := i; j < 5; j++ {
// 	// 		fmt.Print(j, " ")
// 	// 	}
	
// 	// 	fmt.Println()
// 	// }

// }

/*
Pada kode di atas, untuk pertama kalinya fungsi fmt.Println() dipanggil tanpa disisipkan parameter. Cara seperti ini bisa digunakan untuk menampilkan baris baru. Kegunaannya sama seperti output dari statement fmt.Print("\n").
*/


// 7. Pemanfaatan Label Dalam Perulangan
/*
Di perulangan bersarang, break dan continue akan berlaku pada blok perulangan di mana ia digunakan saja. Ada cara agar kedua keyword ini bisa tertuju pada perulangan terluar atau perulangan tertentu, yaitu dengan memanfaatkan teknik pemberian label.

Program untuk memunculkan matriks berikut merupakan contoh penerapan label perulangan.
*/

package main

import "fmt"

func main() {
	outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 4 {
				break outerLoop
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}
}

/*
Tepat sebelum keyword for terluar, terdapat baris kode outerLoop:. Maksud dari kode tersebut adalah disiapkan sebuah label bernama outerLoop untuk for di bawahnya. Nama label bisa diganti dengan nama lain (dan harus diakhiri dengan tanda titik dua atau colon (:) ).

Pada for bagian dalam, terdapat seleksi kondisi untuk pengecekan nilai i. Ketika nilai tersebut sama dengan 3, maka break dipanggil dengan target adalah perulangan yang dilabeli outerLoop, perulangan tersebut akan dihentikan.
*/