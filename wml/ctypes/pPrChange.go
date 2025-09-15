package ctypes

import (
	"encoding/xml"
)

// 17.13.5.29 pPrChange (Revision Information for Paragraph Properties)
type PPrChange struct {
	TrackChange

	ParaProp *ParagraphProp `xml:"pPr"`
}

func (p PPrChange) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:pPrChange"

	if err := p.TrackChange.MarshalXmlAttrs(&start); err != nil {
		return err
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	if p.ParaProp != nil {
		if err := e.EncodeElement(p.ParaProp, xml.StartElement{
			Name: xml.Name{Local: "w:pPr"},
		}); err != nil {
			return err
		}
	}

	return e.EncodeToken(start.End())
}

func (p *PPrChange) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	p.TrackChange = TrackChange{}

	if err := p.TrackChange.UnmarshalXMLAttrs(start.Attr); err != nil {
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
			case "pPr":
				p.ParaProp = &ParagraphProp{}
				if err = d.DecodeElement(p.ParaProp, &el); err != nil {
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
