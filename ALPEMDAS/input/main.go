package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func InputFmtScan() {
	//pengimputan menggunakan fmt.Scan()
	//dimana penginputannya berhenti di spasi
	var name string
	fmt.Print("Masukan Nama Anda : ")
	fmt.Scan(&name)
	fmt.Printf("Data yang dimasukan adalah %s", name)
}

func InputFmtScanln() {
	//pengimputan mengguakan fmt.Scanln()
	//dimana akhir dari penginputanya berhenti di spasi dan enter
	var name string
	fmt.Print("Masukan Nama Anda : ")
	fmt.Scanln(&name)
	fmt.Printf("Data yang dimasukan adalah %s", name)
}

func InputBufioNewReader() {
	//penginputan menggunakan bufio.NewReader(os.Stdin)
	//dengan menggukanan metode ini, inputan yang memiliki spasi dapat terbaca
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukan Nama Anda : ")

	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Printf("Data yang dimasukan adalah %s", name)
}

func InputBufioScanner() {
	//penginputan menggunakan bufio.NewScanner()
	//dengan menggukaan cara ini, spasi dapat dibaca
	reader := bufio.NewScanner(os.Stdin)

	fmt.Print("Masukan Nama Anda : ")
	reader.Scan()
	name := reader.Text()
	fmt.Printf("Data yang dimasukan adalah %s", name)
}

func main() {
}
