package stypes

import (
	"encoding/xml"
)

type DecimalNumber string

func DecimalNumberFromStr(value string) (DecimalNumber, error) {
	return DecimalNumber(value), nil
}

func (d *DecimalNumber) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := DecimalNumberFromStr(attr.Value)
	if err != nil {
		return err
	}

	*d = val

	return nil
}
