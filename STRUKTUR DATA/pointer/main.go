package main

import "fmt"

//deklarasi pointer
func Deklarasi() {
	data := "Rafli"

	var ptr *string //deklarasi pointer tanpa mengisi nilai

	ptr0 := new(string) //deklarasi pointer yang diisi alamat yang berdiri sendiri

	//deklarasi pointer dan mengisi nilai dengan alamat dari variabel data
	var ptr1 *string = &data
	var ptr2 = &data
	ptr3 := &data

	fmt.Print(ptr, ptr0, ptr1, ptr2, ptr3)

	//penggunaan pointer untuk merubah data
	ptr4 := &data
	*ptr4 = "Radiansyah"

	fmt.Print(data)

	//merubah nilai pointer
	ptr5 := &data
	data2 := "tehyung"

	ptr5 = &data2

	print(ptr5)
}

func main() {
	Deklarasi()
}
