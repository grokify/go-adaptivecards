package adaptivecards

type BackgroundImage struct {
	URL      string        `json:"url"`
	FillMode ImageFillMode `json:"value"`
}

type ImageFillMode string

const (
	ImageFillModeCover              ImageFillMode = "covers"
	ImageFillModeRepeatHorizontally               = "repeatHorizontally"
	ImageFillModeRepeatVertically                 = "repeatVertically"
	ImageFillModeRepeat                           = "repeat"
)
