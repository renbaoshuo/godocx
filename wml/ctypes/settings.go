package ctypes

import (
	"encoding/xml"
	"fmt"

	"github.com/gomutex/godocx/wml/stypes"
)

var defaultDocumentSettingsNSAttrs = map[string]string{
	"xmlns:mc":     "http://schemas.openxmlformats.org/markup-compatibility/2006",
	"xmlns:o":      "urn:schemas-microsoft-com:office:office",
	"xmlns:r":      "http://schemas.openxmlformats.org/officeDocument/2006/relationships",
	"xmlns:m":      "http://schemas.openxmlformats.org/officeDocument/2006/math",
	"xmlns:v":      "urn:schemas-microsoft-com:vml",
	"xmlns:w10":    "urn:schemas-microsoft-com:office:word",
	"xmlns:w":      "http://schemas.openxmlformats.org/wordprocessingml/2006/main",
	"xmlns:w14":    "http://schemas.microsoft.com/office/word/2010/wordml",
	"xmlns:sl":     "http://schemas.openxmlformats.org/schemaLibrary/2006/main",
	"mc:Ignorable": "w14",
}

// Settings represents the document-level settings.
// It corresponds to the <w:settings> element in WordprocessingML.
type DocumentSettings struct {
	RelativePath string `xml:"-"`
	Attr         []xml.Attr

	// 17.15.1.1 activeWritingStyle (Grammar Checking Settings)
	// ActiveWritingStyle *WritingStyle `xml:"activeWritingStyle,omitempty"`

	// 17.15.1.2 alignBordersAndEdges (Align Paragraph and Table Borders with Page Border)
	AlignBordersAndEdges *OnOff `xml:"alignBordersAndEdges,omitempty"`

	// 17.15.1.3 alwaysMergeEmptyNamespace (Do Not Mark Custom XML Elements With No Namespace As Invalid)
	AlwaysMergeEmptyNamespace *OnOff `xml:"alwaysMergeEmptyNamespace,omitempty"`

	// 17.15.1.4 alwaysShowPlaceholderText (Use Custom XML Element Names as Default Placeholder Text)
	AlwaysShowPlaceholderText *OnOff `xml:"alwaysShowPlaceholderText,omitempty"`

	// 17.15.1.5 attachedSchema (Attached Custom XML Schema)
	AttachedSchema []*CTString `xml:"attachedSchema,omitempty"`

	// 17.15.1.6 attachedTemplate (Attached Document Template)
	AttachedTemplate *Markup `xml:"attachedTemplate,omitempty"`

	// 17.15.1.7 autoCaption (Single Automatic Captioning Setting)
	// AutoCaption *AutoCaption `xml:"autoCaption,omitempty"`

	// 17.15.1.8 autoCaptions (Automatic Captioning Settings)
	// AutoCaptions *AutoCaptions `xml:"autoCaptions,omitempty"`

	// 17.15.1.9 autoFormatOverride (Allow Automatic Formatting to Override Formatting Protection Settings)
	AutoFormatOverride *OnOff `xml:"autoFormatOverride,omitempty"`

	// 17.15.1.10 autoHyphenation (Automatically Hyphenate Document Contents When Displayed)
	AutoHyphenation *OnOff `xml:"autoHyphenation,omitempty"`

	// 17.15.1.11 bookFoldPrinting (Book Fold Printing)
	BookFoldPrinting *OnOff `xml:"bookFoldPrinting,omitempty"`

	// 17.15.1.12 bookFoldPrintingSheets (Number of Pages Per Booklet)
	BookFoldPrintingSheets *GenSingleStrVal[stypes.DecimalNumber] `xml:"bookFoldPrintingSheets,omitempty"`

	// 17.15.1.13 bookFoldRevPrinting (Reverse Book Fold Printing)
	BookFoldRevPrinting *OnOff `xml:"bookFoldRevPrinting,omitempty"`

	// 17.15.1.14 bordersDoNotSurroundFooter (Page Border Excludes Footer)
	BordersDoNotSurroundFooter *OnOff `xml:"bordersDoNotSurroundFooter,omitempty"`

	// 17.15.1.15 bordersDoNotSurroundHeader (Page Border Excludes Header)
	BordersDoNotSurroundHeader *OnOff `xml:"bordersDoNotSurroundHeader,omitempty"`

	// 17.15.1.16 caption (Single Caption Type Definition)
	// Caption *Caption `xml:"caption,omitempty"`

	// 17.15.1.17 captions (Caption Settings)
	// Captions *Captions `xml:"captions,omitempty"`

	// 17.15.1.18 characterSpacingControl (Character-Level Whitespace Compression)
	CharacterSpacingControl *GenSingleStrVal[stypes.CharacterSpacing] `xml:"characterSpacingControl,omitempty"`

	// 17.15.1.19 clickAndTypeStyle (Paragraph Style Applied to Automatically Generated Paragraphs)
	ClickAndTypeStyle *CTString `xml:"clickAndTypeStyle,omitempty"`

	// 17.15.1.20 clrSchemeMapping (Theme Color Mappings)
	// ClrSchemeMapping *ColorSchemeMapping `xml:"clrSchemeMapping,omitempty"`

	// 17.15.1.21 compat (Compatibility Settings)
	// Compat *Compat `xml:"compat,omitempty"`

	// 17.15.1.22 consecutiveHyphenLimit (Maximum Number of Consecutively Hyphenated Lines)
	ConsecutiveHyphenLimit *GenSingleStrVal[stypes.DecimalNumber] `xml:"consecutiveHyphenLimit,omitempty"`

	// 17.15.1.23 decimalSymbol (Radix Point for Field Code Evaluation)
	DecimalSymbol *CTString `xml:"decimalSymbol,omitempty"`

	// 17.15.1.24 defaultTableStyle (Default Table Style for Newly Inserted Tables)
	DefaultTableStyle *CTString `xml:"defaultTableStyle,omitempty"`

	// 17.15.1.25 defaultTabStop (Distance Between Automatic Tab Stops)
	// DefaultTabStop *GenSingleStrVal[stypes.TwipsMeasure] `xml:"defaultTabStop,omitempty"`

	// 17.15.1.26 displayBackgroundShape (Display Background Objects When Displaying Document)
	DisplayBackgroundShape *OnOff `xml:"displayBackgroundShape,omitempty"`

	// 17.15.1.27 displayHorizontalDrawingGridEvery (Distance between Horizontal Gridlines)
	DisplayHorizontalDrawingGridEvery *GenSingleStrVal[stypes.DecimalNumber] `xml:"displayHorizontalDrawingGridEvery,omitempty"`

	// 17.15.1.28 displayVerticalDrawingGridEvery (Distance between Vertical Gridlines)
	DisplayVerticalDrawingGridEvery *GenSingleStrVal[stypes.DecimalNumber] `xml:"displayVerticalDrawingGridEvery,omitempty"`

	// 17.15.1.29 documentProtection (Document Editing Restrictions)
	// DocumentProtection *DocProtect `xml:"documentProtection,omitempty"`

	// 17.15.1.30 documentType (Document Classification)
	DocumentType *GenSingleStrVal[stypes.DocType] `xml:"documentType,omitempty"`

	// 17.15.1.31 docVar (Single Document Variable)
	// DocVar *DocVar `xml:"docVar,omitempty"`

	// 17.15.1.32 docVars (Document Variables)
	// DocVars *DocVars `xml:"docVars,omitempty"`

	// 17.15.1.33 doNotAutoCompressPictures (Do Not Automatically Compress Images)
	DoNotAutoCompressPictures *OnOff `xml:"doNotAutoCompressPictures,omitempty"`

	// 17.15.1.34 doNotDemarcateInvalidXml (Do Not Show Visual Indicator For Invalid Custom XML Markup)
	DoNotDemarcateInvalidXml *OnOff `xml:"doNotDemarcateInvalidXml,omitempty"`

	// 17.15.1.35 doNotDisplayPageBoundaries (Do Not Display Visual Boundary For Header/Footer or Between Pages)
	DoNotDisplayPageBoundaries *OnOff `xml:"doNotDisplayPageBoundaries,omitempty"`

	// 17.15.1.36 doNotEmbedSmartTags (Remove Smart Tags When Saving)
	DoNotEmbedSmartTags *OnOff `xml:"doNotEmbedSmartTags,omitempty"`

	// 17.15.1.37 doNotHyphenateCaps (Do Not Hyphenate Words in ALL CAPITAL LETTERS)
	DoNotHyphenateCaps *OnOff `xml:"doNotHyphenateCaps,omitempty"`

	// 17.15.1.38 doNotIncludeSubdocsInStats (Do Not Include Content in Text Boxes, Footnotes, and Endnotes in Document Statistics)
	DoNotIncludeSubdocsInStats *OnOff `xml:"doNotIncludeSubdocsInStats,omitempty"`

	// 17.15.1.39 doNotShadeFormData (Do Not Show Visual Indicator For Form Fields)
	DoNotShadeFormData *OnOff `xml:"doNotShadeFormData,omitempty"`

	// 17.15.1.40 doNotTrackFormatting (Do Not Track Formatting Revisions When Tracking Revisions)
	DoNotTrackFormatting *OnOff `xml:"doNotTrackFormatting,omitempty"`

	// 17.15.1.41 doNotTrackMoves (Do Not Use Move Syntax When Tracking Revisions)
	DoNotTrackMoves *OnOff `xml:"doNotTrackMoves,omitempty"`

	// 17.15.1.42 doNotUseMarginsForDrawingGridOrigin (Do Not Use Margins for Drawing Grid Origin)
	DoNotUseMarginsForDrawingGridOrigin *OnOff `xml:"doNotUseMarginsForDrawingGridOrigin,omitempty"`

	// 17.15.1.43 doNotValidateAgainstSchema (Do Not Validate Custom XML Markup Against Schemas)
	DoNotValidateAgainstSchema *OnOff `xml:"doNotValidateAgainstSchema,omitempty"`

	// 17.15.1.44 drawingGridHorizontalOrigin (Drawing Grid Horizontal Origin Point)
	// DrawingGridHorizontalOrigin *GenSingleStrVal[stypes.TwipsMeasure] `xml:"drawingGridHorizontalOrigin,omitempty"`

	// 17.15.1.45 drawingGridHorizontalSpacing (Drawing Grid Horizontal Grid Unit Size)
	// DrawingGridHorizontalSpacing *GenSingleStrVal[stypes.TwipsMeasure] `xml:"drawingGridHorizontalSpacing,omitempty"`

	// 17.15.1.46 drawingGridVerticalOrigin (Drawing Grid Vertical Origin Point)
	// DrawingGridVerticalOrigin *GenSingleStrVal[stypes.TwipsMeasure] `xml:"drawingGridVerticalOrigin,omitempty"`

	// 17.15.1.47 drawingGridVerticalSpacing (Drawing Grid Vertical Grid Unit Size)
	// DrawingGridVerticalSpacing *GenSingleStrVal[stypes.TwipsMeasure] `xml:"drawingGridVerticalSpacing,omitempty"`

	// 17.15.1.48 forceUpgrade (Upgrade Document on Open)
	ForceUpgrade *Empty `xml:"forceUpgrade,omitempty"`

	// 17.15.1.49 formsDesign (Structured Document Tag Placeholder Text Should be Resaved)
	FormsDesign *OnOff `xml:"formsDesign,omitempty"`

	// 17.15.1.50 gutterAtTop (Position Gutter At Top of Page)
	GutterAtTop *OnOff `xml:"gutterAtTop,omitempty"`

	// 17.15.1.51 hideGrammaticalErrors (Do Not Display Visual Indication of Grammatical Errors)
	HideGrammaticalErrors *OnOff `xml:"hideGrammaticalErrors,omitempty"`

	// 17.15.1.52 hideSpellingErrors (Do Not Display Visual Indication of Spelling Errors)
	HideSpellingErrors *OnOff `xml:"hideSpellingErrors,omitempty"`

	// 17.15.1.53 hyphenationZone (Hyphenation Zone)
	// HyphenationZone *GenSingleStrVal[stypes.TwipsMeasure] `xml:"hyphenationZone,omitempty"`

	// 17.15.1.54 ignoreMixedContent (Ignore Mixed Content When Validating Custom XML Markup)
	IgnoreMixedContent *OnOff `xml:"ignoreMixedContent,omitempty"`

	// 17.15.1.55 linkStyles (Automatically Update Styles From Document Template)
	LinkStyles *OnOff `xml:"linkStyles,omitempty"`

	// 17.15.1.56 listSeparator (List Separator for Field Code Evaluation)
	ListSeparator *CTString `xml:"listSeparator,omitempty"`

	// 17.15.1.57 mirrorMargins (Mirror Page Margins)
	MirrorMargins *OnOff `xml:"mirrorMargins,omitempty"`

	// 17.15.1.58 noLineBreaksAfter (Custom Set of Characters Which Cannot End a Line)
	// NoLineBreaksAfter *Kinsoku `xml:"noLineBreaksAfter,omitempty"`

	// 17.15.1.59 noLineBreaksBefore (Custom Set Of Characters Which Cannot Begin A Line)
	// NoLineBreaksBefore *Kinsoku `xml:"noLineBreaksBefore,omitempty"`

	// 17.15.1.60 noPunctuationKerning (Never Kern Punctuation Characters)
	NoPunctuationKerning *OnOff `xml:"noPunctuationKerning,omitempty"`

	// 17.15.1.61 printFormsData (Only Print Form Field Content)
	PrintFormsData *OnOff `xml:"printFormsData,omitempty"`

	// 17.15.1.62 printFractionalCharacterWidth (Print Fractional Character Widths)
	PrintFractionalCharacterWidth *OnOff `xml:"printFractionalCharacterWidth,omitempty"`

	// 17.15.1.63 printPostScriptOverText (Print PostScript Codes With Document Text)
	PrintPostScriptOverText *OnOff `xml:"printPostScriptOverText,omitempty"`

	// 17.15.1.64 printTwoOnOne (Print Two Pages Per Sheet)
	PrintTwoOnOne *OnOff `xml:"printTwoOnOne,omitempty"`

	// 17.15.1.65 proofState (Spelling and Grammatical Checking State)
	// ProofState *Proof `xml:"proofState,omitempty"`

	// 17.15.1.66 readModeInkLockDown (Freeze Document Layout)
	// ReadModeInkLockDown *ReadModeInkLockDown `xml:"readModeInkLockDown,omitempty"`

	// 17.15.1.67 removeDateAndTime (Remove Date and Time from Annotations)
	RemoveDateAndTime *OnOff `xml:"removeDateAndTime,omitempty"`

	// 17.15.1.68 removePersonalInformation (Remove Personal Information from Document Properties)
	RemovePersonalInformation *OnOff `xml:"removePersonalInformation,omitempty"`

	// 17.15.1.69 revisionView (Visibility of Annotation Types)
	// RevisionView *TrackChangesView `xml:"revisionView,omitempty"`

	// 17.15.1.70 rsid (Single Session Revision Save ID)
	Rsid *GenSingleStrVal[stypes.LongHexNum] `xml:"rsid,omitempty"`

	// 17.15.1.71 rsidRoot (Original Document Revision Save ID)
	RsidRoot *GenSingleStrVal[stypes.LongHexNum] `xml:"rsidRoot,omitempty"`

	// 17.15.1.72 rsids (Listing of All Revision Save ID Values)
	// Rsids *DocRsids `xml:"rsids,omitempty"`

	// 17.15.1.73 saveFormsData (Only Save Form Field Content)
	SaveFormsData *OnOff `xml:"saveFormsData,omitempty"`

	// 17.15.1.74 saveInvalidXml (Allow Saving Document As XML File When Custom XML Markup Is Invalid)
	SaveInvalidXml *OnOff `xml:"saveInvalidXml,omitempty"`

	// 17.15.1.75 savePreviewPicture (Generate Thumbnail For Document On Save)
	SavePreviewPicture *OnOff `xml:"savePreviewPicture,omitempty"`

	// 17.15.1.76 saveThroughXslt (Custom XSL Transform To Use When Saving As XML File)
	// SaveThroughXslt *SaveThroughXslt `xml:"saveThroughXslt,omitempty"`

	// 17.15.1.77 saveXmlDataOnly (Only Save Custom XML Markup)
	SaveXmlDataOnly *OnOff `xml:"saveXmlDataOnly,omitempty"`

	// 17.15.1.78 settings (Document Settings)
	// Settings *Settings `xml:"settings,omitempty"`

	// 17.15.1.79 showEnvelope (Show E-Mail Message Header)
	ShowEnvelope *OnOff `xml:"showEnvelope,omitempty"`

	// 17.15.1.80 showXMLTags (Show Visual Indicators for Custom XML Markup Start/End Locations)
	ShowXMLTags *OnOff `xml:"showXMLTags,omitempty"`

	// 17.15.1.81 smartTagType (Supplementary Smart Tag Information)
	// SmartTagType []*SmartTagType `xml:"smartTagType,omitempty"`

	// 17.15.1.82 strictFirstAndLastChars (Use Strict Kinsoku Rules for Japanese Text)
	StrictFirstAndLastChars *OnOff `xml:"strictFirstAndLastChars,omitempty"`

	// 17.15.1.83 styleLockQFSet (Prevent Replacement of Styles Part)
	StyleLockQFSet *OnOff `xml:"styleLockQFSet,omitempty"`

	// 17.15.1.84 styleLockTheme (Prevent Modification of Themes Part)
	StyleLockTheme *OnOff `xml:"styleLockTheme,omitempty"`

	// 17.15.1.85 stylePaneFormatFilter (Suggested Filtering for List of Document Styles)
	// StylePaneFormatFilter *StylePaneFilter `xml:"stylePaneFormatFilter,omitempty"`

	// 17.15.1.86 stylePaneSortMethod (Suggested Sorting for List of Document Styles)
	// StylePaneSortMethod *GenSingleStrVal[stypes.StyleSort] `xml:"stylePaneSortMethod,omitempty"`

	// 17.15.1.87 summaryLength (Percentage of Document to Use When Generating Summary)
	SummaryLength *CTString `xml:"summaryLength,omitempty"`

	// 17.15.1.88 themeFontLang (Theme Font Languages)
	// ThemeFontLang *Language `xml:"themeFontLang,omitempty"`

	// 17.15.1.89 trackRevisions (Track Revisions to Document)
	TrackRevisions *OnOff `xml:"trackRevisions,omitempty"`

	// 17.15.1.90 updateFields (Automatically Recalculate Fields on Open)
	UpdateFields *OnOff `xml:"updateFields,omitempty"`

	// 17.15.1.91 useXSLTWhenSaving (Save Document as XML File through Custom XSL Transform)
	UseXSLTWhenSaving *OnOff `xml:"useXSLTWhenSaving,omitempty"`

	// 17.15.1.92 view (Document View Setting)
	// View *GenSingleStrVal[stypes.View] `xml:"view,omitempty"`

	// 17.15.1.93 writeProtection (Write Protection)
	// WriteProtection *WriteProtection `xml:"writeProtection,omitempty"`

	// 17.15.1.94 zoom (Magnification Setting)
	// Zoom *Zoom `xml:"zoom,omitempty"`
}

