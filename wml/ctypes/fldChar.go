package ctypes

import (
	"encoding/xml"

	"github.com/gomutex/godocx/wml/stypes"
)

// FldChar represents a complex field character in a document run.
// This element specifies the presence of a complex field character at the current location in the run content.
type FldChar struct {
	// The type of complex field character represented by this element.
	FldCharType stypes.FldCharType `xml:"fldCharType,attr"`

	// Specifies that this field shall be recalculated before the document is displayed.
	Dirty *bool `xml:"dirty,attr,omitempty"`

	// Specifies that the parent field shall be locked from updates.
	FldLock *bool `xml:"fldLock,attr,omitempty"`
}

// NewFldChar creates a new FldChar with the specified field character type.
func NewFldChar(fldCharType stypes.FldCharType) *FldChar {
	return &FldChar{
		FldCharType: fldCharType,
	}
}

// SetDirty sets the dirty flag for the field character.
func (f *FldChar) SetDirty(dirty bool) *FldChar {
	f.Dirty = &dirty
	return f
}

// SetFldLock sets the field lock flag for the field character.
func (f *FldChar) SetFldLock(lock bool) *FldChar {
	f.FldLock = &lock
	return f
}

func (f FldChar) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:fldChar"

	start.Attr = append(start.Attr, xml.Attr{
		Name:  xml.Name{Local: "w:fldCharType"},
		Value: string(f.FldCharType),
	})

	if f.Dirty != nil {
		value := "0"
		if *f.Dirty {
			value = "1"
		}
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "w:dirty"},
			Value: value,
		})
	}

	if f.FldLock != nil {
		value := "0"
		if *f.FldLock {
			value = "1"
		}
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "w:fldLock"},
			Value: value,
		})
	}

	return e.EncodeElement("", start)
}

func (f *FldChar) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "fldCharType":
			fldCharType, err := stypes.FldCharTypeFromStr(attr.Value)
			if err != nil {
				return err
			}
			f.FldCharType = fldCharType

		case "dirty":
			dirty := attr.Value == "1" || attr.Value == "true"
			f.Dirty = &dirty

		case "fldLock":
			lock := attr.Value == "1" || attr.Value == "true"
			f.FldLock = &lock
		}
	}

	return d.Skip()
}
