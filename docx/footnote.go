package docx

import (
	"github.com/gomutex/godocx/wml/ctypes"
)

type Footnote struct {
	root *RootDoc         // root is the root document to which this footnote belongs.
	ct   *ctypes.Footnote // ct is the underlying footnote element from the wml/ctypes package.
}

func newFootnote(root *RootDoc, ct *ctypes.Footnote) *Footnote {
	return &Footnote{
		root: root,
		ct:   ct,
	}
}

func (rd *RootDoc) GetFootnoteByID(footnoteID int) *ctypes.Footnote {
	if rd.Footnotes == nil {
		return nil
	}

	for _, footnote := range rd.Footnotes.Footnotes {
		if footnote.ID == footnoteID {
			return footnote
		}
	}

	return nil
}

func (rd *RootDoc) AddFootnote() *Footnote {
	if rd.Footnotes == nil {
		rd.Footnotes = &ctypes.Footnotes{}
	}
	footnote := &ctypes.Footnote{
		ID: rd.Document.IncFootnoteID(),
	}
	rd.Footnotes.Footnotes = append(rd.Footnotes.Footnotes, footnote)
	return newFootnote(rd, footnote)
}

func (fn *Footnote) GetCT() *ctypes.Footnote {
	return fn.ct
}

func (fn *Footnote) GetID() int {
	if fn.ct == nil {
		return -1
	}
	return fn.ct.ID
}

func (fn *Footnote) AddParagraph() *Paragraph {
	if fn.ct.Children == nil {
		fn.ct.Children = make([]ctypes.FootnoteChild, 0)
	}
	p := &ctypes.Paragraph{}
	fn.ct.Children = append(fn.ct.Children, ctypes.FootnoteChild{Paragraph: p})
	dp := &Paragraph{
		root: fn.root,
		ct:   p,
	}
	dp.Style("FootnoteText")
	return dp
}
