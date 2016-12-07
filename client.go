package luis

import (
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	StreamContent string = "application/octet-stream"
	JsonContent   string = "application/json"
)

type Client struct {
	key string
}

// New oxford client based on key
func NewClient(key string) *Client {
	c := new(Client)
	c.key = key
	return c
}

// Connect with API url and data, return response byte or error if http.Status is not OK
func (c *Client) Connect(mode string, url string, data *bytes.Buffer, useJson bool) ([]byte, *ErrorResponse) {
	client := &http.Client{}
	r, _ := http.NewRequest(mode, url, data)

	if useJson {
		r.Header.Add("Content-Type", JsonContent)
	} else {
		r.Header.Add("Content-Type", StreamContent)
	}

	r.Header.Add("Ocp-Apim-Subscription-Key", c.key)
	ret := new(ErrorResponse)
	resp, err := client.Do(r)
	if err != nil {
		log.Println("er:", err)
		ret.Err = err
		return nil, ret
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("er:", err)
		ret.Err = err
		return nil, ret
	}

	if resp.StatusCode != http.StatusOK {
		ret.ErrorCode = resp.StatusCode
		ret.Err = errors.New("Error on:" + string(body))
		log.Println("Error happen! body:", string(body))
		return body, ret
	}

	return body, nil
}
