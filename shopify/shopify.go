package shopify

import (
  "fmt"
  "encoding/json"
  "github.com/kitacode/go-store/store"
)

type Shopify struct {
    store.Store
    Apikey        string
    ApiPassword   string
}

func (shopify *Shopify) GetOrders() Orders {
  url := shopify.buildUrl()
  json_str := shopify.Get(url)

  var orders Orders
  json.Unmarshal([]byte(json_str), &orders)
  // a, _ := json.Marshal(orders)
  // fmt.Println(string(a))

  // return string(a)
  return orders
}

func (shopify *Shopify) GetOrder() {
  fmt.Println("test get order")
}

func (shopify *Shopify) buildUrl() string {
  return "https://" + shopify.Apikey + ":" + shopify.ApiPassword + "@clodeodev.myshopify.com/admin/orders.json"
}
