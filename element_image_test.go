package adaptivecards

import (
	"testing"
)

var adaptiveCardTest = AdaptiveCard{
	Body: []Element{
		ElementImage{
			ID: "masterImage",
			Fallback: ElementImage{
				ID: "fallbackImage",
			},
			URL: "https://example.com/image.png",
		},
	},
}

var elementImageTests = []struct {
	card    AdaptiveCard `json:"adaptiveCard"`
	element Element      `json:"element"`
	image   ElementImage `json:"image"`
	id      string
}{
	{
		card: adaptiveCardTest,
		element: ElementImage{
			ID: "masterImage",
			Fallback: ElementImage{
				ID: "fallbackImage",
			},
			URL: "https://example.com/image.png",
		},
		image: ElementImage{
			ID: "masterImage1",
			Fallback: ElementImage{
				ID: "fallbackImage1",
			},
			URL: "https://example.com/image1.png",
		},
		id: "masterImage",
	},
}

func TestElementImage(t *testing.T) {
	for _, tt := range elementImageTests {
		if tt.element.ElementID() != tt.id {
			t.Errorf("Error")
		}
	}
}
