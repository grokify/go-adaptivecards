package adaptivecards

type ElementMedia struct {
	Type    string               `json:"type"`
	AltText string               `json:"altText,omitempty"`
	Poster  string               `json:"poster,omitempty"`
	Sources []ElementMediaSource `json:"sources,omitempty"`

	Fallback  ElementOrFallbackOption `json:"fallback,omitoption"`
	Height    BlockElementHeight      `json:"height,omitempty"`
	ID        string                  `json:"id,omitempty"`
	IsVisible bool                    `json:"isVisible"`
	Separator bool                    `json:"separator,omitempty"`
	Spacing   Spacing                 `json:"spacing,omitempty"`
}

func (el ElementMedia) ElementID() string { return el.ID }

func (el ElementMedia) FallbackOption() bool { return true }

func (el ElementMedia) SetVisibility(visible bool) {
	el.IsVisible = visible
}
