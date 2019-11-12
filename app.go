package main

import "os"
import "fmt"
import "net/http"

func index(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	fmt.Fprintf(w, "Hello JavaDays 2019 ")
	fmt.Fprintf(w, hostname)
	fmt.Fprintf(w, "\n")
}

func main() {
	http.HandleFunc("/", index)
	fmt.Println("Server started.")
	http.ListenAndServe(":80", nil)
}
