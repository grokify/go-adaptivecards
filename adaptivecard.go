package adaptivecards

type AdaptiveCard struct {
	Type         string    `json:"type"`
	Version      string    `json:"version"`
	Body         []Element `json:"body"`
	FallbackText string    `json:"fallbackText,omitempty"`
	Lang         string    `json:"lang,omitempty"`
	MinHeight    string    `json:"minHeight,omitempty"`
	Schema       string    `json:"$schema,omitempty"`
	Speak        string    `json:"speak,omitempty"`
}
