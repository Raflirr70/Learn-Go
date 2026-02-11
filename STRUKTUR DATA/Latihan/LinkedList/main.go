package main

import "fmt"

type Gerbong struct { //pembuatan struct Gerbong
	nama  string
	bobot int
	next  *Gerbong
}

var head *Gerbong //pembuatan head dari gerbong, yang bertujuan untuk penunjuk linklist paling depan

// pembuatan function untuk mengecek apakah linklist tersebut kosong atau tidak
func IsEmpty() bool {
	return head == nil
}

//misalkan kita mau membuat function yang dapat menambahkan gerbong dari depan dan belakang

// tambah dari depan
func CreateGerbongDepan(nama string, bobot int) {
	baru := new(Gerbong)
	baru.nama = nama
	baru.bobot = bobot

	if IsEmpty() {
		head = baru
	} else {
		baru.next = head
		head = baru
	}
}

// tambah dari belakang
func CreateGerbongBelakang(nama string, bobot int) {
	baru := new(Gerbong)
	baru.nama = nama
	baru.bobot = bobot

	if IsEmpty() {
		head = baru
	} else {
		bantu := head
		for bantu.next != nil {
			bantu = bantu.next
		}
		bantu.next = baru
	}
}

// funtion untuk menghapus Gerbong di Depan
func DeleteGerbongDepan() {
	if IsEmpty() {
		fmt.Print("Gerbong Masih Kosong")
		return
	}
	head = head.next
}

// function untuk menghapus gerbong di belakang
func DeleteGerbongBelakang() {
	if IsEmpty() {
		fmt.Print("Gerbong Masih Kosong")
		return
	}
	if head.next == nil {
		head = nil
		return
	}

	bantu := head
	for bantu.next.next != nil {
		bantu = bantu.next
	}
	bantu.next = nil
}

// function untuk menampilkan seluruh gerbong
func View() {
	bantu := head
	i := 1
	for bantu != nil {
		fmt.Printf("%d. Gerbong %s dengan bobot %d\n", i, bantu.nama, bantu.bobot)
		i++
		bantu = bantu.next
	}
}

func main() {
	CreateGerbongDepan("mamat", 2)
	CreateGerbongDepan("jajang", 5)
	CreateGerbongBelakang("tehyung", 10)
	DeleteGerbongBelakang()
	View()
}
