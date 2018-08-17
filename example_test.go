package luis_test

import (
	"fmt"
	"os"

	"github.com/kkdai/luis"
	"github.com/prometheus/common/log"
)

func Example() {
	var API_KEY string
	var APPID string

	APPID = os.Getenv("APP_ID")
	API_KEY = os.Getenv("SUB_KEY")

	if API_KEY == "" {
		fmt.Println("Please export your key to environment first, `export SUB_KEY=12234 && export APP_ID=5678`")
		return
	}

	e := luis.NewLuis(API_KEY, APPID)

	res, err := e.IntelList()
	if err != nil {
		log.Error("Error happen on :", err.Err)
	}

	fmt.Println("Got response:", string(res))
	result := luis.NewIntentListResponse(res)
	fmt.Println("Luis Intent Ret", result)

	//Add utterances
	ex := luis.ExampleJson{ExampleText: "test", SelectedIntentName: "test2"}
	res, err = e.AddLabel(ex)

	//Train it
	res, err = e.Train()
	if err != nil {
		log.Error("Error happen on :", err.Err)
	}

	//Predict it, once you have train your models.
	res, err = e.Predict("test string")

	if err != nil {
		log.Error("Error happen on :", err.Err)
	}
	fmt.Println("Got response:", string(res))
	fmt.Println("Get the best predict result:", luis.GetBestScoreIntent(luis.NewPredictResponse(res)))
}

func ExampleIntelList() {
	var API_KEY string
	var APPID string

	APPID = os.Getenv("APP_ID")
	API_KEY = os.Getenv("SUB_KEY")

	if API_KEY == "" {
		fmt.Println("Please export your key to environment first, `export SUB_KEY=12234 && export APP_ID=5678`")
		return
	}

	e := luis.NewLuis(API_KEY, APPID)

	_, err := e.IntelList()
	if err != nil {
		log.Error("Error happen on :", err.Err)
	}
}
