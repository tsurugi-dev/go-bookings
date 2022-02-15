package main

import (
	"html/template"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	render(w, "home.page.tmpl")
}

func about(w http.ResponseWriter, r *http.Request) {
	render(w, "about.page.tmpl")
}

func render(w http.ResponseWriter, name string) {
	t, err := template.ParseFiles("./templates/" + name)
	if err != nil {
		log.Fatal("テンプレートファイルが正しく読み取りません")
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Fatal("テンプレートファイルレンダリング失敗しました。")
	}
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("サーバー起動失敗しました", err)
	}
}
