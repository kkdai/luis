package luis

import (
	"bytes"
	"errors"
)

type Emotion struct {
	client *Client
}

func NewEmotion(key string) *Emotion {
	if len(key) == 0 {
		return nil
	}

	e := new(Emotion)
	e.client = NewClient(key)
	return e
}

func (e *Emotion) detect(data *bytes.Buffer, useJson bool) ([]byte, *ErrorResponse) {
	url := getEmotionURL()
	return e.client.Connect("POST", url, data, useJson)
}

//Detect Emotion with input URL
func (e *Emotion) EmotionUrl(url string) ([]byte, *ErrorResponse) {
	data := getUrlByteBuffer(url)
	return e.detect(data, true)
}

//Detect Emotion with input image file path
func (e *Emotion) EmotionFile(filePath string) ([]byte, *ErrorResponse) {
	data, err := getFileByteBuffer(filePath)
	if err != nil {
		return nil, &ErrorResponse{Err: errors.New("File path err:" + err.Error())}
	}
	return e.detect(data, false)
}
