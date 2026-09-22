# tmark

`tmark` is a document markup language inspired by Telegram Rich Messages and Instant View.

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
- Extensibility should be supported but constrained enough to keep documents
  portable.
- Parsers and serializers should be as strict as possible to avoid ambiguity.

## Document Model

A `tmark` document is a tree of nodes.

A node can be:

- text
- a semantic element with typed children and attributes

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
- date/time, anchors, references, inline icons.

## Comparison with HTML

### Inline types

| tmark type     | HTML                |
|----------------|---------------------|
| `Text`         | text node           |
| `Link`         | `<a>`               |
| `AnchorLink`   | `<a href="#...">`   |
| `Reference`    | reference text      |
| `ReferenceLink` | reference link      |
| `Bold`         | `<b>`               |
| `Italic`       | `<i>`               |
| `Marked`       | `<mark>`            |
| `Underline`    | `<u>`               |
| `Strikethrough` | `<s>`               |
| `Spoiler`      | CSS: foreground color matches background |
| `Subscript`    | `<sub>`             |
| `Superscript`  | `<sup>`             |
| `DateTime`     | `<time>`            |
| `Code`         | `<code>`            |
| `Math`         | `<math>`            |
| `Icon`         | custom icon/emoji   |

### Block types

| tmark type     | HTML                |
|----------------|---------------------|
| `Paragraph`    | `<p>`               |
| `Header`       | `<h1>`–`<h6>`       |
| `Preformatted` | `<pre>`             |
| `MathBlock`    | display math        |
| `Anchor`       | named anchor        |
| `Divider`      | `<hr>`              |
| `Blockquote`   | `<blockquote>`      |
| `PullQuote`    | `<aside>`           |
| `Collage`      | media collage       |
| `List`         | `<ol/ul>`           |
| `ListItem`     | `<li>`              |
| `Map`          | map embed           |
| `Image`        | image               |
| `Video`        | `<video>`           |
| `Audio`        | `<audio>`           |
| `Slideshow`    | media slideshow     |
| `Table`        | `<table>`           |
| `Cell`         | `<td>` / `<th>`     |
| `Details`      | `<details>`         |

### Document-level types

- `Document` — document root
- `AttachedMedia` — embedded media data

## Extensibility and backward compatibility

The format defines a broad set of types intended to cover most use cases within
its target domain.

All standard types are defined in the [specification](spec.md). A new type that
is useful beyond a single application may be standardized in the next
compatibility level.

Applications may define custom types for internal use and support them through
renderer extensions. Custom type tags SHOULD use an `x-` or `appname-` prefix
to avoid collisions with current or future standard types.

To preserve compatibility, every implementation MUST provide a fallback for
unknown but syntactically valid types. An unknown node inherits the abstract
display type required by its containing array. An implementation MUST render an
unknown node in an inline array as `Code` and an unknown node in a block array
as `Preformatted`. It MUST preserve the node's raw content and display it
exactly as it appeared in the source document.

## Implementations
