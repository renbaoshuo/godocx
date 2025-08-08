package docx

import (
	"github.com/gomutex/godocx/wml/ctypes"
)

func (rd *RootDoc) GetFootnoteByID(footnoteID int) *ctypes.Footnote {
	if rd.Footnotes == nil {
		return nil
	}

	for _, footnote := range rd.Footnotes.Footnotes {
		if footnote.ID == footnoteID {
			return &footnote
		}
	}

	return nil
}

// TODO: refactor to use a better way to manage footnote content
func (rd *RootDoc) AddFootnote(footnote ctypes.Footnote) int {
	if rd.Footnotes == nil {
		rd.Footnotes = &ctypes.Footnotes{}
	}
	id := rd.Document.IncFootnoteID()
	footnote.ID = id
	rd.Footnotes.Footnotes = append(rd.Footnotes.Footnotes, footnote)
	return id
}
