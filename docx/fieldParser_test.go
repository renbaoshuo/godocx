package docx

import (
	"testing"

	"github.com/gomutex/godocx/wml/ctypes"
	"github.com/gomutex/godocx/wml/stypes"
)

func TestFieldParser_ParseHyperlinkField(t *testing.T) {
	// Create a test document
	root := &RootDoc{
		Document: &Document{
			DocRels: Relationships{
				Relationships: []*Relationship{},
			},
			RID: 0,
		},
	}

	// Create a paragraph with a HYPERLINK field
	para := &Paragraph{
		root: root,
		ct: &ctypes.Paragraph{
			Children: []ctypes.ParagraphChild{
				// w:r with fldChar begin
				{
					Run: &ctypes.Run{
						Children: []ctypes.RunChild{
							{
								FldChar: &ctypes.FldChar{
									FldCharType: stypes.FldCharTypeBegin,
								},
							},
						},
					},
				},
				// w:r with instrText
				{
					Run: &ctypes.Run{
						Children: []ctypes.RunChild{
							{
								InstrText: ctypes.TextFromString(` HYPERLINK "http://www.example.com" `),
							},
						},
					},
				},
				// w:r with fldChar separate
				{
					Run: &ctypes.Run{
						Children: []ctypes.RunChild{
							{
								FldChar: &ctypes.FldChar{
									FldCharType: stypes.FldCharTypeSeparate,
								},
							},
						},
					},
				},
				// w:r with display text
				{
					Run: &ctypes.Run{
						Property: &ctypes.RunProperty{
							Style: &ctypes.CTString{
								Val: "Hyperlink",
							},
						},
						Children: []ctypes.RunChild{
							{
								Text: ctypes.TextFromString("示例链接"),
							},
						},
					},
				},
				// w:r with fldChar end
				{
					Run: &ctypes.Run{
						Children: []ctypes.RunChild{
							{
								FldChar: &ctypes.FldChar{
									FldCharType: stypes.FldCharTypeEnd,
								},
							},
						},
					},
				},
			},
		},
	}

	// Create field parser and process the paragraph
	parser := NewFieldParser(root)
	err := parser.ParseParagraphFields(para)
	if err != nil {
		t.Fatalf("Failed to parse paragraph fields: %v", err)
	}

	// Verify that the field was converted to a hyperlink
	if len(para.ct.Children) != 1 {
		t.Errorf("Expected 1 child after field conversion, got %d", len(para.ct.Children))
	}

	if para.ct.Children[0].Link == nil {
		t.Error("Expected a hyperlink after field conversion")
	}

	hyperlink := para.ct.Children[0].Link
	if hyperlink.ID == "" {
		t.Error("Expected hyperlink to have an ID")
	}

	// Check that hyperlink has children (runs with display text)
	if len(hyperlink.Children) == 0 {
		t.Error("Expected hyperlink to have children runs")
	}

	// Get the first run from hyperlink children
	firstRun := hyperlink.Children[0].Run
	if firstRun == nil {
		t.Error("Expected first hyperlink child to be a run")
	}

	if len(firstRun.Children) != 1 {
		t.Errorf("Expected hyperlink run to have 1 child, got %d", len(firstRun.Children))
	}

	if firstRun.Children[0].Text == nil {
		t.Error("Expected hyperlink run child to have text")
	}

	if firstRun.Children[0].Text.Text != "示例链接" {
		t.Errorf("Expected hyperlink text to be '示例链接', got '%s'", firstRun.Children[0].Text.Text)
	}

	// Verify that a relationship was created
	if len(root.Document.DocRels.Relationships) != 1 {
		t.Errorf("Expected 1 relationship after field conversion, got %d", len(root.Document.DocRels.Relationships))
	}

	rel := root.Document.DocRels.Relationships[0]
	if rel.Target != "http://www.example.com" {
		t.Errorf("Expected relationship target to be 'http://www.example.com', got '%s'", rel.Target)
	}

	if rel.TargetMode != "External" {
		t.Errorf("Expected relationship target mode to be 'External', got '%s'", rel.TargetMode)
	}
}

