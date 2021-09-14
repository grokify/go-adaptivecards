package adaptivecards

type ElementTextBlock struct {
	Type                string              `json:"type"`
	Text                string              `json:"text"`
	Color               Colors              `json:"color,omitempty"`
	FontType            FontType            `json:"fontType,omitempty"`
	HorizontalAlignment HorizontalAlignment `json:"horizontalAlignment,omitempty"`
	IsSubtle            bool                `json:"isSubtle,omitempty"`
	MaxLines            int                 `json:"maxLines,omitempty"`
	Size                FontSize            `json:"size,omitempty"`
	Weight              FontWeight          `json:"weight,omitempty"`
	Wrap                bool                `json:"wrap,omitempty"`

	Fallback  ElementOrFallbackOption `json:"fallback,omitempty"`
	Height    BlockElementHeight      `json:"height,omitempty"`
	ID        string                  `json:"id,omitempty"`
	IsVisible bool                    `json:"isVisible"`
	Separator bool                    `json:"separator,omitempty"`
	Spacing   Spacing                 `json:"spacing,omitempty"`
	Requires  map[string]string       `json:"requires,omitempty"`
	Width     string                  `json:"width,omitempty"`
}

func (el ElementTextBlock) GetType() string { return el.Type }

func (el ElementTextBlock) ElementID() string { return el.ID }

func (el ElementTextBlock) FallbackOption() bool { return true }

func (el ElementTextBlock) SetVisibility(visible bool) {
	el.IsVisible = visible
}
