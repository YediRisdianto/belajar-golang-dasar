// 1. Tipe Data Numerik Non-Desimal
/*
Tipe data numerik non-desimal atau non floating point di Go ada beberapa jenis. Secara umum ada 2 tipe data kategori ini yang perlu diketahui.

uint, tipe data untuk bilangan cacah (bilangan positif).
int, tipe data untuk bilangan bulat (bilangan negatif dan positif).
*/

// package main

// import "fmt"

// func main() {
// 	var positiveNumber uint8 = 89
// 	var negativeNumber = -1243423644

// 	fmt.Printf("bilangan positif: %d\n", positiveNumber)
// 	fmt.Printf("bilangan negatif: %d\n", negativeNumber)
// }

// String format %d pada fmt.Printf() digunakan untuk memformat data numerik non-desimal.

// 2. Tipe Data Numerik Desimal

/*
Tipe data numerik desimal yang perlu diketahui ada 2, float32 dan float64. Perbedaan kedua tipe data tersebut berada di lebar cakupan nilai desimal yang bisa ditampung. Untuk lebih jelasnya bisa merujuk ke spesifikasi IEEE-754 32-bit floating-point numbers.
*/

// package main

// import "fmt"

// func main() {
// 	var decimalNumber = 2.62

// 	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
// 	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)
// }

/*
String format %f digunakan untuk memformat data numerik desimal menjadi string. Digit desimal yang akan dihasilkan adalah 6 digit. Pada contoh di atas, hasil format variabel decimalNumber adalah 2.620000. Jumlah digit yang muncul bisa dikontrol menggunakan %.nf, tinggal ganti n dengan angka yang diinginkan. Contoh: %.3f maka akan menghasilkan 3 digit desimal, %.10f maka akan menghasilkan 10 digit desimal.
*/


// 3. Tipe Data bool (Boolean)
// Tipe data bool berisikan hanya 2 variansi nilai, true dan false. Tipe data ini biasa dimanfaatkan dalam seleksi kondisi dan perulangan

// package main

// import "fmt"

// func main() {
// 	var exist bool = true
//     fmt.Printf("exist? %t \n", exist)
// }

// Gunakan %t untuk memformat data bool menggunakan fungsi fmt.Printf().


// 4. Tipe Data string

// package main

// import "fmt"

// func main() {
// 	var message string = "Halo"
//     fmt.Printf("message: %s \n", message)
// }

/*
Selain menggunakan tanda quote, deklarasi string juga bisa dengan tanda grave accent/backticks (`), tanda ini terletak di sebelah kiri tombol 1. Keistimewaan string yang dideklarasikan menggunakan backtics adalah membuat semua karakter di dalamnya tidak di escape, termasuk \n, tanda petik dua dan tanda petik satu, baris baru, dan lainnya. Semua akan terdeteksi sebagai string.
*/

package main

import "fmt"

func main() {
	var message = `Nama saya "Yedi!".
    Salam kenal.
    Mari belajar "Golang".`

    fmt.Println(message)
}


// 5. Nilai nil & Zero Value

/*
nil bukan merupakan tipe data, melainkan sebuah nilai. Variabel yang isi nilainya nil berarti memiliki nilai kosong.

Semua tipe data yang sudah dibahas di atas memiliki zero value (nilai default tipe data). Artinya meskipun variabel dideklarasikan dengan tanpa nilai awal, tetap akan ada nilai default-nya.

Zero value dari string adalah "" (string kosong).
Zero value dari bool adalah false.
Zero value dari tipe numerik non-desimal adalah 0.
Zero value dari tipe numerik desimal adalah 0.0.
Selain tipe data yang disebutkan di atas, ada juga tipe data lain yang zero value-nya adalah nil. Nil merepresentasikan nilai kosong, benar-benar kosong. nil tidak bisa digunakan pada tipe data yang sudah dibahas di atas.

Beberapa tipe data yang bisa di-set nilainya dengan nil, di antaranya:

pointer
tipe data fungsi
slice
map
channel
interface kosong atau any (yang merupakan alias dari interface{})
Nantinya kita akan sering bertemu dengan nilai nil setelah masuk pada pembahasan-pembahasan tersebut.
*/