package adaptivecards

type Element interface {
	ElementID() string
	GetType() string
	SetVisibility(visiblity bool)
}
