package luis

import (
	"encoding/json"
	"log"
)

//Storage all HTTP response include http.response code and error code
type ErrorResponse struct {
	ErrorCode int
	Err       error
}

//Face Response struct
type FaceResponse []struct {
	Faceid        string `json:"faceId"`
	Facerectangle struct {
		Top    int `json:"top"`
		Left   int `json:"left"`
		Width  int `json:"width"`
		Height int `json:"height"`
	} `json:"faceRectangle"`
	Facelandmarks struct {
		Pupilleft struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"pupilLeft"`
		Pupilright struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"pupilRight"`
		Nosetip struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"noseTip"`
		Mouthleft struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"mouthLeft"`
		Mouthright struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"mouthRight"`
		Eyebrowleftouter struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"eyebrowLeftOuter"`
		Eyebrowleftinner struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"eyebrowLeftInner"`
		Eyeleftouter struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"eyeLeftOuter"`
		Eyelefttop struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"eyeLeftTop"`
		Eyeleftbottom struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"eyeLeftBottom"`
		Eyeleftinner struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"eyeLeftInner"`
		Eyebrowrightinner struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"eyebrowRightInner"`
		Eyebrowrightouter struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"eyebrowRightOuter"`
		Eyerightinner struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"eyeRightInner"`
		Eyerighttop struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"eyeRightTop"`
		Eyerightbottom struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"eyeRightBottom"`
		Eyerightouter struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"eyeRightOuter"`
		Noserootleft struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"noseRootLeft"`
		Noserootright struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"noseRootRight"`
		Noseleftalartop struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"noseLeftAlarTop"`
		Noserightalartop struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"noseRightAlarTop"`
		Noseleftalarouttip struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"noseLeftAlarOutTip"`
		Noserightalarouttip struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"noseRightAlarOutTip"`
		Upperliptop struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"upperLipTop"`
		Upperlipbottom struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"upperLipBottom"`
		Underliptop struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"underLipTop"`
		Underlipbottom struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"underLipBottom"`
	} `json:"faceLandmarks"`
	Faceattributes struct {
		Smile    float64 `json:"smile"`
		Headpose struct {
			Pitch float64 `json:"pitch"`
			Roll  float64 `json:"roll"`
			Yaw   float64 `json:"yaw"`
		} `json:"headPose"`
		Gender     string  `json:"gender"`
		Age        float64 `json:"age"`
		Facialhair struct {
			Moustache float64 `json:"moustache"`
			Beard     float64 `json:"beard"`
			Sideburns float64 `json:"sideburns"`
		} `json:"facialHair"`
	} `json:"faceAttributes"`
}

func NewFaceResponse(raw []byte) FaceResponse {
	ret := FaceResponse{}
	err := json.Unmarshal(raw, &ret)
	if err != nil {
		log.Println("json unmarshal err:", err)
		return nil
	}
	return ret
}

type SimilarResponse []struct {
	Faceid     string  `json:"faceId"`
	Confidence float64 `json:"confidence"`
}

func NewSimilarResponse(raw []byte) SimilarResponse {
	ret := SimilarResponse{}
	err := json.Unmarshal(raw, &ret)
	if err != nil {
		log.Println("json unmarshal err:", err)
		return nil
	}
	return ret
}

type GroupResponse struct {
	Groups     [][]string `json:"groups"`
	Messygroup []string   `json:"messyGroup"`
}

func NewGroupResponse(raw []byte) *GroupResponse {
	ret := new(GroupResponse)
	err := json.Unmarshal(raw, ret)
	if err != nil {
		log.Println("json unmarshal err:", err)
		return nil
	}
	return ret
}

type VerifyResponse struct {
	IsIdentical bool    `json:"isIdentical"`
	Confidence  float64 `json:"confidence"`
}

func NewVerifyResponse(raw []byte) VerifyResponse {
	ret := VerifyResponse{}
	err := json.Unmarshal(raw, &ret)
	if err != nil {
		log.Println("json unmarshal err:", err)
		return ret
	}
	return ret
}
