package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/waus/tmark-go"
)

const (
	staticURL = "https://tmark.waus.app/static/"
	audioURL  = staticURL + "mountain-ambient.mp3"
	videoURL  = staticURL + "Big_Buck_Bunny_720_10s_1MB.mp4"
)

type example struct {
	filename string
	document tmark.Document
}

func main() {
	outDir := filepath.Join("examples")
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		panic(err)
	}

	for _, ex := range examples() {
		data, err := tmark.Marshal(ex.document)
		if err != nil {
			panic(fmt.Errorf("%s: %w", ex.filename, err))
		}
		path := filepath.Join(outDir, ex.filename)
		if err := os.WriteFile(path, append(data, '\n'), 0o644); err != nil {
			panic(err)
		}
		fmt.Println(path)
	}
}

func asset(name string) tmark.Text {
	return tmark.Text(staticURL + name)
}

func examples() []example {
	return []example{
		{"all_widgets.tmark", document("all_widgets.tmark", "All widgets", allWidgets())},
		{"inline_formatting.tmark", document("inline_formatting.tmark", "Inline formatting", inlineFormatting())},
		{"media_gallery.tmark", document("media_gallery.tmark", "Media gallery", mediaGallery())},
		{"tables.tmark", document("tables.tmark", "Tables", tables())},
		{"lists_and_tasks.tmark", document("lists_and_tasks.tmark", "Lists and tasks", listsAndTasks())},
		{"quotes_and_details.tmark", document("quotes_and_details.tmark", "Quotes and details", quotesAndDetails())},
		{"math_and_code.tmark", document("math_and_code.tmark", "Math and code", mathAndCode())},
		{"references_and_anchors.tmark", document("references_and_anchors.tmark", "References and anchors", referencesAndAnchors())},
		{"maps_and_events.tmark", document("maps_and_events.tmark", "Maps and events", mapsAndEvents())},
		{"edge_empties.tmark", document("edge_empties.tmark", "Empty elements", edgeEmpties())},
		{"edge_cases.tmark", document("edge_cases.tmark", "Other edge cases", edgeCases())},
	}
}

func document(filename string, title tmark.Text, content tmark.RichBlocks) tmark.Document {
	return tmark.Document{
		URL:     tmark.T("https://tmark.waus.app/examples/" + filename),
		Title:   title,
		Content: content,
	}
}

