package mundipagg

import "github.com/lusantisuper/mundipagg/internal/req"

// NewMundipagg Create a Mundipagg object
func NewMundipagg(basicSecretAuthKey string) (*req.Login, error) {
	return req.NewMundipagg(basicSecretAuthKey)
}

func main() {}
