package render

import (
	"html/template"
	"log"
	"net/http"
)

func Render(w http.ResponseWriter, name string) {
	t, err := template.ParseFiles("./templates/" + name)
	if err != nil {
		log.Fatal("テンプレートファイルが正しく読み取りません")
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Fatal("テンプレートファイルレンダリング失敗しました。")
	}
}