func allWidgets() tmark.RichBlocks {
	return tmark.NewRichBlocks(
		tmark.NewHeader(3, tmark.T("Text and formatting")),
		tmark.NewHeader(4, tmark.T("Inline formatting")),
		tmark.P(
			tmark.T("Plain text, "),
			tmark.B(tmark.T("bold text")),
			tmark.T(", "),
			tmark.I(tmark.T("italic text")),
			tmark.T(", "),
			tmark.NewMarked(tmark.T("marked text")),
			tmark.T(", "),
			tmark.NewUnderline(tmark.T("underlined text")),
			tmark.T(", "),
			tmark.NewStrikethrough(tmark.T("struck text")),
			tmark.T(", "),
			tmark.NewSpoiler(tmark.T("spoiler text")),
			tmark.T(", "),
			tmark.NewSubscript(tmark.T("sub")),
			tmark.T(" and "),
			tmark.NewSuperscript(tmark.T("sup")),
			tmark.T(", "),
			tmark.NewCode("inline code"),
			tmark.T(", "),
			tmark.NewMath("E=mc^2"),
			tmark.T(", "),
			tmark.NewDateTime(1767225600, "UTC"),
			tmark.T(", "),
			tmark.NewLink("https://example.com", tmark.T("external link")),
			tmark.T(", and an icon: "),
			tmark.NewIcon(asset("c-logo.png")).WithAlternativeText("C language logo"),
			tmark.T("."),
		),
		tmark.P(tmark.T("A paragraph can contain line breaks.\nThis line is still in the same paragraph.")),
		tmark.NewPreformatted("package main\n\nfunc main() {\n\tprintln(\"Hello, tmark\")\n}").WithLanguage("go"),
		tmark.NewMathBlock("\\int_0^1 x^2 dx = \\frac{1}{3}"),
		tmark.NewAnchor("section-anchor"),
		tmark.NewPullQuote(tmark.T("A pull quote highlights a short passage of text.")).WithCredit(tmark.T("Example author")),
		tmark.NewBlockquote(
			tmark.NewHeader(4, tmark.T("This is a blockquote")),
			tmark.P(tmark.T("It can contain paragraphs, images, and other blocks."))).
			WithCredit(tmark.T("Example author")),
		tmark.NewList(
			tmark.NewListItem(tmark.P(tmark.T("A plain bullet"))),
			tmark.NewListItem(tmark.P(tmark.T("Another bullet"))),
		),
		tmark.NewList(
			tmark.NewListItem(tmark.P(tmark.T("A numbered item"))).WithOrder(1),
			tmark.NewListItem(tmark.P(tmark.T("An item labeled with a letter"))).WithOrder(2).WithType("A"),
		),
		tmark.NewList(
			tmark.NewListItem(tmark.P(tmark.T("A completed task"))).WithType("checkbox").WithChecked(true),
		),
		tmark.NewDivider(),
		tmark.NewMap(52.520008, 13.404954).WithZoom(12).WithCaption(tmark.NewCaption(tmark.T("Berlin"))),
		tmark.NewImage(asset("mountains-wide.webp")).WithCaption(tmark.NewCaption(tmark.T("Mountain panorama")).WithCredit(tmark.T("Unknown author"))),
		tmark.NewCollage(
			tmark.NewImage(asset("mountains-lake.webp")),
			tmark.NewImage(asset("mountains-portrait.webp")),
			tmark.NewImage(asset("mountains-summit.webp")),
		).WithCaption(tmark.NewCaption(tmark.T("Image collage"))),
		tmark.NewSlideshow(
			tmark.NewImage(asset("mountains-rockies.webp")),
			tmark.NewImage(asset("mountains-valley.webp")),
			tmark.NewVideo(videoURL),
		).WithCaption(tmark.NewCaption(tmark.T("Image and video slideshow"))),
		tmark.NewTable(
			tmark.NewTableRow(tmark.NewCell(tmark.T("")).Header(), tmark.NewCell(tmark.T("Tmark")).Header().WithAlign(tmark.TableCellAlignCenter), tmark.NewCell(tmark.T("Markdown")).Header().WithAlign(tmark.TableCellAlignCenter), tmark.NewCell(tmark.T("HTML")).Header().WithAlign(tmark.TableCellAlignCenter), tmark.NewCell(tmark.T("LaTeX")).Header().WithAlign(tmark.TableCellAlignCenter)),
			tmark.NewTableRow(tmark.NewCell(tmark.T("Strict")).Header(), tmark.NewCell(tmark.T("***")).WithAlign(tmark.TableCellAlignCenter), tmark.NewCell(tmark.T("*")).WithAlign(tmark.TableCellAlignCenter), tmark.NewCell(tmark.T("**")).WithAlign(tmark.TableCellAlignCenter), tmark.NewCell(tmark.T("***")).WithAlign(tmark.TableCellAlignCenter)),
			tmark.NewTableRow(tmark.NewCell(tmark.T("Batteries")).Header(), tmark.NewCell(tmark.T("***")).WithAlign(tmark.TableCellAlignCenter), tmark.NewCell(tmark.T("**")).WithAlign(tmark.TableCellAlignCenter), tmark.NewCell(tmark.T("****")).WithAlign(tmark.TableCellAlignCenter), tmark.NewCell(tmark.T("****")).WithAlign(tmark.TableCellAlignCenter)),
			tmark.NewTableRow(tmark.NewCell(tmark.T("Extensibility")).Header(), tmark.NewCell(tmark.T("***")).WithAlign(tmark.TableCellAlignCenter), tmark.NewCell(tmark.T("*")).WithAlign(tmark.TableCellAlignCenter), tmark.NewCell(tmark.T("**")).WithAlign(tmark.TableCellAlignCenter), tmark.NewCell(tmark.T("****")).WithAlign(tmark.TableCellAlignCenter)),
			tmark.NewTableRow(tmark.NewCell(tmark.T("Complexity")).Header(), tmark.NewCell(tmark.T("*")).WithAlign(tmark.TableCellAlignCenter), tmark.NewCell(tmark.T("***")).WithAlign(tmark.TableCellAlignCenter), tmark.NewCell(tmark.T("***")).WithAlign(tmark.TableCellAlignCenter), tmark.NewCell(tmark.T("*****")).WithAlign(tmark.TableCellAlignCenter)),
		).WithCaption(tmark.T("Table")).WithBordered().WithStriped(),
		tmark.NewDetails(tmark.P(tmark.T("Details can also contain several paragraphs or other blocks."))).WithSummary(tmark.T("More details")),
		tmark.NewAudio(audioURL).WithCaption(tmark.NewCaption(tmark.T("Audio"))),
	)
}