// NewDocumentSettings creates a new DocumentSettings with default values.
func NewDocumentSettings() DocumentSettings {
	return DocumentSettings{}
}

// MarshalXML marshals DocumentSettings to XML.
func (ds DocumentSettings) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "w:settings"

	if len(ds.Attr) == 0 {
		for key, value := range defaultDocumentSettingsNSAttrs {
			attr := xml.Attr{Name: xml.Name{Local: key}, Value: value}
			start.Attr = append(start.Attr, attr)
		}
	} else {
		start.Attr = ds.Attr
	}

	if err := e.EncodeToken(start); err != nil {
		return err
	}

	// Group 1: OnOff elements - Document behavior settings
	onOffElements1 := []struct {
		elem    *OnOff
		xmlName string
	}{
		{ds.AlignBordersAndEdges, "w:alignBordersAndEdges"},                               // 17.15.1.2
		{ds.AlwaysMergeEmptyNamespace, "w:alwaysMergeEmptyNamespace"},                     // 17.15.1.3
		{ds.AlwaysShowPlaceholderText, "w:alwaysShowPlaceholderText"},                     // 17.15.1.4
		{ds.AutoFormatOverride, "w:autoFormatOverride"},                                   // 17.15.1.9
		{ds.AutoHyphenation, "w:autoHyphenation"},                                         // 17.15.1.10
		{ds.BookFoldPrinting, "w:bookFoldPrinting"},                                       // 17.15.1.11
		{ds.BookFoldRevPrinting, "w:bookFoldRevPrinting"},                                 // 17.15.1.13
		{ds.BordersDoNotSurroundFooter, "w:bordersDoNotSurroundFooter"},                   // 17.15.1.14
		{ds.BordersDoNotSurroundHeader, "w:bordersDoNotSurroundHeader"},                   // 17.15.1.15
		{ds.DisplayBackgroundShape, "w:displayBackgroundShape"},                           // 17.15.1.26
		{ds.DoNotAutoCompressPictures, "w:doNotAutoCompressPictures"},                     // 17.15.1.33
		{ds.DoNotDemarcateInvalidXml, "w:doNotDemarcateInvalidXml"},                       // 17.15.1.34
		{ds.DoNotDisplayPageBoundaries, "w:doNotDisplayPageBoundaries"},                   // 17.15.1.35
		{ds.DoNotEmbedSmartTags, "w:doNotEmbedSmartTags"},                                 // 17.15.1.36
		{ds.DoNotHyphenateCaps, "w:doNotHyphenateCaps"},                                   // 17.15.1.37
		{ds.DoNotIncludeSubdocsInStats, "w:doNotIncludeSubdocsInStats"},                   // 17.15.1.38
		{ds.DoNotShadeFormData, "w:doNotShadeFormData"},                                   // 17.15.1.39
		{ds.DoNotTrackFormatting, "w:doNotTrackFormatting"},                               // 17.15.1.40
		{ds.DoNotTrackMoves, "w:doNotTrackMoves"},                                         // 17.15.1.41
		{ds.DoNotUseMarginsForDrawingGridOrigin, "w:doNotUseMarginsForDrawingGridOrigin"}, // 17.15.1.42
		{ds.DoNotValidateAgainstSchema, "w:doNotValidateAgainstSchema"},                   // 17.15.1.43
		{ds.FormsDesign, "w:formsDesign"},                                                 // 17.15.1.49
		{ds.GutterAtTop, "w:gutterAtTop"},                                                 // 17.15.1.50
		{ds.HideGrammaticalErrors, "w:hideGrammaticalErrors"},                             // 17.15.1.51
		{ds.HideSpellingErrors, "w:hideSpellingErrors"},                                   // 17.15.1.52
		{ds.IgnoreMixedContent, "w:ignoreMixedContent"},                                   // 17.15.1.54
		{ds.LinkStyles, "w:linkStyles"},                                                   // 17.15.1.55
		{ds.MirrorMargins, "w:mirrorMargins"},                                             // 17.15.1.57
		{ds.NoPunctuationKerning, "w:noPunctuationKerning"},                               // 17.15.1.60
		{ds.PrintFormsData, "w:printFormsData"},                                           // 17.15.1.61
		{ds.PrintFractionalCharacterWidth, "w:printFractionalCharacterWidth"},             // 17.15.1.62
		{ds.PrintPostScriptOverText, "w:printPostScriptOverText"},                         // 17.15.1.63
		{ds.PrintTwoOnOne, "w:printTwoOnOne"},                                             // 17.15.1.64
		{ds.RemoveDateAndTime, "w:removeDateAndTime"},                                     // 17.15.1.67
		{ds.RemovePersonalInformation, "w:removePersonalInformation"},                     // 17.15.1.68
		{ds.SaveFormsData, "w:saveFormsData"},                                             // 17.15.1.73
		{ds.SaveInvalidXml, "w:saveInvalidXml"},                                           // 17.15.1.74
		{ds.SavePreviewPicture, "w:savePreviewPicture"},                                   // 17.15.1.75
		{ds.SaveXmlDataOnly, "w:saveXmlDataOnly"},                                         // 17.15.1.77
		{ds.ShowEnvelope, "w:showEnvelope"},                                               // 17.15.1.79
		{ds.ShowXMLTags, "w:showXMLTags"},                                                 // 17.15.1.80
		{ds.StrictFirstAndLastChars, "w:strictFirstAndLastChars"},                         // 17.15.1.82
		{ds.StyleLockQFSet, "w:styleLockQFSet"},                                           // 17.15.1.83
		{ds.StyleLockTheme, "w:styleLockTheme"},                                           // 17.15.1.84
		{ds.TrackRevisions, "w:trackRevisions"},                                           // 17.15.1.89
		{ds.UpdateFields, "w:updateFields"},                                               // 17.15.1.90
		{ds.UseXSLTWhenSaving, "w:useXSLTWhenSaving"},                                     // 17.15.1.91
	}

	for _, entry := range onOffElements1 {
		if entry.elem == nil {
			continue
		}
		if err := entry.elem.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: entry.xmlName},
		}); err != nil {
			return fmt.Errorf("error in marshaling document settings `%s`: %w", entry.xmlName, err)
		}
	}

	// 17.15.1.5 attachedSchema (Attached Custom XML Schema)
	for _, schema := range ds.AttachedSchema {
		if schema != nil {
			if err := schema.MarshalXML(e, xml.StartElement{
				Name: xml.Name{Local: "w:attachedSchema"},
			}); err != nil {
				return fmt.Errorf("attachedSchema: %w", err)
			}
		}
	}

	// 17.15.1.6 attachedTemplate (Attached Document Template)
	if ds.AttachedTemplate != nil {
		if err := ds.AttachedTemplate.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:attachedTemplate"},
		}); err != nil {
			return fmt.Errorf("attachedTemplate: %w", err)
		}
	}

	// 17.15.1.12 bookFoldPrintingSheets (Number of Pages Per Booklet)
	if ds.BookFoldPrintingSheets != nil {
		if err := ds.BookFoldPrintingSheets.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:bookFoldPrintingSheets"},
		}); err != nil {
			return fmt.Errorf("bookFoldPrintingSheets: %w", err)
		}
	}

	// 17.15.1.18 characterSpacingControl (Character-Level Whitespace Compression)
	if ds.CharacterSpacingControl != nil {
		if err := ds.CharacterSpacingControl.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:characterSpacingControl"},
		}); err != nil {
			return fmt.Errorf("characterSpacingControl: %w", err)
		}
	}

	// 17.15.1.19 clickAndTypeStyle (Paragraph Style Applied to Automatically Generated Paragraphs)
	if ds.ClickAndTypeStyle != nil {
		if err := ds.ClickAndTypeStyle.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:clickAndTypeStyle"},
		}); err != nil {
			return fmt.Errorf("clickAndTypeStyle: %w", err)
		}
	}

	// 17.15.1.22 consecutiveHyphenLimit (Maximum Number of Consecutively Hyphenated Lines)
	if ds.ConsecutiveHyphenLimit != nil {
		if err := ds.ConsecutiveHyphenLimit.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:consecutiveHyphenLimit"},
		}); err != nil {
			return fmt.Errorf("consecutiveHyphenLimit: %w", err)
		}
	}

	// 17.15.1.23 decimalSymbol (Radix Point for Field Code Evaluation)
	if ds.DecimalSymbol != nil {
		if err := ds.DecimalSymbol.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:decimalSymbol"},
		}); err != nil {
			return fmt.Errorf("decimalSymbol: %w", err)
		}
	}

	// 17.15.1.24 defaultTableStyle (Default Table Style for Newly Inserted Tables)
	if ds.DefaultTableStyle != nil {
		if err := ds.DefaultTableStyle.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:defaultTableStyle"},
		}); err != nil {
			return fmt.Errorf("defaultTableStyle: %w", err)
		}
	}

	// 17.15.1.27 displayHorizontalDrawingGridEvery (Distance between Horizontal Gridlines)
	if ds.DisplayHorizontalDrawingGridEvery != nil {
		if err := ds.DisplayHorizontalDrawingGridEvery.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:displayHorizontalDrawingGridEvery"},
		}); err != nil {
			return fmt.Errorf("displayHorizontalDrawingGridEvery: %w", err)
		}
	}

	// 17.15.1.28 displayVerticalDrawingGridEvery (Distance between Vertical Gridlines)
	if ds.DisplayVerticalDrawingGridEvery != nil {
		if err := ds.DisplayVerticalDrawingGridEvery.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:displayVerticalDrawingGridEvery"},
		}); err != nil {
			return fmt.Errorf("displayVerticalDrawingGridEvery: %w", err)
		}
	}

	// 17.15.1.30 documentType (Document Classification)
	if ds.DocumentType != nil {
		if err := ds.DocumentType.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:documentType"},
		}); err != nil {
			return fmt.Errorf("documentType: %w", err)
		}
	}

	// 17.15.1.48 forceUpgrade (Upgrade Document on Open)
	if ds.ForceUpgrade != nil {
		if err := ds.ForceUpgrade.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:forceUpgrade"},
		}); err != nil {
			return fmt.Errorf("forceUpgrade: %w", err)
		}
	}

	// 17.15.1.56 listSeparator (List Separator for Field Code Evaluation)
	if ds.ListSeparator != nil {
		if err := ds.ListSeparator.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:listSeparator"},
		}); err != nil {
			return fmt.Errorf("listSeparator: %w", err)
		}
	}

	// 17.15.1.70 rsid (Single Session Revision Save ID)
	if ds.Rsid != nil {
		if err := ds.Rsid.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:rsid"},
		}); err != nil {
			return fmt.Errorf("rsid: %w", err)
		}
	}

	// 17.15.1.71 rsidRoot (Original Document Revision Save ID)
	if ds.RsidRoot != nil {
		if err := ds.RsidRoot.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:rsidRoot"},
		}); err != nil {
			return fmt.Errorf("rsidRoot: %w", err)
		}
	}

	// 17.15.1.87 summaryLength (Percentage of Document to Use When Generating Summary)
	if ds.SummaryLength != nil {
		if err := ds.SummaryLength.MarshalXML(e, xml.StartElement{
			Name: xml.Name{Local: "w:summaryLength"},
		}); err != nil {
			return fmt.Errorf("summaryLength: %w", err)
		}
	}

	return e.EncodeToken(start.End())
}

