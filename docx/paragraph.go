package docx

import (
	"encoding/xml"
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/gomutex/godocx/common/constants"
	"github.com/gomutex/godocx/common/units"
	"github.com/gomutex/godocx/dml"
	"github.com/gomutex/godocx/dml/dmlct"
	"github.com/gomutex/godocx/dml/dmlpic"
	"github.com/gomutex/godocx/internal"
	"github.com/gomutex/godocx/wml/ctypes"
	"github.com/gomutex/godocx/wml/stypes"
)

// Paragraph represents a paragraph in a DOCX document.
type Paragraph struct {
	root *RootDoc          // root is a reference to the root document.
	ct   *ctypes.Paragraph // ct holds the underlying Paragraph Complex Type.
}

func NewParagraph(root *RootDoc, ct *ctypes.Paragraph) *Paragraph {
	return &Paragraph{
		root: root,
		ct:   ct,
	}
}

func (p *Paragraph) unmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return p.ct.UnmarshalXML(d, start)
}

type paraOpts struct {
	text string
}

func newParaOpts() *paraOpts {
	return &paraOpts{}
}

// paraOption defines a type for functions that can configure a Paragraph.
type paraOption func(*Paragraph)

// newParagraph creates and initializes a new Paragraph instance with given options.
func newParagraph(root *RootDoc, opts ...paraOption) *Paragraph {
	p := &Paragraph{
		root: root,
		ct: &ctypes.Paragraph{
			Children: make([]ctypes.ParagraphChild, 0),
		},
	}
	for _, opt := range opts {
		opt(p)
	}
	return p
}

// paraWithText is an option for adding text to a Paragraph.
func paraWithText(text string) paraOption {
	return func(p *Paragraph) {
		p.AddText(text)
	}
}

func (p *Paragraph) ensureProp() {
	if p.ct.Property == nil {
		p.ct.Property = ctypes.DefaultParaProperty()
	}
}

// GetCT returns a pointer to the underlying Paragraph Complex Type.
func (p *Paragraph) GetCT() *ctypes.Paragraph {
	return p.ct
}

// AddParagraph adds a new paragraph with the specified text to the document.
// It returns the created Paragraph instance.
//
// Parameters:
//   - text: The text to be added to the paragraph.
//
// Returns:
//   - p: The created Paragraph instance.
func (rd *RootDoc) AddParagraph(text string) *Paragraph {
	p := rd.AddEmptyParagraph()

	p.AddText(text)
	return p
}

/*
@param before: Spacing Above Paragraph in twips
@param after: Spacing Below Paragraph in twips
*/
func (p *Paragraph) Spacing(before uint64, after uint64) {
	p.ensureProp()
	p.ct.Property.Spacing = ctypes.NewParagraphSpacing(before, after)
}

// Style sets the paragraph style.
//
// Parameters:
//   - value: A string representing the style value. It can be any valid style defined in the WordprocessingML specification.
//
// Example:
//
//	p1 := document.AddParagraph("Example para")
//	paragraph.Style("List Number")
func (p *Paragraph) Style(value string) {
	p.ensureProp()
	p.ct.Property.Style = ctypes.NewParagraphStyle(value)
}

// Justification sets the paragraph justification type.
//
// Parameters:
//   - value: A value of type stypes.Justification representing the justification type.
//     It can be one of the Justification type values defined in the stypes package.
//
// Example:
//
//	p1 := document.AddParagraph("Example justified para")
//	p1.Justification(stypes.JustificationCenter) // Center justification
func (p *Paragraph) Justification(value stypes.Justification) {
	p.ensureProp()

	p.ct.Property.Justification = ctypes.NewGenSingleStrVal(value)
}

// Numbering sets the paragraph numbering properties.
//
// This function assigns a numbering definition ID and a level to the paragraph,
// which affects how numbering is displayed in the document.
//
// Parameters:
//   - id: An integer representing the numbering definition ID.
//   - level: An integer representing the level within the numbering definition.
//
// Example:
//
//	p1 := document.AddParagraph("Example numbered para")
//	p1.Numbering(1, 0)
//
// In this example, the paragraph p1 is assigned the numbering properties
// defined by numbering definition ID 1 and level 0.
func (p *Paragraph) Numbering(id int, level int) {

	p.ensureProp()

	if p.ct.Property.NumProp == nil {
		p.ct.Property.NumProp = &ctypes.NumProp{}
	}

	p.ct.Property.NumProp.NumID = ctypes.NewDecimalNum(id)
	p.ct.Property.NumProp.ILvl = ctypes.NewDecimalNum(level)
}

