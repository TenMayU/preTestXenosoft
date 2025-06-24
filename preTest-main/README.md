# preTest
preTest Xenosoft

วิธี Run 
1.ติดตั้ง Docker desktop(Window)
2.เข้าไปยัง Root ของ Project 
3.ใช้คำสั่ง docker-compose up --build

Frontend Path: http://localhost:3000/
Backend Path   http://localhost:8080/api


หากไม่ได้ใช้ Docker

1.cd ไปยัง frontend แล้วจึงติดตั้ง dependency ด้วย npm i
2.ใช้คำสั่ง npm run dev
3.cd ไปยัง backend แล้วจึงติดตั้ง dependency ด้วย go mod tidy
4.ใช้คำสั่ง go run src/cmd/server/server.go

Frontend Path: http://localhost:3000/
Backend Path   http://localhost:8080/api
