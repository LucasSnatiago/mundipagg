package mundipagg

import "time"

// Customer Constructing a customer
type Customer struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	// Refence client code in the store
	Code string `json:"code,omitempty"`
	// CPF or CNPJ
	Document string     `json:"document,omitempty"`
	Type     string     `json:"type,omitempty"`
	Gender   string     `json:"gender,omitempty"`
	Address  *Address   `json:"address,omitempty"`
	Phones   *Phones    `json:"phones,omitempty"`
	Birthday *time.Time `json:"birthday,omitempty"`
	// Metadata Add some extra information
	Metadata interface{} `json:"metadata,omitempty"`
}

// Address Constructing a address for the customers
type Address struct {
	ID string `json:"id,omitempty"`
	// O atributo line_1 deverá seguir o formato: Número, Rua, Bairro - nesta ordem e separados por vírgula, em virtude do antifraude ou boleto com registro.
	Line1 string `json:"line_1,omitempty"`
	// O atributo line_2 poderá conter informações complementares do endereço, tais como andar, apto, sala, etc.
	Line2     string     `json:"line_2,omitempty"`
	Zipcode   string     `json:"zip_code,omitempty"`
	City      string     `json:"city,omitempty"`
	State     string     `json:"state,omitempty"`
	Country   string     `json:"country,omitempty"`
	Status    string     `json:"status,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

// Phones Struct to contain the data of the phones
type Phones struct {
	HomePhone   *Phone `json:"home_phone,omitempty"`
	MobilePhone *Phone `json:"mobile_phone,omitempty"`
}

// Phone Data about the phone
type Phone struct {
	CountryCode string `json:"country_code,omitempty"`
	AreaCode    string `json:"area_code,omitempty"`
	Number      string `json:"number,omitempty"`
}
