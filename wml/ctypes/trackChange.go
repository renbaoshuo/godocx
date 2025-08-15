package ctypes

import (
	"encoding/xml"
	"strconv"
)

// TrackChange represents the complex type for track change
type TrackChange struct {
	ID     int     `xml:"id,attr"`
	Author string  `xml:"author,attr"`
	Date   *string `xml:"date,attr,omitempty"`
}

func (t TrackChange) MarshalXmlAttrs(start *xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:id"}, Value: strconv.Itoa(t.ID)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:author"}, Value: t.Author})

	if t.Date != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:date"}, Value: *t.Date})
	}

	return nil
}

func (t *TrackChange) UnmarshalXMLAttrs(attrs []xml.Attr) error {
	for _, attr := range attrs {
		switch attr.Name.Local {
		case "id":
			id, err := strconv.Atoi(attr.Value)
			if err != nil {
				return err
			}
			t.ID = id
		case "author":
			t.Author = attr.Value
		case "date":
			t.Date = &attr.Value
		}
	}
	return nil
}

func (t TrackChange) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := t.MarshalXmlAttrs(&start); err != nil {
		return err
	}

	return e.EncodeElement("", start)
}
