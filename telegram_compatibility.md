# Telegram Rich API compatibility

Source: <https://core.telegram.org/bots/api>, Bot API 10.1 rich messages.

| Telegram api type | tmark type                                                                                              |
| --- |----------------------------------------------------------------------------------------------------------|
| **RichTextBold** | **Bold**                                                                                                 |
| **RichTextItalic** | **Italic**                                                                                               |
| **RichTextUnderline** | **Underline**                                                                                            |
| **RichTextStrikethrough** | **Strikethrough**                                                                                        |
| **RichTextSpoiler** | **Spoiler**                                                                                              |
| RichTextDateTime | Partially supported as `DateTime`: tmark stores UNIX seconds and timezone, but not custom display text. |
| RichTextTextMention | Telegram-specific. Use `Link` for external usage.                                                        |
| **RichTextSubscript** | **Subscript**                                                                                            |
| **RichTextSuperscript** | **Superscript**                                                                                          |
| **RichTextMarked** | **Marked**                                                                                               |
| **RichTextCode** | **Code**                                                                                                 |
| **RichTextCustomEmoji** | **Icon**                                                                                                 |
| **RichTextMathematicalExpression** | **Math**                                                                                                 |
| **RichTextUrl** | **Link**                                                                                                 |
| RichTextEmailAddress | Not a separate tmark type. Store as `Text`; renderers may auto-detect/highlight.                        |
| RichTextPhoneNumber | Not a separate tmark type. Store as `Text`; renderers may auto-detect/highlight.                        |
| RichTextBankCardNumber | Not a separate tmark type. Store as `Text`; renderers may auto-detect/highlight.                        |
| RichTextMention | Telegram-specific. Use `Link` for external usage, or store as `Text` with optional auto-detection.       |
| RichTextHashtag | Not a separate tmark type. Store as `Text`; renderers may auto-detect/highlight.                        |
| RichTextCashtag | Not a separate tmark type. Store as `Text`; renderers may auto-detect/highlight.                        |
| RichTextBotCommand | Not a separate tmark type. Store as `Text`; renderers may auto-detect/highlight.                        |
| **RichTextAnchor** | **Anchor**                                                                                               |
| **RichTextAnchorLink** | **AnchorLink**                                                                                           |
| **RichTextReference** | **Reference**                                                                                            |
| **RichTextReferenceLink** | **ReferenceLink**                                                                                        |
| **RichBlockParagraph** | **Paragraph**                                                                                            |
| **RichBlockSectionHeading** | **Header**                                                                                               |
| **RichBlockPreformatted** | **Preformatted**                                                                                         |
| **RichBlockFooter** | **Paragraph**                                                                                            |
| **RichBlockDivider** | **Divider**                                                                                              |
| **RichBlockMathematicalExpression** | **MathBlock**                                                                                            |
| **RichBlockAnchor** | **Anchor**                                                                                               |
| **RichBlockList** | **List**                                                                                                 |
| **RichBlockListItem** | **ListItem**                                                                                             |
| **RichBlockBlockQuotation** | **Blockquote**                                                                                           |
| **RichBlockPullQuotation** | **PullQuote**                                                                                            |
| **RichBlockCollage** | **Collage**                                                                                              |
| **RichBlockSlideshow** | **Slideshow**                                                                                            |
| **RichBlockTable** | **Table**                                                                                                |
| **RichBlockDetails** | **Details**                                                                                              |
| **RichBlockMap** | **Map**                                                                                                  |
| RichBlockAnimation | Stored as `Video`.                                                                                       |
| **RichBlockAudio** | **Audio**                                                                                                |
| **RichBlockPhoto** | **Image**                                                                                                |
| **RichBlockVideo** | **Video**                                                                                                |
| RichBlockVoiceNote | Stored as `Audio`.                                                                                       |
| RichBlockThinking | Not supported: ephemeral Telegram-only block.                                                            |
