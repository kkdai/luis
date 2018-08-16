package luis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

//Luis :
type Luis struct {
	appid     string
	versionid string
	client    *Client
}

//NewLuis :
func NewLuis(key string, appid string) *Luis {
	if len(key) == 0 {
		return nil
	}

	e := new(Luis)
	e.appid = appid
	e.client = NewClient(key)
	res, err := e.GetAppInfo()
	if err != nil {
		fmt.Println(err, "Cannot get app info")
		return nil
	}
	app := NewAppInfo(res)
	fmt.Println(app)

	e.versionid = app.Endpoints.PRODUCTION.VersionID
	fmt.Println(e.versionid)
	return e
}

//GetAppInfo :Get App detail info and refresh app version id.
func (l *Luis) GetAppInfo() ([]byte, *ErrorResponse) {
	url := getAppInfo(l.appid)
	fmt.Println("app detail URL=", url)

	return l.client.Connect("GET", url, nil, true)
}

//IntelList :Get All Intent List of this app
//Retreives information about the intent models.
func (l *Luis) IntelList() ([]byte, *ErrorResponse) {
	url := getIntentListURL(l.appid, l.versionid)
	fmt.Println("intent list URL=", url)
	return l.client.Connect("GET", url, nil, true)
}

//ActionChannels :get Available Action Channels
//gets a list of all available action channels for the application
func (l *Luis) ActionChannels() ([]byte, *ErrorResponse) {
	url := getActionChannels(l.appid, l.versionid)
	return l.client.Connect("GET", url, nil, true)
}

//Predict :get Train Model Predictions
//gets the trained model predictions for the input examples
func (l *Luis) Predict(utterance string) ([]byte, *ErrorResponse) {
	url := getPredictURL(l.appid, l.versionid)
	data := getStringDataByteBuffer(utterance)
	return l.client.Connect("POST", url, data, true)
}

//Train :Training Status
//gets the training status of all models of the specified application
func (l *Luis) Train() ([]byte, *ErrorResponse) {
	url := getTrainURL(l.appid, l.versionid)
	return l.client.Connect("POST", url, nil, true)
}

//AddLabel :Add Label
// Adds a labeled example to the specified application
func (l *Luis) AddLabel(example ExampleJson) ([]byte, *ErrorResponse) {
	url := getAddExampleURL(l.appid, l.versionid)
	bytExample, err := json.Marshal(example)
	if err != nil {
		log.Println("err:", err)
		return nil, nil
	}
	return l.client.Connect("POST", url, bytes.NewBuffer(bytExample), true)
}