// Indent sets the paragraph indentation properties.
//
// This function assigns an indent definition to the paragraph,
// which affects how exactly the paragraph is going to be indented.
//
// Parameters:
//   - indentProp: A ctypes.Indent instance pointer representing exact indentation
//     measurements to use.
//
// Example:
//
//	var size360 int = 360
//	var sizeu420 uint64 = 420
//	indent360 := ctypes.Indent{Left: &size360, Hanging: &sizeu420}
//
//	p1 := document.AddParagraph("Example indented para")
//	p1.Indent(&indent360)
func (p *Paragraph) Indent(indentProp *ctypes.Indent) {

	p.ensureProp()

	p.ct.Property.Indent = indentProp
}

// Appends a new text to the Paragraph.
// Example:
//
//	paragraph := AddParagraph()
//	modifiedParagraph := paragraph.AddText("Hello, World!")
//
// Parameters:
//   - text: A string representing the text to be added to the Paragraph.
//
// Returns:
//   - *Run: The newly created Run instance added to the Paragraph.
func (p *Paragraph) AddText(text string) *Run {
	t := ctypes.TextFromString(text)

	runChildren := []ctypes.RunChild{}
	runChildren = append(runChildren, ctypes.RunChild{
		Text: t,
	})
	run := &ctypes.Run{
		Children: runChildren,
	}

	p.ct.Children = append(p.ct.Children, ctypes.ParagraphChild{Run: run})

	return newRun(p.root, run)
}

// AddEmptyParagraph adds a new empty paragraph to the document.
// It returns the created Paragraph instance.
//
// Returns:
//   - p: The created Paragraph instance.
func (rd *RootDoc) AddEmptyParagraph() *Paragraph {
	p := newParagraph(rd)

	bodyElem := DocumentChild{
		Para: p,
	}
	rd.Document.Body.Children = append(rd.Document.Body.Children, bodyElem)

	return p
}

// AddRun adds a new Run to the Paragraph.
//
// Returns:
//   - run: The newly created Run instance added to the Paragraph.
func (p *Paragraph) AddRun() *Run {
	run := &ctypes.Run{}

	p.ct.Children = append(p.ct.Children, ctypes.ParagraphChild{Run: run})

	return newRun(p.root, run)
}

// AddIns adds a new inserted RunTrackChange to the Paragraph.
//
// Returns:
//   - *RunTrackChange: The newly created RunTrackChange instance added to the Paragraph.
func (p *Paragraph) AddIns(author string, date *string) *RunTrackChange {
	id := p.root.Document.IncAnnotationID()
	ins := ctypes.NewRunTrackChangeIns(id, author, date)
	p.ct.Children = append(p.ct.Children, ctypes.ParagraphChild{
		Change: ins,
	})
	return newRunTrackChange(p.root, ins)
}

// AddDel adds a new deleted RunTrackChange to the Paragraph.
//
// Returns:
//   - *RunTrackChange: The newly created RunTrackChange instance added to the Paragraph.
func (p *Paragraph) AddDel(author string, date *string) *RunTrackChange {
	id := p.root.Document.IncAnnotationID()
	del := ctypes.NewRunTrackChangeDel(id, author, date)
	p.ct.Children = append(p.ct.Children, ctypes.ParagraphChild{
		Change: del,
	})
	return newRunTrackChange(p.root, del)
}

// GetStyle retrieves the style information applied to the Paragraph.
//
// Returns:
//   - *ctypes.Style: The style information of the Paragraph.
//   - error: An error if the style information is not found.
func (p *Paragraph) GetStyle() (*ctypes.Style, error) {
	if p.ct.Property == nil || p.ct.Property.Style == nil {
		return nil, errors.New("No property for the style")
	}

	style := p.root.GetStyleByID(p.ct.Property.Style.Val, stypes.StyleTypeParagraph)
	if style == nil {
		return nil, errors.New("No style found for the paragraph")
	}

	return style, nil
}

// AddLink adds a external hyperlink to the Paragraph.
//
// Parameters:
//   - text: The text to be displayed for the hyperlink.
//   - link: The URL or location of the hyperlink.
//
// Returns:
//   - *Hyperlink: The created Hyperlink instance representing the added link.
func (p *Paragraph) AddLink(text string, link string) *Hyperlink {
	rId := p.root.Document.addLinkRelation(link)

	runChildren := []ctypes.RunChild{}
	runChildren = append(runChildren, ctypes.RunChild{
		Text: ctypes.TextFromString(text),
	})
	run := &ctypes.Run{
		Children: runChildren,
		Property: &ctypes.RunProperty{
			Style: &ctypes.CTString{
				Val: constants.HyperLinkStyle,
			},
		},
	}

	hyperLink := &ctypes.Hyperlink{
		ID:  rId,
		Run: run,
	}

	p.ct.Children = append(p.ct.Children, ctypes.ParagraphChild{Link: hyperLink})

	return newHyperlink(p.root, hyperLink)
}

