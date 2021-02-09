package adaptivecards

type Element interface {
	ElementID() string
}

type ElementOrFallbackOption interface {
	ElementID() string
	FallbackOption() bool
}
