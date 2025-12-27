package questpack

import (
	"fmt"
)

type QuestionType int

const (
	CommonQuestion   QuestionType = 0
	CatInBagQuestion QuestionType = 1
	AuctionQuestion  QuestionType = 2
)

var questTypes = map[QuestionType]interface{}{
	CommonQuestion:   nil,
	CatInBagQuestion: nil,
	AuctionQuestion:  nil,
}

type Question struct {
	Text          string         `json:"text"`
	MediaMetadata *MediaMetadata `json:"media_metadata"`
	Value         int            `json:"value"`
	Type          QuestionType   `json:"type"`
}

func newQuestion(text string, value int, questionType QuestionType) (*Question, error) {
	if _, ok := questTypes[questionType]; !ok {
		return nil, fmt.Errorf("unknown question type: %v", questionType)
	}

	return &Question{
		Text:  text,
		Value: value,
		Type:  questionType,
	}, nil
}

func (q *Question) NewQuestion(text string, value int, questionType QuestionType) (*Question, error) {
	return newQuestion(text, value, questionType)
}

func NewMultimediaQuestion(text string, value int, metadata *MediaMetadata, questionType QuestionType) (*Question, error) {
	quest, err := newQuestion(text, value, questionType)
	if err != nil {
		return nil, fmt.Errorf("new multimedia question: %v", err)
	}
	quest.MediaMetadata = metadata

	return quest, nil
}
