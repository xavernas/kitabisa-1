package main

import (
	"fmt"
	"strconv"
)

// Batasan: Tidak ada kurung
func stringMath(angka string) int {
	// Kalau angka biasa yang diinput, tidak perlu diproses
	if converted_number, err := strconv.Atoi(angka); err == nil {
		fmt.Println(true)
		return converted_number
	}
	index := 0
	operator := ""
	angkastring := ""
	// angka minus tidak akan diproses
	var hasil_akhir int
	var indexsementara int
	for index < len(angka) {
		indexsementara = index
		if int(angka[indexsementara]) > 47 && int(angka[indexsementara]) < 58 {
			for indexsementara < len(angka) && int(angka[indexsementara]) > 47 && int(angka[indexsementara]) < 58 {
				angkastring += string(angka[indexsementara])
				indexsementara += 1
			}
			// Untuk assign angka pertama
			if hasil_akhir == 0 && len(string(hasil_akhir)) == 1 {
				if converted_string, err := strconv.Atoi(angkastring); err == nil {
					hasil_akhir = converted_string
				}
				angkastring = ""
			}
		} else if int(angka[index]) == 43 {
			operator += "+"
		} else if int(angka[index]) == 45 {
			operator += "-"
		}
		if angkastring != "" && operator != "" {
			if converted_string, err := strconv.Atoi(angkastring); err == nil {
				fmt.Println(angkastring, " ", operator)
				angkaconverted := converted_string
				if operator == "+" {
					hasil_akhir += angkaconverted
				} else if operator == "-" {
					hasil_akhir -= angkaconverted
				}
				angkastring = ""
				operator = ""
			}
		}
		if indexsementara != index {
			fmt.Println(index, " ke ", indexsementara)
			index = indexsementara + 1
		} else {
			fmt.Println(index)
			index += 1
		}
	}
	return hasil_akhir
}
func main() {
	// fmt.Println(stringMath("3252"))
	// fmt.Println()
	fmt.Println(stringMath("    3427 +      2 "))
	// fmt.Println()
	fmt.Println(stringMath("5 + 2 + 1"))  // 8
	fmt.Println(stringMath("7 - 4 + 21")) // 24
	// fmt.Println(stringMath("15 + 241 - 741 + -24")) // -509
	// fmt.Println(stringMath("274 - 3561 + 80 - 74")) // -3281
}
