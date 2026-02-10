package main

import "fmt"

//pembuatan function dengan nilai balik berupa integer
func Tambah(a, b int) int {
	return a + b
}

//pembuatan function dengan parameter yang dinamis
func TambahT[T int | float32 | float64](a, b T) T {
	return a + b
}

//pembuatan function yang memiliki lebih dari 1 nilai balik
func NilaiAkhir(mtk, penjas, kalkulus float32) (float32, string) {
	total := ((mtk + penjas + kalkulus) / 3)
	var ip string

	if total > 90 {
		ip = "A"
	} else if total > 80 {
		ip = "B"
	} else if total > 70 {
		ip = "BC"
	} else if total > 60 {
		ip = "C"
	} else {
		ip = "E"
	}
	return total, ip
}

func main() {
	//implementasi function yang memiliki nilai return
	c, x := NilaiAkhir(12, 53, 12)
	fmt.Printf("Anda Mendapatkan Nilai Akhir : %0.2f \nAnda Mendapatkan : %s", c, x)
}
