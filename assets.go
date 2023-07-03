package main

import "net/http"

func asset() {
	fs := http.FileServer(http.Dir("assets/"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":3002", nil)
}
