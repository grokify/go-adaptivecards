package adaptivecards

type Element interface {
	ElementID() string
}

type ElementFallbackOption interface {
	ElementID() string
	FallbackOption() bool
}
