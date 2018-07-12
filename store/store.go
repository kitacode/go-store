package store

import (
  "fmt"
  "net/http"
  "io/ioutil"
)

type Store struct {}

func (store *Store) GetOrders() {
  fmt.Println("tes")
}

func (store *Store) Get(url string) string {
  resp, err := http.Get(url)
  if err != nil {
  	// handle error
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  text := string(body)
  return text
}
