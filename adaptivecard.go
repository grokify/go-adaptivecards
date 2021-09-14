package adaptivecards

import (
	"strings"
)

const (
	TypeAdaptiveCard = "AdaptiveCard"
	SchemaURL        = "http://adaptivecards.io/schemas/adaptive-card.json"
	SchemaVersion13  = "1.3"
)

type AdaptiveCard struct {
	Type         string    `json:"type"`
	Version      string    `json:"version"`
	Body         []Element `json:"body,omitempty"`
	Actions      []Action  `json:"actions,omitempty"`
	FallbackText string    `json:"fallbackText,omitempty"`
	Lang         string    `json:"lang,omitempty"`
	MinHeight    string    `json:"minHeight,omitempty"`
	Schema       string    `json:"$schema,omitempty"`
	Speak        string    `json:"speak,omitempty"`
}

func NewAdaptiveCard() AdaptiveCard {
	return AdaptiveCard{
		Type:    TypeAdaptiveCard,
		Schema:  SchemaURL,
		Version: SchemaVersion13}
}

func (card *AdaptiveCard) Inflate() {
	if len(strings.TrimSpace(card.Type)) == 0 {
		card.Type = TypeAdaptiveCard
	}
	if len(strings.TrimSpace(card.Schema)) == 0 {
		card.Schema = SchemaURL
	}
	if len(strings.TrimSpace(card.Version)) == 0 {
		card.Version = SchemaVersion13
	}
}
