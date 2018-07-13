package bukalapak

import (
  "fmt"
  "encoding/json"
  "github.com/kitacode/go-store/store"
)

type Bukalapak struct {
    store.Store
    Orders
    UserId        string
    Token         string
}

func (bukalapak *Bukalapak) GetOrders() {
  url := "https://api.bukalapak.com/v2/transactions.json"
  // API auth
  bukalapak.AuthType = "basic"
  bukalapak.AuthUsername = bukalapak.UserId
  bukalapak.AuthPassword = bukalapak.Token
  // access API
  json_str := bukalapak.Get(url)

  var orders Orders
  json.Unmarshal([]byte(json_str), &orders)
  // fmt.Println(orders.Orders)
  a, _ := json.Marshal(orders)
  fmt.Println(string(a))
}
