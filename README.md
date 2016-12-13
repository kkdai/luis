LUIS.ai for Golang
======================
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/kkdai/luis/master/LICENSE)  [![GoDoc](https://godoc.org/github.com/kkdai/luis?status.svg)](https://godoc.org/github.com/kkdai/luis)  [![Build Status](https://travis-ci.org/kkdai/luis.svg)](https://travis-ci.org/kkdai/luis)
 


###Language Understanding Intelligent Service (LUIS)

LUIS lets your app understand language

- LUIS is in beta and free to use
- Supported browsers: Internet Explorer 10/11, Chrome
 
![](https://www.luis.ai/Content/images/CreateLanguageModels.png)

Create language understanding models.

![](https://www.luis.ai/Content/images/UsePrebuiltModels.png)

Use pre-built, world-class models from Bing and Cortana.

![](https://www.luis.ai/Content/images/DeployModels.png)

Deploy your models to an HTTP endpoint.

![](https://www.luis.ai/Content/images/ConsumeModels.png)

Activate models on any device.
 

Installation
---------------

```
go get github.com/kkdai/luis
```

How to use it
---------------


```go

var API_KEY string
var APPID string

func main() {
	APPID = os.Getenv("APP_ID")
	if API_KEY == "" {
		fmt.Println("Please export your key to environment first, `export SUB_KEY=12234 && export APP_ID=5678`")

	if API_KEY == "" {
		return
	}

	e := getLuis(t)

	res, err := e.IntelList()

	if err != nil {
		log.Error("Error happen on :", err.Err)
	}
	fmt.Println("Got response:", string(res))
	result := NewIntentListResponse(res)
	fmt.Println("Luis Intent Ret", result)
	
	//Add utterances
	ex := ExampleJson{ExampleText: "test", SelectedIntentName: "test2"}
	res, err = e.AddLabel(ex)

	//Train it
	res, err = e.Train()

	//Predict it, once you have train your models.
		res, err := e.Predict("test string")

	if err != nil {
		log.Error("Error happen on :", err.Err)
	}
	fmt.Println("Got response:", string(res))
	fmt.Println("Get the best predict result:", GetBestScoreIntent(NewPredictResponse(res)))

	//If you don't have any model
	//Get the best predict result: {None <nil> 0.28}
}
```

Implemented APIs
---------------

- [actionChannels](https://dev.projectoxford.ai/docs/services/56d95961e597ed0f04b76e58/operations/5739a8c71984550500affdfa)
- [intents](https://dev.projectoxford.ai/docs/services/56d95961e597ed0f04b76e58/operations/56f8a55119845511c81de467)
- [predict](https://dev.projectoxford.ai/docs/services/56d95961e597ed0f04b76e58/operations/56f8a55119845511c81de479)
- [train](https://dev.projectoxford.ai/docs/services/56d95961e597ed0f04b76e58/operations/56f8a55119845511c81de483)
- [example](https://dev.projectoxford.ai/docs/services/56d95961e597ed0f04b76e58/operations/56f8a55119845511c81de461)


Unimplement APIs (Yet)
---------------

Need your help to send your PR.

Contribute
---------------

Please open up an issue on GitHub before you put a lot efforts on pull request.
The code submitting to PR must be filtered with `gofmt`

License
---------------

This package is licensed under MIT license. See LICENSE for details.


[![Bitdeli Badge](https://d2weczhvl823v0.cloudfront.net/kkdai/mstranslator/trend.png)](https://bitdeli.com/free "Bitdeli Badge")

