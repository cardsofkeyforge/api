package tts

type TransformTTS struct {
	posX   float32
	posY   float32
	posZ   float32
	rotX   float32
	rotY   float32
	rotZ   float32
	scaleX float32
	scaleY float32
	scaleZ float32
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

func defaultDeckTTS() DeckTTS {
	deckTTS := DeckTTS{
		Name:             "DeckCustom",
		ContainedObjects: nil,
		DeckIDs:          nil,
		CustomDeck:       nil,
		Transform: TransformTTS{
			posX:   0,
			posY:   1,
			posZ:   0,
			rotX:   0,
			rotY:   180,
			rotZ:   180,
			scaleX: 1.5,
			scaleY: 1,
			scaleZ: 1.5,
		},
	}
	return deckTTS
}

func defaultCardTTS(id int, name string, text string) CardTTS {
	cardTTS := CardTTS{
		CardID:      id,
		Name:        "Card",
		Nickname:    name,
		Description: text,
		Transform: TransformTTS{
			posX:   0,
			posY:   0,
			posZ:   0,
			rotX:   0,
			rotY:   180,
			rotZ:   180,
			scaleX: 1,
			scaleY: 1,
			scaleZ: 1,
		},
	}
	return cardTTS
}

func defaultCardDataTTS(face string, back string) CardDataTTS {
	cardDataTTS := CardDataTTS{
		FaceURL:      face,
		BackURL:      back,
		NumHeight:    1,
		NumWidth:     1,
		BackIsHidden: true,
	}
	return cardDataTTS
}
