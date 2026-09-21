# tmark serialization spec

| tmark type | display | tag name | line mode |
|---|---|---|---|
| `Text` | inline | plain text | none |
| `Link` | inline | `a` | none |
| `AnchorLink` | inline | `anchor-link` | none |
| `Reference` | inline | `ref` | none |
| `ReferenceLink` | inline | `ref-link` | none |
| `Bold` | inline | `b` | none |
| `Italic` | inline | `i` | none |
| `Marked` | inline | `m` | none |
| `Underline` | inline | `u` | none |
| `Strikethrough` | inline | `s` | none |
| `Spoiler` | inline | `spoiler` | none |
| `Subscript` | inline | `sub` | none |
| `Superscript` | inline | `sup` | none |
| `DateTime` | inline | `datetime` | none |
| `Code` | inline | `code` | none |
| `Math` | inline | `math` | none |
| `Icon` | inline | `icon` | none |
| `Paragraph` | block | `p` | none |
| `Header` | block | `h` | none |
| `Preformatted` | block | `pre` | always |
| `MathBlock` | block | `math-block` | always |
| `Anchor` | block | `anchor` | none |
| `Divider` | block | `hr` | none |
| `Blockquote` | block | `q` | multiple |
| `PullQuote` | block | `as` | none |
| `Collage` | block | `collage` | always |
| `List` | block | `list` | always |
| `ListItem` | block | `li` | multiple |
| `Map` | block | `map` | none |
| `Image` | block | `img` | none |
| `Video` | block | `video` | none |
| `Audio` | block | `audio` | none |
| `Slideshow` | block | `slideshow` | always |
| `Table` | block | `table` | always |
| `Cell` | block | `td` | none |
| `Details` | block | `details` | multiple |
| `AttachedMedia` | document field | `attached` | always |
| `Document` | root | `document` | always |

## Rules

Tag name and line mode are type-level properties.

| rule | meaning |
|---|---|
| tag name | serialize after `{` and terminate with `;` |
| named field | serialize as `#name{value}` |
| unnamed field | serialize as the node primary value |
| struct in a named field | serialize only its fields, without an extra node tag |
| line mode `none` | serialize fields without extra newlines |
| line mode `always` | serialize every field on a separate line |
| line mode `multiple` | serialize every field on a separate line only when the primary sequence has more than one value |

Serialized values must not exceed 16 nested tmark nodes.

## Scalar values

| value | serialization |
|---|---|
| text | UTF-8 without BOM; escape `#`, `{`, `}`, and `\` with `\` |
| integer | decimal integer |
| float | decimal float without exponential notation |
| boolean | `t` or `f` |
