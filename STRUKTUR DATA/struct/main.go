package main

import "fmt"

//deklarasi sebuah struct
type Node struct {
	name  string
	kelas int
	npm   string
}

//Pembuatan Objek Struct
func Deklarasi() {
	var obj Node       //deklarasi objek dari Node
	obj.name = "mamat" //pengisian atribut dari objek obj

	obj1 := obj           //deklarasi objek yang diisi oleh data objek obj
	obj1.name = "tehyung" //jika mengubah nama menggunakan obj1, maka data dari obj tidak akan berubah
	fmt.Print(obj.name)

	var objptr *Node     //deklarasi pointer yang menunjuk alamat dari objek Node
	objptr1 := &obj1     //deklarasi pointer yang menunjuk alamat dari objek obj1
	objptr3 := new(Node) //deklarasi pointer objek yang diisi alamat dari Objek Node yang berdiri sendiri

	objptr1.name = "jajang" //jika mengubah data menggunakan pointer maka objek yang ditunjuk oleh pointer tersebut juga akan berubah datanya

	print(objptr, objptr3) //agar tidak erro

}

func main() {
	Deklarasi()
}
