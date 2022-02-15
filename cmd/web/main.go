package main

import (
	"github.com/tsurugi-dev/go-bookings/pkg/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/about", handler.About)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("サーバー起動失敗しました", err)
	}
}
