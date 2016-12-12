package luis

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
)

func getBooleanString(b bool) string {
	if b {
		return "true"
	} else {
		return "false"
	}
}

func getFileByteBuffer(path string) (*bytes.Buffer, error) {
	fileData, err := ioutil.ReadFile(path)

	if err != nil {
		log.Println("File open err:", err)
		return nil, err
	}
	return bytes.NewBuffer(fileData), nil
}

func getUrlByteBuffer(url string) *bytes.Buffer {
	byteData := []byte(fmt.Sprintf(`{"url":"%s"}`, url))
	return bytes.NewBuffer(byteData)

}

func getUserDataByteBuffer(userData string) *bytes.Buffer {
	byteData := []byte(fmt.Sprintf(`{"userData":"%s"}`, userData))
	return bytes.NewBuffer(byteData)

}

func getStringDataByteBuffer(userData string) *bytes.Buffer {
	byteData := []byte(fmt.Sprintf(`["%s"]`, userData))
	return bytes.NewBuffer(byteData)

}

//ExampleJson :
type ExampleJson struct {
	ExampleText        string `json:"ExampleText"`
	SelectedIntentName string `json:"SelectedIntentName"`
	EntityLabels       []struct {
		EntityType string `json:"EntityType"`
		StartToken int    `json:"StartToken"`
		EndToken   int    `json:"StartToken"`
	} `json:"EntityLabels,omitempty"`
}
