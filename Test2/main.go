package main

import (
	"fmt"
	"strconv"
)

func main() {
	var symbols string
	fmt.Print("input : ")
	fmt.Scanln(&symbols)
	//symbols := "LLRR="
	result := make([]int, len(symbols)+1)

	// วนลูปผ่านสัญลักษณ์ตัวอักษร
	loop := 0
SyntaxLoop: //หากรูปแบบข้อมูลผิดจะ break
	for loop < len(symbols) {

	OuterLoop: //break เพื่อเรียงเงื่อนไขใหม่
		for i, s := range symbols {
			switch s {
			case 'L':
				if result[i] <= result[i+1] {
					result[i]++
					loop = 0
					break OuterLoop
				} else {
					loop++
				}
			case 'R':
				if result[i+1] <= result[i] {
					result[i+1]++
					loop = 0
					break OuterLoop
				} else {
					loop++
				}
			case '=':
				if result[i+1] < result[i] {
					result[i+1] = result[i]
					loop = 0
					break OuterLoop
				} else if result[i+1] > result[i] {
					result[i] = result[i+1]
					loop = 0
					break OuterLoop
				} else {
					loop++
				}
			default:
				fmt.Println("Error Syntax")
				break SyntaxLoop
			}

		}
	}
	var str1 string
	// วนลูปผ่านอาร์เรย์และแปลงแต่ละตัวเลขเป็นสตริง
	for _, num := range result {
		// แปลงตัวเลขเป็นสตริง
		str := strconv.Itoa(num)
		str1 += str
	}
	fmt.Println(str1)
}
