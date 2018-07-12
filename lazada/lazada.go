package lazada

import (
  "github.com/kitacode/go-store/store"
  "fmt"
)

type Lazada struct {
    store.Store
}

func (lazada *Lazada) GetOrders() {
  fmt.Println("lazada")
}
