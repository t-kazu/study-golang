# Fortune Go -星座占い管理ツール-

## 概要
[Web ad Fortune](http://jugemkey.jp/api/waf/api_free.php)から星座占いを取得し、保存や共有を行います

## 環境
- 検証は以下の環境で行っております
  - OS:Ubuntu18.04
  - version：1.10
- goのバージョン確認
```
$ go versions
go version go1.10 {OS}
```

## 実行方法
```
$ go run server.go
```
現在指定しているポートは8080番です

## コード整形
```
$ go fmt {ファイル名}
```

## TODO
- 日にちの指定（Post or 今日の日付）
- 星座のPOST（インデックスと星座が対応しています）
- 占いの共有（DBの保存。今日の出来事をコメントして共有しよう）
- 見た目をもっとキレイに  
etc
