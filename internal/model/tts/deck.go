package tts

type TransformTTS struct {
	PosX   float32 `json:"posX"`
	PosY   float32 `json:"posY"`
	PosZ   float32 `json:"posZ"`
	RotX   float32 `json:"rotX"`
	RotY   float32 `json:"rotY"`
	RotZ   float32 `json:"rotZ"`
	ScaleX float32 `json:"scaleX"`
	ScaleY float32 `json:"scaleY"`
	ScaleZ float32 `json:"scaleZ"`
}

type CardTTS struct {
	CardID      int
	Name        string
	Nickname    string
	Description string
	Transform   TransformTTS
}

type CardDataTTS struct {
	FaceURL      string
	BackURL      string
	NumHeight    int
	NumWidth     int
	BackIsHidden bool
}

type DeckTTS struct {
	Name             string
	ContainedObjects []CardTTS
	DeckIDs          []int
	CustomDeck       map[string]CardDataTTS
	Transform        TransformTTS
}

type ObjectTTS struct {
	ObjectStates []DeckTTS
}

func DefaultDeckTTS() DeckTTS {
	deckTTS := DeckTTS{
		Name:             "DeckCustom",
		ContainedObjects: nil,
		DeckIDs:          nil,
		CustomDeck:       nil,
		Transform: TransformTTS{
			PosX:   0,
			PosY:   1,
			PosZ:   0,
			RotX:   0,
			RotY:   180,
			RotZ:   180,
			ScaleX: 1.5,
			ScaleY: 1,
			ScaleZ: 1.5,
		},
	}
	return deckTTS
}

func DefaultCardTTS(id int, name string, text string) CardTTS {
	cardTTS := CardTTS{
		CardID:      id,
		Name:        "Card",
		Nickname:    name,
		Description: text,
		Transform: TransformTTS{
			PosX:   0,
			PosY:   0,
			PosZ:   0,
			RotX:   0,
			RotY:   180,
			RotZ:   180,
			ScaleX: 1,
			ScaleY: 1,
			ScaleZ: 1,
		},
	}
	return cardTTS
}

func DefaultCardDataTTS(face string, back string) CardDataTTS {
	cardDataTTS := CardDataTTS{
		FaceURL:      face,
		BackURL:      back,
		NumHeight:    1,
		NumWidth:     1,
		BackIsHidden: true,
	}
	return cardDataTTS
}