func inlineFormatting() tmark.RichBlocks {
	return tmark.NewRichBlocks(
		tmark.NewHeader(2, tmark.T("Combine inline styles")),
		tmark.P(
			tmark.T("Nested: "),
			tmark.B(tmark.I(tmark.NewUnderline(tmark.T("bold, italic, and underlined")))),
			tmark.T("; hidden: "),
			tmark.NewSpoiler(tmark.T("spoiler")),
			tmark.T("; chemistry: H"),
			tmark.NewSubscript(tmark.T("2")),
			tmark.T("O; power: x"),
			tmark.NewSuperscript(tmark.T("2")),
			tmark.T("."),
		),
		tmark.P(
			tmark.NewLink("https://example.com/docs", tmark.T("External link")),
			tmark.T(" and "),
			tmark.NewAnchorLink("formatting-notes", tmark.T("link to the note below")),
			tmark.T(" can be mixed with "),
			tmark.NewCode("code"),
			tmark.T(" and "),
			tmark.NewMath("a^2+b^2=c^2"),
			tmark.T("."),
		),
		tmark.NewAnchor("formatting-notes"),
		tmark.P(tmark.T("This note is the target of the anchor link above.")),
	)
}

func mediaGallery() tmark.RichBlocks {
	return tmark.NewRichBlocks(
		tmark.NewHeader(2, tmark.T("Images")),
		tmark.NewImage(asset("mountains-rockies.webp")).WithCaption(tmark.NewCaption(tmark.T("Rocky Mountains"))),
		tmark.NewImage(asset("mountains-wide.webp")).WithCaption(tmark.NewCaption(tmark.T("Hidden mountain panorama"))).WithSpoiler(),
		tmark.NewCollage(
			tmark.NewImage(asset("mountains-summit.webp")),
			tmark.NewImage(asset("mountains-teton.webp")),
			tmark.NewImage(asset("mountains-portrait.webp")),
			tmark.NewImage(asset("mountains-lake.webp")),
		).WithCaption(tmark.NewCaption(tmark.T("Image collage"))),
		tmark.NewSlideshow(
			tmark.NewImage(asset("mountains-hero.webp")),
			tmark.NewVideo(videoURL).WithSpoiler(),
			tmark.NewImage(asset("mountains-valley.webp")),
		).WithCaption(tmark.NewCaption(tmark.T("Image and video slideshow"))),
		tmark.NewAudio(audioURL).WithCaption(tmark.NewCaption(tmark.T("Ambient mountain audio"))),
	)
}

