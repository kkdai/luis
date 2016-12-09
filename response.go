package luis

import (
	"encoding/json"
	"log"
)

//IntentListResponse :
type IntentListResponse []struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

//NewIntentListResponse :
func NewIntentListResponse(raw []byte) *IntentListResponse {
	ret := &IntentListResponse{}
	err := json.Unmarshal(raw, &ret)
	if err != nil {
		log.Println("json unmarshal err:", err)
		return ret
	}
	return ret
}

//ErrorResponse :
//Storage all HTTP response include http.response code and error code
type ErrorResponse struct {
	ErrorCode int
	Err       error
}

//ActionChannelsResponse :
type ActionChannelsResponse []struct {
	Name       string `json:"Name"`
	Version    int    `json:"Version"`
	SupportURI string `json:"SupportUri"`
	Actions    []struct {
		Name       string `json:"Name"`
		Parameters []struct {
			Name            string `json:"Name"`
			DataType        string `json:"DataType"`
			PlaceHolderText string `json:"PlaceHolderText"`
			Required        bool   `json:"Required"`
		} `json:"Parameters"`
		SupportedAuthenticationTypes []int `json:"SupportedAuthenticationTypes"`
	} `json:"Actions"`
}
