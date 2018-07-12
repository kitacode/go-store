package store

import (
  "fmt"
  "net/http"
  "io/ioutil"
)

type Store struct {
  Auth
}

func (store *Store) GetOrders() {
  fmt.Println("tes")
}

func (store *Store) Get(url string) string {
  client := &http.Client{}
  req, _ := http.NewRequest("GET", url, nil)

  switch store.AuthType {
  case "basic":
    req.SetBasicAuth(store.AuthUsername, store.AuthPassword)
    fmt.Println("auth basic")
  }

  resp, _ := client.Do(req)
  // if err != nil {
  // 	// handle error
  // }
  defer resp.Body.Close()
  body, _ := ioutil.ReadAll(resp.Body)
  text := string(body)
  return text
}