func tables() tmark.RichBlocks {
	return tmark.NewRichBlocks(
		tmark.NewHeader(2, tmark.T("Sales by quarter")),
		tmark.NewTable(
			tmark.NewTableRow(tmark.NewCell(tmark.T("Quarter")).Header(), tmark.NewCell(tmark.T("Revenue")).Header().WithAlign(tmark.TableCellAlignRight), tmark.NewCell(tmark.T("Status")).Header()),
			tmark.NewTableRow(tmark.NewCell(tmark.T("Q1")), tmark.NewCell(tmark.T("$12,400")).WithAlign(tmark.TableCellAlignRight), tmark.NewCell(tmark.T("closed"))),
			tmark.NewTableRow(tmark.NewCell(tmark.T("Q2")), tmark.NewCell(tmark.T("$11,400")).WithAlign(tmark.TableCellAlignRight), tmark.NewCell(tmark.T("closed"))),
			tmark.NewTableRow(tmark.NewCell(tmark.T("Q3")), tmark.NewCell(tmark.T("$22,000")).WithAlign(tmark.TableCellAlignRight), tmark.NewCell(tmark.T("closed"))),
			tmark.NewTableRow(tmark.NewCell(tmark.T("Q4")), tmark.NewCell(tmark.T("$15,100")).WithAlign(tmark.TableCellAlignRight), tmark.NewCell(tmark.T("forecast"))),
		).WithCaption(tmark.T("Quarterly revenue (striped rows)")).WithStriped(),
		tmark.NewTable(
			tmark.NewTableRow(tmark.NewCell(tmark.T("Merged header")).Header().Col(3)),
			tmark.NewTableRow(tmark.NewCell(tmark.T("Left")).WithValign(tmark.TableCellValignTop), tmark.NewCell(tmark.T("Center")).WithAlign(tmark.TableCellAlignCenter), tmark.NewCell(tmark.T("Right")).WithAlign(tmark.TableCellAlignRight)),
			tmark.NewTableRow(tmark.NewCell(tmark.T("Tall")).Row(2), tmark.NewCell(tmark.T("A")), tmark.NewCell(tmark.T("B"))),
			tmark.NewTableRow(tmark.NewCell(tmark.T("C")), tmark.NewCell(tmark.T("D"))),
		).WithCaption(tmark.T("Merged cells and text alignment")).WithBordered(),
	)
}

func listsAndTasks() tmark.RichBlocks {
	return tmark.NewRichBlocks(
		tmark.NewHeader(2, tmark.T("Bullets and nested lists")),
		tmark.NewList(
			tmark.NewListItem(tmark.P(tmark.T("A plain bullet"))),
			tmark.NewListItem(
				tmark.P(tmark.T("A bullet with a nested list")),
				tmark.NewList(
					tmark.NewListItem(tmark.P(tmark.T("First nested bullet"))),
					tmark.NewListItem(tmark.P(tmark.T("Second nested bullet"))),
				),
			),
		),
		tmark.NewHeader(3, tmark.T("Numbered and lettered items")),
		tmark.NewList(
			tmark.NewListItem(tmark.P(tmark.T("Starts at ten"))).WithOrder(10),
			tmark.NewListItem(tmark.P(tmark.T("Continues at eleven"))).WithOrder(11),
		),
		tmark.NewList(
			tmark.NewListItem(tmark.P(tmark.T("Lowercase letters"))).WithType("a").WithOrder(1),
			tmark.NewListItem(tmark.P(tmark.T("The next lowercase letter"))).WithType("a").WithOrder(2),
			tmark.NewListItem(tmark.P(tmark.T("Uppercase letters"))).WithType("A").WithOrder(3),
			tmark.NewListItem(tmark.P(tmark.T("Letters continue past Z"))).WithType("A").WithOrder(27),
		),
		tmark.NewList(
			tmark.NewListItem(tmark.P(tmark.T("Lowercase Roman numerals"))).WithType("i").WithOrder(4),
			tmark.NewListItem(tmark.P(tmark.T("Uppercase Roman numerals"))).WithType("I").WithOrder(9),
			tmark.NewListItem(tmark.P(tmark.T("Explicit decimal numbering"))).WithType("1").WithOrder(12),
		),
		tmark.NewHeader(3, tmark.T("Tasks")),
		tmark.NewList(
			tmark.NewListItem(tmark.P(tmark.T("An unfinished task"))).WithType("checkbox").WithChecked(false),
			tmark.NewListItem(tmark.P(tmark.T("A finished task"))).WithType("checkbox").WithChecked(true),
		),
	)
}

