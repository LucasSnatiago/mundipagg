package req

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/lusantisuper/mundipagg/internal/utils"
)

// Request the high-level struct to do a request
type Request struct {
	Data           interface{} `json:"data,omitempty"`
	IdempotencyKey string      `json:"idempotency_key,omitempty"`
}

// MakePostRequest do the low-level request
func (l Login) MakePostRequest() Response {
	url := SUBSCRIPTIONURL

	postData, err := json.Marshal(l.Request.Data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(postData))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(postData))

	// Setting the headers to make the request
	req.Header.Set("Authorization", "Basic "+utils.ToBase64(l.BasicSecretAuthKey+":"))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Idempotency-Key", l.Request.IdempotencyKey)

	// Running the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Results of the request
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	// Saving the result
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}

	return response
}
