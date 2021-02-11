package adaptivecards

type InputChoice struct {
	Title string `json:"title"`
	Value string `json:"value"`
}

type InputChoiceSet struct {
	Type          string        `json:"type"` // must be "Input.Date"
	Choices       []InputChoice `json:"choices"`
	ID            string        `json:"id"`
	IsMultiSelect bool          `json:"isMultiSelect,omitempty"`
	Style         string        `json:"style,omitempty"`
	Value         string        `json:"value,omitempty"`
	Placeholder   string        `json:"placeholder,omitempty"`
	Wrap          bool          `json:"wrap,omitempty"`

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

type ChoiceInputStyle string

const (
	ChoiceInputStyleCompact  ChoiceInputStyle = "compact"
	ChoiceInputStyleExpanded                  = "expanded"
)
