package ctypes

import (
	"encoding/xml"
)

type RPrChange struct {
	TrackChange

	Prop *RunProperty `xml:"rPr"`
}

func (r RPrChange) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:rPrChange"

	if err := r.TrackChange.MarshalXmlAttrs(&start); err != nil {
		return err
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	if r.Prop != nil {
		if err := e.EncodeElement(r.Prop, xml.StartElement{
			Name: xml.Name{Local: "w:rPr"},
		}); err != nil {
			return err
		}
	}

	return e.EncodeToken(start.End())
}

func (r *RPrChange) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	r.TrackChange = TrackChange{}

	if err := r.TrackChange.UnmarshalXMLAttrs(start.Attr); err != nil {
		return err
	}

	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}

		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "rPr":
				r.Prop = &RunProperty{}
				if err = d.DecodeElement(r.Prop, &el); err != nil {
					return err
				}
			}
		case xml.EndElement:
			if el.Name.Local == start.Name.Local {
				return nil
			}
		}
	}
}
