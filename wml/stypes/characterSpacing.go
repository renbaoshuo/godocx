package stypes

import (
	"encoding/xml"
	"errors"
)

// 17.18.7 ST_CharacterSpacing (Character-Level Whitespace Compression Settings)
type CharacterSpacing string

const (
	CharacterSpacingCompressPunctuation                CharacterSpacing = "compressPunctuation"
	CharacterSpacingCompressPunctuationAndJapaneseKana CharacterSpacing = "compressPunctuationAndJapaneseKana"
	CharacterSpacingDoNotCompress                      CharacterSpacing = "doNotCompress"
)

func CharacterSpacingFromStr(value string) (CharacterSpacing, error) {
	switch value {
	case "compressPunctuation":
		return CharacterSpacingCompressPunctuation, nil
	case "compressPunctuationAndJapaneseKana":
		return CharacterSpacingCompressPunctuationAndJapaneseKana, nil
	case "doNotCompress":
		return CharacterSpacingDoNotCompress, nil
	default:
		return "", errors.New("invalid CharacterSpacing value")
	}
}

func (c *CharacterSpacing) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := CharacterSpacingFromStr(attr.Value)
	if err != nil {
		return err
	}

	*c = val

	return nil
}
