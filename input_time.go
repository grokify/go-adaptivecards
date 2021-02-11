package adaptivecards

type InputTime struct {
	Type        string `json:"type"` // must be "Input.Date"
	ID          string `json:"id"`
	Max         string `json:"max,omitempty"`
	Min         string `json:"min,omitempty"`
	Placeholder string `json:"placeholder,omitempty"`
	Value       string `json:"value,omitempty"`

	ErrorMessage string                  `json:"errorMessage,omitempty"`
	IsRequired   bool                    `json:"isRequired"`
	Label        string                  `json:"label,omitempty"`
	Fallback     ElementOrFallbackOption `json:"fallback,omitempty"`
	Height       BlockElementHeight      `json:"height,omitempty"`
	Separator    bool                    `json:"separator,omitempty"`
	Spacing      Spacing                 `json:"spacing,omitempty"`
	IsVisible    bool                    `json:"isVisible"`
	Requires     map[string]string       `json:"requires,omitempty"`
}
