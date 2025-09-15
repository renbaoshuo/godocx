package stypes

import (
	"encoding/xml"
	"errors"
)

// 17.18.19 ST_DocType (Document Classification Values)
type DocType string

const (
	DocTypeEMail        DocType = "eMail"
	DocTypeLetter       DocType = "letter"
	DocTypeNotSpecified DocType = "notSpecified"
)

func DocTypeFromStr(value string) (DocType, error) {
	switch value {
	case "eMail":
		return DocTypeEMail, nil
	case "letter":
		return DocTypeLetter, nil
	case "notSpecified":
		return DocTypeNotSpecified, nil
	default:
		return "", errors.New("invalid DocType value")
	}
}

func (d *DocType) UnmarshalXMLAttr(attr xml.Attr) error {
	val, err := DocTypeFromStr(attr.Value)
	if err != nil {
		return err
	}

	*d = val

	return nil
}
