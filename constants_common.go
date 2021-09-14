package adaptivecards

const (
	TypeTextBlock = "TextBlock"
)

type BlockElementHeight string

const (
	BlockElementHeightAuto    BlockElementHeight = "auto"
	BlockElementHeightStretch BlockElementHeight = "stretch"
)

type Colors string

const (
	ColorDefault   Colors = ""
	ColorDark      Colors = "Dark"
	ColorLight     Colors = "Light"
	ColorAccent    Colors = "Accent"
	ColorGood      Colors = "Good"
	ColorWarning   Colors = "Warning"
	ColorAttention Colors = "Attention"
)

type FontSize string

const (
	FontSizeDefault    FontSize = ""
	FontSizeSmall      FontSize = "Small"
	FontSizeMedium     FontSize = "Medium"
	FontSizeLarge      FontSize = "Large"
	FontSizeExtraLarge FontSize = "ExtraLarge"
)

type FontType string

const (
	FontTypeDefault   FontType = ""
	FontTypeMonospace FontType = "Monospace"
)

type FontWeight string

const (
	FontWeightDefault FontWeight = ""
	FontWeightLighter FontWeight = "Lighter"
	FontWeightBolder  FontWeight = "Bolder"
)

type HorizontalAlignment string

const (
	HorizontalAlignmentLeft   ImageSize = "left"
	HorizontalAlignmentCenter ImageSize = "center"
	HorizontalAlignmentRight  ImageSize = "right"
)

type ISelectAction string

const (
	ISelectActionOpenURL          ISelectAction = "Action.OpenUrl"
	ISelectActionSubmit           ISelectAction = "Action.Submit"
	ISelectActionToggleVisibility ISelectAction = "Action.ToggleVisibility"
)