func TestFieldParser_ParseHyperlinkURL(t *testing.T) {
	parser := &FieldParser{}

	tests := []struct {
		name      string
		instrText string
		expected  string
	}{
		{
			name:      "URL with quotes",
			instrText: ` HYPERLINK "http://www.example.com" `,
			expected:  "http://www.example.com",
		},
		{
			name:      "URL without quotes",
			instrText: "HYPERLINK http://www.google.com",
			expected:  "http://www.google.com",
		},
		{
			name:      "Complex URL with parameters",
			instrText: ` HYPERLINK "https://www.example.com/page?param=value&other=123" `,
			expected:  "https://www.example.com/page?param=value&other=123",
		},
		{
			name:      "Invalid instruction",
			instrText: "NOT_A_HYPERLINK",
			expected:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parser.parseHyperlinkURL(tt.instrText)
			if result != tt.expected {
				t.Errorf("parseHyperlinkURL() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestFieldParser_ParseInternalHyperlinkField(t *testing.T) {
	// Create a test document
	root := &RootDoc{
		Document: &Document{
			DocRels: Relationships{
				Relationships: []*Relationship{},
			},
			RID: 0,
		},
	}

	// Create a paragraph with an internal HYPERLINK field
	para := &Paragraph{
		root: root,
		ct: &ctypes.Paragraph{
			Children: []ctypes.ParagraphChild{
				// w:r with fldChar begin
				{
					Run: &ctypes.Run{
						Children: []ctypes.RunChild{
							{
								FldChar: &ctypes.FldChar{
									FldCharType: stypes.FldCharTypeBegin,
								},
							},
						},
					},
				},
				// w:r with instrText for internal link
				{
					Run: &ctypes.Run{
						Children: []ctypes.RunChild{
							{
								InstrText: ctypes.TextFromString(` HYPERLINK \l "bookmark1" `),
							},
						},
					},
				},
				// w:r with fldChar separate
				{
					Run: &ctypes.Run{
						Children: []ctypes.RunChild{
							{
								FldChar: &ctypes.FldChar{
									FldCharType: stypes.FldCharTypeSeparate,
								},
							},
						},
					},
				},
				// w:r with display text
				{
					Run: &ctypes.Run{
						Property: &ctypes.RunProperty{
							Style: &ctypes.CTString{
								Val: "Hyperlink",
							},
						},
						Children: []ctypes.RunChild{
							{
								Text: ctypes.TextFromString("跳转到书签"),
							},
						},
					},
				},
				// w:r with fldChar end
				{
					Run: &ctypes.Run{
						Children: []ctypes.RunChild{
							{
								FldChar: &ctypes.FldChar{
									FldCharType: stypes.FldCharTypeEnd,
								},
							},
						},
					},
				},
			},
		},
	}

	// Create field parser and process the paragraph
	parser := NewFieldParser(root)
	err := parser.ParseParagraphFields(para)
	if err != nil {
		t.Fatalf("Failed to parse paragraph fields: %v", err)
	}

	// Verify that the field was converted to an internal hyperlink
	if len(para.ct.Children) != 1 {
		t.Errorf("Expected 1 child after field conversion, got %d", len(para.ct.Children))
	}

	if para.ct.Children[0].Link == nil {
		t.Error("Expected a hyperlink after field conversion")
	}

	hyperlink := para.ct.Children[0].Link
	if hyperlink.Anchor != "bookmark1" {
		t.Errorf("Expected hyperlink anchor to be 'bookmark1', got '%s'", hyperlink.Anchor)
	}

	if hyperlink.ID != "" {
		t.Error("Expected internal hyperlink to have no ID (no external relationship)")
	}

	// Check that hyperlink has children (runs with display text)
	if len(hyperlink.Children) == 0 {
		t.Error("Expected hyperlink to have children runs")
	}

	// Get the first run from hyperlink children
	firstRun := hyperlink.Children[0].Run
	if firstRun == nil {
		t.Error("Expected first hyperlink child to be a run")
	}

	if firstRun.Children[0].Text.Text != "跳转到书签" {
		t.Errorf("Expected hyperlink text to be '跳转到书签', got '%s'", firstRun.Children[0].Text.Text)
	}

	// Verify that no external relationship was created for internal link
	if len(root.Document.DocRels.Relationships) != 0 {
		t.Errorf("Expected 0 relationships for internal link, got %d", len(root.Document.DocRels.Relationships))
	}
}

func TestFieldParser_ParseHyperlinkAnchor(t *testing.T) {
	parser := &FieldParser{}

	tests := []struct {
		name      string
		instrText string
		expected  string
	}{
		{
			name:      "Anchor with quotes",
			instrText: ` HYPERLINK \l "bookmark1" `,
			expected:  "bookmark1",
		},
		{
			name:      "Anchor without quotes",
			instrText: `HYPERLINK \l bookmark2`,
			expected:  "bookmark2",
		},
		{
			name:      "Complex anchor name",
			instrText: ` HYPERLINK \l "my_bookmark_123" `,
			expected:  "my_bookmark_123",
		},
		{
			name:      "Invalid instruction",
			instrText: "HYPERLINK http://example.com",
			expected:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parser.parseHyperlinkAnchor(tt.instrText)
			if result != tt.expected {
				t.Errorf("parseHyperlinkAnchor() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestFieldParser_IsInternalLink(t *testing.T) {
	parser := &FieldParser{}

	tests := []struct {
		name      string
		instrText string
		expected  bool
	}{
		{
			name:      "Internal link with \\l",
			instrText: ` HYPERLINK \l "bookmark1" `,
			expected:  true,
		},
		{
			name:      "External link",
			instrText: ` HYPERLINK "http://example.com" `,
			expected:  false,
		},
		{
			name:      "Mixed case internal link",
			instrText: `HYPERLINK \L bookmark2`,
			expected:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parser.isInternalLink(tt.instrText)
			if result != tt.expected {
				t.Errorf("isInternalLink() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestFieldParser_MultiRunStyling(t *testing.T) {
	// Create a test document
	root := &RootDoc{
		Document: &Document{
			DocRels: Relationships{
				Relationships: []*Relationship{},
			},
			RID: 0,
		},
	}

	// Create a paragraph with a HYPERLINK field that has multiple runs with different styles
	para := &Paragraph{
		root: root,
		ct: &ctypes.Paragraph{
			Children: []ctypes.ParagraphChild{
				// w:r with fldChar begin
				{
					Run: &ctypes.Run{
						Children: []ctypes.RunChild{
							{
								FldChar: &ctypes.FldChar{
									FldCharType: stypes.FldCharTypeBegin,
								},
							},
						},
					},
				},
				// w:r with instrText
				{
					Run: &ctypes.Run{
						Children: []ctypes.RunChild{
							{
								InstrText: ctypes.TextFromString(` HYPERLINK "http://www.example.com" `),
							},
						},
					},
				},
				// w:r with fldChar separate
				{
					Run: &ctypes.Run{
						Children: []ctypes.RunChild{
							{
								FldChar: &ctypes.FldChar{
									FldCharType: stypes.FldCharTypeSeparate,
								},
							},
						},
					},
				},
				// First display run with bold text
				{
					Run: &ctypes.Run{
						Property: &ctypes.RunProperty{
							Bold: ctypes.OnOffFromBool(true),
						},
						Children: []ctypes.RunChild{
							{
								Text: ctypes.TextFromString("粗体"),
							},
						},
					},
				},
				// Second display run with italic text
				{
					Run: &ctypes.Run{
						Property: &ctypes.RunProperty{
							Italic: ctypes.OnOffFromBool(true),
						},
						Children: []ctypes.RunChild{
							{
								Text: ctypes.TextFromString("斜体"),
							},
						},
					},
				},
				// Third display run with underline text
				{
					Run: &ctypes.Run{
						Property: &ctypes.RunProperty{
							Underline: &ctypes.GenSingleStrVal[stypes.Underline]{
								Val: stypes.UnderlineSingle,
							},
						},
						Children: []ctypes.RunChild{
							{
								Text: ctypes.TextFromString("下划线"),
							},
						},
					},
				},
				// w:r with fldChar end
				{
					Run: &ctypes.Run{
						Children: []ctypes.RunChild{
							{
								FldChar: &ctypes.FldChar{
									FldCharType: stypes.FldCharTypeEnd,
								},
							},
						},
					},
				},
			},
		},
	}

	// Create field parser and process the paragraph
	parser := NewFieldParser(root)
	err := parser.ParseParagraphFields(para)
	if err != nil {
		t.Fatalf("Failed to parse paragraph fields: %v", err)
	}

	// Verify that the field was converted to a hyperlink
	if len(para.ct.Children) != 1 {
		t.Errorf("Expected 1 child after field conversion, got %d", len(para.ct.Children))
	}

	if para.ct.Children[0].Link == nil {
		t.Error("Expected a hyperlink after field conversion")
	}

	hyperlink := para.ct.Children[0].Link
	if hyperlink.ID == "" {
		t.Error("Expected hyperlink to have an ID")
	}

	// Check that hyperlink preserves all 3 styled runs
	if len(hyperlink.Children) != 3 {
		t.Errorf("Expected hyperlink to have 3 children runs, got %d", len(hyperlink.Children))
	}

	// Verify first run (bold)
	firstRun := hyperlink.Children[0].Run
	if firstRun == nil {
		t.Error("Expected first hyperlink child to be a run")
		return
	}
	if firstRun.Property == nil || firstRun.Property.Bold == nil {
		t.Error("Expected first run to have bold property")
		return
	}
	if firstRun.Property.Bold.Val == nil || *firstRun.Property.Bold.Val != stypes.OnOffTrue {
		t.Error("Expected first run to be bold")
	}
	if len(firstRun.Children) == 0 || firstRun.Children[0].Text == nil {
		t.Error("Expected first run to have text")
		return
	}
	if firstRun.Children[0].Text.Text != "粗体" {
		t.Errorf("Expected first run text to be '粗体', got '%s'", firstRun.Children[0].Text.Text)
	}

	// Verify second run (italic)
	secondRun := hyperlink.Children[1].Run
	if secondRun == nil {
		t.Error("Expected second hyperlink child to be a run")
		return
	}
	if secondRun.Property == nil || secondRun.Property.Italic == nil {
		t.Error("Expected second run to have italic property")
		return
	}
	if secondRun.Property.Italic.Val == nil || *secondRun.Property.Italic.Val != stypes.OnOffTrue {
		t.Error("Expected second run to be italic")
	}
	if len(secondRun.Children) == 0 || secondRun.Children[0].Text == nil {
		t.Error("Expected second run to have text")
		return
	}
	if secondRun.Children[0].Text.Text != "斜体" {
		t.Errorf("Expected second run text to be '斜体', got '%s'", secondRun.Children[0].Text.Text)
	}

	// Verify third run (underline)
	thirdRun := hyperlink.Children[2].Run
	if thirdRun == nil {
		t.Error("Expected third hyperlink child to be a run")
		return
	}
	if thirdRun.Property == nil || thirdRun.Property.Underline == nil {
		t.Error("Expected third run to have underline property")
		return
	}
	if thirdRun.Property.Underline.Val != stypes.UnderlineSingle {
		t.Error("Expected third run to have single underline")
	}
	if len(thirdRun.Children) == 0 || thirdRun.Children[0].Text == nil {
		t.Error("Expected third run to have text")
		return
	}
	if thirdRun.Children[0].Text.Text != "下划线" {
		t.Errorf("Expected third run text to be '下划线', got '%s'", thirdRun.Children[0].Text.Text)
	}

	// Verify that a relationship was created
	if len(root.Document.DocRels.Relationships) != 1 {
		t.Errorf("Expected 1 relationship after field conversion, got %d", len(root.Document.DocRels.Relationships))
	}
}
