package adaptivecards

const (
	TypeTextBlock = "TextBlock"
)

type BlockElementHeight string

const (
	BlockElementHeightAuto    BlockElementHeight = "auto"
	BlockElementHeightStretch                    = "stretch"
)

type Colors string

const (
	ColorDefault   Colors = ""
	ColorDark             = "Dark"
	ColorLight            = "Light"
	ColorAccent           = "Accent"
	ColorGood             = "Good"
	ColorWarning          = "Warning"
	ColorAttention        = "Attention"
)

type FontSize string

const (
	FontSizeDefault    FontType = ""
	FontSizeSmall               = "Small"
	FontSizeMedium              = "Medium"
	FontSizeLarge               = "Large"
	FontSizeExtraLarge          = "ExtraLarge"
)

type FontType string

const (
	FontTypeDefault   FontType = "default"
	FontTypeMonospace          = "monospace"
)

type FontWeight string

const (
	FontWeightDefault FontWeight = ""
	FontWeightLighter            = "Lighter"
	FontWeightBolder             = "Bolder"
)

type HorizontalAlignment string

const (
	HorizontalAlignmentLeft   ImageSize = "left"
	HorizontalAlignmentCenter           = "center"
	HorizontalAlignmentRight            = "right"
)

type ISelectAction string

const (
	ISelectActionOpenURL          ISelectAction = "Action.OpenUrl"
	ISelectActionSubmit                         = "Action.Submit"
	ISelectActionToggleVisibility               = "Action.ToggleVisibility"
)
