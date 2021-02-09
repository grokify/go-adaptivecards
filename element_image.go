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

func (el ElementImage) ElementID() string { return el.ID }

func (el ElementImage) FallbackOption() bool { return true }

type ImageSize string

const (
	ImageSizeAuto    ImageSize = "auto"
	ImageSizeStretch           = "stretch"
	ImageSizeSmall             = "small"
	ImageSizeMedium            = "medium"
	ImageSizeLarge             = "large"
)

type ImageStyle string

const (
	ImageSizeDefault ImageSize = "default"
	ImageSizePerson            = "person"
)

type ISelectAction string

const (
	ISelectActionOpenURL          ISelectAction = "Action.OpenUrl"
	ISelectActionSubmit                         = "Action.Submit"
	ISelectActionToggleVisibility               = "Action.ToggleVisibility"
)

type Spacing string

const (
	SpacingDefault    Spacing = "default"
	SpacingNone               = "none"
	SpacingSmall              = "small"
	SpacingMedium             = "medium"
	SpacingLarge              = "large"
	SpacingExtraLarge         = "extraLarge"
	SpacingPadding            = "padding"
)
