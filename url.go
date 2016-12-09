package luis

const (
	//LuisURL :Basic URL
	LuisURL string = "https://api.projectoxford.ai/luis/v1.0/prog/apps/"
	//LuisAPIIntents :API Intent List
	LuisAPIIntents string = "intents"
	//LuisAPIActionChannels :API Action Channels
	LuisAPIActionChannels string = "actionChannels"
	//LuisAPIPredict :API Predict
	LuisAPIPredict string = "predict"
	//LuisAPITrain :API Predict
	LuisAPITrain string = "train"
)

func getIntentListURL(apid string) string {
	return LuisURL + apid + "/" + LuisAPIIntents
}

func getActionChannels(apid string) string {
	return LuisURL + apid + "/" + LuisAPIActionChannels
}

func getPredictURL(apid string) string {
	return LuisURL + apid + "/" + LuisAPIPredict
}

func getTrainURL(apid string) string {
	return LuisURL + apid + "/" + LuisAPITrain
}
