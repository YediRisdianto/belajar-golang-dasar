// 1. Operator Aritmatika
/*
Operator aritmatika adalah operator yang digunakan untuk operasi yang sifatnya perhitungan. Seperti +,-,*,/,%
*/

// package main

// import "fmt"

// func main() {
// 	var value = (((2+6)%3)*4 - 2) / 3
// 	fmt.Printf(value, "\n")
// }


// 2. Operator perbandingan
/*
Operator perbandingan digunakan untuk menentukan kebenaran suatu kondisi. Hasilnya berupa nilai boolean, true atau false.

==	apakah nilai kiri sama dengan nilai kanan
!=	apakah nilai kiri tidak sama dengan nilai kanan
<	apakah nilai kiri lebih kecil daripada nilai kanan
<=	apakah nilai kiri lebih kecil atau sama dengan nilai kanan
>	apakah nilai kiri lebih besar dari nilai kanan
>=	apakah nilai kiri lebih besar atau sama dengan nilai kanan

*/

// package main

// import "fmt"

// func main() {
// 	var value = (((2+6)%3)*4 - 2) / 3
// 	var isEqual = (value == 2)

// 	fmt.Printf("nilai %d (%t) \n", value, isEqual)
// }


// 3. Operator Logika
/*
Operator ini digunakan untuk mencari benar tidaknya kombinasi data bertipe bool (bisa berupa variabel bertipe bool, atau hasil dari operator perbandingan).

Beberapa operator logika standar yang bisa digunakan:

&&	kiri dan kanan (and)
||	kiri atau kanan (or)
!	negasi / nilai kebalikan (not)
*/

package main

import "fmt"

func main() {
	var left = false
	var right = true

	var leftAndRight = left && right
	fmt.Printf("left && right \t(%t) \n", leftAndRight)

	var leftOrRight = left || right
	fmt.Printf("left || right \t(%t) \n", leftOrRight)

	var leftReverse = !left
	fmt.Printf("!left \t\t(%t) \n", leftReverse)
}