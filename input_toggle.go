package adaptivecards

type InputToggle struct {
	Type     string `json:"type"` // must be "Input.Date"
	Title    string `json:"title"`
	ID       string `json:"id"`
	Value    string `json:"value,omitempty"`
	ValueOff string `json:"valueOff,omitempty"`
	ValueOn  string `json:"valueOn,omitempty"`
	Wrap     string `json:"wrap,omitempty"`

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
