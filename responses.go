package mundipagg

import "time"

// Response struct that holds a response from Mundipagg
type Response struct {
	ID         string     `json:"id,omitempty"`
	Name       string     `json:"name,omitempty"`
	Email      string     `json:"email,omitempty"`
	Delinquent bool       `json:"delinquent,omitempty"`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
}
