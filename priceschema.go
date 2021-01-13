package mundipagg

// PriceSchema schemas to precify the items for the products
type PriceSchema struct {
	Price        int64 `json:"price,omitempty"`
	MinimumPrice int64 `json:"minimum_price,omitempty"`

	// Schema types
	// If choose unit please provide a price and a minimum_price
	SchemaType string `json:"schema_type,omitempty"`

	// Only works with package, volume e tier.
	PriceBrackets *[]PriceBracket `json:"price_brackets,omitempty"`
}

/* SchemaTypes
1 - unit
2 - package
3 - volume
4 - tier
*/
func (p *PriceSchema) SchemaTypes(schemaTypes int) {
	schemas := map[int]string{
		1: "unit",
		2: "package",
		3: "volume",
		4: "tier",
	}

	p.SchemaType = schemas[schemaTypes]
}

// PriceBracket Struct to hold all the intervals for counting the final price
type PriceBracket struct {
	StartedQuantity int64 `json:"start_quantity,omitempty"`
	EndQuantity     int64 `json:"end_quantity,omitempty"`
	OveragePrice    int64 `json:"overage_price,omitempty"`
	Price           int64 `json:"price,omitempty"`
}
