//Similar to get

package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:]) //Direct access from URL
	vat := r.Header.Get("authorization")
	fmt.Println(vat)
	data, err := base64.StdEncoding.DecodeString(vat[6:])
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("%q\n", data)
	

}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
