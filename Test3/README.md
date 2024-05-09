3. พาย ไฟ ได - Pie Fire Dire
   
`go run main.go` server จะทำการ start ที่ `port:8080`

ให้ทดสอบโดยการเรียก api ไปที่ `http://localhost:8080/beef/summary`

การทำงานของโปรแกรมคร่าวๆ เมื่อเรียก api /beef/summary
  
  1.จะเป็นการดึงข้อมูลจาก link ที่ให้มา
  
  2.การดึงชนิดเนื้อจาก default-custom.txt มาเพื่อคัดชนิดเนื้อในข้อความ
  
  3.ทำการคัดข้อความ จัดjson และส่งออกข้อมูล