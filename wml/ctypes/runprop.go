package ctypes

import (
	"encoding/xml"
	"fmt"

	"github.com/gomutex/godocx/wml/stypes"
)

// RunProperty represents the properties of a run of text within a paragraph.
type RunProperty struct {
	//1. Referenced Character Style
	Style *CTString `xml:"rStyle,omitempty"`

	//2. Run Fonts
	Fonts *RunFonts `xml:"rFonts,omitempty"`

	//3. Bold
	Bold *OnOff `xml:"b,omitempty"`

	//4.Complex Script Bold
	BoldCS *OnOff `xml:"bCs,omitempty"`

	// 5.Italics
	Italic *OnOff `xml:"i,omitempty"`

	//6.Complex Script Italics
	ItalicCS *OnOff `xml:"iCs,omitempty"`

	//7.Display All Characters As Capital Letters
	Caps *OnOff `xml:"caps,omitempty"`

	//8.Small Caps
	SmallCaps *OnOff `xml:"smallCaps,omitempty"`

	//9.Single Strikethrough
	Strike *OnOff `xml:"strike,omitempty"`

	//10.Double Strikethrough
	DoubleStrike *OnOff `xml:"dstrike,omitempty"`

	//11.Display Character Outline
	Outline *OnOff `xml:"outline,omitempty"`

	//12.Shadow
	Shadow *OnOff `xml:"shadow,omitempty"`

	//13.Embossing
	Emboss *OnOff `xml:"emboss,omitempty"`

	//14.Imprinting
	Imprint *OnOff `xml:"imprint,omitempty"`

	//15.Do Not Check Spelling or Grammar
	NoGrammar *OnOff `xml:"noProof,omitempty"`

	//16.Use Document Grid Settings For Inter-Character Spacing
	SnapToGrid *OnOff `xml:"snapToGrid,omitempty"`

	//17.Hidden Text
	Vanish *OnOff `xml:"vanish,omitempty"`

	//18.Web Hidden Text
	WebHidden *OnOff `xml:"webHidden,omitempty"`

	//19.Run Content Color
	Color *Color `xml:"color,omitempty"`

	//20. Character Spacing Adjustment
	Spacing *DecimalNum `xml:"spacing,omitempty"`

	//21.Expanded/Compressed Text
	ExpaComp *ExpaComp `xml:"w,omitempty"`

	//22.Font Kerning
	Kern *Uint64Elem `xml:"kern,omitempty"`

	//23. Vertically Raised or Lowered Text
	Position *DecimalNum `xml:"position,omitempty"`

	//24.Font Size
	Size *FontSize `xml:"sz,omitempty"`

	//25.Complex Script Font Size
	SizeCs *FontSizeCS `xml:"szCs,omitempty"`

	//26.Text Highlighting
	Highlight *CTString `xml:"highlight,omitempty"`

	//27.Underline
	Underline *GenSingleStrVal[stypes.Underline] `xml:"u,omitempty"`

	//28.Animated Text Effect
	Effect *Effect `xml:"effect,omitempty"`

	//29.Text Border
	Border *Border `xml:"bdr,omitempty"`

	//30.Run Shading
	Shading *Shading `xml:"shd,omitempty"`

	//31.Manual Run Width
	FitText *FitText `xml:"fitText,omitempty"`

	//32.Subscript/Superscript Text
	VertAlign *GenSingleStrVal[stypes.VerticalAlignRun] `xml:"vertAlign,omitempty"`

	//33.Right To Left Text
	RightToLeft *OnOff `xml:"rtl,omitempty"`

	//34.Use Complex Script Formatting on Run
	CSFormat *OnOff `xml:"cs,omitempty"`

	//35.Emphasis Mark
	Em *GenSingleStrVal[stypes.Em] `xml:"em,omitempty"`

	//36.Languages for Run Content
	Lang *Lang `xml:"lang,omitempty"`

	//37.East Asian Typography Settings
	EALayout *EALayout `xml:"eastAsianLayout,omitempty"`

	//38.Paragraph Mark Is Always Hidden
	SpecVanish *OnOff `xml:"specVanish,omitempty"`

	//39.Office Open XML Math
	OMath *OnOff `xml:"oMath,omitempty"`

	// 17.13.5.31 rPrChange (Revision Information for Run Properties)
	RPrChange *RPrChange `xml:"rPrChange,omitempty"`
}

