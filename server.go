package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"
)

// https://mholt.github.io/json-to-go/ で自動生成しています・
type AutoGenerated struct {
	Horoscope struct {
		Two018053 []struct {
			Content string      `json:"content"`
			Item    string      `json:"item"`
			Money   int         `json:"money"`
			Total   int         `json:"total"`
			Job     int         `json:"job"`
			Color   string      `json:"color"`
			Day     json.Number `json:"day"` //Stringとintの可能性があります
			Love    int         `json:"love"`
			Rank    int         `json:"rank"`
			Sign    string      `json:"sign"`
		} `json:"2018/05/11"` //指定する必要があります
	} `json:"horoscope"`
}

// 自動生成終わり

// 表示するパラメータ
type Page struct {
	Day     string
	Content string
	Rank    int
	Sign    string
}

func fortuneTellingAPI() Page {
	day := "2018/5/11"                                        //TODO 日付をPostあるいは今日の日付で指定する。
	url := "http://api.jugemkey.jp/api/horoscope/free/" + day //TODO 日付の指定

	resp, _ := http.Get(url) //TODO エラーハンドリグ
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body) //TODO エラーハンドリング
	jsonBytes := ([]byte)(string(byteArray))
	data := new(AutoGenerated)
	if err := json.Unmarshal(jsonBytes, data); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return Page{"err", "err", 0, "err"} //TODO いい文言に変える
	}

	//TODO 星座はPostされたものを指定する
	page := Page{day, data.Horoscope.Two018053[0].Content, data.Horoscope.Two018053[0].Rank, data.Horoscope.Two018053[0].Sign}
	return page
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	page := fortuneTellingAPI()
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
	http.HandleFunc("/", viewHandler)
	http.ListenAndServe(":8080", nil)
}
