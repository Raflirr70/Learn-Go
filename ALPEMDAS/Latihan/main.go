package main

import "fmt"

/*
	buatlah program untuk membuat sebuah bintang
	*
	* *
	* * *
	* * * *
*/
func Bintang1() {
	for a := 1; a < 5; a++ {
		for b := 1; b <= a; b++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

/*
buatlah program untuk membuat sebuah bintang
* * * *
* * *
* *
*
*/
func Bintang2() {
	for a := 1; a < 5; a++ {
		for b := 5; b > a; b-- {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

/*
buatlah program untuk membuat sebuah bintang
* * * * * * * *
* * *     * * *
* *         * *
*             *
*             *
* *         * *
* * *     * * *
* * * * * * * *
*/
func Bintang3() {
	for a := 1; a < 5; a++ {
		for b := 5; b > a; b-- {
			fmt.Print("* ")
		}
		for b := 1; b < a; b++ {
			fmt.Print("  ")
		}
		for b := 1; b < a; b++ {
			fmt.Print("  ")
		}
		for b := 5; b > a; b-- {
			fmt.Print("* ")
		}
		fmt.Println()
	}
	for a := 1; a < 5; a++ {
		for b := 1; b <= a; b++ {
			fmt.Print("* ")
		}
		for b := 4; b > a; b-- {
			fmt.Print("  ")
		}
		for b := 4; b > a; b-- {
			fmt.Print("  ")
		}
		for b := 1; b <= a; b++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func main() {
	Bintang3()
}
