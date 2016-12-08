package luis_test

import (
	"fmt"
	"os"
	"testing"

	. "github.com/kkdai/luis"
)

var API_KEY string
var APPID string

const ()

func init() {
	API_KEY = os.Getenv("SUB_KEY")
	APPID = os.Getenv("APP_ID")
	if API_KEY == "" {
		fmt.Println("Please export your key to environment first, `export SUB_KEY=12234 && export APP_ID=5678`")
	}
}

func TestIntentList(t *testing.T) {
	if API_KEY == "" {
		return
	}

	e := NewLuis(API_KEY, APPID)
	if e == nil {
		t.Error("Cannot connect to server")
	}

	res, err := e.IntelList()

	if err != nil {
		t.Error("Error happen on :", err.Err)
	}
	fmt.Println("Got response:", string(res))
}
