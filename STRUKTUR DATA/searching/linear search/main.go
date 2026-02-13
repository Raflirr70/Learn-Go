package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// array biasa
func low() {
	//input
	scn := bufio.NewScanner(os.Stdin)
	arrs := []int{6, 2, 1, 10, 7, 9}

	for a := 0; a < len(arrs); a++ {
		fmt.Print(arrs[a], " ")
	}
	fmt.Print("\n\n")

	fmt.Print("Masukan Angka Yang Dicari : ")
	scn.Scan()
	input := scn.Text()

	// konversi string ke int
	data, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Input harus berupa angka!")
		return
	}

	ketemu := false
	for a := 0; a < len(arrs); a++ {
		if data == arrs[a] {
			ketemu = true
			break
		}
	}
	fmt.Print("\n\n")
	fmt.Printf("Data : %d", data)
	if ketemu {
		fmt.Print(",ditemukan")
	} else {
		fmt.Print(", Tidak ditemukan")
	}
}

//linked list

// persiapkan structnya
type Node struct {
	data int
	next *Node
}

var head *Node

func create(data int) {
	baru := new(Node)
	baru.data = data
	if head == nil {
		head = baru
		return
	}
	baru.next = head
	head = baru
}

func searchLL() {
	create(1)
	create(10)
	create(9)
	create(5)
	create(7)
	create(2)

	bantu := head
	for bantu != nil {
		fmt.Print(bantu.data, " ")
		bantu = bantu.next
	}
	fmt.Print("\n Cari Data : ")

	read := bufio.NewScanner(os.Stdin)
	read.Scan()
	input := read.Text()

	data, err := strconv.Atoi(input)
	if err != nil {
		fmt.Print("Masukan integer")
		return
	}

	bantu = head
	ketemu := false
	for bantu != nil {
		if bantu.data == data {
			ketemu = true
		}
		bantu = bantu.next
	}
	if ketemu {
		fmt.Print("Data Ditemukan")
		return
	}
	fmt.Print("Data Tidak Ditemukan")
}

func main() {
	searchLL()
}
