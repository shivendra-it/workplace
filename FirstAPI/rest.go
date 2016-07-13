package main
import (
  "encoding/json"
  "fmt"
  "net/http"
)

type Payload struct {
  Stuff Data
}

type Data struct {
  Fruit Fruits
  Veggies Vegetables
}

type Fruits map[string]int
type Vegetables map[string]int

func serveRest(w http.ResponseWriter, r *http.Request){
  response,err := getJsonResponce()
  if err != nil {
    panic(err)
  }
  fmt.Fprintf(w, string(response))
}

func getJsonResponce() ([]byte,error){
  fruits := make(map[string]int)
  fruits["Apple"] = 15
  fruits["Oranges"] = 11
  vegetables := make(map[string]int)
  vegetables["Carrots"] = 21
  vegetables["Peppers"] = 0

  d := Data{fruits,vegetables}
  p := Payload{d}

  return json.MarshalIndent(p, "", "  ")

}

func main(){
    http.HandleFunc("/",serveRest)
    http.ListenAndServe("localhost:1337",nil)
}