// NewRunProperty creates a new RunProperty with default values.
func NewRunProperty() RunProperty {
	return RunProperty{}
}

// MarshalXML marshals RunProperty to XML.
func (rp RunProperty) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:rPr"
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	// 1. Referenced Character Style
	if rp.Style != nil {
		if err = rp.Style.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:rStyle"},
		}); err != nil {
			return fmt.Errorf("style: %w", err)
		}
	}

	//2.Run Fonts
	if rp.Fonts != nil {
		if err = rp.Fonts.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("Fonts: %w", err)
		}
	}

	set1 := []struct {
		elem    *OnOff
		XMLName string
	}{
		{rp.Bold, "w:b"},                //3.Bold
		{rp.BoldCS, "w:bCs"},            //4.Complex Script Bold
		{rp.Italic, "w:i"},              //5.Italics
		{rp.ItalicCS, "w:iCs"},          //6.Complex Script Italics
		{rp.Caps, "w:caps"},             //7.Display All Characters As Capital Letters
		{rp.SmallCaps, "w:smallCaps"},   //8.Small Caps
		{rp.Strike, "w:strike"},         //9.Single Strikethrough
		{rp.DoubleStrike, "w:dstrike"},  //10.Double Strikethrough
		{rp.Outline, "w:outline"},       //11.Display Character Outline
		{rp.Shadow, "w:shadow"},         //12.Shadow
		{rp.Emboss, "w:emboss"},         //13.Embossing
		{rp.Imprint, "w:imprint"},       //14.Imprinting
		{rp.NoGrammar, "w:noProof"},     //15.Do Not Check Spelling or Grammar
		{rp.SnapToGrid, "w:snapToGrid"}, //16.Use Document Grid Settings For Inter-Character Spacing
		{rp.Vanish, "w:vanish"},         //17.Hidden Text
		{rp.WebHidden, "w:webHidden"},   //18.Web Hidden Text
	}

	for _, entry := range set1 {
		if entry.elem == nil {
			continue
		}
		if err = entry.elem.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: entry.XMLName},
		}); err != nil {
			return fmt.Errorf("error in marshaling run property `%s`: %w", entry.XMLName, err)
		}
	}

	//19.Run Content Color
	if rp.Color != nil {
		if err = rp.Color.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("color: %w", err)
		}
	}

	//20. Character Spacing Adjustment
	if rp.Spacing != nil {
		if err = rp.Spacing.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:spacing"},
		}); err != nil {
			return fmt.Errorf("spacing: %w", err)
		}
	}

	//21.Expanded/Compressed Text
	if rp.ExpaComp != nil {
		if err = rp.ExpaComp.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:w"},
		}); err != nil {
			return fmt.Errorf("expand/compression text: %w", err)
		}
	}

	//22.Font Kerning
	if rp.Kern != nil {
		if err = rp.Kern.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:kern"},
		}); err != nil {
			return fmt.Errorf("kern: %w", err)
		}
	}

	//23. Vertically Raised or Lowered Text
	if rp.Position != nil {
		if err = rp.Position.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:position"},
		}); err != nil {
			return fmt.Errorf("position: %w", err)
		}
	}

	//24.Font Size
	if rp.Size != nil {
		if err = rp.Size.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("size: %w", err)
		}
	}

	//25.Complex Script Font Size
	if rp.SizeCs != nil {
		if err = rp.SizeCs.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("size complex script: %w", err)
		}
	}

	//26.Text Highlighting
	if rp.Highlight != nil {
		if err = rp.Highlight.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:highlight"},
		}); err != nil {
			return fmt.Errorf("highlight: %w", err)
		}
	}

	//27.Underline
	if rp.Underline != nil {
		if err = rp.Underline.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:u"},
		}); err != nil {
			return fmt.Errorf("underline: %w", err)
		}
	}

	//28.Animated Text Effect
	if rp.Effect != nil {
		if err = rp.Effect.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:effect"},
		}); err != nil {
			return fmt.Errorf("effect: %w", err)
		}
	}

	//29.Text Border
	if rp.Border != nil {
		if err = rp.Border.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:bdr"},
		}); err != nil {
			return fmt.Errorf("border: %w", err)
		}
	}

	//30.Run Shading
	if rp.Shading != nil {
		if err = rp.Shading.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("shading: %w", err)
		}
	}

	//31.Manual Run Width
	if rp.FitText != nil {
		if err = rp.FitText.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("fit text: %w", err)
		}
	}

	//32.Subscript/Superscript Text
	if rp.VertAlign != nil {
		if err = rp.VertAlign.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:vertAlign"},
		}); err != nil {
			return fmt.Errorf("vertical align: %w", err)
		}
	}

	//33.Right To Left Text
	if rp.RightToLeft != nil {
		if err = rp.RightToLeft.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:rtl"},
		}); err != nil {
			return fmt.Errorf("error in marshaling run property `%s`: %w", "RightToLeft", err)
		}
	}

	//34.Use Complex Script Formatting on Run
	if rp.CSFormat != nil {
		if err = rp.CSFormat.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:cs"},
		}); err != nil {
			return fmt.Errorf("error in marshaling run property `%s`: %w", "CSFormat", err)
		}
	}

	//35.Emphasis Mark
	if rp.Em != nil {
		if err = rp.Em.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:em"},
		}); err != nil {
			return fmt.Errorf("emphasis mark: %w", err)
		}
	}

	//36.Languages for Run Content
	if rp.Lang != nil {
		if err = rp.Lang.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("languages for Run Content: %w", err)
		}
	}

	//37.East Asian Typography Settings
	if rp.EALayout != nil {
		if err = rp.EALayout.MarshalXML(e, xml.StartElement{}); err != nil {
			return fmt.Errorf("East Asian Typography Settings: %w", err)
		}
	}

	//38.Paragraph Mark Is Always Hidden
	if rp.SpecVanish != nil {
		if err = rp.SpecVanish.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:specVanish"},
		}); err != nil {
			return fmt.Errorf("error in marshaling run property `%s`: %w", "specVanish", err)
		}
	}

	//39.Office Open XML Math
	if rp.OMath != nil {
		if err = rp.OMath.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:oMath"},
		}); err != nil {
			return fmt.Errorf("error in marshaling run property `%s`: %w", "oMath", err)
		}
	}

	//rPrChange
	if rp.RPrChange != nil {
		if err = rp.RPrChange.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:rPrChange"},
		}); err != nil {
			return fmt.Errorf("rPrChange: %w", err)
		}
	}

	return e.EncodeToken(start.End())
}

