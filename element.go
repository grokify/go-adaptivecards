package adaptivecards

const (
	TypeImage         = "Image"
	TypeMedia         = "Media"
	TypeRichTextBlock = "RichTextBlock"
	TypeTextBlock     = "TextBlock"
)

type Element interface {
	ElementID() string
	GetType() string
	SetVisibility(visiblity bool)
}
