package adaptivecards

type ElementImage struct {
	AltText             string              `json:"altText,omitempty"`
	BackgroundColor     string              `json:"backgroundColor,omitempty"`
	Height              string              `json:"height,omitempty"`
	HorizontalAlignment HorizontalAlignment `json:"horizontalAlignment,omitempty"`
	SelectAction        ISelectAction       `json:"selectAction,omitempty"`
	Size                ImageSize           `json:"size,omitempty"`
	Style               ImageStyle          `json:"style,omitempty"`
	Type                string              `json:"type"`
	URL                 string              `json:"url,omitempty"`
	Width               string              `json:"width,omitempty"`

	Fallback  ElementOrFallbackOption `json:"fallback,omitoption"`
	ID        string                  `json:"id,omitempty"`
	IsVisible bool                    `json:"isVisible"`
	Requires  map[string]string       `json:"requires,omitempty"`
	Separator bool                    `json:"separator,omitempty"`
	Spacing   Spacing                 `json:"spacing,omitempty"`
}

func (el ElementImage) GetType() string { return el.Type }

func (el ElementImage) ElementID() string { return el.ID }

func (el ElementImage) FallbackOption() bool { return true }

func (el ElementImage) SetVisibility(visible bool) {
	el.IsVisible = visible
}

type ImageSize string

const (
	ImageSizeAuto    ImageSize = "auto"
	ImageSizeStretch ImageSize = "stretch"
	ImageSizeSmall   ImageSize = "small"
	ImageSizeMedium  ImageSize = "medium"
	ImageSizeLarge   ImageSize = "large"
)

type ImageStyle string

const (
	ImageSizeDefault ImageSize = "default"
	ImageSizePerson  ImageSize = "person"
)

type Spacing string

const (
	SpacingDefault    Spacing = "default"
	SpacingNone       Spacing = "none"
	SpacingSmall      Spacing = "small"
	SpacingMedium     Spacing = "medium"
	SpacingLarge      Spacing = "large"
	SpacingExtraLarge Spacing = "extraLarge"
	SpacingPadding    Spacing = "padding"
)
