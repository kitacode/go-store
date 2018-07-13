package bukalapak


type Orders struct {
  Orders      []*Order          `json:"transactions"`
}

type Order struct {
  Id          int               `json:"id"`
}
