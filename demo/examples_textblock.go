package demo

import (
	"fmt"
	"strings"

	ac "github.com/grokify/go-adaptivecards"
)

func TestTextBlocks() []ac.ElementTextBlock {
	sizes := FontSizes()
	colors := Colors()
	weights := FontWeights()
	types := FontTypes()
	blocks := []ac.ElementTextBlock{}
	for _, size := range sizes {
		for _, color := range colors {
			for _, weight := range weights {
				for _, fontType := range types {
					blocks = append(blocks, ac.ElementTextBlock{
						Type:     ac.ElementTypeTextBlock,
						Size:     size,
						Color:    color,
						Weight:   weight,
						FontType: fontType,
						IsSubtle: false})
					blocks = append(blocks, ac.ElementTextBlock{
						Type:     ac.ElementTypeTextBlock,
						Size:     size,
						Color:    color,
						Weight:   weight,
						FontType: fontType,
						IsSubtle: true,
					})
				}
			}
		}
	}
	return blocks
}

func FontString(eb ac.ElementTextBlock) string {
	defaultStr := "default"
	delim := ": "
	sep := "; "
	parts := []string{}
	size := strings.TrimSpace(string(eb.Size))
	if len(size) == 0 {
		size = defaultStr
	}
	parts = append(parts, "Size"+delim+size)
	color := strings.TrimSpace(string(eb.Color))
	if len(color) == 0 {
		color = defaultStr
	}
	parts = append(parts, "Color"+delim+color)
	weight := strings.TrimSpace(string(eb.Weight))
	if len(color) == 0 {
		weight = defaultStr
	}
	parts = append(parts, "Weight"+delim+weight)
	fontType := strings.TrimSpace(string(eb.FontType))
	if len(color) == 0 {
		fontType = defaultStr
	}
	parts = append(parts, "FontType"+delim+fontType)
	parts = append(parts, fmt.Sprintf("Subtle%s%v", delim, eb.IsSubtle))
	return strings.Join(parts, sep)
}
