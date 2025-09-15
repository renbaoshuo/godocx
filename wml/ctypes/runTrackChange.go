package ctypes

import (
	"encoding/xml"
)

type RunTrackChangeType int

const (
	RunTrackChangeTypeInsert RunTrackChangeType = iota
	RunTrackChangeTypeDelete
)

func (t RunTrackChangeType) Tag() string {
	switch t {
	case RunTrackChangeTypeInsert:
		return "w:ins"
	case RunTrackChangeTypeDelete:
		return "w:del"
	default:
		return "w:ins" // Default to insert if unknown type
	}
}

func TagLocalNameToRunTrackChangeType(tag string) RunTrackChangeType {
	switch tag {
	case "ins":
		return RunTrackChangeTypeInsert
	case "del":
		return RunTrackChangeTypeDelete
	default:
		return RunTrackChangeTypeInsert // Default to insert if unknown tag
	}
}

type RunTrackChange struct {
	TrackChange

	Type     RunTrackChangeType `xml:"-"`
	Children []RunTrackChangeChild
}

type RunTrackChangeChild struct {
	Run *Run
}

func NewRunTrackChangeIns(id int, author string, date *string) *RunTrackChange {
	return &RunTrackChange{
		TrackChange: TrackChange{
			ID:     id,
			Author: author,
			Date:   date,
		},
		Type:     RunTrackChangeTypeInsert,
		Children: []RunTrackChangeChild{},
	}
}

func NewRunTrackChangeDel(id int, author string, date *string) *RunTrackChange {
	return &RunTrackChange{
		TrackChange: TrackChange{
			ID:     id,
			Author: author,
			Date:   date,
		},
		Type:     RunTrackChangeTypeDelete,
		Children: []RunTrackChangeChild{},
	}
}

func (rtc RunTrackChange) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = rtc.Type.Tag()

	if err := rtc.TrackChange.MarshalXmlAttrs(&start); err != nil {
		return err
	}

	if err := e.EncodeToken(start); err != nil {
		return err
	}

	for _, child := range rtc.Children {
		if child.Run != nil {
			if err := e.EncodeElement(child.Run, xml.StartElement{Name: xml.Name{Local: "w:r"}}); err != nil {
				return err
			}
		}
	}

	return e.EncodeToken(start.End())
}

func (rtc *RunTrackChange) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	rtc.Type = TagLocalNameToRunTrackChangeType(start.Name.Local)

	rtc.TrackChange = TrackChange{}
	if err := rtc.TrackChange.UnmarshalXMLAttrs(start.Attr); err != nil {
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
			case "r":
				r := NewRun()
				if err = d.DecodeElement(r, &el); err != nil {
					return err
				}
				rtc.Children = append(rtc.Children, RunTrackChangeChild{Run: r})
			}
		case xml.EndElement:
			if el.Name.Local == start.Name.Local {
				return nil
			}
		}
	}
}
