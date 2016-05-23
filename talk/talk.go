package talk

import (
	cabocha "github.com/ledyba/go-cabocha"
)

type textContent interface {
	Text() string
	SetText(string) error
}

// TextContent express all text type content. This include dependent field because it's used at each topics, so for usefulness.
type TextContent struct {
	text string
}

// Text method retuns the value of the text field.
func (content TextContent) Text() string {
	return content.text
}

// SetText method set a value to the text field.
func (content *TextContent) SetText(text string) error {
	content.text = text
	return nil
}

// TextContentWithDependent interface will be used in topic. The object fit this interface fit also endpoint.TextContent, so you can use the object as to send.
type TextContentWithDependent interface {
	Text() string
	Dependent() cabocha.Sentence
}

type textContentWithDependent struct {
	TextContent
	dependent cabocha.Sentence
}

func (content *textContentWithDependent) SetText(text string) error {
	content.text = text

	_cabocha := cabocha.MakeCabocha()
	sentence, err := _cabocha.Parse(text)
	if err != nil {
		return err
	}
	content.dependent = *sentence
	return nil
}

func (content textContentWithDependent) Dependent() cabocha.Sentence {
	return content.dependent
}

// AddDependentInfo returns new TextContent interface added "dependent" field, setter and getter.
func AddDependentInfo(content textContent) (TextContentWithDependent, error) {
	newTextContent := textContentWithDependent{}
	err := newTextContent.SetText(content.Text())
	return newTextContent, err
}
