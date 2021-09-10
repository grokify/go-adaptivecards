package adaptivecards

type ActionOpenUrl struct {
	Type  string `json:"type,omitempty"`
	Title string `json:"title,omitempty"`
	URL   string `json:"url,omitempty"`
}

func NewActionOpenUrl(url, title string) ActionOpenUrl {
	return ActionOpenUrl{
		Type:  TypeActionOpenUrl,
		Title: title,
		URL:   url}
}

func (action ActionOpenUrl) Action() string {
	return action.Type
}
