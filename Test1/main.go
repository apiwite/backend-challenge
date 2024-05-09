package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	//data_test := [][]int{{59}, {73, 41}, {52, 40, 53}, {26, 53, 6, 34}}

	url := "https://raw.githubusercontent.com/7-solutions/backend-challenge/main/files/hard.json"

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("ไม่สามารถเรียกข้อมูล:", err)
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("ไม่สามารถอ่านข้อมูล:", err)
		return
	}

	var data [][]int
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("ไม่สามารถแปลงข้อมูล JSON:", err)
		return
	}

	maxSum := findMaxSum(data)
	fmt.Println("ผลรวมที่มากที่สุด:", maxSum)
}

func findMaxSum(data [][]int) int {
	n := len(data)
	dp := make([]int, n)

	// กำหนดค่าใน dp ให้เป็นค่าของโหนดในระดับล่างสุดของต้นไม้
	copy(dp, data[n-1])

	// หาผลรวมที่มากที่สุดของแต่ละโหนดในแต่ละระดับของต้นไม้โดยจะเรียงจากล่างขึ้นบน
	for i := n - 2; i >= 0; i-- {
		for j := 0; j < len(data[i]); j++ {
			//fmt.Println("<<", dp)
			//fmt.Println("=", data[i][j], "+", dp[j], ",", dp[j+1])
			dp[j] = data[i][j] + max(dp[j], dp[j+1])
			//fmt.Println("=>", dp[j])

		}
	}
	//fmt.Println(dp)

	return dp[0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
