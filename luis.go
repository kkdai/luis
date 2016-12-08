package luis

import "fmt"

type Luis struct {
	appid  string
	client *Client
}

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
func (e *Luis) IntelList() ([]byte, *ErrorResponse) {
	url := getIntentListURL(e.appid)
	fmt.Println(url)
	return e.client.Connect("GET", url, nil, true)
}
