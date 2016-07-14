//Similar to get

package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
	robots, err := ioutil.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
	var dat map[string]interface{}
	if err := json.Unmarshal(robots, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
