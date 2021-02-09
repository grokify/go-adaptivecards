package adaptivecards

type ElementTextRun struct {
	Type          string        `json:"type"`
	Text          string        `json:"text"`
	Color         Colors        `json:"color,omitempty"`
	FontType      FontType      `json:"fontType,omitempty"`
	Highlight     bool          `json:"highlight,omitempty"`
	IsSubtle      bool          `json:"isSubtle,omitempty"`
	Italic        bool          `json:"italic,omitempty"`
	SelectAction  ISelectAction `json:"selectAction,omitempty"`
	Size          FontSize      `json:"fontSize,omitempty"`
	Strikethrough bool          `json:"strikethrough,omitempty"`
	Underline     bool          `json:"underline,omitempty"`
	Weight        FontWeight    `json:"fontWeight,omitempty"`
}
