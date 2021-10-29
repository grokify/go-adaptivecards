package adaptivecards

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

const (
	ElementTypeImage         = "Image"
	ElementTypeMedia         = "Media"
	ElementTypeRichTextBlock = "RichTextBlock"
	ElementTypeTextBlock     = "TextBlock"
	ElementTypeTextRun       = "TextRun"
)

type Elements []Element

type Element interface {
	ElementID() string
	GetType() string
	SetVisibility(visiblity bool)
}

func (els *Elements) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	} else if string(data) == "[]" {
		return nil
	}
	els1 := []interface{}{}
	els2 := []Element{}

	err := json.Unmarshal(data, &els1)
	if err != nil {
		return err
	}
	for _, el := range els1 {
		el1, ok := el.(elementInspect)
		if !ok {
			return errors.New("element `type` property is missing or empty")
		}
		switch strings.TrimSpace(el1.Type) {
		case ElementTypeTextBlock:
			el2, ok := el.(ElementTextBlock)
			if !ok {
				return errors.New("cannot parse type=`TextBlock`")
			}
			els2 = append(els2, el2)
		case ElementTypeImage:
			el2, ok := el.(ElementImage)
			if !ok {
				return errors.New("cannot parse type=`TextBlock`")
			}
			els2 = append(els2, el2)
		case ElementTypeMedia:
			el2, ok := el.(ElementMedia)
			if !ok {
				return errors.New("cannot parse type=`TextBlock`")
			}
			els2 = append(els2, el2)
		default:
			return fmt.Errorf("cannot parse type=`%s`", el1.Type)
		}
	}
	*els = els2
	return nil
}

type elementInspect struct {
	Type string `json:"type"`
}
