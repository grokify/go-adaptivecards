package demo

import (
	ac "github.com/grokify/go-adaptivecards"
)

func Colors() []ac.Colors {
	return []ac.Colors{
		ac.ColorDefault, ac.ColorDark, ac.ColorLight, ac.ColorAccent,
		ac.ColorGood, ac.ColorWarning, ac.ColorAttention}
}

func FontSizes() []ac.FontSize {
	return []ac.FontSize{ac.FontSizeDefault,
		ac.FontSizeSmall, ac.FontSizeMedium, ac.FontSizeLarge, ac.FontSizeExtraLarge}
}

func FontTypes() []ac.FontType {
	return []ac.FontType{ac.FontTypeDefault, ac.FontTypeMonospace}
}

func FontWeights() []ac.FontWeight {
	return []ac.FontWeight{
		ac.FontWeightDefault, ac.FontWeightLighter, ac.FontWeightBolder}
}