func quotesAndDetails() tmark.RichBlocks {
	return tmark.NewRichBlocks(
		tmark.NewHeader(2, tmark.T("Quoted content")),
		tmark.NewBlockquote(
			tmark.P(tmark.T("A blockquote can contain several blocks.")),
			tmark.NewList(
				tmark.NewListItem(tmark.P(tmark.T("A list item inside the quote"))),
			),
		).WithCredit(tmark.T("Example source")),
		tmark.NewPullQuote(tmark.T("A pull quote draws attention to one short thought.")).WithCredit(tmark.T("Editor")),
		tmark.NewDetails(
			tmark.P(tmark.T("This paragraph is hidden until you expand the section.")),
			tmark.NewImage(asset("mountains-valley.webp")),
		).WithSummary(tmark.B(tmark.T("Collapsed section"))),
		tmark.NewDetails(tmark.P(tmark.T("This paragraph is visible by default."))).WithSummary(tmark.T("Open section")).Open(),
	)
}

func mathAndCode() tmark.RichBlocks {
	return tmark.NewRichBlocks(
		tmark.NewHeader(2, tmark.T("Formulas and source code")),
		tmark.P(tmark.T("An inline formula: "), tmark.NewMath("\\sqrt{a^2+b^2}"), tmark.T(". Inline code: "), tmark.NewCode("const ok = true"), tmark.T(".")),
		tmark.NewMathBlock("\\sum_{i=1}^{n} i = \\frac{n(n+1)}{2}"),
		tmark.NewPreformatted("type Example struct {\n\tName string `json:\"name\"`\n}\n\nfunc (e Example) Empty() bool {\n\treturn e.Name == \"\"\n}").WithLanguage("go"),
		tmark.NewPreformatted("line with #hash\nline with {braces}\nline with backslash \\\\"),
	)
}

func referencesAndAnchors() tmark.RichBlocks {
	return tmark.NewRichBlocks(
		tmark.NewHeader(2, tmark.T("Links, references, and anchors")),
		tmark.P(tmark.T("Lets start with simple external link. It should opens a browser: "), tmark.NewLink("https://example.com/spec", tmark.T("External spec"))),
		tmark.NewAnchor("intro"),
		tmark.P(tmark.T(`For example we have some long text:
Bitcoin (abbreviation: BTC; sign: ₿) is the first decentralized cryptocurrency. 
Based on a free-market ideology, bitcoin was invented in 2008 when an unknown person published a white paper under the pseudonym of Satoshi Nakamoto. Use of bitcoin as a currency began in 2009`),
			tmark.NewReferenceLink("first-usage", tmark.T("[1]")), tmark.T(` with the release of its open-source implementation. 
From 2021 to 2025, El Salvador adopted it as legal tender currency.`),
			tmark.NewReferenceLink("el-salvador", tmark.T("[2]")), tmark.T(` As bitcoin is pseudonymous, its use by criminals has attracted the attention of regulators, leading to its ban by several countries.`),
		),
		tmark.NewImage(asset("bit1.png")).WithCaption(tmark.NewCaption(tmark.T("Bitcoin origins and peer-to-peer adoption"))),
		tmark.P(tmark.T("Bitcoin works through the collaboration of computers, each of which acts as a node in the peer-to-peer bitcoin network. Each node maintains an independent copy of a public distributed ledger of transactions, called a blockchain, without central oversight. Transactions are validated through the use of cryptography, preventing one person from spending another person's bitcoin, as long as the owner of the bitcoin keeps certain sensitive data secret. "),
			tmark.NewReferenceLink("mastering-bitcoin", tmark.T("[3]"))),
		tmark.NewImage(asset("bit2.png")).WithCaption(tmark.NewCaption(tmark.T("Blockchain nodes validating transactions"))),
		tmark.P(tmark.T("Consensus between nodes about the content of the blockchain is achieved using a computationally intensive process based on proof of work, called mining, which is performed by purpose-built computers.")),
		tmark.NewDivider(),
		tmark.NewHeader(4, tmark.T("References")),
		tmark.P(tmark.NewReference("first-usage", tmark.T("[1] ")), tmark.T("Davis, Joshua (10 October 2011). \"The Crypto-Currency: Bitcoin and its mysterious inventor\". "), tmark.NewLink("https://web.archive.org/web/20141101014157/http://www.newyorker.com/magazine/2011/10/10/the-crypto-currency", tmark.T("The New Yorker.")), tmark.I(tmark.T(" Archived from the original on 1 November 2014."))),
		tmark.P(tmark.NewReference("el-salvador", tmark.T("[2] ")), tmark.T("\"El Salvador's dangerous gamble on bitcoin\". "), tmark.NewLink("https://web.archive.org/web/20141101014157/http://www.newyorker.com/magazine/2011/10/10/the-crypto-currency", tmark.T("Financial Times.")), tmark.I(tmark.T(" Archived from the original on 7 September 2021."))),
		tmark.P(tmark.NewReference("mastering-bitcoin", tmark.T("[3] ")), tmark.T("Mastering Bitcoin: Unlocking Digital Crypto-Currencies "), tmark.NewLink("https://www.oreilly.com/library/view/mastering-bitcoin/9781491902639/", tmark.T("O'Reilly Media."))),
	)
}