// AddInternalLink adds a hyperlink to an internal anchor within the document.
//
// Parameters:
//   - text: The text to be displayed for the hyperlink.
//   - anchor: Specifies the name of a bookmark within the document. If it is not a
//     valid bookmark, the default behavior is to navigate to the start of
//     the document. Example bookmark: <w:bookmarkStart w:id="0" w:name="exampleAnchorName"/>
//
// Returns:
//   - *Hyperlink: The created Hyperlink instance representing the internal link.
func (p *Paragraph) AddInternalLink(text string, anchor string) *Hyperlink {
	runChildren := []ctypes.RunChild{}
	runChildren = append(runChildren, ctypes.RunChild{
		Text: ctypes.TextFromString(text),
	})
	run := &ctypes.Run{
		Children: runChildren,
		Property: &ctypes.RunProperty{
			Style: &ctypes.CTString{
				Val: constants.HyperLinkStyle,
			},
		},
	}

	hyperLink := &ctypes.Hyperlink{
		Anchor: anchor,
		Run:    run,
	}

	p.ct.Children = append(p.ct.Children, ctypes.ParagraphChild{Link: hyperLink})

	return newHyperlink(p.root, hyperLink)
}

// AddDrawing adds a new drawing (image) to the Paragraph.
//
// Parameters:
//   - rID: The relationship ID of the image in the document.
//   - imgCount: The count of images in the document.
//   - width: The width of the image in inches.
//   - height: The height of the image in inches.
//
// Returns:
//   - *dml.Inline: The created Inline instance representing the added drawing.
func (p *Paragraph) addDrawing(rID string, imgCount uint, width units.Inch, height units.Inch) *dml.Inline {
	eWidth := width.ToEmu()
	eHeight := height.ToEmu()

	inline := dml.NewInline(
		*dmlct.NewPostvSz2D(eWidth, eHeight),
		dml.DocProp{
			ID:   uint64(imgCount),
			Name: fmt.Sprintf("Image%d", imgCount),
		},
		*dml.NewPicGraphic(dmlpic.NewPic(rID, imgCount, eWidth, eHeight)),
	)

	runChildren := []ctypes.RunChild{}
	drawing := &dml.Drawing{}

	drawing.Inline = append(drawing.Inline, inline)

	runChildren = append(runChildren, ctypes.RunChild{
		Drawing: drawing,
	})

	run := &ctypes.Run{
		Children: runChildren,
	}

	p.ct.Children = append(p.ct.Children, ctypes.ParagraphChild{Run: run})

	return inline
}

func (p *Paragraph) AddPicture(path string, width units.Inch, height units.Inch) (*PicMeta, error) {
	imgBytes, err := internal.FileToByte(path)
	if err != nil {
		return nil, err
	}

	imgExt := filepath.Ext(path)
	p.root.ImageCount += 1
	fileName := fmt.Sprintf("image%d%s", p.root.ImageCount, imgExt)
	fileIdxPath := fmt.Sprintf("%s%s", constants.MediaPath, fileName)

	imgExtStripDot := strings.TrimPrefix(imgExt, ".")
	imgMIME, err := MIMEFromExt(imgExtStripDot)
	if err != nil {
		return nil, err
	}

	err = p.root.ContentType.AddExtension(imgExtStripDot, imgMIME)
	if err != nil {
		return nil, err
	}

	overridePart := fmt.Sprintf("/%s%s", constants.MediaPath, fileName)
	err = p.root.ContentType.AddOverride(overridePart, imgMIME)
	if err != nil {
		return nil, err
	}

	p.root.FileMap.Store(fileIdxPath, imgBytes)

	relName := fmt.Sprintf("media/%s", fileName)

	rID := p.root.Document.addRelation(constants.SourceRelationshipImage, relName)

	inline := p.addDrawing(rID, p.root.ImageCount, width, height)

	return &PicMeta{
		Para:   p,
		Inline: inline,
	}, nil
}

