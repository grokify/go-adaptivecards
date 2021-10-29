package demo

import (
	"encoding/json"
	"testing"

	"github.com/grokify/go-adaptivecards"
)

var cardTests = []struct {
	example string
}{
	{"bulletsmulti"},
	{"bullets"},
}

// TestExampleMessages ensures parse sample correct.
func TestExampleMessages(t *testing.T) {
	for _, tt := range cardTests {
		switch tt.example {
		case "bullets":
			roundTripCheck(t, ExampleCardBullets(""))
		case "bulletsmulti":
			roundTripCheck(t, ExampleCardBulletsMulti())
		}
	}
}

func roundTripCheck(t *testing.T, card *adaptivecards.AdaptiveCard) {
	js, err := json.Marshal(card)
	if err != nil {
		t.Errorf("jsonMarshal error: [%s]", err.Error())
	}
	ac2, err := adaptivecards.Parse(js)
	if err != nil {
		t.Errorf("jsonMarshal error: [%s]", err.Error())
	}
	js2, err := json.Marshal(ac2)
	if err != nil {
		t.Errorf("jsonMarshal error: [%s]", err.Error())
	}
	if string(js) != string(js2) {
		t.Errorf("marshal/unmarshal round trip failure for demo card type: [%s]", "bullets")
	}
}
