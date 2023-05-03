package main

import "net/http"

func main() {
	http.HandleFunc("/", sampleAPI)
	http.ListenAndServe(":8080", nil)
}

func sampleAPI(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello!"))
}
