package bukalapak

import (
  "fmt"
  // "encoding/json"
  "github.com/kitacode/go-store/store"
)

type Bukalapak struct {
    store.Store
    UserId        string
    Token   string
}

func (bukalapak *Bukalapak) GetOrders() {
  url := "https://api.bukalapak.com/v2/transactions.json"
  
  // auth
  bukalapak.AuthType = "basic"
  bukalapak.AuthUsername = bukalapak.UserId
  bukalapak.AuthPassword = bukalapak.Token

  json_str := bukalapak.Get(url)
  fmt.Println(json_str)
  // rawIn := json.RawMessage(json_str)
  // orders, err := rawIn.MarshalJSON()
  // if err != nil {
  //     panic(err)
  // }
  //
  // var o Orders
  // err = json.Unmarshal(orders, &o)
  // if err != nil {
  //     panic(err)
  // }
  //
  // a, _ := json.Marshal(o)
  //
  // fmt.Println(string(a))

  // var orders Orders
  // json.Unmarshal([]byte(json_str), &orders)
  // a, _ := json.Marshal(orders)
  // fmt.Println(string(a))
}
