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
		elMap := el.(map[string]interface{})
		elType, ok := elMap["type"]
		if !ok {
			return errors.New("element `type` property is missing or empty")
		}
		elTypeString := elType.(string)
		switch strings.TrimSpace(elTypeString) {
		case ElementTypeTextBlock:
			el2, err := interfaceToTextBlock(el)
			if err != nil {
				return errors.New("cannot parse type=`TextBlock`:" + err.Error())
			}
			els2 = append(els2, el2)
		case ElementTypeImage:
			el2, err := interfaceToImage(el)
			if err != nil {
				return errors.New("cannot parse type=`Image`:" + err.Error())
			}
			els2 = append(els2, el2)
		case ElementTypeMedia:
			el2, err := interfaceToMedia(el)
			if err != nil {
				return errors.New("cannot parse type=`Media`:" + err.Error())
			}
			els2 = append(els2, el2)
		default:
			return fmt.Errorf("cannot parse type=`%s`", elTypeString)
		}
	}
	*els = els2
	return nil
}

type elementInspect struct {
	Type string `json:"type"`
}

func interfaceToTextBlock(data interface{}) (ElementTextBlock, error) {
	el := ElementTextBlock{}
	j, err := json.Marshal(data)
	if err != nil {
		return el, err
	}
	return el, json.Unmarshal(j, &el)
}

func interfaceToImage(data interface{}) (ElementImage, error) {
	el := ElementImage{}
	j, err := json.Marshal(data)
	if err != nil {
		return el, err
	}
	return el, json.Unmarshal(j, &el)
}

func interfaceToMedia(data interface{}) (ElementMedia, error) {
	el := ElementMedia{}
	j, err := json.Marshal(data)
	if err != nil {
		return el, err
	}
	return el, json.Unmarshal(j, &el)
}