// Clone creates a deep copy of the RunProperty.
// Note: This method ignores any existing rPrChange to avoid creating history of history.
func (rp RunProperty) Clone() *RunProperty {
	clone := &RunProperty{}

	// Clone simple pointer fields
	if rp.Style != nil {
		clone.Style = &CTString{Val: rp.Style.Val}
	}

	// Clone RunFonts
	if rp.Fonts != nil {
		clone.Fonts = &RunFonts{
			Hint:          rp.Fonts.Hint,
			Ascii:         rp.Fonts.Ascii,
			HAnsi:         rp.Fonts.HAnsi,
			EastAsia:      rp.Fonts.EastAsia,
			CS:            rp.Fonts.CS,
			AsciiTheme:    rp.Fonts.AsciiTheme,
			HAnsiTheme:    rp.Fonts.HAnsiTheme,
			EastAsiaTheme: rp.Fonts.EastAsiaTheme,
			CSTheme:       rp.Fonts.CSTheme,
		}
	}

	// Clone OnOff boolean fields
	if rp.Bold != nil {
		clone.Bold = &OnOff{Val: rp.Bold.Val}
	}
	if rp.BoldCS != nil {
		clone.BoldCS = &OnOff{Val: rp.BoldCS.Val}
	}
	if rp.Italic != nil {
		clone.Italic = &OnOff{Val: rp.Italic.Val}
	}
	if rp.ItalicCS != nil {
		clone.ItalicCS = &OnOff{Val: rp.ItalicCS.Val}
	}
	if rp.Caps != nil {
		clone.Caps = &OnOff{Val: rp.Caps.Val}
	}
	if rp.SmallCaps != nil {
		clone.SmallCaps = &OnOff{Val: rp.SmallCaps.Val}
	}
	if rp.Strike != nil {
		clone.Strike = &OnOff{Val: rp.Strike.Val}
	}
	if rp.DoubleStrike != nil {
		clone.DoubleStrike = &OnOff{Val: rp.DoubleStrike.Val}
	}
	if rp.Outline != nil {
		clone.Outline = &OnOff{Val: rp.Outline.Val}
	}
	if rp.Shadow != nil {
		clone.Shadow = &OnOff{Val: rp.Shadow.Val}
	}
	if rp.Emboss != nil {
		clone.Emboss = &OnOff{Val: rp.Emboss.Val}
	}
	if rp.Imprint != nil {
		clone.Imprint = &OnOff{Val: rp.Imprint.Val}
	}
	if rp.NoGrammar != nil {
		clone.NoGrammar = &OnOff{Val: rp.NoGrammar.Val}
	}
	if rp.SnapToGrid != nil {
		clone.SnapToGrid = &OnOff{Val: rp.SnapToGrid.Val}
	}
	if rp.Vanish != nil {
		clone.Vanish = &OnOff{Val: rp.Vanish.Val}
	}
	if rp.WebHidden != nil {
		clone.WebHidden = &OnOff{Val: rp.WebHidden.Val}
	}
	if rp.RightToLeft != nil {
		clone.RightToLeft = &OnOff{Val: rp.RightToLeft.Val}
	}
	if rp.CSFormat != nil {
		clone.CSFormat = &OnOff{Val: rp.CSFormat.Val}
	}
	if rp.SpecVanish != nil {
		clone.SpecVanish = &OnOff{Val: rp.SpecVanish.Val}
	}
	if rp.OMath != nil {
		clone.OMath = &OnOff{Val: rp.OMath.Val}
	}

	// Clone Color
	if rp.Color != nil {
		cloneColor := &Color{
			Val:        rp.Color.Val,
			ThemeColor: rp.Color.ThemeColor,
			ThemeTint:  rp.Color.ThemeTint,
			ThemeShade: rp.Color.ThemeShade,
		}
		clone.Color = cloneColor
	}

	// Clone numeric fields
	if rp.Spacing != nil {
		clone.Spacing = &DecimalNum{Val: rp.Spacing.Val}
	}
	if rp.Kern != nil {
		clone.Kern = &Uint64Elem{Val: rp.Kern.Val}
	}
	if rp.Position != nil {
		clone.Position = &DecimalNum{Val: rp.Position.Val}
	}
	if rp.Size != nil {
		clone.Size = &FontSize{Value: rp.Size.Value}
	}
	if rp.SizeCs != nil {
		clone.SizeCs = &FontSizeCS{Value: rp.SizeCs.Value}
	}

	// Clone string fields
	if rp.Highlight != nil {
		clone.Highlight = &CTString{Val: rp.Highlight.Val}
	}

	// Clone generic single string value fields
	if rp.Underline != nil {
		clone.Underline = &GenSingleStrVal[stypes.Underline]{Val: rp.Underline.Val}
	}
	if rp.VertAlign != nil {
		clone.VertAlign = &GenSingleStrVal[stypes.VerticalAlignRun]{Val: rp.VertAlign.Val}
	}
	if rp.Em != nil {
		clone.Em = &GenSingleStrVal[stypes.Em]{Val: rp.Em.Val}
	}

	// Clone complex fields (if they have Clone methods, use them; otherwise implement basic cloning)
	if rp.ExpaComp != nil {
		clone.ExpaComp = &ExpaComp{Val: rp.ExpaComp.Val}
	}
	if rp.Effect != nil {
		clone.Effect = &Effect{Val: rp.Effect.Val}
	}

	// Clone Border (basic implementation - may need to be enhanced if Border has more complex fields)
	if rp.Border != nil {
		cloneBorder := &Border{
			Val: rp.Border.Val,
		}
		if rp.Border.Color != nil {
			borderColor := *rp.Border.Color
			cloneBorder.Color = &borderColor
		}
		if rp.Border.ThemeColor != nil {
			borderThemeColor := *rp.Border.ThemeColor
			cloneBorder.ThemeColor = &borderThemeColor
		}
		if rp.Border.ThemeTint != nil {
			borderThemeTint := *rp.Border.ThemeTint
			cloneBorder.ThemeTint = &borderThemeTint
		}
		if rp.Border.ThemeShade != nil {
			borderThemeShade := *rp.Border.ThemeShade
			cloneBorder.ThemeShade = &borderThemeShade
		}
		if rp.Border.Size != nil {
			borderSize := *rp.Border.Size
			cloneBorder.Size = &borderSize
		}
		if rp.Border.Space != nil {
			borderSpace := *rp.Border.Space
			cloneBorder.Space = &borderSpace
		}
		if rp.Border.Shadow != nil {
			borderShadow := *rp.Border.Shadow
			cloneBorder.Shadow = &borderShadow
		}
		if rp.Border.Frame != nil {
			borderFrame := *rp.Border.Frame
			cloneBorder.Frame = &borderFrame
		}
		clone.Border = cloneBorder
	}

	// Clone Shading
	if rp.Shading != nil {
		cloneShading := &Shading{
			Val: rp.Shading.Val,
		}
		if rp.Shading.Color != nil {
			shadingColor := *rp.Shading.Color
			cloneShading.Color = &shadingColor
		}
		if rp.Shading.ThemeColor != nil {
			shadingThemeColor := *rp.Shading.ThemeColor
			cloneShading.ThemeColor = &shadingThemeColor
		}
		if rp.Shading.ThemeFill != nil {
			shadingThemeFill := *rp.Shading.ThemeFill
			cloneShading.ThemeFill = &shadingThemeFill
		}
		if rp.Shading.ThemeTint != nil {
			shadingThemeTint := *rp.Shading.ThemeTint
			cloneShading.ThemeTint = &shadingThemeTint
		}
		if rp.Shading.ThemeShade != nil {
			shadingThemeShade := *rp.Shading.ThemeShade
			cloneShading.ThemeShade = &shadingThemeShade
		}
		if rp.Shading.Fill != nil {
			shadingFill := *rp.Shading.Fill
			cloneShading.Fill = &shadingFill
		}
		if rp.Shading.ThemeFillTint != nil {
			shadingThemeFillTint := *rp.Shading.ThemeFillTint
			cloneShading.ThemeFillTint = &shadingThemeFillTint
		}
		if rp.Shading.ThemeFillShade != nil {
			shadingThemeFillShade := *rp.Shading.ThemeFillShade
			cloneShading.ThemeFillShade = &shadingThemeFillShade
		}
		clone.Shading = cloneShading
	}

	// Clone FitText (basic implementation)
	if rp.FitText != nil {
		clone.FitText = &FitText{ID: rp.FitText.ID, Val: rp.FitText.Val}
	}

	// Clone Lang (basic implementation)
	if rp.Lang != nil {
		cloneLang := &Lang{
			Val:      rp.Lang.Val,
			EastAsia: rp.Lang.EastAsia,
			Bidi:     rp.Lang.Bidi,
		}
		clone.Lang = cloneLang
	}

	// Clone EALayout (basic implementation)
	if rp.EALayout != nil {
		cloneEALayout := &EALayout{}
		if rp.EALayout.ID != nil {
			eaLayoutID := *rp.EALayout.ID
			cloneEALayout.ID = &eaLayoutID
		}
		if rp.EALayout.Combine != nil {
			eaLayoutCombine := *rp.EALayout.Combine
			cloneEALayout.Combine = &eaLayoutCombine
		}
		if rp.EALayout.CombineBrkts != nil {
			eaLayoutCombineBrkts := *rp.EALayout.CombineBrkts
			cloneEALayout.CombineBrkts = &eaLayoutCombineBrkts
		}
		if rp.EALayout.Vert != nil {
			eaLayoutVert := *rp.EALayout.Vert
			cloneEALayout.Vert = &eaLayoutVert
		}
		if rp.EALayout.VertCompress != nil {
			eaLayoutVertCompress := *rp.EALayout.VertCompress
			cloneEALayout.VertCompress = &eaLayoutVertCompress
		}
		clone.EALayout = cloneEALayout
	}

	// Note: We intentionally do NOT clone any existing rPrChange to avoid creating history of history
	// This ensures that when this method is used for revision tracking, we get a clean snapshot

	return clone
}
