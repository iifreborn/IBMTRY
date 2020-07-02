package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func htmlHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html")
// 	html := `<!doctype html>
// 							<META http-equiv="Content-Type" content="text/html" charset="utf-8">
// 							<html lang="zh-CN">
// 							<head>
// 								<title>Golang</title>
// 								<meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=0;" />
// 							</head>
// 							<body>
// 								<div id="app">Welcome GoLang Server!</div>
// 							</body>
// 						</html>
// 					`
// 	fmt.Fprint(w, html)
// }

func main() {
	// http.HandleFunc("/", htmlHandler)
	// http.ListenAndServe(":80", nil)
}
