package main

import "net/http"

func main() {
	panic(
		http.ListenAndServe("0.0.0.0:8080",
			http.FileServer(http.Dir("./fif/dist/fif/"))))
}
