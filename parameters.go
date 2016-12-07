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
