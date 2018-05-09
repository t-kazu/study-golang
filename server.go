package main

import (
	"io/ioutil"
	"net/http"
	"text/template"
)

type Page struct {
	Title string
	Count int
}

func fortuneTellingAPI() string {
	url := "http://api.jugemkey.jp/api/horoscope/free/2018/5/3"

	resp, _ := http.Get(url) //TODO エラーハンドリグ
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body) //TODO エラーハンドリング
	return string(byteArray)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	json := fortuneTellingAPI()
	page := Page{json, 1}
	tmpl, err := template.ParseFiles("layout.html") // ParseFilesを使う
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, page)
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", viewHandler) // hello
	http.ListenAndServe(":8080", nil)
}
