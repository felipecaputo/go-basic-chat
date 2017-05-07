package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		html := []byte(`
		<html>
			<head>Chat</head>
			<body>Hello chat!</body>
		</html>
		`)
		w.Write(html)
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
