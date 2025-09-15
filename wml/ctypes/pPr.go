package ctypes

import (
	"encoding/xml"
	"fmt"

	"github.com/gomutex/godocx/wml/stypes"
)

// Numbering Level Associated Paragraph Properties
type ParagraphProp struct {
	// 1. This element specifies the style ID of the paragraph style which shall be used to format the contents of this paragraph.
	Style *CTString `xml:"pStyle,omitempty"`

	// 2. Keep Paragraph With Next Paragraph
	KeepNext *OnOff `xml:"keepNext,omitempty"`

	// 3. Keep All Lines On One Page
	KeepLines *OnOff `xml:"keepLines,omitempty"`

	// 4. Start Paragraph on Next Page
	PageBreakBefore *OnOff `xml:"pageBreakBefore,omitempty"`

	// 5. Text Frame Properties
	FrameProp *FrameProp `xml:"framePr,omitempty"`

	// 6. Allow First/Last Line to Display on a Separate Page
	WindowControl *OnOff `xml:"widowControl,omitempty"`

	// 7. Numbering Definition Instance Reference
	NumProp *NumProp `xml:"numPr,omitempty"`

	// 8. Suppress Line Numbers for Paragraph
	SuppressLineNmbrs *OnOff `xml:"suppressLineNumbers,omitempty"`

	// 9. Paragraph Borders
	Border *ParaBorder `xml:"pBdr,omitempty"`

	// 10. This element specifies the shading applied to the contents of the paragraph.
	Shading *Shading `xml:"shd,omitempty"`

	// 11. Set of Custom Tab Stops
	Tabs Tabs `xml:"tabs,omitempty"`

	// 12. Suppress Hyphenation for Paragraph
	SuppressAutoHyphens *OnOff `xml:"suppressAutoHyphens,omitempty"`

	// 13. Use East Asian Typography Rules for First and Last Character per Line
	Kinsoku *OnOff `xml:"kinsoku,omitempty"`

	// 14. Allow Line Breaking At Character Level
	WordWrap *OnOff `xml:"wordWrap,omitempty"`

	// 15. Allow Punctuation to Extent Past Text Extents
	OverflowPunct *OnOff `xml:"overflowPunct,omitempty"`

	// 16. Compress Punctuation at Start of a Line
	TopLinePunct *OnOff `xml:"topLinePunct,omitempty"`

	// 17. Automatically Adjust Spacing of Latin and East Asian Text
	AutoSpaceDE *OnOff `xml:"autoSpaceDE,omitempty"`

	// 18. Automatically Adjust Spacing of East Asian Text and Numbers
	AutoSpaceDN *OnOff `xml:"autoSpaceDN,omitempty"`

	// 19. Right to Left Paragraph Layout
	Bidi *OnOff `xml:"bidi,omitempty"`

	// 20. Automatically Adjust Right Indent When Using Document Grid
	AdjustRightInd *OnOff `xml:"adjustRightInd,omitempty"`

	// 21. Use Document Grid Settings for Inter-Line Paragraph Spacing
	SnapToGrid *OnOff `xml:"snapToGrid,omitempty"`

	// 22. Spacing Between Lines and Above/Below Paragraph
	Spacing *Spacing `xml:"spacing,omitempty"`

	// 23. Paragraph Indentation
	Indent *Indent `xml:"ind,omitempty"`

	// 24. Ignore Spacing Above and Below When Using Identical Styles
	CtxlSpacing *OnOff `xml:"contextualSpacing,omitempty"`

	// 25. Use Left/Right Indents as Inside/Outside Indents
	MirrorIndents *OnOff `xml:"mirrorIndents,omitempty"`

	// 26. Prevent Text Frames From Overlapping
	SuppressOverlap *OnOff `xml:"suppressOverlap,omitempty"`

	// 27. Paragraph Alignment
	Justification *GenSingleStrVal[stypes.Justification] `xml:"jc,omitempty"`

	// 28. Paragraph Text Flow Direction
	TextDirection *GenSingleStrVal[stypes.TextDirection] `xml:"textDirection,omitempty"`

	// 29. Vertical Character Alignment on Line
	TextAlignment *GenSingleStrVal[stypes.TextAlign] `xml:"textAlignment,omitempty"`

	// 30.Allow Surrounding Paragraphs to Tight Wrap to Text Box Contents
	TextboxTightWrap *GenSingleStrVal[stypes.TextboxTightWrap] `xml:"textboxTightWrap,omitempty"`

	// 31. Associated Outline Level
	OutlineLvl *DecimalNum `xml:"outlineLvl,omitempty"`

	// 32. Associated HTML div ID
	DivID *DecimalNum `xml:"divId,omitempty"`

	// 33. Paragraph Conditional Formatting
	CnfStyle *CTString `xml:"cnfStyle,omitempty"`

	// 34. Run Properties for the Paragraph Mark
	RunProperty *RunProperty `xml:"rPr,omitempty"`

	// 35. Section Properties
	SectPr *SectionProp `xml:"sectPr,omitempty"`

	// 36. Revision Information for Paragraph Properties
	PPrChange *PPrChange `xml:"pPrChange,omitempty"`
}

