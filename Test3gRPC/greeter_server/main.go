/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"regexp"
	"strings"

	pb "helloworld/Meetstock"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)
var BeefCounts map[string]int32

// server is used to implement GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements GreeterServer
func (s *server) Meetstock(ctx context.Context, in *pb.StockRequest) (*pb.StockReply, error) {
	log.Printf("Received: %v", in.GetName())
	data := &pb.StockReply{
		Beef: BeefCounts,
	}

	//return &pb.StockReply{Message: "Hello " + in.GetName()}, nil
	return data, nil
}

func updateBeefCounts(text string, words []string, patternRegex *regexp.Regexp) {
	BeefCountsInt := countWordsInText(text, patternRegex)
	BeefCounts = make(map[string]int32)
	for word, count := range BeefCountsInt {
		BeefCounts[word] = int32(count)
	}
}

func main() {
	//ข้อความที่ต้องการใช้
	url := "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text"

	// เรียกใช้งานฟังก์ชันเพื่อดึงข้อความจาก URL
	text, err := fetchTextFromURL(url)
	if err != nil {
		log.Fatal(err)
	}

	//ชนิดเนื้อใน .txt
	filePath := "default-custom.txt"
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("ไม่สามารถอ่านไฟล์ชนิดเนื้อได้:", err)
		return
	}

	// แยกข้อความเป็นคำๆ
	words := strings.Fields(string(data))
	//fmt.Println(words)

	// กำหนดค่าพรีคอมไพล์ของรูปแบบพาเทิร์นสำหรับการค้นหาคำ
	patternRegex := regexp.MustCompile(strings.Join(words, "|"))

	//update ชนิดเนื้อที่พบ
	updateBeefCounts(text, words, patternRegex)

	//แสดงชนิดเนื้อที่แยก
	fmt.Println("Count", BeefCounts)

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func fetchTextFromURL(url string) (string, error) {
	// สร้าง channel เพื่อรอรับข้อมูล
	resultChan := make(chan string)
	errChan := make(chan error)

	// เริ่ม Go routine เพื่อดึงข้อมูล URL
	go func() {
		resp, err := http.Get(url)
		if err != nil {
			errChan <- err
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			errChan <- err
			return
		}

		resultChan <- string(body)
	}()

	// รอรับข้อมูลจาก channel
	select {
	case result := <-resultChan:
		return result, nil
	case err := <-errChan:
		return "", fmt.Errorf("could not fetch data from URL: %v", err)
	}
}

func countWordsInText(text string, patternRegex *regexp.Regexp) map[string]int {
	BeefCounts := make(map[string]int)

	// ค้นหาคำที่ต้องการในข้อความโดยใช้พาเทิร์นที่พรีคอมไพล์ไว้
	matches := patternRegex.FindAllString(text, -1)

	// นับจำนวนครั้งที่คำปรากฏและเพิ่มใน map
	for _, match := range matches {
		BeefCounts[match]++
	}

	return BeefCounts
}