func mapsAndEvents() tmark.RichBlocks {
	return tmark.NewRichBlocks(
		tmark.NewHeader(2, tmark.T("Apollo 11 launch")),
		tmark.P(tmark.T("Apollo 11 launched on "), tmark.NewDateTime(-14552880, "America/New_York"), tmark.T(" from Launch Pad 39A at Kennedy Space Center, Florida (9:32 a.m. EDT).")),
		tmark.NewMap(28.608402, -80.604201).WithZoom(14).WithCaption(tmark.NewCaption(tmark.T("Apollo 11 launch site: Pad 39A"))),
		tmark.NewMap(-33.86882, 151.20929).WithCaption(tmark.NewCaption(tmark.B(tmark.T("Sydney")), tmark.T(" without explicit zoom"))),
	)
}

func edgeEmpties() tmark.RichBlocks {
	return tmark.NewRichBlocks(
		tmark.NewHeader(4, tmark.T("Empty paragraph")),
		tmark.P(),
		tmark.NewHeader(4, tmark.T("Empty list")),
		tmark.NewList(),
		tmark.NewHeader(4, tmark.T("Empty list item")),
		tmark.NewList(tmark.NewListItem()),
		tmark.NewHeader(4, tmark.T("Empty table")),
		tmark.NewTable(),
		tmark.NewHeader(4, tmark.T("Empty table row")),
		tmark.NewTable(tmark.NewTableRow()),
		tmark.NewHeader(4, tmark.T("Empty blockquote")),
		tmark.NewBlockquote(),
		tmark.NewHeader(4, tmark.T("Empty pull quote")),
		tmark.NewPullQuote(),
		tmark.NewHeader(4, tmark.T("Empty details")),
		tmark.NewDetails().WithSummary(tmark.T("Empty details")).Open(),
		tmark.NewHeader(4, tmark.T("Empty collage")),
		tmark.NewCollage(),
		tmark.NewHeader(4, tmark.T("Empty slideshow")),
		tmark.NewSlideshow(),
		tmark.NewHeader(4, tmark.T("Empty code block")),
		tmark.NewPreformatted(""),
		tmark.NewHeader(4, tmark.T("Empty math block")),
		tmark.NewMathBlock(""),
		tmark.NewHeader(4, tmark.T("Contentless divider")),
		tmark.NewDivider(),
	)
}

