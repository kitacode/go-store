package main

import (
    // "./store/lazada"
    "github.com/kitacode/go-store/shopify"
)

func main() {
    // lzd := lazada.Lazada{}
    spf := shopify.Shopify{
      Apikey: "86d42b670bd5822538b165b8d8846e8b",
      ApiPassword: "a9e28fc74fdef411d823bcbcc2207c43",
    }

    // lzd.GetOrders()
    spf.GetOrders()
    // spf.GetOrder()
}
