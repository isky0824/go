package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("文件服务启动...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		handler := http.FileServer(http.Dir("./"))
		handler.ServeHTTP(w, r)
	})

	err := http.ListenAndServe(":8020", nil)
	fmt.Println(err)
}
