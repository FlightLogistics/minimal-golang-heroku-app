package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<h1>Hello World! I am alive and gay</h1>")
	})

	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