// UnmarshalXML unmarshals DocumentSettings from XML.
func (ds *DocumentSettings) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

loop:
	for {
		currentToken, err := d.Token()
		if err != nil {
			return err
		}

		switch elem := currentToken.(type) {
		case xml.StartElement:
			switch elem.Name.Local {
			// 17.15.1.2 alignBordersAndEdges
			case "alignBordersAndEdges":
				onOff := OnOff{}
				if err := d.DecodeElement(&onOff, &elem); err != nil {
					return err
				}
				ds.AlignBordersAndEdges = &onOff

			// 17.15.1.3 alwaysMergeEmptyNamespace
			case "alwaysMergeEmptyNamespace":
				onOff := OnOff{}
				if err := d.DecodeElement(&onOff, &elem); err != nil {
					return err
				}
				ds.AlwaysMergeEmptyNamespace = &onOff

			// 17.15.1.4 alwaysShowPlaceholderText
			case "alwaysShowPlaceholderText":
				onOff := OnOff{}
				if err := d.DecodeElement(&onOff, &elem); err != nil {
					return err
				}
				ds.AlwaysShowPlaceholderText = &onOff

			// 17.15.1.5 attachedSchema
			case "attachedSchema":
				schema := CTString{}
				if err := d.DecodeElement(&schema, &elem); err != nil {
					return err
				}
				ds.AttachedSchema = append(ds.AttachedSchema, &schema)

			// 17.15.1.6 attachedTemplate
			case "attachedTemplate":
				template := Markup{}
				if err := d.DecodeElement(&template, &elem); err != nil {
					return err
				}
				ds.AttachedTemplate = &template

			// 17.15.1.9 autoFormatOverride
			case "autoFormatOverride":
				onOff := OnOff{}
				if err := d.DecodeElement(&onOff, &elem); err != nil {
					return err
				}
				ds.AutoFormatOverride = &onOff

			// 17.15.1.10 autoHyphenation
			case "autoHyphenation":
				onOff := OnOff{}
				if err := d.DecodeElement(&onOff, &elem); err != nil {
					return err
				}
				ds.AutoHyphenation = &onOff

			// 17.15.1.11 bookFoldPrinting
			case "bookFoldPrinting":
				onOff := OnOff{}
				if err := d.DecodeElement(&onOff, &elem); err != nil {
					return err
				}
				ds.BookFoldPrinting = &onOff

			// 17.15.1.12 bookFoldPrintingSheets
			case "bookFoldPrintingSheets":
				sheets := GenSingleStrVal[stypes.DecimalNumber]{}
				if err := d.DecodeElement(&sheets, &elem); err != nil {
					return err
				}
				ds.BookFoldPrintingSheets = &sheets

			// 17.15.1.13 bookFoldRevPrinting
			case "bookFoldRevPrinting":
				onOff := OnOff{}
				if err := d.DecodeElement(&onOff, &elem); err != nil {
					return err
				}
				ds.BookFoldRevPrinting = &onOff

			// 17.15.1.14 bordersDoNotSurroundFooter
			case "bordersDoNotSurroundFooter":
				onOff := OnOff{}
				if err := d.DecodeElement(&onOff, &elem); err != nil {
					return err
				}
				ds.BordersDoNotSurroundFooter = &onOff

			// 17.15.1.15 bordersDoNotSurroundHeader
			case "bordersDoNotSurroundHeader":
				onOff := OnOff{}
				if err := d.DecodeElement(&onOff, &elem); err != nil {
					return err
				}
				ds.BordersDoNotSurroundHeader = &onOff

			// 17.15.1.18 characterSpacingControl
			case "characterSpacingControl":
				control := GenSingleStrVal[stypes.CharacterSpacing]{}
				if err := d.DecodeElement(&control, &elem); err != nil {
					return err
				}
				ds.CharacterSpacingControl = &control

			// 17.15.1.19 clickAndTypeStyle
			case "clickAndTypeStyle":
				style := CTString{}
				if err := d.DecodeElement(&style, &elem); err != nil {
					return err
				}
				ds.ClickAndTypeStyle = &style

			// 17.15.1.22 consecutiveHyphenLimit
			case "consecutiveHyphenLimit":
				limit := GenSingleStrVal[stypes.DecimalNumber]{}
				if err := d.DecodeElement(&limit, &elem); err != nil {
					return err
				}
				ds.ConsecutiveHyphenLimit = &limit

			// 17.15.1.23 decimalSymbol
			case "decimalSymbol":
				symbol := CTString{}
				if err := d.DecodeElement(&symbol, &elem); err != nil {
					return err
				}
				ds.DecimalSymbol = &symbol

			// 17.15.1.24 defaultTableStyle
			case "defaultTableStyle":
				style := CTString{}
				if err := d.DecodeElement(&style, &elem); err != nil {
					return err
				}
				ds.DefaultTableStyle = &style

			// 17.15.1.26 displayBackgroundShape
			case "displayBackgroundShape":
				onOff := OnOff{}
				if err := d.DecodeElement(&onOff, &elem); err != nil {
					return err
				}
				ds.DisplayBackgroundShape = &onOff

			// 17.15.1.27 displayHorizontalDrawingGridEvery
			case "displayHorizontalDrawingGridEvery":
				grid := GenSingleStrVal[stypes.DecimalNumber]{}
				if err := d.DecodeElement(&grid, &elem); err != nil {
					return err
				}
				ds.DisplayHorizontalDrawingGridEvery = &grid

			// 17.15.1.28 displayVerticalDrawingGridEvery
			case "displayVerticalDrawingGridEvery":
				grid := GenSingleStrVal[stypes.DecimalNumber]{}
				if err := d.DecodeElement(&grid, &elem); err != nil {
					return err
				}
				ds.DisplayVerticalDrawingGridEvery = &grid

			// 17.15.1.30 documentType
			case "documentType":
				docType := GenSingleStrVal[stypes.DocType]{}
				if err := d.DecodeElement(&docType, &elem); err != nil {
					return err
				}
				ds.DocumentType = &docType

			// Continue with remaining OnOff elements
			case "doNotAutoCompressPictures":
				onOff := OnOff{}
				if err := d.DecodeElement(&onOff, &elem); err != nil {
					return err
				}
				ds.DoNotAutoCompressPictures = &onOff

			case "doNotDemarcateInvalidXml":
				onOff := OnOff{}
				if err := d.DecodeElement(&onOff, &elem); err != nil {
					return err
				}
				ds.DoNotDemarcateInvalidXml = &onOff

			case "doNotDisplayPageBoundaries":
				onOff := OnOff{}
				if err := d.DecodeElement(&onOff, &elem); err != nil {
					return err
				}
				ds.DoNotDisplayPageBoundaries = &onOff

			case "doNotEmbedSmartTags":
				onOff := OnOff{}
				if err := d.DecodeElement(&onOff, &elem); err != nil {
					return err
				}
				ds.DoNotEmbedSmartTags = &onOff

			case "doNotHyphenateCaps":
				onOff := OnOff{}
				if err := d.DecodeElement(&onOff, &elem); err != nil {
					return err
				}
				ds.DoNotHyphenateCaps = &onOff

			case "doNotIncludeSubdocsInStats":
				onOff := OnOff{}
				if err := d.DecodeElement(&onOff, &elem); err != nil {
					return err
				}
				ds.DoNotIncludeSubdocsInStats = &onOff

			case "doNotShadeFormData":
				onOff := OnOff{}
				if err := d.DecodeElement(&onOff, &elem); err != nil {
					return err
				}
				ds.DoNotShadeFormData = &onOff

			case "doNotTrackFormatting":
				onOff := OnOff{}
				if err := d.DecodeElement(&onOff, &elem); err != nil {
					return err
				}
				ds.DoNotTrackFormatting = &onOff

			case "doNotTrackMoves":
				onOff := OnOff{}
				if err := d.DecodeElement(&onOff, &elem); err != nil {
					return err
				}
				ds.DoNotTrackMoves = &onOff

			case "doNotUseMarginsForDrawingGridOrigin":
				onOff := OnOff{}
				if err := d.DecodeElement(&onOff, &elem); err != nil {
					return err
				}
				ds.DoNotUseMarginsForDrawingGridOrigin = &onOff

			case "doNotValidateAgainstSchema":
				onOff := OnOff{}
				if err := d.DecodeElement(&onOff, &elem); err != nil {
					return err
				}
				ds.DoNotValidateAgainstSchema = &onOff

			case "forceUpgrade":
				empty := Empty{}
				if err := d.DecodeElement(&empty, &elem); err != nil {
					return err
				}
				ds.ForceUpgrade = &empty

			case "formsDesign":
				onOff := OnOff{}
				if err := d.DecodeElement(&onOff, &elem); err != nil {
					return err
				}
				ds.FormsDesign = &onOff

			case "gutterAtTop":
				onOff := OnOff{}
				if err := d.DecodeElement(&onOff, &elem); err != nil {
					return err
				}
				ds.GutterAtTop = &onOff

			case "hideGrammaticalErrors":
				onOff := OnOff{}
				if err := d.DecodeElement(&onOff, &elem); err != nil {
					return err
				}
				ds.HideGrammaticalErrors = &onOff

			case "hideSpellingErrors":
				onOff := OnOff{}
				if err := d.DecodeElement(&onOff, &elem); err != nil {
					return err
				}
				ds.HideSpellingErrors = &onOff

			case "ignoreMixedContent":
				onOff := OnOff{}
				if err := d.DecodeElement(&onOff, &elem); err != nil {
					return err
				}
				ds.IgnoreMixedContent = &onOff

			case "linkStyles":
				onOff := OnOff{}
				if err := d.DecodeElement(&onOff, &elem); err != nil {
					return err
				}
				ds.LinkStyles = &onOff

			case "listSeparator":
				separator := CTString{}
				if err := d.DecodeElement(&separator, &elem); err != nil {
					return err
				}
				ds.ListSeparator = &separator

			case "mirrorMargins":
				onOff := OnOff{}
				if err := d.DecodeElement(&onOff, &elem); err != nil {
					return err
				}
				ds.MirrorMargins = &onOff

			case "noPunctuationKerning":
				onOff := OnOff{}
				if err := d.DecodeElement(&onOff, &elem); err != nil {
					return err
				}
				ds.NoPunctuationKerning = &onOff

			case "printFormsData":
				onOff := OnOff{}
				if err := d.DecodeElement(&onOff, &elem); err != nil {
					return err
				}
				ds.PrintFormsData = &onOff

			case "printFractionalCharacterWidth":
				onOff := OnOff{}
				if err := d.DecodeElement(&onOff, &elem); err != nil {
					return err
				}
				ds.PrintFractionalCharacterWidth = &onOff

			case "printPostScriptOverText":
				onOff := OnOff{}
				if err := d.DecodeElement(&onOff, &elem); err != nil {
					return err
				}
				ds.PrintPostScriptOverText = &onOff

			case "printTwoOnOne":
				onOff := OnOff{}
				if err := d.DecodeElement(&onOff, &elem); err != nil {
					return err
				}
				ds.PrintTwoOnOne = &onOff

			case "removeDateAndTime":
				onOff := OnOff{}
				if err := d.DecodeElement(&onOff, &elem); err != nil {
					return err
				}
				ds.RemoveDateAndTime = &onOff

			case "removePersonalInformation":
				onOff := OnOff{}
				if err := d.DecodeElement(&onOff, &elem); err != nil {
					return err
				}
				ds.RemovePersonalInformation = &onOff

			case "rsid":
				rsid := GenSingleStrVal[stypes.LongHexNum]{}
				if err := d.DecodeElement(&rsid, &elem); err != nil {
					return err
				}
				ds.Rsid = &rsid

			case "rsidRoot":
				rsidRoot := GenSingleStrVal[stypes.LongHexNum]{}
				if err := d.DecodeElement(&rsidRoot, &elem); err != nil {
					return err
				}
				ds.RsidRoot = &rsidRoot

			case "saveFormsData":
				onOff := OnOff{}
				if err := d.DecodeElement(&onOff, &elem); err != nil {
					return err
				}
				ds.SaveFormsData = &onOff

			case "saveInvalidXml":
				onOff := OnOff{}
				if err := d.DecodeElement(&onOff, &elem); err != nil {
					return err
				}
				ds.SaveInvalidXml = &onOff

			case "savePreviewPicture":
				onOff := OnOff{}
				if err := d.DecodeElement(&onOff, &elem); err != nil {
					return err
				}
				ds.SavePreviewPicture = &onOff

			case "saveXmlDataOnly":
				onOff := OnOff{}
				if err := d.DecodeElement(&onOff, &elem); err != nil {
					return err
				}
				ds.SaveXmlDataOnly = &onOff

			case "showEnvelope":
				onOff := OnOff{}
				if err := d.DecodeElement(&onOff, &elem); err != nil {
					return err
				}
				ds.ShowEnvelope = &onOff

			case "showXMLTags":
				onOff := OnOff{}
				if err := d.DecodeElement(&onOff, &elem); err != nil {
					return err
				}
				ds.ShowXMLTags = &onOff

			case "strictFirstAndLastChars":
				onOff := OnOff{}
				if err := d.DecodeElement(&onOff, &elem); err != nil {
					return err
				}
				ds.StrictFirstAndLastChars = &onOff

			case "styleLockQFSet":
				onOff := OnOff{}
				if err := d.DecodeElement(&onOff, &elem); err != nil {
					return err
				}
				ds.StyleLockQFSet = &onOff

			case "styleLockTheme":
				onOff := OnOff{}
				if err := d.DecodeElement(&onOff, &elem); err != nil {
					return err
				}
				ds.StyleLockTheme = &onOff

			case "summaryLength":
				length := CTString{}
				if err := d.DecodeElement(&length, &elem); err != nil {
					return err
				}
				ds.SummaryLength = &length

			case "trackRevisions":
				onOff := OnOff{}
				if err := d.DecodeElement(&onOff, &elem); err != nil {
					return err
				}
				ds.TrackRevisions = &onOff

			case "updateFields":
				onOff := OnOff{}
				if err := d.DecodeElement(&onOff, &elem); err != nil {
					return err
				}
				ds.UpdateFields = &onOff

			case "useXSLTWhenSaving":
				onOff := OnOff{}
				if err := d.DecodeElement(&onOff, &elem); err != nil {
					return err
				}
				ds.UseXSLTWhenSaving = &onOff

			default:
				// Skip unknown elements
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break loop
		}
	}

	return nil
}
