package adaptivecards

type ElementOrFallbackOption interface {
	ElementID() string
	FallbackOption() bool
}
