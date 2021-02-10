package adaptivecards

type InputText struct {
	Type         string         `json:"type"` // must be "Input.Date"
	ID           string         `json:"id"`
	IsMultiline  string         `json:"isMultiline,omitempty"`
	MaxLength    int            `json:"maxLength,omitempty"`
	Placeholder  string         `json:"placeholder,omitempty"`
	Regex        string         `json:"regex,omitempty"`
	Style        TextInputStyle `json:"style,omitempty"`
	InlineAction ISelectAction  `json:"inlineAction,omitempty"`
	Value        string         `json:"value,omitempty"`

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

type TextInputStyle string

const (
	TextInputStyleText  TextInputStyle = "text"
	TextInputStyleTel                  = "tel"
	TextInputStyleURL                  = "url"
	TextInputStyleEmail                = "email"
)
