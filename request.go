package mundipagg

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/lusantisuper/mundipagg/internal/utils"
)

// MakePostRequest do the low-level request
func MakePostRequest(data interface{}, secretKey string, indepotencyKey string, url string, requestType ...string) (*Response, error) {
	postData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	requestMethod := "POST"
	if len(requestType) > 0 {
		requestMethod = requestType[0]
	}
	req, err := http.NewRequest(requestMethod, url, bytes.NewBuffer(postData))

	// Setting the headers to make the request
	req.Header.Set("Authorization", "Basic "+utils.ToBase64(secretKey+":"))
	req.Header.Set("Content-Type", "application/json")
	if !utils.IsStringEmpty(indepotencyKey) {
		req.Header.Set("Idempotency-Key", indepotencyKey)
	}

	// Running the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Results of the request
	body, _ := ioutil.ReadAll(resp.Body)
	if body == nil || string(body) == "" {
		body = []byte("{}")
	}

	// Saving the result
	response := Response{MundipaggJSONAnswer: string(body)}

	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		err = errors.New("Invalid Request:\nSended:\n" + string(postData) + "Received:\n" + string(body))
	}

	if err == nil {
		err = json.Unmarshal(body, &response)
	}

	return &response, err
}
