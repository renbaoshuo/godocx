package docx

import (
	"github.com/gomutex/godocx/common/units"
	"github.com/gomutex/godocx/dml"
	"github.com/gomutex/godocx/dml/dmlpic"
)

type PicMeta struct {
	Para   *Paragraph
	Inline *dml.Inline
}

// AddPicture adds a new image to the document.
//
// Example usage:
//
//	// Add a picture to the document
//	_, err = document.AddPicture("gopher.png", units.Inch(2.9), units.Inch(2.9))
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// Parameters:
//   - path: The path of the image file to be added.
//   - width: The width of the image in inches.
//   - height: The height of the image in inches.
//
// Returns:
//   - *PicMeta: Metadata about the added picture, including the Paragraph instance and Inline element.
//   - error: An error, if any occurred during the process.
func (rd *RootDoc) AddPicture(path string, width units.Inch, height units.Inch) (*PicMeta, error) {

	p := newParagraph(rd)

	bodyElem := DocumentChild{
		Para: p,
	}
	rd.Document.Body.Children = append(rd.Document.Body.Children, bodyElem)

	return p.AddPicture(path, width, height)
}

func (pm *PicMeta) SetAltText(title, desc string) {
	if pm.Inline == nil {
		return
	}

	// 1. Set docPr attributes
	// This is the main storage of alt text in WordML
	pm.Inline.DocProp.Name = title
	pm.Inline.DocProp.Description = desc

	// 2. Set title in graphic data
	// This is used by some applications (e.g., PowerPoint) to show alt text
	if pm.Inline.Graphic.Data == nil {
		pm.Inline.Graphic.Data = &dml.GraphicData{}
	}
	if pm.Inline.Graphic.Data.Pic == nil {
		pm.Inline.Graphic.Data.Pic = &dmlpic.Pic{}
	}
	pm.Inline.Graphic.Data.Pic.NonVisualPicProp.CNvPr.Name = title
	pm.Inline.Graphic.Data.Pic.NonVisualPicProp.CNvPr.Description = desc
}
