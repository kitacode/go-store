package shopify

import(
  "encoding/json"
)
type Orders struct {
  Orders      []*Order          `json:"orders"`
}

type Order struct {
  Id          int               `json:"id"`
  Email       string            `json:"email"`
  ClosedAt    json.RawMessage   `json:"closed_at"`
  LineItems   []*LineItems      `json:"line_items"`
}

type LineItems struct {
  Id          int               `json:"id"`
  VariantId   int               `json:"variant_id"`
  Title       string            `json:"title"`
  Quantity    int               `json:"quantity"`
  Price       json.Number       `json:"price"`
}
