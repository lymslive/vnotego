package markdown

import (
	"github.com/gomarkdown/markdown"
)

func ByteToHTML(input []byte) []byte {
	output := markdown.ToHTML(input, nil, nil)
	return output
}
