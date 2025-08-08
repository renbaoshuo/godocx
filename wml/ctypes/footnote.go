package ctypes

import (
	"encoding/xml"
	"io"
	"strconv"
)

var defaultFootnoteNSAttrs = map[string]string{
	"xmlns:wpc":      "http://schemas.microsoft.com/office/word/2010/wordprocessingCanvas",
	"xmlns:cx":       "http://schemas.microsoft.com/office/drawing/2014/chartex",
	"xmlns:cx1":      "http://schemas.microsoft.com/office/drawing/2015/9/8/chartex",
	"xmlns:cx2":      "http://schemas.microsoft.com/office/drawing/2015/10/21/chartex",
	"xmlns:cx3":      "http://schemas.microsoft.com/office/drawing/2016/5/9/chartex",
	"xmlns:cx4":      "http://schemas.microsoft.com/office/drawing/2016/5/10/chartex",
	"xmlns:cx5":      "http://schemas.microsoft.com/office/drawing/2016/5/11/chartex",
	"xmlns:cx6":      "http://schemas.microsoft.com/office/drawing/2016/5/12/chartex",
	"xmlns:cx7":      "http://schemas.microsoft.com/office/drawing/2016/5/13/chartex",
	"xmlns:cx8":      "http://schemas.microsoft.com/office/drawing/2016/5/14/chartex",
	"xmlns:mc":       "http://schemas.openxmlformats.org/markup-compatibility/2006",
	"xmlns:aink":     "http://schemas.microsoft.com/office/drawing/2016/ink",
	"xmlns:am3d":     "http://schemas.microsoft.com/office/drawing/2017/model3d",
	"xmlns:o":        "urn:schemas-microsoft-com:office:office",
	"xmlns:oel":      "http://schemas.microsoft.com/office/2019/extlst",
	"xmlns:r":        "http://schemas.openxmlformats.org/officeDocument/2006/relationships",
	"xmlns:m":        "http://schemas.openxmlformats.org/officeDocument/2006/math",
	"xmlns:v":        "urn:schemas-microsoft-com:vml",
	"xmlns:wp14":     "http://schemas.microsoft.com/office/word/2010/wordprocessingDrawing",
	"xmlns:wp":       "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing",
	"xmlns:w10":      "urn:schemas-microsoft-com:office:word",
	"xmlns:w":        "http://schemas.openxmlformats.org/wordprocessingml/2006/main",
	"xmlns:w14":      "http://schemas.microsoft.com/office/word/2010/wordml",
	"xmlns:w15":      "http://schemas.microsoft.com/office/word/2012/wordml",
	"xmlns:w16cex":   "http://schemas.microsoft.com/office/word/2018/wordml/cex",
	"xmlns:w16cid":   "http://schemas.microsoft.com/office/word/2016/wordml/cid",
	"xmlns:w16":      "http://schemas.microsoft.com/office/word/2018/wordml",
	"xmlns:w16du":    "http://schemas.microsoft.com/office/word/2023/wordml/word16du",
	"xmlns:w16sdtdh": "http://schemas.microsoft.com/office/word/2020/wordml/sdtdatahash",
	"xmlns:w16sdtfl": "http://schemas.microsoft.com/office/word/2024/wordml/sdtformatlock",
	"xmlns:w16se":    "http://schemas.microsoft.com/office/word/2015/wordml/symex",
	"xmlns:wpg":      "http://schemas.microsoft.com/office/word/2010/wordprocessingGroup",
	"xmlns:wpi":      "http://schemas.microsoft.com/office/word/2010/wordprocessingInk",
	"xmlns:wne":      "http://schemas.microsoft.com/office/word/2006/wordml",
	"xmlns:wps":      "http://schemas.microsoft.com/office/word/2010/wordprocessingShape",
	"mc:Ignorable":   "w14 w15 w16se w16cid w16 w16cex w16sdtdh w16sdtfl w16du wp14",
}

type Footnotes struct {
	RelativePath string `xml:"-"`
	Attr         []xml.Attr

	Footnotes []Footnote
}

func (fns *Footnotes) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:footnotes"

	if len(fns.Attr) == 0 {
		for k, v := range defaultFootnoteNSAttrs {
			start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: k}, Value: v})
		}
	} else {
		start.Attr = fns.Attr
	}

	if err := e.EncodeToken(start); err != nil {
		return err
	}

	for _, footnote := range fns.Footnotes {
		if err := e.EncodeElement(footnote, xml.StartElement{Name: xml.Name{Local: "w:footnote"}}); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (fns *Footnotes) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for {
		token, err := d.Token()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		switch t := token.(type) {
		case xml.StartElement:
			if t.Name.Local == "footnote" {
				footnote := Footnote{}
				if err := d.DecodeElement(&footnote, &t); err != nil {
					return err
				}
				fns.Footnotes = append(fns.Footnotes, footnote)
			}

		case xml.EndElement:
			if t.Name == start.Name {
				return nil
			}
		}
	}

	return nil
}

type Footnote struct {
	ID   int    `xml:"id,attr"`
	Type string `xml:"type,attr,omitempty"`

	Children []FootnoteChild
}

type FootnoteChild struct {
	Paragraph *Paragraph
	Table     *Table
}

func (fn Footnote) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:footnote"

	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:id"}, Value: strconv.Itoa(fn.ID)})
	if fn.Type != "" {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:type"}, Value: fn.Type})
	}

	if err := e.EncodeToken(start); err != nil {
		return err
	}

	for _, child := range fn.Children {
		if child.Paragraph != nil {
			if err := e.EncodeElement(child.Paragraph, xml.StartElement{Name: xml.Name{Local: "w:p"}}); err != nil {
				return err
			}
		} else if child.Table != nil {
			if err := e.EncodeElement(child.Table, xml.StartElement{Name: xml.Name{Local: "w:tbl"}}); err != nil {
				return err
			}
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (fn *Footnote) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "id":
			id, err := strconv.Atoi(attr.Value)
			if err != nil {
				return err
			}
			fn.ID = id
		case "type":
			fn.Type = attr.Value
		}
	}

	for {
		token, err := d.Token()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		switch t := token.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "p":
				p := &Paragraph{}
				if err := d.DecodeElement(p, &t); err != nil {
					return err
				}
				fn.Children = append(fn.Children, FootnoteChild{Paragraph: p})
			case "tbl":
				tbl := &Table{}
				if err := d.DecodeElement(tbl, &t); err != nil {
					return err
				}
				fn.Children = append(fn.Children, FootnoteChild{Table: tbl})
			default:
				if err := d.Skip(); err != nil {
					return err
				}
			}

		case xml.EndElement:
			if t.Name == start.Name {
				return nil
			}
		}
	}

	return nil
}
