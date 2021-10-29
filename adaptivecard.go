package adaptivecards

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

const (
	TypeAdaptiveCard = "AdaptiveCard"
	SchemaURL        = "http://adaptivecards.io/schemas/adaptive-card.json"
	SchemaVersion13  = "1.3"
)

type AdaptiveCard struct {
	Type         string   `json:"type"`
	Version      string   `json:"version"`
	Body         Elements `json:"body,omitempty"`
	Actions      []Action `json:"actions,omitempty"`
	FallbackText string   `json:"fallbackText,omitempty"`
	Lang         string   `json:"lang,omitempty"`
	MinHeight    string   `json:"minHeight,omitempty"`
	Schema       string   `json:"$schema,omitempty"`
	Speak        string   `json:"speak,omitempty"`
}

func NewAdaptiveCard() *AdaptiveCard {
	return &AdaptiveCard{
		Type:    TypeAdaptiveCard,
		Schema:  SchemaURL,
		Version: SchemaVersion13}
}

func ReadFile(filename string) (*AdaptiveCard, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return Parse(bytes)
}

func Parse(card []byte) (*AdaptiveCard, error) {
	ac := NewAdaptiveCard()
	return ac, json.Unmarshal(card, ac)
}

func (card *AdaptiveCard) Inflate(makeVisible bool) {
	if len(strings.TrimSpace(card.Type)) == 0 {
		card.Type = TypeAdaptiveCard
	}
	if len(strings.TrimSpace(card.Schema)) == 0 {
		card.Schema = SchemaURL
	}
	if len(strings.TrimSpace(card.Version)) == 0 {
		card.Version = SchemaVersion13
	}
	if makeVisible {
		for i := range card.Body {
			switch card.Body[i].GetType() {
			case ElementTypeImage:
				elt := card.Body[i].(ElementImage)
				elt.IsVisible = true
				card.Body[i] = elt
			case ElementTypeMedia:
				elt := card.Body[i].(ElementMedia)
				elt.IsVisible = true
				card.Body[i] = elt
			case ElementTypeTextBlock:
				elt := card.Body[i].(ElementTextBlock)
				elt.IsVisible = true
				card.Body[i] = elt
			}
		}
	}
}
