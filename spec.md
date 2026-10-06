# tmark serialization specification

## Scalar values

| Value | Serialization |
|---|---|
| Text | UTF-8 without a BOM; escape `#`, `{`, `}`, and `\` with `\` |
| Number | Canonical decimal notation as defined below |
| Boolean | `t` or `f` |

`Number` represents a finite decimal number and imposes no range or precision
limit at the serialization layer. Its integer part is either `0` or a nonzero
digit followed by zero or more digits. An optional fractional part consists of
`.` followed by one or more digits and MUST end in a nonzero digit. A negative
value starts with `-`; zero is always serialized as `0`. A field definition may
impose additional constraints appropriate to its semantics. Exponential
notation, `NaN`, and infinities are not supported by any type at the current
compatibility level.

## Type system

Every tmark building block, or _type_, defines:

- a name used in libraries and documentation;
- for non-scalar types, an encoding tag used to identify the type in a
  serialized document;
- zero or more named, statically typed fields;
- at most one unnamed statically typed field;
- a _display type_;
- a _line mode_.

These are properties of the type, not of an individual serialized value.

Every field has a statically defined type. An array field MUST define an element
type. Its element type may be a concrete tmark type or one of the abstract
`inline` and `block` types. A field may further restrict which concrete types
are allowed in an abstract array. Every element MUST match its declared element
type. In particular, an array MUST NOT mix inline and block nodes.

Scalar types are serialized without tags. Non-scalar type tags and field names
MUST contain between 1 and 32 characters from
`abcdefghijklmnopqrstuvwxyz0123456789-_`.

The display type defines how a value is **rendered**: either inline, as part of a
line of text, or as a block following the preceding block.

| Display type | Description |
|---|---|
| `inline` | Rendered within a line of text |
| `block` | Rendered as a block following the preceding block |
| `none` | Neither inline nor block; used by special types such as the `Document` root |

The line mode controls **serialization** whitespace to improve readability and
produce cleaner diffs. It does not affect the document model.

| Line mode | Description |
|---|---|
| `none` | Serialize fields without additional line breaks; this is the only mode available to inline types |
| `always` | Start every field and the closing brace on a new line |
| `multiple` | Behave like `always` when the unnamed array contains more than one element; otherwise behave like `none` |

## Encoding rules

The general serialized form of a tmark value is:

```text
{tag-name;#named-field-1{field value}#named-field-2{field value}unnamed-field-value}
```

Every tagged value begins with `{` and ends with `}`. Its tag MUST immediately
follow the opening brace and MUST end with `;`.

Serializers MUST use LF (`U+000A`) for line breaks. Parsers MUST also accept
CRLF (`U+000D U+000A`) and treat it as LF. A standalone CR is invalid.

A named field consists of `#`, the field name, and the serialized field value
enclosed in braces: `#name{value}`.

Field names MUST be unique within a value. A parser MUST reject a value that
contains duplicate field names. For example, the following value is invalid:

```text
{a;#href{https://example.com/}#href{https://example.com/}Some page}
```

The unnamed field, if present, MUST follow all named fields. It has no marker or
wrapper of its own; its contents are serialized directly.

A value whose concrete type is defined by its containing field MUST omit its
type tag. The field's braces delimit the value. For example:

```text
{audio;#caption{#credit{Waus}My best track}https://tmark.waus.app/static/mountain-ambient.mp3}
```

not:

```text
{audio;#caption{caption;#credit{Waus}My best track}https://tmark.waus.app/static/mountain-ambient.mp3}
```

Each element of an array with a concrete element type MUST also omit its type
tag. The element retains its enclosing braces, which delimit it from adjacent
elements. For example, the elements of a `ListItem` array are serialized as:

```text
{list;
{#type{checkbox}#checked{t}{p;Typed nodes}}
{#type{checkbox}#checked{f}{p;More nodes}}
}
```

An array whose element type is the abstract `inline` or `block` type determines
each element's concrete type at runtime. Every non-scalar element in such an
array MUST include its tag. Scalar values remain untagged in every context. For
example, the `Document` block array retains both tags:

```text
{document;
{p;A paragraph}
{list;...}
}
```

Examples of each line mode follow.

`none`:

```text
{p;{b;Some bold} text}
```

Line breaks inside field values are preserved and do not change the line mode:

```text
{h;#s{1}Even with a line break
inside}
```

`always`:

```text
{slideshow;
#caption{Every field starts on a new line. The closing \} does too.}
{img;https://tmark.waus.app/static/mountains-wide1.jpg}
{img;https://tmark.waus.app/static/mountains-wide2.jpg}
{img;https://tmark.waus.app/static/mountains-wide3.jpg}
}
```

`multiple`, with one element in the unnamed array:

```text
{q;{p;A quote containing one paragraph}}
```

With more than one value, it behaves like `always`:

```text
{q;
{p;The first paragraph of the quote}
{p;The second paragraph of the quote}
}
```

The root node has depth 0. Every child node has a depth one greater than its
parent. Nodes at depth 16 or greater are invalid; the maximum valid depth is 15.

### Array serialization

Array elements are serialized consecutively. Elements with a concrete element
type are delimited by braces. Elements of an abstract `inline` or `block` array
use their normal tagged serialization. When an array is the unnamed field of a
type whose line mode resolves to `always`, each element starts on a new line.

Boundaries between adjacent `Text` elements are not preserved. They are
equivalent to a single `Text` element containing their concatenated values.

### Bidirectional text

tmark serializes bidirectional text in Unicode logical order. Bidirectional
control characters are valid `Text` content and MUST be preserved verbatim.
Parsing does not perform directional processing. Renderers are responsible for
applying the [Unicode Bidirectional Algorithm](https://www.unicode.org/reports/tr9/).

## Compatibility level 1 types

The following table lists serialized field names and their types. `inline[]`
and `block[]` are arrays of the corresponding abstract type. A named field is
written as `#name{value}`; the unnamed field has no field name in the document.

| tmark type | Display type | Tag | Line mode | Named fields | Unnamed field | Special conditions |
|---|---|---|---|---|---|---|
| `Text` | inline | none | none | none | none | none |
| `Caption` | none | `caption` | none | `credit`: `inline[]` | `inline[]` | none |
| `Link` | inline | `a` | none | `href`: `Text` | `inline[]` | none |
| `AnchorLink` | inline | `anchor-link` | none | `name`: `Text` | `inline[]` | none |
| `Reference` | inline | `ref` | none | `name`: `Text` | `inline[]` | none |
| `ReferenceLink` | inline | `ref-link` | none | `name`: `Text` | `inline[]` | none |
| `Bold` | inline | `b` | none | none | `inline[]` | none |
| `Italic` | inline | `i` | none | none | `inline[]` | none |
| `Marked` | inline | `m` | none | none | `inline[]` | none |
| `Underline` | inline | `u` | none | none | `inline[]` | none |
| `Strikethrough` | inline | `s` | none | none | `inline[]` | none |
| `Spoiler` | inline | `spoiler` | none | none | `inline[]` | none |
| `Subscript` | inline | `sub` | none | none | `inline[]` | none |
| `Superscript` | inline | `sup` | none | none | `inline[]` | none |
| `DateTime` | inline | `datetime` | none | `timezone`: `Text` | `Number` | none |
| `Code` | inline | `code` | none | none | `Text` | none |
| `Math` | inline | `math` | none | none | `Text` | none |
| `Icon` | inline | `icon` | none | `alt`: `Text` | `Text` | none |
| `Paragraph` | block | `p` | none | none | `inline[]` | none |
| `Header` | block | `h` | none | `s`: `Number` | `inline[]` | `s`: integer from 1 to 6 |
| `Preformatted` | block | `pre` | always | `language`: `Text` | `Text` | none |
| `MathBlock` | block | `math-block` | always | none | `Text` | none |
| `Anchor` | block | `anchor` | none | none | `Text` | none |
| `Divider` | block | `hr` | none | none | none | none |
| `Blockquote` | block | `q` | multiple | `credit`: `inline[]` | `block[]` | none |
| `PullQuote` | block | `as` | none | `credit`: `inline[]` | `inline[]` | none |
| `Collage` | block | `collage` | always | `caption`: `Caption` | `block[]` | Children: `Image` or `Video` only |
| `List` | block | `list` | always | none | `ListItem[]` | none |
| `ListItem` | block | `li` | multiple | `type`: `Text`; `order`: `Number (integer)`; `checked`: `Boolean` | `block[]` | `type`: `a`, `A`, `i`, `I`, or `1` for ordered lists; `checkbox` for checklists |
| `Map` | block | `map` | none | `lat`: `Number`; `lon`: `Number`; `zoom`: `Number (integer)`; `caption`: `Caption` | none | none |
| `Image` | block | `img` | none | `caption`: `Caption`; `has_spoiler`: `Boolean` | `Text` | none |
| `Video` | block | `video` | none | `caption`: `Caption`; `has_spoiler`: `Boolean`; `preview`: `Text`; `loop`: `Boolean` | `Text` | `preview` is required |
| `Audio` | block | `audio` | none | `caption`: `Caption` | `Text` | none |
| `Slideshow` | block | `slideshow` | always | `caption`: `Caption` | `block[]` | Children: `Image` or `Video` only |
| `Table` | block | `table` | always | `caption`: `inline[]`; `bordered`: `Boolean`; `striped`: `Boolean` | `TableRow[]` | none |
| `TableRow` | block | `tr` | none | none | `Cell[]` | none |
| `Cell` | block | `td` | none | `header`: `Boolean`; `colspan`: `Number (integer)`; `rowspan`: `Number (integer)`; `align`: `Text`; `valign`: `Text` | `inline[]` | `align`: `left`, `center`, or `right`; `valign`: `top`, `middle`, or `bottom` |
| `Details` | block | `details` | multiple | `summary`: `inline[]`; `open`: `Boolean` | `block[]` | none |
| `AttachedMedia` | none | `attached` | always | `hash`: `Text` | `block[]` | none |
| `Document` | none | `document` | always | `url`: `Text`; `title`: `Text`; `description`: `Text`; `author_name`: `Text`; `author_url`: `Text`; `image_url`: `Text`; `attached_media`: `AttachedMedia[]` | `block[]` | none |

## Media formats

At compatibility level 1, renderers MUST support the following media formats
for `Image` and `Audio` values, including those nested in other types.
Support means decoding and displaying images, playing Lottie animations, and
playing audio. 

### Images — MUST supported

| Format | Representation | Example extension |
|---|---|---|
| JPEG | JPEG image | `.jpg`, `.jpeg` |
| WebP | WebP image | `.webp` |
| AVIF | AVIF image | `.avif` |
| PNG | PNG image | `.png` |
| Lottie | Self-contained Lottie JSON animation | `.json` |

### Audio — MUST supported

| Codec | Required container | Profile | Example extension |
|---|---|---|---|
| MP3 | Native MP3 bitstream | — | `.mp3` |
| FLAC | Native FLAC container | — | `.flac` |
| Opus | Ogg | — | `.ogg`, `.opus` |
| AAC | ISO BMFF / MP4 | AAC-LC | `.m4a`, `.mp4` |

A supported codec in another container does not satisfy these requirements.
In particular, AAC in ADTS, Opus in WebM, FLAC in Ogg, and AAC profiles other
than AAC-LC are outside the required set.
