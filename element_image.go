package adaptivecards

type ElementImage struct {
	AltText             string                `json:"altText,omitempty"`
	BackgroundColor     string                `json:"backgroundColor,omitempty"`
	Fallback            ElementFallbackOption `json:"fallback,omitoption"`
	Height              string                `json:"height,omitempty"`
	HorizontalAlignment HorizontalAlignment   `json:"horizontalAlignment,omitempty"`
	ID                  string                `json:"id,omitempty"`
	IsVisible           bool                  `json:"isVisible"`
	Requires            map[string]string     `json:"requires,omitempty"`
	SelectAction        ISelectAction         `json:"selectAction,omitempty"`
	Separator           bool                  `json:"separator,omitempty"`
	Size                ImageSize             `json:"size,omitempty"`
	Spacing             Spacing               `json:"spacing,omitempty"`
	Style               ImageStyle            `json:"style,omitempty"`
	Type                string                `json:"type"`
	URL                 string                `json:"url,omitempty"`
	Width               string                `json:"width,omitempty"`
}

func (img ElementImage) ElementID() string {
	return img.ID
}

func (img ElementImage) FallbackOption() bool {
	return true
}

type HorizontalAlignment string

const (
	HorizontalAlignmentLeft   ImageSize = "left"
	HorizontalAlignmentCenter           = "center"
	HorizontalAlignmentRight            = "right"
)

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
