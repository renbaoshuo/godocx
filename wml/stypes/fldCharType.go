package stypes

import (
	"encoding/xml"
	"errors"
)

// FldCharType represents the type of complex field character.
type FldCharType string

const (
	// FldCharTypeBegin specifies that this character is the field begin character.
	FldCharTypeBegin FldCharType = "begin"

	// FldCharTypeSeparate specifies that this character is the field separator character.
	FldCharTypeSeparate FldCharType = "separate"

	// FldCharTypeEnd specifies that this character is the field end character.
	FldCharTypeEnd FldCharType = "end"
)

// FldCharTypeFromStr converts a string to FldCharType
func FldCharTypeFromStr(val string) (FldCharType, error) {
	switch val {
	case "begin":
		return FldCharTypeBegin, nil
	case "separate":
		return FldCharTypeSeparate, nil
	case "end":
		return FldCharTypeEnd, nil
	default:
		return "", errors.New("invalid field character type")
	}
}

// UnmarshalXMLAttr implements xml.UnmarshalerAttr interface
func (f *FldCharType) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := FldCharTypeFromStr(attr.Value)
	if err != nil {
		return err
	}
	*f = val
	return nil
}

// String returns the string representation of FldCharType
func (f FldCharType) String() string {
	return string(f)
}