type binElems struct {
	elem    *OnOff
	XMLName string
}

func (pp ParagraphProp) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	elem := xml.StartElement{Name: xml.Name{Local: "w:pPr"}}

	// Opening <w:pPr> element
	if err = e.EncodeToken(elem); err != nil {
		return err
	}

	// 1. PStyle
	if pp.Style != nil {
		if err = pp.Style.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:pStyle"},
		}); err != nil {
			return fmt.Errorf("style: %w", err)
		}
	}

	bElems1 := []binElems{
		{pp.KeepNext, "w:keepNext"},               //2
		{pp.KeepLines, "w:keepLines"},             //3
		{pp.PageBreakBefore, "w:pageBreakBefore"}, //4
	}

	for _, entry := range bElems1 {
		if entry.elem == nil {
			continue
		}
		if err = entry.elem.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: entry.XMLName},
		}); err != nil {
			return fmt.Errorf("error in marshaling paragraph property `%s`: %w", entry.XMLName, err)
		}
	}

	// 5. FrameProp
	if pp.FrameProp != nil {
		if err = pp.FrameProp.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("FrameProp: %w", err)
		}
	}

	// 6. WindowControl
	if pp.WindowControl != nil {
		if err = pp.WindowControl.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:widowControl"},
		}); err != nil {
			return fmt.Errorf("WindowControl: %w", err)
		}
	}

	// 7. NumProp
	if pp.NumProp != nil {
		if err = pp.NumProp.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("NumberingProperty: %w", err)
		}
	}

	// 8. SuppressLineNmbrs
	if pp.SuppressLineNmbrs != nil {
		if err = pp.SuppressLineNmbrs.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:suppressLineNumbers"},
		}); err != nil {
			return fmt.Errorf("SuppressLineNmbrs: %w", err)
		}
	}

	// 9.Border
	if pp.Border != nil {
		if err = pp.Border.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("Border: %w", err)
		}
	}

	// 10. Shading
	if pp.Shading != nil {
		if err = pp.Shading.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:shd"},
		}); err != nil {
			return fmt.Errorf("Shading: %w", err)
		}
	}

	// 11. Tabs
	if err = pp.Tabs.MarshalXML(e, xml.StartElement{}); err != nil {
		return fmt.Errorf("Tabs: %w", err)
	}

	bElems2 := []binElems{
		{pp.SuppressAutoHyphens, "w:suppressAutoHyphens"}, //12
		{pp.Kinsoku, "w:kinsoku"},                         //13
		{pp.WordWrap, "w:wordWrap"},                       //4
		{pp.OverflowPunct, "w:overflowPunct"},             //15
		{pp.TopLinePunct, "w:topLinePunct"},               //16
		{pp.AutoSpaceDE, "w:autoSpaceDE"},                 //17
		{pp.AutoSpaceDN, "w:autoSpaceDN"},                 //18
		{pp.Bidi, "w:bidi"},                               //19
		{pp.AdjustRightInd, "w:adjustRightInd"},           //20
		{pp.SnapToGrid, "w:snapToGrid"},                   //21
	}

	for _, entry := range bElems2 {
		if entry.elem == nil {
			continue
		}
		if err = entry.elem.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: entry.XMLName},
		}); err != nil {
			return fmt.Errorf("error in marshaling paragraph property `%s`: %w", entry.XMLName, err)
		}
	}

	// 22. Spacing
	if pp.Spacing != nil {
		if err = pp.Spacing.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("Spacing: %w", err)
		}
	}

	// 23. Indent
	if pp.Indent != nil {
		if err = pp.Indent.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("Indent: %w", err)
		}
	}

	bElems3 := []binElems{
		{pp.CtxlSpacing, "w:contextualSpacing"},   //24
		{pp.MirrorIndents, "w:mirrorIndents"},     //25
		{pp.SuppressOverlap, "w:suppressOverlap"}, //26
	}

	for _, entry := range bElems3 {
		if entry.elem == nil {
			continue
		}
		if err = entry.elem.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: entry.XMLName},
		}); err != nil {
			return fmt.Errorf("error in marshaling paragraph property `%s`: %w", entry.XMLName, err)
		}
	}

	// 27. Justification
	if pp.Justification != nil {
		if err = pp.Justification.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:jc"},
		}); err != nil {
			return fmt.Errorf("Justification: %w", err)
		}
	}

	// 28. TextDirection
	if pp.TextDirection != nil {
		if err = pp.TextDirection.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:textDirection"},
		}); err != nil {
			return fmt.Errorf("TextDirection: %w", err)
		}
	}

	// 29. TextAlignment
	if pp.TextAlignment != nil {
		if err = pp.TextAlignment.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:textAlignment"},
		}); err != nil {
			return fmt.Errorf("TextAlignment: %w", err)
		}
	}

	// 30. TextboxTightWrap
	if pp.TextboxTightWrap != nil {
		if err = pp.TextboxTightWrap.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:textboxTightWrap"},
		}); err != nil {
			return fmt.Errorf("TextboxTightWrap: %w", err)
		}
	}

	// 31. OutlineLvl
	if pp.OutlineLvl != nil {
		if err = pp.OutlineLvl.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:outlineLvl"},
		}); err != nil {
			return fmt.Errorf("OutlineLvl: %w", err)
		}
	}

	// 32. DivID
	if pp.DivID != nil {
		if err = pp.DivID.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:divId"},
		}); err != nil {
			return fmt.Errorf("DivID: %w", err)
		}
	}

	// 33. CnfStyle
	if pp.CnfStyle != nil {
		if err = pp.CnfStyle.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:cnfStyle"},
		}); err != nil {
			return fmt.Errorf("CnfStyle: %w", err)
		}
	}

	// 34. RunProperty
	if pp.RunProperty != nil {
		propsElement := xml.StartElement{Name: xml.Name{Local: "w:rPr"}}
		if err = e.EncodeElement(pp.RunProperty, propsElement); err != nil {
			return err
		}
	}

	if pp.SectPr != nil {
		if err = pp.SectPr.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:sectPr"},
		}); err != nil {
			return fmt.Errorf("PPrChange: %w", err)
		}
	}

	//36. PPrChange
	if pp.PPrChange != nil {
		if err = pp.PPrChange.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:pPrChange"},
		}); err != nil {
			return fmt.Errorf("PPrChange: %w", err)
		}
	}

	return e.EncodeToken(elem.End())
}

