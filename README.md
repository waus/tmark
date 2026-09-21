# tmark

`tmark` is a document markup language inspired by Telegram Markup.

The goal is a simple, portable format for exchanging structured documents where
the author describes **what the document contains**, not **how it must be
rendered pixel by pixel**.

## Motivation

Most popular markup formats mix content, structure, and presentation control.

HTML and DOCX expose too many ways to directly affect visual output. That makes
custom renderers, sanitizers, and converters complex and fragile.

Markdown looks simpler, but in practice it is a family of incompatible dialects.
It partially relies on HTML, has many extensions, and real parser behavior often
differs from the documentation.

AsciiDoc and TeX are powerful, but they bring large specialized stacks and fit a
narrower set of use cases.

`tmark` chooses a different trade-off: the format gives authors almost no
control over fonts, spacing, markers, grids, or other styling details. Instead,
it captures document semantics and leaves rendering to the target application
and its styles.

## Principles

- Semantics over presentation.
- Rendering belongs to the application, not to the document.
- The format should cover common document needs out of the box.
- Parsers and internal models should be simple to implement across languages.
- Extensibility should be limited so documents remain portable.
- Parsers and serializers should be as strict as possible to avoid ambiguity.



## Document Model

A `tmark` document is a tree of nodes.

A node can be:

- text;
- a semantic element with typed children and attributes.


## Format Surface

Block-level content:

- paragraphs and section headings;
- blockquotes, asides, expandable details;
- lists, tables, separators, anchors;
- images, videos, audio, slideshows;
- maps and attached media.

Inline content:

- links;
- bold, italic, marked, underline, strikethrough, spoiler;
- subscript, superscript, inline code, math;
- date/time, anchors, references, inline icon.

## Type Surface

| tmark type       | html                |
|------------------|---------------------|
| `Text`           | text node           |
| `Link`           | `<a>`               |
| `AnchorLink`     | `<a href="#...">`   |
| `Reference`      | reference text      |
| `ReferenceLink`  | reference link      |
| `Bold`           | `<b>`               |
| `Italic`         | `<i>`               |
| `Marked`         | `<mark>`            |
| `Underline`      | `<u>`               |
| `Strikethrough`  | `<s>`               |
| `Spoiler`        | css: color=bg-color |
| `Subscript`      | `<sub>`             |
| `Superscript`    | `<sup>`             |
| `DateTime`       | `<time>`            |
| `Code`           | `<code>`            |
| `Math`           | `<math>`            |
| `Icon`           | custom icon/emoji   |
| `Paragraph`      | `<p>`               |
| `Header`         | `<h1>`...`<h6>`     |
| `Preformatted`   | `<pre>`             |
| `MathBlock`      | display math        |
| `Anchor`         | named anchor        |
| `Divider`        | `<hr>`              |
| `Blockquote`     | `<blockquote>`      |
| `PullQuote`      | `<aside>`           |
| `Collage`        | media collage       |
| `List`           | `<ol/ul>`           |
| `ListItem`       | `<li>`              |
| `Map`            | map embed           |
| `Image`          | image               |
| `Video`          | `<video>`           |
| `Audio`          | `<audio>`           |
| `Slideshow`      | media slideshow     |
| `Table`          | `<table>`           |
| `Cell`           | `<td>` / `<th>`     |
| `Details`        | `<details>`         |
| `Document`       | document root       |
| `AttachedMedia`  | embedded media data |


## Implementations

