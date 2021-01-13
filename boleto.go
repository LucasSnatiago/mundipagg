package mundipagg

import "time"

// Boleto struct that holds all nedded information to make a boleto
type Boleto struct {
	// Choose a receiver bank
	Bank string `json:"bank,omitempty"`
	// Instructions in boleto. Max: 256 characters
	Instructions string `json:"instructions,omitempty"`
	// DueAt Validity of the boleto (Optional)
	DueAt *time.Time `json:"due_at,omitempty"`
	// NossoNumero Number to indentify a boleto
	NossoNumero string `json:"nosso_numero,omitempty"`
	// Type of the boleto DM e BDP
	Type string `json:"type,omitempty"`
	// Metadata extra information
	Metadata interface{} `json:"metadata,omitempty"`
	// DocumentNumber Boleto indentifier
	DocumentNumber string `json:"document_number,omitempty"`
}

/* Bank Types
1 - Banco do Brasil
2 - Santander
3 - Caixa Econômica Federal
4 - Bradesco
5 - Itau
6 - Citibank
*/
func (b *Boleto) BankTypes(bankType int) {
	bank := map[int]string{
		1: "001",
		2: "033",
		3: "104",
		4: "237",
		5: "341",
		6: "745",
	}

	b.Bank = bank[bankType]
}

/* BoletoTypes
1 - DM (Duplicata Mercantil)
2 - BDP (Boleto de proposta)
*/
func (b *Boleto) BoletoTypes(boletoType int) {
	boleto := map[int]string{
		1: "DM",
		2: "BDP",
	}

	b.Type = boleto[boletoType]
}
