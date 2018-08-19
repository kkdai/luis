package luis

const (
	//LuisURL :Basic URL (V1.0)
	LuisURL string = "https://westus.api.cognitive.microsoft.com/luis/api/v2.0/apps/"
	// {ap/pId}/versions/{versionId}/example

	//LuisAPIIntents :API Intent List
	LuisAPIIntents string = "intents"
	//LuisAPIActionChannels :API Action Channels
	LuisAPIActionChannels string = "actionChannels"
	//LuisAPIPredict :API Predict
	LuisAPIPredict string = "predict"
	//LuisAPITrain :API Train
	LuisAPITrain string = "train"
	//LuisAPIAddExample :API Add Label
	LuisAPIAddExample string = "example"
	//LuisAPIPublish :API Publish your application
	LuisAPIPublish string = "publish"
	//LuisAPIVersion :API Get application version
	LuisAPIVersion string = "version"
)

func getAppInfo(appid string) string {
	return LuisURL + appid
}
func getIntentListURL(apid string, versionid string) string {
	return LuisURL + apid + "/versions/" + versionid + "/" + LuisAPIIntents
}

func getActionChannels(apid string, versionid string) string {
	return LuisURL + apid + "/versions/" + versionid + "/" + LuisAPIActionChannels
}

func getPredictURL(apid string, versionid string) string {
	return "https://westus.api.cognitive.microsoft.com/luis/v2.0/apps/" + apid
}

func getTrainURL(apid string, versionid string) string {
	return LuisURL + apid + "/versions/" + versionid + "/" + LuisAPITrain
}

func getAddExampleURL(apid string, versionid string) string {
	return LuisURL + apid + "/versions/" + versionid + "/" + LuisAPIAddExample
}

func getPublishURL(apid string, versionid string) string {
	return LuisURL + apid + "/" + LuisAPIPublish
}

func getVersionURL(apid string, versionid string) string {
	return LuisURL + apid + "/versions/"
}
