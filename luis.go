package luis

//Luis :
type Luis struct {
	appid  string
	client *Client
}

//NewLuis :
func NewLuis(key string, appid string) *Luis {
	if len(key) == 0 {
		return nil
	}

	e := new(Luis)
	e.appid = appid
	e.client = NewClient(key)
	return e
}

//IntelList :Get All Intent List of this app
//Retreives information about the intent models.
func (l *Luis) IntelList() ([]byte, *ErrorResponse) {
	url := getIntentListURL(l.appid)
	return l.client.Connect("GET", url, nil, true)
}

//ActionChannels :get Available Action Channels
//gets a list of all available action channels for the application
func (l *Luis) ActionChannels() ([]byte, *ErrorResponse) {
	url := getActionChannels(l.appid)
	return l.client.Connect("GET", url, nil, true)
}

//Predict :get Train Model Predictions
//gets the trained model predictions for the input examples
func (l *Luis) Predict(utterance string) ([]byte, *ErrorResponse) {
	url := getPredictURL(l.appid)
	data := getStringDataByteBuffer(utterance)
	return l.client.Connect("POST", url, data, true)
}

//Train :Training Status
//gets the training status of all models of the specified application
func (l *Luis) Train(utterance string) ([]byte, *ErrorResponse) {
	url := getTrainURL(l.appid)
	return l.client.Connect("POST", url, nil, true)
}
