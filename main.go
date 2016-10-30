package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ref:"+r.FormValue("ref"))
	fmt.Fprintf(w, "commits:"+r.FormValue("commits"))
	fmt.Fprintf(w, "head_commit:"+r.FormValue("head_commit"))
	fmt.Fprintf(w, "repository:"+r.FormValue("repository"))
	fmt.Fprintf(w, "pusher:"+r.FormValue("pusher"))
	fmt.Fprintf(w, "sender:"+r.FormValue("sender"))

	fmt.Printf("ref:" + r.FormValue("ref"))
	fmt.Printf("commits:" + r.FormValue("commits"))
	fmt.Printf("head_commit:" + r.FormValue("head_commit"))
	fmt.Printf("repository:" + r.FormValue("repository"))
	fmt.Printf("pusher:" + r.FormValue("pusher"))
	fmt.Printf("sender:" + r.FormValue("sender"))
}

func main() {
	http.HandleFunc("/", handler)
	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, nil)
}
