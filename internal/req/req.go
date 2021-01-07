package req

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Request the high-level struct to do a request
type Request struct {
	Data           []byte `json:"data"`
	IdempotencyKey string
}

// MakePostRequest do the low-level request
func (r Request) MakePostRequest(l *Login, json []byte) {
	url := BASEURL
	fmt.Println("URL:>", url)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))

	req.SetBasicAuth(l.BasicSecretAuthKey, "")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Idempotency-key", r.IdempotencyKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
