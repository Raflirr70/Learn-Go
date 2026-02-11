package main

import "fmt"

//di bahasa Go tidak ada perulangan menggunakan while
//tapi perulangan for bisa digunakan mirip seperti perulangan while
//dan jika ingin menggunakan perulangan do while bisa diakali dengan menggunakan perulangan for infinit dan perkondisian untuk menghentikannya

//perulangan biasa
func ForLoop() {
	//for <inisialisasi> ; <terminasi> ; <iterasi> {statement}
	for a := 1; a <= 10; a++ { //membuat perulangan dari 1 - 10
		fmt.Print(a, " ")
	}
	for a := 10; a > 0; a-- { //membuat perulangan dari 10 - 1
		fmt.Print(a, " ")
	}
}

//perulangan mirip while
func WhileLoop() {
	i := 5       //inisiaslisasi
	for i == 1 { //terminasi
		//statement
		i-- //iterasi
	}
}

//perulangan mirip do while
func DoWhileLoop() {
	i := 5 //inisialisi
	for {
		//statement
		i++
		if i == 10 {
			break
		}
	}
}

func main() {

}
