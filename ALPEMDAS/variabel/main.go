package main

import "fmt"

func main() {
	//type data char di golang tidak ada, dan diganti dengan rune
	// var huruf rune

	var nama string           //deklarasi variabel harus disertakan dengan type
	nama = "Rafli Radiansyah" //pengisian variabel

	//deklarasi dan implementasi variabel
	var kelas string = "IF C" //cara 1 : deklarasi dengan tipe jelas
	var npm = "01928321"      //cara 2 : deklarasi tanpa tipe
	umur := 17                //cara 3 : deklarasi langsung
	ipk := 3.92

	fmt.Printf("%s merupakan nama saya \n", nama)
	fmt.Printf("Saya ada di kelas %s \n", kelas)
	fmt.Printf("npm saya %s \n", npm)
	fmt.Printf("dan umur saya masih %d tahun\n", umur)
	fmt.Printf("Sedangkan IPK saya %0.2f", ipk)

}