// NewParagraphStyle creates a new ParagraphStyle.
func NewParagraphStyle(val string) *CTString {
	return &CTString{Val: val}
}

// DefaultParagraphStyle creates the default ParagraphStyle with the value "Normal".
func DefaultParagraphStyle() *CTString {
	return &CTString{Val: "Normal"}
}

func DefaultParaProperty() *ParagraphProp {
	return &ParagraphProp{}
}

// Clone creates a deep copy of the ParagraphProp.
// Note: This method ignores any existing pPrChange to avoid creating history of history.
func (pp *ParagraphProp) Clone() *ParagraphProp {
	if pp == nil {
		return nil
	}

	clone := &ParagraphProp{}

	// Clone StyleID
	if pp.Style != nil {
		clone.Style = &CTString{Val: pp.Style.Val}
	}

	// Clone KeepNext
	if pp.KeepNext != nil {
		clone.KeepNext = &OnOff{Val: pp.KeepNext.Val}
	}

	// Clone KeepLines
	if pp.KeepLines != nil {
		clone.KeepLines = &OnOff{Val: pp.KeepLines.Val}
	}

	// Clone PageBreakBefore
	if pp.PageBreakBefore != nil {
		clone.PageBreakBefore = &OnOff{Val: pp.PageBreakBefore.Val}
	}

	// Clone FrameProp
	if pp.FrameProp != nil {
		cloneFrame := &FrameProp{}
		if pp.FrameProp.Width != nil {
			width := *pp.FrameProp.Width
			cloneFrame.Width = &width
		}
		if pp.FrameProp.Height != nil {
			height := *pp.FrameProp.Height
			cloneFrame.Height = &height
		}
		if pp.FrameProp.DropCap != nil {
			dropCap := *pp.FrameProp.DropCap
			cloneFrame.DropCap = &dropCap
		}
		if pp.FrameProp.Lines != nil {
			lines := *pp.FrameProp.Lines
			cloneFrame.Lines = &lines
		}
		if pp.FrameProp.VSpace != nil {
			vSpace := *pp.FrameProp.VSpace
			cloneFrame.VSpace = &vSpace
		}
		if pp.FrameProp.HSpace != nil {
			hSpace := *pp.FrameProp.HSpace
			cloneFrame.HSpace = &hSpace
		}
		if pp.FrameProp.Wrap != nil {
			wrap := *pp.FrameProp.Wrap
			cloneFrame.Wrap = &wrap
		}
		if pp.FrameProp.HAnchor != nil {
			hAnchor := *pp.FrameProp.HAnchor
			cloneFrame.HAnchor = &hAnchor
		}
		if pp.FrameProp.VAnchor != nil {
			vAnchor := *pp.FrameProp.VAnchor
			cloneFrame.VAnchor = &vAnchor
		}
		if pp.FrameProp.AbsHPos != nil {
			absHPos := *pp.FrameProp.AbsHPos
			cloneFrame.AbsHPos = &absHPos
		}
		if pp.FrameProp.AbsVPos != nil {
			absVPos := *pp.FrameProp.AbsVPos
			cloneFrame.AbsVPos = &absVPos
		}
		if pp.FrameProp.XAlign != nil {
			xAlign := *pp.FrameProp.XAlign
			cloneFrame.XAlign = &xAlign
		}
		if pp.FrameProp.YAlign != nil {
			yAlign := *pp.FrameProp.YAlign
			cloneFrame.YAlign = &yAlign
		}
		if pp.FrameProp.HRule != nil {
			hRule := *pp.FrameProp.HRule
			cloneFrame.HRule = &hRule
		}
		if pp.FrameProp.AnchorLock != nil {
			anchorLock := *pp.FrameProp.AnchorLock
			cloneFrame.AnchorLock = &anchorLock
		}
		clone.FrameProp = cloneFrame
	}

	// Clone WidowControl
	if pp.WindowControl != nil {
		clone.WindowControl = &OnOff{Val: pp.WindowControl.Val}
	}

	// Clone NumProp
	if pp.NumProp != nil {
		cloneNum := &NumProp{}
		if pp.NumProp.ILvl != nil {
			cloneNum.ILvl = &DecimalNum{Val: pp.NumProp.ILvl.Val}
		}
		if pp.NumProp.NumID != nil {
			cloneNum.NumID = &DecimalNum{Val: pp.NumProp.NumID.Val}
		}
		if pp.NumProp.NumChange != nil {
			cloneNumChange := &TrackChangeNum{
				ID:     pp.NumProp.NumChange.ID,
				Author: pp.NumProp.NumChange.Author,
			}
			if pp.NumProp.NumChange.Date != nil {
				date := *pp.NumProp.NumChange.Date
				cloneNumChange.Date = &date
			}
			cloneNum.NumChange = cloneNumChange
		}
		if pp.NumProp.Ins != nil {
			cloneIns := &TrackChange{
				ID:     pp.NumProp.Ins.ID,
				Author: pp.NumProp.Ins.Author,
			}
			if pp.NumProp.Ins.Date != nil {
				date := *pp.NumProp.Ins.Date
				cloneIns.Date = &date
			}
			cloneNum.Ins = cloneIns
		}
		clone.NumProp = cloneNum
	}

	// Clone SuppressLineNmbrs
	if pp.SuppressLineNmbrs != nil {
		clone.SuppressLineNmbrs = &OnOff{Val: pp.SuppressLineNmbrs.Val}
	}

	// Clone Border
	if pp.Border != nil {
		cloneBorder := &ParaBorder{}
		if pp.Border.Top != nil {
			cloneTop := &Border{Val: pp.Border.Top.Val}
			if pp.Border.Top.Color != nil {
				color := *pp.Border.Top.Color
				cloneTop.Color = &color
			}
			if pp.Border.Top.ThemeColor != nil {
				themeColor := *pp.Border.Top.ThemeColor
				cloneTop.ThemeColor = &themeColor
			}
			if pp.Border.Top.ThemeTint != nil {
				themeTint := *pp.Border.Top.ThemeTint
				cloneTop.ThemeTint = &themeTint
			}
			if pp.Border.Top.ThemeShade != nil {
				themeShade := *pp.Border.Top.ThemeShade
				cloneTop.ThemeShade = &themeShade
			}
			if pp.Border.Top.Size != nil {
				size := *pp.Border.Top.Size
				cloneTop.Size = &size
			}
			if pp.Border.Top.Space != nil {
				space := *pp.Border.Top.Space
				cloneTop.Space = &space
			}
			if pp.Border.Top.Shadow != nil {
				shadow := *pp.Border.Top.Shadow
				cloneTop.Shadow = &shadow
			}
			if pp.Border.Top.Frame != nil {
				frame := *pp.Border.Top.Frame
				cloneTop.Frame = &frame
			}
			cloneBorder.Top = cloneTop
		}
		if pp.Border.Left != nil {
			cloneLeft := &Border{Val: pp.Border.Left.Val}
			if pp.Border.Left.Color != nil {
				color := *pp.Border.Left.Color
				cloneLeft.Color = &color
			}
			if pp.Border.Left.ThemeColor != nil {
				themeColor := *pp.Border.Left.ThemeColor
				cloneLeft.ThemeColor = &themeColor
			}
			if pp.Border.Left.ThemeTint != nil {
				themeTint := *pp.Border.Left.ThemeTint
				cloneLeft.ThemeTint = &themeTint
			}
			if pp.Border.Left.ThemeShade != nil {
				themeShade := *pp.Border.Left.ThemeShade
				cloneLeft.ThemeShade = &themeShade
			}
			if pp.Border.Left.Size != nil {
				size := *pp.Border.Left.Size
				cloneLeft.Size = &size
			}
			if pp.Border.Left.Space != nil {
				space := *pp.Border.Left.Space
				cloneLeft.Space = &space
			}
			if pp.Border.Left.Shadow != nil {
				shadow := *pp.Border.Left.Shadow
				cloneLeft.Shadow = &shadow
			}
			if pp.Border.Left.Frame != nil {
				frame := *pp.Border.Left.Frame
				cloneLeft.Frame = &frame
			}
			cloneBorder.Left = cloneLeft
		}
		if pp.Border.Right != nil {
			cloneRight := &Border{Val: pp.Border.Right.Val}
			if pp.Border.Right.Color != nil {
				color := *pp.Border.Right.Color
				cloneRight.Color = &color
			}
			if pp.Border.Right.ThemeColor != nil {
				themeColor := *pp.Border.Right.ThemeColor
				cloneRight.ThemeColor = &themeColor
			}
			if pp.Border.Right.ThemeTint != nil {
				themeTint := *pp.Border.Right.ThemeTint
				cloneRight.ThemeTint = &themeTint
			}
			if pp.Border.Right.ThemeShade != nil {
				themeShade := *pp.Border.Right.ThemeShade
				cloneRight.ThemeShade = &themeShade
			}
			if pp.Border.Right.Size != nil {
				size := *pp.Border.Right.Size
				cloneRight.Size = &size
			}
			if pp.Border.Right.Space != nil {
				space := *pp.Border.Right.Space
				cloneRight.Space = &space
			}
			if pp.Border.Right.Shadow != nil {
				shadow := *pp.Border.Right.Shadow
				cloneRight.Shadow = &shadow
			}
			if pp.Border.Right.Frame != nil {
				frame := *pp.Border.Right.Frame
				cloneRight.Frame = &frame
			}
			cloneBorder.Right = cloneRight
		}
		if pp.Border.Bottom != nil {
			cloneBottom := &Border{Val: pp.Border.Bottom.Val}
			if pp.Border.Bottom.Color != nil {
				color := *pp.Border.Bottom.Color
				cloneBottom.Color = &color
			}
			if pp.Border.Bottom.ThemeColor != nil {
				themeColor := *pp.Border.Bottom.ThemeColor
				cloneBottom.ThemeColor = &themeColor
			}
			if pp.Border.Bottom.ThemeTint != nil {
				themeTint := *pp.Border.Bottom.ThemeTint
				cloneBottom.ThemeTint = &themeTint
			}
			if pp.Border.Bottom.ThemeShade != nil {
				themeShade := *pp.Border.Bottom.ThemeShade
				cloneBottom.ThemeShade = &themeShade
			}
			if pp.Border.Bottom.Size != nil {
				size := *pp.Border.Bottom.Size
				cloneBottom.Size = &size
			}
			if pp.Border.Bottom.Space != nil {
				space := *pp.Border.Bottom.Space
				cloneBottom.Space = &space
			}
			if pp.Border.Bottom.Shadow != nil {
				shadow := *pp.Border.Bottom.Shadow
				cloneBottom.Shadow = &shadow
			}
			if pp.Border.Bottom.Frame != nil {
				frame := *pp.Border.Bottom.Frame
				cloneBottom.Frame = &frame
			}
			cloneBorder.Bottom = cloneBottom
		}
		if pp.Border.Between != nil {
			cloneBetween := &Border{Val: pp.Border.Between.Val}
			if pp.Border.Between.Color != nil {
				color := *pp.Border.Between.Color
				cloneBetween.Color = &color
			}
			if pp.Border.Between.ThemeColor != nil {
				themeColor := *pp.Border.Between.ThemeColor
				cloneBetween.ThemeColor = &themeColor
			}
			if pp.Border.Between.ThemeTint != nil {
				themeTint := *pp.Border.Between.ThemeTint
				cloneBetween.ThemeTint = &themeTint
			}
			if pp.Border.Between.ThemeShade != nil {
				themeShade := *pp.Border.Between.ThemeShade
				cloneBetween.ThemeShade = &themeShade
			}
			if pp.Border.Between.Size != nil {
				size := *pp.Border.Between.Size
				cloneBetween.Size = &size
			}
			if pp.Border.Between.Space != nil {
				space := *pp.Border.Between.Space
				cloneBetween.Space = &space
			}
			if pp.Border.Between.Shadow != nil {
				shadow := *pp.Border.Between.Shadow
				cloneBetween.Shadow = &shadow
			}
			if pp.Border.Between.Frame != nil {
				frame := *pp.Border.Between.Frame
				cloneBetween.Frame = &frame
			}
			cloneBorder.Between = cloneBetween
		}
		if pp.Border.Bar != nil {
			cloneBar := &Border{Val: pp.Border.Bar.Val}
			if pp.Border.Bar.Color != nil {
				color := *pp.Border.Bar.Color
				cloneBar.Color = &color
			}
			if pp.Border.Bar.ThemeColor != nil {
				themeColor := *pp.Border.Bar.ThemeColor
				cloneBar.ThemeColor = &themeColor
			}
			if pp.Border.Bar.ThemeTint != nil {
				themeTint := *pp.Border.Bar.ThemeTint
				cloneBar.ThemeTint = &themeTint
			}
			if pp.Border.Bar.ThemeShade != nil {
				themeShade := *pp.Border.Bar.ThemeShade
				cloneBar.ThemeShade = &themeShade
			}
			if pp.Border.Bar.Size != nil {
				size := *pp.Border.Bar.Size
				cloneBar.Size = &size
			}
			if pp.Border.Bar.Space != nil {
				space := *pp.Border.Bar.Space
				cloneBar.Space = &space
			}
			if pp.Border.Bar.Shadow != nil {
				shadow := *pp.Border.Bar.Shadow
				cloneBar.Shadow = &shadow
			}
			if pp.Border.Bar.Frame != nil {
				frame := *pp.Border.Bar.Frame
				cloneBar.Frame = &frame
			}
			cloneBorder.Bar = cloneBar
		}
		clone.Border = cloneBorder
	}

	// Clone Shading
	if pp.Shading != nil {
		cloneShading := &Shading{
			Val: pp.Shading.Val,
		}
		if pp.Shading.Color != nil {
			color := *pp.Shading.Color
			cloneShading.Color = &color
		}
		if pp.Shading.ThemeColor != nil {
			themeColor := *pp.Shading.ThemeColor
			cloneShading.ThemeColor = &themeColor
		}
		if pp.Shading.ThemeTint != nil {
			themeTint := *pp.Shading.ThemeTint
			cloneShading.ThemeTint = &themeTint
		}
		if pp.Shading.ThemeShade != nil {
			themeShade := *pp.Shading.ThemeShade
			cloneShading.ThemeShade = &themeShade
		}
		if pp.Shading.Fill != nil {
			fill := *pp.Shading.Fill
			cloneShading.Fill = &fill
		}
		if pp.Shading.ThemeFill != nil {
			themeFill := *pp.Shading.ThemeFill
			cloneShading.ThemeFill = &themeFill
		}
		if pp.Shading.ThemeFillTint != nil {
			themeFillTint := *pp.Shading.ThemeFillTint
			cloneShading.ThemeFillTint = &themeFillTint
		}
		if pp.Shading.ThemeFillShade != nil {
			themeFillShade := *pp.Shading.ThemeFillShade
			cloneShading.ThemeFillShade = &themeFillShade
		}
		clone.Shading = cloneShading
	}

	// Clone Tabs
	clone.Tabs = Tabs{}
	if len(pp.Tabs.Tab) > 0 {
		clone.Tabs.Tab = make([]Tab, len(pp.Tabs.Tab))
		for i, tab := range pp.Tabs.Tab {
			cloneTab := Tab{
				Val:      tab.Val,
				Position: tab.Position,
			}
			if tab.LeaderChar != nil {
				leaderChar := *tab.LeaderChar
				cloneTab.LeaderChar = &leaderChar
			}
			clone.Tabs.Tab[i] = cloneTab
		}
	}

	// Clone SuppressAutoHyphens
	if pp.SuppressAutoHyphens != nil {
		clone.SuppressAutoHyphens = &OnOff{Val: pp.SuppressAutoHyphens.Val}
	}

	// Clone Kinsoku
	if pp.Kinsoku != nil {
		clone.Kinsoku = &OnOff{Val: pp.Kinsoku.Val}
	}

	// Clone WordWrap
	if pp.WordWrap != nil {
		clone.WordWrap = &OnOff{Val: pp.WordWrap.Val}
	}

	// Clone OverflowPunct
	if pp.OverflowPunct != nil {
		clone.OverflowPunct = &OnOff{Val: pp.OverflowPunct.Val}
	}

	// Clone TopLinePunct
	if pp.TopLinePunct != nil {
		clone.TopLinePunct = &OnOff{Val: pp.TopLinePunct.Val}
	}

	// Clone AutoSpaceDE
	if pp.AutoSpaceDE != nil {
		clone.AutoSpaceDE = &OnOff{Val: pp.AutoSpaceDE.Val}
	}

	// Clone AutoSpaceDN
	if pp.AutoSpaceDN != nil {
		clone.AutoSpaceDN = &OnOff{Val: pp.AutoSpaceDN.Val}
	}

	// Clone BiDi
	if pp.Bidi != nil {
		clone.Bidi = &OnOff{Val: pp.Bidi.Val}
	}

	// Clone AdjustRightInd
	if pp.AdjustRightInd != nil {
		clone.AdjustRightInd = &OnOff{Val: pp.AdjustRightInd.Val}
	}

	// Clone SnapToGrid
	if pp.SnapToGrid != nil {
		clone.SnapToGrid = &OnOff{Val: pp.SnapToGrid.Val}
	}

	// Clone Spacing
	if pp.Spacing != nil {
		cloneSpacing := &Spacing{}
		if pp.Spacing.Before != nil {
			before := *pp.Spacing.Before
			cloneSpacing.Before = &before
		}
		if pp.Spacing.BeforeLines != nil {
			beforeLines := *pp.Spacing.BeforeLines
			cloneSpacing.BeforeLines = &beforeLines
		}
		if pp.Spacing.BeforeAutospacing != nil {
			beforeAutospacing := *pp.Spacing.BeforeAutospacing
			cloneSpacing.BeforeAutospacing = &beforeAutospacing
		}
		if pp.Spacing.After != nil {
			after := *pp.Spacing.After
			cloneSpacing.After = &after
		}
		if pp.Spacing.AfterAutospacing != nil {
			afterAutospacing := *pp.Spacing.AfterAutospacing
			cloneSpacing.AfterAutospacing = &afterAutospacing
		}
		if pp.Spacing.Line != nil {
			line := *pp.Spacing.Line
			cloneSpacing.Line = &line
		}
		if pp.Spacing.LineRule != nil {
			lineRule := *pp.Spacing.LineRule
			cloneSpacing.LineRule = &lineRule
		}
		clone.Spacing = cloneSpacing
	}

	// Clone Indent
	if pp.Indent != nil {
		cloneIndent := &Indent{}
		if pp.Indent.Left != nil {
			left := *pp.Indent.Left
			cloneIndent.Left = &left
		}
		if pp.Indent.LeftChars != nil {
			leftChars := *pp.Indent.LeftChars
			cloneIndent.LeftChars = &leftChars
		}
		if pp.Indent.Right != nil {
			right := *pp.Indent.Right
			cloneIndent.Right = &right
		}
		if pp.Indent.RightChars != nil {
			rightChars := *pp.Indent.RightChars
			cloneIndent.RightChars = &rightChars
		}
		if pp.Indent.Hanging != nil {
			hanging := *pp.Indent.Hanging
			cloneIndent.Hanging = &hanging
		}
		if pp.Indent.HangingChars != nil {
			hangingChars := *pp.Indent.HangingChars
			cloneIndent.HangingChars = &hangingChars
		}
		if pp.Indent.FirstLine != nil {
			firstLine := *pp.Indent.FirstLine
			cloneIndent.FirstLine = &firstLine
		}
		if pp.Indent.FirstLineChars != nil {
			firstLineChars := *pp.Indent.FirstLineChars
			cloneIndent.FirstLineChars = &firstLineChars
		}
		clone.Indent = cloneIndent
	}

	// Clone ContextualSpacing
	if pp.CtxlSpacing != nil {
		clone.CtxlSpacing = &OnOff{Val: pp.CtxlSpacing.Val}
	}

	// Clone MirrorIndents
	if pp.MirrorIndents != nil {
		clone.MirrorIndents = &OnOff{Val: pp.MirrorIndents.Val}
	}

	// Clone SuppressOverlap
	if pp.SuppressOverlap != nil {
		clone.SuppressOverlap = &OnOff{Val: pp.SuppressOverlap.Val}
	}

	// Clone JC
	if pp.Justification != nil {
		clone.Justification = &GenSingleStrVal[stypes.Justification]{Val: pp.Justification.Val}
	}

	// Clone TextDirection
	if pp.TextDirection != nil {
		clone.TextDirection = &GenSingleStrVal[stypes.TextDirection]{Val: pp.TextDirection.Val}
	}

	// Clone TextAlignment
	if pp.TextAlignment != nil {
		clone.TextAlignment = &GenSingleStrVal[stypes.TextAlign]{Val: pp.TextAlignment.Val}
	}

	// Clone TextboxTightWrap
	if pp.TextboxTightWrap != nil {
		clone.TextboxTightWrap = &GenSingleStrVal[stypes.TextboxTightWrap]{Val: pp.TextboxTightWrap.Val}
	}

	// Clone OutlineLvl
	if pp.OutlineLvl != nil {
		clone.OutlineLvl = &DecimalNum{Val: pp.OutlineLvl.Val}
	}

	// Clone DivID
	if pp.DivID != nil {
		clone.DivID = &DecimalNum{Val: pp.DivID.Val}
	}

	// Clone CnfStyle
	if pp.CnfStyle != nil {
		clone.CnfStyle = &CTString{Val: pp.CnfStyle.Val}
	}

	// Clone RunProperty
	if pp.RunProperty != nil {
		clone.RunProperty = pp.RunProperty.Clone()
	}

	// Clone SectionProp
	if pp.SectPr != nil {
		cloneSectionProp := &SectionProp{}
		if pp.SectPr.HeaderReference != nil {
			cloneHeaderRef := &HeaderReference{}
			*cloneHeaderRef = *pp.SectPr.HeaderReference
			cloneSectionProp.HeaderReference = cloneHeaderRef
		}
		if pp.SectPr.FooterReference != nil {
			cloneFooterRef := &FooterReference{}
			*cloneFooterRef = *pp.SectPr.FooterReference
			cloneSectionProp.FooterReference = cloneFooterRef
		}
		if pp.SectPr.PageSize != nil {
			clonePageSize := &PageSize{}
			*clonePageSize = *pp.SectPr.PageSize
			cloneSectionProp.PageSize = clonePageSize
		}
		if pp.SectPr.Type != nil {
			cloneSectionProp.Type = &GenSingleStrVal[stypes.SectionMark]{Val: pp.SectPr.Type.Val}
		}
		if pp.SectPr.PageMargin != nil {
			clonePageMargin := &PageMargin{}
			*clonePageMargin = *pp.SectPr.PageMargin
			cloneSectionProp.PageMargin = clonePageMargin
		}
		if pp.SectPr.PageNum != nil {
			clonePageNum := &PageNumbering{}
			*clonePageNum = *pp.SectPr.PageNum
			cloneSectionProp.PageNum = clonePageNum
		}
		if pp.SectPr.FormProt != nil {
			cloneSectionProp.FormProt = &GenSingleStrVal[stypes.OnOff]{Val: pp.SectPr.FormProt.Val}
		}
		if pp.SectPr.TitlePg != nil {
			cloneSectionProp.TitlePg = &GenSingleStrVal[stypes.OnOff]{Val: pp.SectPr.TitlePg.Val}
		}
		if pp.SectPr.TextDir != nil {
			cloneSectionProp.TextDir = &GenSingleStrVal[stypes.TextDirection]{Val: pp.SectPr.TextDir.Val}
		}
		if pp.SectPr.DocGrid != nil {
			cloneDocGrid := &DocGrid{}
			*cloneDocGrid = *pp.SectPr.DocGrid
			cloneSectionProp.DocGrid = cloneDocGrid
		}
		clone.SectPr = cloneSectionProp
	}

	// Note: We intentionally skip PPrChange to avoid creating history of history

	return clone
}