func edgeCases() tmark.RichBlocks {
	nested := tmark.NewDetails(
		tmark.P(tmark.T("Detail content 12")),
		tmark.P(tmark.T("This paragraph is inside twelve nested details.")),
	).WithSummary(tmark.T("Level 12"))
	for level := 11; level >= 1; level-- {
		label := tmark.T(fmt.Sprintf("Level %d", level))
		nested = tmark.NewDetails(tmark.P(tmark.T(fmt.Sprintf("Detail content %d", level))), nested).WithSummary(label)
	}
	nestedQuotes := tmark.NewBlockquote(
		tmark.P(tmark.T("Level 12")),
		tmark.P(tmark.T("This paragraph is inside twelve nested blockquotes.")),
	)
	for level := 11; level >= 1; level-- {
		nestedQuotes = tmark.NewBlockquote(tmark.P(tmark.T(fmt.Sprintf("Level %d", level))), nestedQuotes)
	}

	return tmark.NewRichBlocks(
		tmark.NewHeader(4, tmark.T("Reserved characters")),
		tmark.P(tmark.T("Reserved characters: # { } \\ should be escaped.")),
		tmark.NewHeader(4, tmark.T("Semicolon and line break")),
		tmark.P(tmark.T("A semicolon (;) stays literal; line breaks\nstay in the text.")),
		tmark.NewHeader(4, tmark.T("Special characters in a URL")),
		tmark.P(tmark.NewLink("https://example.com/search?q={tmark}#top", tmark.T("URL with braces and hash"))),
		tmark.NewHeader(4, tmark.T("Raw markup in a code block")),
		tmark.NewPreformatted("{document;\n#url{u}\n#title{t}\n{p;raw-ish # text}\n}").WithLanguage("tmark"),
		tmark.NewHeader(4, tmark.T("Combined formatting")),
		tmark.P(tmark.B(tmark.I(tmark.T("This text is bold and italic.")))),
		tmark.P(tmark.NewStrikethrough(tmark.B(tmark.T("This text is struck through and bold.")))),
		tmark.P(tmark.NewStrikethrough(tmark.NewMarked(tmark.I(tmark.T("This text is struck through, highlighted, and italic."))))),
		tmark.P(tmark.NewUnderline(tmark.B(tmark.T("This text is underlined and bold.")))),
		tmark.P(tmark.NewMarked(tmark.NewUnderline(tmark.T("This text is highlighted and underlined.")))),
		tmark.P(tmark.B(tmark.I(tmark.NewUnderline(tmark.T("This text is bold, italic, and underlined."))))),
		tmark.P(tmark.NewSpoiler(tmark.B(tmark.T("This spoiler contains bold text.")))),
		tmark.P(tmark.NewLink("https://example.com", tmark.I(tmark.NewUnderline(tmark.T("This link is italic and underlined."))))),
		tmark.P(tmark.T("A sentence with "), tmark.B(tmark.NewMarked(tmark.T("bold highlighting"))), tmark.T(" and "), tmark.I(tmark.NewStrikethrough(tmark.T("italic strikethrough"))), tmark.T(".")),
		tmark.P(tmark.T("Bold inline code: "), tmark.B(tmark.NewCode("const value = 1"))),
		tmark.NewHeader(4, tmark.T("Long unbroken text")),
		tmark.P(tmark.T(strings.Repeat("long-word-", 20)), tmark.NewCode(tmark.Text(strings.Repeat("x", 80)))),
		tmark.NewHeader(4, tmark.T("Unknown block")),
		tmark.P(tmark.T("Paragraph before an unknown block.")),
		tmark.Unknown{Raw: []byte("{colorpicker;#from{ad34f1}#to{000000}My palette}")},
		tmark.P(tmark.T("Paragraph after an unknown block.")),
		tmark.NewHeader(4, tmark.T("Small image")),
		tmark.NewImage(asset("c-logo.png")).WithCaption(tmark.NewCaption(tmark.T("C language logo"))),
		tmark.NewHeader(4, tmark.T("Twelve nested details")),
		nested,
		tmark.NewHeader(4, tmark.T("Twelve nested blockquotes")),
		nestedQuotes,
	)
}
