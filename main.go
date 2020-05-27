package main

import (
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(PlayerServer)
	if err := http.ListenAndServe(":5000", handler); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}

// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func myWeb(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "这是一个开始")
// }

// func main() {
// 	http.HandleFunc("/", myWeb)

// 	fmt.Println("服务器即将开启，访问地址 http://localhost:8080")

// 	err := http.ListenAndServe(":8080", nil)
// 	if err != nil {
// 		fmt.Println("服务器开启错误: ", err)
// 	}
// }
