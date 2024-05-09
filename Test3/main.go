package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	fmt.Println("Server start at", "http://localhost:8080")
	fmt.Println("http://localhost:8080/beef/summary")
	http.HandleFunc("/beef/summary", beefSummaryHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}

}

func beefSummaryHandler(w http.ResponseWriter, r *http.Request) {
	//ข้อความที่ต้องการใช้
	url := "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text"

	// เรียกใช้งานฟังก์ชันเพื่อดึงข้อความจาก URL
	text, err := fetchTextFromURL(url)
	if err != nil {
		log.Fatal(err)
	}

	filePath := "default-custom.txt"
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("ไม่สามารถอ่านไฟล์ชนิดเนื้อได้:", err)
		return
	}

	// แยกข้อความเป็นคำๆ
	words := strings.Fields(string(data))

	//update ชนิดเนื้อที่พบ
	BeefCounts := countWordsInText(text, words)

	BeefSum := map[string]map[string]int{
		"beef": BeefCounts,
	}

	// แปลงข้อมูล JSON เป็น bytes
	jsonData, err := json.Marshal(BeefSum)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("GET : /beef/summary")
	// ตั้งค่า Content-Type เป็น application/json
	w.Header().Set("Content-Type", "application/json")
	// ส่ง JSON กลับไปยัง client
	w.Write(jsonData)
}

func fetchTextFromURL(url string) (string, error) {
	// สร้าง HTTP request เพื่อเรียกข้อมูลจาก URL
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("could not get data from URL: %v", err)
	}
	defer resp.Body.Close()

	// อ่านข้อมูลที่ได้รับจาก response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("could not read response body: %v", err)
	}

	return string(body), nil
}

func countWordsInText(text string, words []string) map[string]int {
	BeefCounts := make(map[string]int)

	// ตรวจสอบแต่ละคำใน array
	for _, word := range words {
		// นับจำนวนครั้งที่คำปรากฏในข้อความ
		count := strings.Count(text, word)
		// เพิ่มใน map แต่ถ้าจำนวนครั้งเป็น 0 จะไม่เพิ่ม
		if count > 0 {
			BeefCounts[word] = count
		}
	}
	return BeefCounts
}