// AddBookmarkStart adds a bookmark start to the Paragraph.
//
// Parameters:
//   - name: The name of the bookmark.
//
// Returns:
//   - int: The ID of the bookmark start.
//     This ID is used to uniquely identify the bookmark within the document.
func (p *Paragraph) AddBookmarkStart(name string) int {
	id := p.root.Document.IncBookmarkID()

	bookmarkStart := &ctypes.BookmarkStart{
		ID:   id, // ID is not used in the current implementation
		Name: name,
	}

	p.ct.Children = append(p.ct.Children, ctypes.ParagraphChild{
		Bookmark: &ctypes.Bookmark{
			Start: bookmarkStart,
		},
	})

	return id
}

// AddBookmarkEnd adds a bookmark end to the Paragraph.
//
// Parameters:
//   - id: The ID of the bookmark to end.
func (p *Paragraph) AddBookmarkEnd(id int) {
	bookmarkEnd := &ctypes.BookmarkEnd{
		ID: id, // ID is not used in the current implementation
	}

	p.ct.Children = append(p.ct.Children, ctypes.ParagraphChild{
		Bookmark: &ctypes.Bookmark{
			End: bookmarkEnd,
		},
	})
}

// AddFootnoteRef adds a footnote reference mark to the Paragraph.
//
// 17.11.13 footnoteRef (Footnote Reference Mark)
//
// This method should be used when you want to insert a reference to a
// footnote content. For document body's paragraphs, please use
// [AddFootnoteReference] instead.
func (p *Paragraph) AddFootnoteRef() {
	run := &ctypes.Run{
		Property: &ctypes.RunProperty{
			Style: &ctypes.CTString{
				Val: "FootnoteReference",
			},
		},
		Children: []ctypes.RunChild{
			{
				FootnoteRef: &ctypes.Empty{},
			},
		},
	}

	p.ct.Children = append(p.ct.Children, ctypes.ParagraphChild{Run: run})
}

// AddFootnoteReference adds a footnote reference to the paragraph.
func (p *Paragraph) AddFootnoteReference(id int) {
	run := &ctypes.Run{
		Property: &ctypes.RunProperty{
			Style: &ctypes.CTString{
				Val: "FootnoteReference",
			},
		},
		Children: []ctypes.RunChild{
			{
				FootnoteReference: &ctypes.Markup{
					ID: id,
				},
			},
		},
	}

	p.ct.Children = append(p.ct.Children, ctypes.ParagraphChild{Run: run})
}

// ClonePropertyToRevision creates a deep copy of the current paragraph properties
// and assigns it to a PPrChange structure. If there's an existing PPrChange,
// it will be replaced with the new cloned properties.
//
// This method is useful for revision tracking where you want to preserve the
// current state of paragraph properties before making changes.
//
// Parameters:
//   - id: The revision ID for tracking purposes
//   - author: The author of the revision
//   - date: Optional date string for the revision (can be nil)
//
// Example:
//
//	paragraph.Style("Heading1")
//	paragraph.Justification(stypes.JustificationCenter)
//	// Now clone current properties to revision
//	paragraph.ClonePropertyToRevision(1, "John Doe", nil)
//	// Make changes to current properties
//	paragraph.Style("Normal")
//	paragraph.Justification(stypes.JustificationLeft)
func (p *Paragraph) ClonePropertyToRevision(author string, date *string) {
	p.ensureProp()

	id := p.root.Document.IncAnnotationID()

	// Clone current properties (excluding existing PPrChange)
	clonedProps := p.ct.Property.Clone()

	// Create new PPrChange with cloned properties
	p.ct.Property.PPrChange = &ctypes.PPrChange{
		TrackChange: ctypes.TrackChange{
			ID:     id,
			Author: author,
			Date:   date,
		},
		ParaProp: clonedProps,
	}
}

func (p *Paragraph) scanBookmarkIds() {
	for _, child := range p.ct.Children {
		if child.Bookmark != nil {
			if child.Bookmark.Start != nil {
				p.root.Document.UpdateBookmarkID(child.Bookmark.Start.ID)
			}
			if child.Bookmark.End != nil {
				p.root.Document.UpdateBookmarkID(child.Bookmark.End.ID)
			}
		}
		if child.Link != nil {
			l := Hyperlink{
				root: p.root,
				ct:   child.Link,
			}
			l.scanBookmarkIds()
		}
	}
}

func (p *Paragraph) scanAnnotationIds() {
	for _, child := range p.ct.Children {
		if child.Run != nil {
			if child.Run.Property != nil && child.Run.Property.RPrChange != nil {
				p.root.Document.updateAnnotationID(child.Run.Property.RPrChange.ID)
			}
		}
		if child.Change != nil {
			p.root.Document.updateAnnotationID(child.Change.ID)
		}
	}
}
