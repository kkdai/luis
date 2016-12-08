package luis

const (
	LUIS_URL         string = "https://api.projectoxford.ai/luis/v1.0/prog/apps/"
	LUIS_API_INTENTS string = "intents"
)

func getIntentListURL(apid string) string {
	return LUIS_URL + apid + "/" + LUIS_API_INTENTS
}
