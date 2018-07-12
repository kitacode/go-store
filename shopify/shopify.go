package shopify

import (
  "fmt"
  "encoding/json"
  "github.com/kitacode/go-store/store"
)

type Shopify struct {
    store.Store
    Order
    Apikey        string
    ApiPassword   string
}

func (shopify *Shopify) GetOrders() {
  url := shopify.buildUrl()
  json_str := shopify.Get(url)

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

  var orders Orders
  json.Unmarshal([]byte(json_str), &orders)
  a, _ := json.Marshal(orders)
  fmt.Println(string(a))
}

func (shopify *Shopify) GetOrder() {
  fmt.Println("test get order")
}

func (shopify *Shopify) buildUrl() string {
  return "https://" + shopify.Apikey + ":" + shopify.ApiPassword + "@clodeodev.myshopify.com/admin/orders.json"
}
