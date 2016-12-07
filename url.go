package luis

const (
	EMOTION_URL string = "https://api.projectoxford.ai/emotion/v1.0/"
	EMOTION_API string = "recognize"
)

func getEmotionURL() string {
	return EMOTION_URL + EMOTION_API
}
