package main

import "fmt"

func IfFlow() {

	kondisi := true
	//pembuatan perkondisian if sederhana
	if kondisi {
		fmt.Print("Masuk Kedalam control flow if")
	}

	//pembuatan perkondisian if else sederhana
	if kondisi {
		fmt.Print("Kondisi 1 Terpenuhi")
	} else {
		fmt.Print("Kondisi 1 tidak terpenuhi")
	}

	//pembuatan perkondisian if | else if | else
	nilai := 10
	if nilai > 5 {
		fmt.Print("Kondisi 1 terpenuhi")
	} else if nilai == 5 {
		fmt.Print("Kondisi 2 Terpenuhi")
	} else {
		fmt.Print("Kondisi 1 dan kondisi 2 tidak terpenuhi")
	}
}

func SwitchFlow() {
	kondisi := 2

	//pembuatan control flow menggunakan switch (switch di go tidak memerlukan break untuk menghentikan control flow)
	//sehingga jika case 1 terpenuhi maka control flow switch akan berakhir
	switch kondisi {
	case 1:
		fmt.Print("Case 1 Terpenuhi")
	case 2:
		fmt.Print("Case 2 Terpenuhi")
	case 3:
		fmt.Print("Case 3 Terpenuhi")
	default:
		fmt.Print("Case 1,2,3 tidak Terpenuhi")
	}

	//pembuatan control flow menggunakan switch dengan menggunakan fallthrough
	//sehingga jika case 1 terpenuhi maka program akan masuk ke case 2 dan seterusnya
	switch kondisi {
	case 1:
		fmt.Print("Case 1 Terpenuhi")
		fallthrough
	case 2:
		fmt.Print("Case 2 Terpenuhi")
		fallthrough
	case 3:
		fmt.Print("Case 3 Terpenuhi")
		fallthrough
	default:
	}
	fmt.Print("Case 1,2,3 tidak Terpenuhi")
}

func main() {

}
