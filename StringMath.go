package main

import (
	"fmt"
	"strconv"
)

func stringMath(angka string) []int {
	// Diconvert ke string
	var slice_string []string
	for _, value := range angka {
		slice_string = append(slice_string, string(value))
	}
	fmt.Print(slice_string)
	// Kalau tidak ada yang sama, diconvert ke integer
	var slice_integer []int
	// var inget_posisi []int
	for _, nilai := range slice_string {
		for _, eval := range slice_string {
			fmt.Print(nilai, eval)
			if i, err := strconv.Atoi(nilai); err == nil {
				slice_integer = append(slice_integer, i)
			}
		}
	}

	fruit := []int{
		20,
		5,
		5, // comma added
	}
	return fruit
}
func main() {
	fmt.Println(stringMath("5 + 2 + 1"))            // 8
	fmt.Println(stringMath("7 - 4 + 21"))           // 24
	fmt.Println(stringMath("15 + 241 - 741 + -24")) // 8
	fmt.Println(stringMath("5 + 2 + 1"))            // 8
}
