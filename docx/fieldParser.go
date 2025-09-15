package docx

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gomutex/godocx/common/constants"
	"github.com/gomutex/godocx/wml/ctypes"
	"github.com/gomutex/godocx/wml/stypes"
)

// FieldParser processes complex fields in a paragraph and converts them to appropriate formats
type FieldParser struct {
	root *RootDoc
}

// NewFieldParser creates a new field parser
func NewFieldParser(root *RootDoc) *FieldParser {
	return &FieldParser{root: root}
}

// ParseParagraphFields processes all complex fields in a paragraph and converts them
func (fp *FieldParser) ParseParagraphFields(para *Paragraph) error {
	if para == nil || para.ct.Children == nil {
		return nil
	}

	fields := fp.extractFields(para.ct.Children)

	for i := len(fields) - 1; i >= 0; i-- {
		field := fields[i]
		if err := fp.convertField(para, field); err != nil {
			return fmt.Errorf("failed to convert field: %w", err)
		}
	}

	return nil
}

// FieldSequence represents a complete field from begin to end
type FieldSequence struct {
	BeginIndex    int
	SeparateIndex int
	EndIndex      int
	InstrText     string
	DisplayRuns   []*ctypes.Run
}

func (fp *FieldParser) extractFields(children []ctypes.ParagraphChild) []FieldSequence {
	var fields []FieldSequence
	var currentField *FieldSequence

	for i, child := range children {
		if child.Run == nil {
			continue
		}

		for _, runChild := range child.Run.Children {
			if runChild.FldChar != nil {
				switch runChild.FldChar.FldCharType {
				case stypes.FldCharTypeBegin:
					currentField = &FieldSequence{
						BeginIndex:    i,
						SeparateIndex: -1,
						EndIndex:      -1,
					}

				case stypes.FldCharTypeSeparate:
					if currentField != nil {
						currentField.SeparateIndex = i
					}

				case stypes.FldCharTypeEnd:
					if currentField != nil {
						currentField.EndIndex = i

						fp.extractFieldContent(children, currentField)

						fields = append(fields, *currentField)
						currentField = nil
					}
				}
			}
		}
	}

	return fields
}

func (fp *FieldParser) extractFieldContent(children []ctypes.ParagraphChild, field *FieldSequence) {
	var instrText strings.Builder
	separateIndex := field.SeparateIndex
	if separateIndex == -1 {
		separateIndex = field.EndIndex
	}

	for i := field.BeginIndex + 1; i < separateIndex; i++ {
		if children[i].Run != nil {
			for _, runChild := range children[i].Run.Children {
				if runChild.InstrText != nil && runChild.InstrText.Text != "" {
					instrText.WriteString(runChild.InstrText.Text)
				}
			}
		}
	}
	field.InstrText = strings.TrimSpace(instrText.String())

	var displayRuns []*ctypes.Run

	if field.SeparateIndex != -1 {
		for i := field.SeparateIndex + 1; i < field.EndIndex; i++ {
			if children[i].Run != nil {
				displayRuns = append(displayRuns, children[i].Run)
			}
		}
	}

	field.DisplayRuns = displayRuns
}

func (fp *FieldParser) convertField(para *Paragraph, field FieldSequence) error {
	// Check if this is a HYPERLINK field
	if strings.HasPrefix(strings.ToUpper(field.InstrText), "HYPERLINK") {
		return fp.convertHyperlinkField(para, field)
	}

	// Add support for other field types here in the future

	return nil
}

func (fp *FieldParser) convertHyperlinkField(para *Paragraph, field FieldSequence) error {
	if fp.isInternalLink(field.InstrText) {
		anchor := fp.parseHyperlinkAnchor(field.InstrText)
		// if anchor == "" {
		// 	return fmt.Errorf("failed to parse anchor from internal HYPERLINK field: %s", field.InstrText)
		// }

		fp.removeFieldRuns(para, field)

		hyperlink := fp.createInternalHyperlink(anchor, field.DisplayRuns)

		insertIndex := field.BeginIndex
		para.ct.Children = append(
			para.ct.Children[:insertIndex],
			append([]ctypes.ParagraphChild{{Link: hyperlink}}, para.ct.Children[insertIndex:]...)...,
		)

		return nil
	}

	// Handle external links
	url := fp.parseHyperlinkURL(field.InstrText)
	// if url == "" {
	// 	return fmt.Errorf("failed to parse URL from HYPERLINK field: %s", field.InstrText)
	// }

	fp.removeFieldRuns(para, field)

	hyperlink := fp.createExternalHyperlink(url, field.DisplayRuns)

	insertIndex := field.BeginIndex
	para.ct.Children = append(
		para.ct.Children[:insertIndex],
		append([]ctypes.ParagraphChild{{Link: hyperlink}}, para.ct.Children[insertIndex:]...)...,
	)

	return nil
}

