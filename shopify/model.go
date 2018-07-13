package shopify

import(
  "encoding/json"
)

type Orders struct {
  Orders      []*Order          `json:"orders"`
}

type Order struct {
  Id                    int               `json:"id"`
  Number                int               `json:"number"`
  TotalWeight           int               `json:"total_weight"`
  UserId                int               `json:"user_id"`
  LocationId            int               `json:"location_id"`
  AppId                 int               `json:"app_id"`
  OrderNumber           int               `json:"order_number"`
  Name                  string            `json:"name"`
  Email                 string            `json:"email"`
  Note                  string            `json:"note"`
  Token                 string            `json:"token"`
  Gateway               string            `json:"gateway"`
  Currency              string            `json:"currency"`
  FinancialStatus       string            `json:"financial_status"`
  CartToken             string            `json:"cart_token"`
  Phone                 string            `json:"phone"`
  TaxesIncluded         bool              `json:"taxes_included"`
  Confirmed             bool              `json:"confirmed"`
  Test                  bool              `json:"test"`
  BuyerAcceptsMarketing bool              `json:"buyer_accepts_marketing"`
  ClosedAt              json.RawMessage   `json:"closed_at"`
  CreatedAt             json.RawMessage   `json:"created_at"`
  UpdatedAt             json.RawMessage   `json:"updated_at"`
  CancelledAt           json.RawMessage   `json:"cancelled_at"`
  ProcessedAt           json.RawMessage   `json:"processed_at"`
  ReferringSite         json.RawMessage   `json:"referring_site"`
  LandingSite           json.RawMessage   `json:"landing_site"`
  CancelReason          json.RawMessage   `json:"cancel_reason"`
  CheckoutToken         json.RawMessage   `json:"checkout_token"`
  Reference             json.RawMessage   `json:"reference"`
  SourceIdentifier      json.RawMessage   `json:"source_identifier"`
  SourceUrl             json.RawMessage   `json:"source_url"`
  DeviceId              json.RawMessage   `json:"device_id"`
  CustomerLocale        json.RawMessage   `json:"customer_locale"`
  BrowserIp             json.RawMessage   `json:"browser_ip"`
  LandingSiteRef        json.RawMessage   `json:"landing_site_ref"`
  TotalPrice            json.Number       `json:"total_price"`
  SubtotalPrice         json.Number       `json:"subtotal_price"`
  TotalTax              json.Number       `json:"total_tax"`
  TotalDiscounts        json.Number       `json:"total_discounts"`
  TotalLineItemsPrice   json.Number       `json:"total_line_items_price"`
  LineItems             []*LineItems      `json:"line_items"`
}

type LineItems struct {
  Id                    int               `json:"id"`
  VariantId             int               `json:"variant_id"`
  Quantity              int               `json:"quantity"`
  Title                 string            `json:"title"`
  Price                 json.Number       `json:"price"`
}
