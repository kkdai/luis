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

type IntentsResultType struct {
	Name  string      `json:"Name"`
	Label interface{} `json:"label"`
	Score float64     `json:"score"`
}

//PredictResponse :
type PredictResponse []struct {
	IntentsResults  []IntentsResultType `json:"IntentsResults"`
	EntitiesResults []interface{}       `json:"EntitiesResults"`
	UtteranceText   string              `json:"utteranceText"`
	TokenizedText   []string            `json:"tokenizedText"`
	ExampleID       string              `json:"exampleId"`
	Metadata        interface{}         `json:"metadata"`
}

//NewPredictResponse :
func NewPredictResponse(raw []byte) *PredictResponse {
	ret := &PredictResponse{}
	err := json.Unmarshal(raw, &ret)
	if err != nil {
		log.Println("json unmarshal err:", err)
		return ret
	}
	return ret
}

//GetBestScoreIntent :Get the best score intent prediction
func GetBestScoreIntent(preResult *PredictResponse) IntentsResultType {
	bestResultIndex := 0
	bestResultScore := -1.0

	for k, v := range (*preResult)[0].IntentsResults {
		if bestResultScore < v.Score {
			bestResultScore = v.Score
			bestResultIndex = k
		}
	}
	return (*preResult)[0].IntentsResults[bestResultIndex]
}
