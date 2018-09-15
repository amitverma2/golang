package main

import (
	"fmt"
	"log"
	"net/http"
        "golang.org/x/talks/2016/applicative/google"
        "time"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		log.Println("serving", req.URL)
		fmt.Fprintln(w, "Hello, World")
	})
	http.HandleFunc("/search", func(w http.ResponseWriter, req *http.Request) {
		log.Println("serving", req.URL)
		query := req.FormValue("q")
		if query == "" {
			http.Error(w, `missing "q" URL parameter`, http.StatusBadRequest)
			return
		}
                start := time.Now()
                results, err := google.Search(query)
                elapsed := time.Since(start)
                if err != nil {
                    http.Error(w, err.Error(), http.StatusInternalServerError)
                    return
                }
                log.Printf("results type %T elapsed time %v\n", results, elapsed)
	})
	fmt.Println("serving on http://localhost:7777/hello")
	log.Fatal(http.ListenAndServe("localhost:7777", nil))
}