func (fp *FieldParser) parseHyperlinkURL(instrText string) string {
	re := regexp.MustCompile(`HYPERLINK\s+"([^"]+)"`)
	matches := re.FindStringSubmatch(instrText)

	if len(matches) >= 2 {
		return matches[1]
	}

	re2 := regexp.MustCompile(`HYPERLINK\s+(\S+)`)
	matches2 := re2.FindStringSubmatch(instrText)

	if len(matches2) >= 2 {
		if matches2[1] == "\\h" {
			return ""
		}
		return matches2[1]
	}

	return ""
}

func (fp *FieldParser) parseHyperlinkAnchor(instrText string) string {
	re := regexp.MustCompile(`(?i)HYPERLINK\s+\\l\s+"([^"]+)"`)
	matches := re.FindStringSubmatch(instrText)

	if len(matches) >= 2 {
		return matches[1]
	}

	re2 := regexp.MustCompile(`(?i)HYPERLINK\s+\\l\s+(\S+)`)
	matches2 := re2.FindStringSubmatch(instrText)

	if len(matches2) >= 2 {
		if matches2[1] == "\\h" {
			return ""
		}
		return matches2[1]
	}

	return ""
}

func (fp *FieldParser) isInternalLink(instrText string) bool {
	return strings.Contains(strings.ToUpper(instrText), "\\L")
}

func (fp *FieldParser) removeFieldRuns(para *Paragraph, field FieldSequence) {
	for i := field.EndIndex; i >= field.BeginIndex; i-- {
		para.ct.Children = append(para.ct.Children[:i], para.ct.Children[i+1:]...)
	}
}

func (fp *FieldParser) createExternalHyperlink(url string, displayRuns []*ctypes.Run) *ctypes.Hyperlink {
	rId := fmt.Sprintf("rId%d", fp.root.Document.IncRelationID())

	rel := &Relationship{
		ID:         rId,
		Type:       constants.SourceRelationshipHyperLink,
		Target:     url,
		TargetMode: "External",
	}
	fp.root.Document.DocRels.Relationships = append(fp.root.Document.DocRels.Relationships, rel)

	hyperlink := &ctypes.Hyperlink{
		ID: rId,
		Run: &ctypes.Run{
			Children: []ctypes.RunChild{{
				Text: ctypes.TextFromString(""),
			}},
			Property: &ctypes.RunProperty{
				Style: &ctypes.CTString{
					Val: constants.HyperLinkStyle,
				},
			},
		},
	}

	if len(displayRuns) > 0 {
		for _, run := range displayRuns {
			runCopy := &ctypes.Run{
				Children: make([]ctypes.RunChild, len(run.Children)),
				Property: run.Property,
			}
			copy(runCopy.Children, run.Children)

			hyperlink.Children = append(hyperlink.Children, ctypes.ParagraphChild{
				Run: runCopy,
			})
		}
	}

	return hyperlink
}

func (fp *FieldParser) createInternalHyperlink(anchor string, displayRuns []*ctypes.Run) *ctypes.Hyperlink {
	hyperlink := &ctypes.Hyperlink{
		Anchor: anchor,
		Run: &ctypes.Run{
			Children: []ctypes.RunChild{{
				Text: ctypes.TextFromString(""),
			}},
			Property: &ctypes.RunProperty{
				Style: &ctypes.CTString{
					Val: constants.HyperLinkStyle,
				},
			},
		},
	}

	if len(displayRuns) > 0 {
		for _, run := range displayRuns {
			runCopy := &ctypes.Run{
				Children: make([]ctypes.RunChild, len(run.Children)),
				Property: run.Property,
			}
			copy(runCopy.Children, run.Children)

			hyperlink.Children = append(hyperlink.Children, ctypes.ParagraphChild{
				Run: runCopy,
			})
		}
	}

	return hyperlink
}
