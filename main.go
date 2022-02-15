package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprint(w, "Hello golang!")
		if err != nil {
			log.Println("レスポンス書き込みエラー", err)
		}
		log.Printf("レスポンス書き込みサイズ: %d\n", n)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("サーバー起動失敗しました", err)
	}
}
