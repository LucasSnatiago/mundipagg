package mundipagg

// Precification schemas to precify the items for the products
type Precification struct {
	Price         int64          `json:"price,omitempty"`
	MinimumPrice  int64          `json:"minimum_price,omitempty"`
	SchemaType    string         `json:"schema_type,omitempty"`
	PriceBrackets []PriceBracket `json:"price_brackets,omitempty"`
}

// PriceBracket Struct to hold all the intervals for counting the final price
type PriceBracket struct {
	StartedQuantity int64 `json:"start_quantity,omitempty"`
	EndQuantity     int64 `json:"end_quantity,omitempty"`
	OveragePrice    int64 `json:"overage_price,omitempty"`
	Price           int64 `json:"price,omitempty"`
}
