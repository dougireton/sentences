package sentences

import (
	prose "gopkg.in/jdkato/prose.v2"
)

// ParseText parses one or more paragraphs of text
// and returns a slice of sentences.
func ParseText(s string) []string {

	var sentences []string

	doc, _ := prose.NewDocument(s)
	for _, sent := range doc.Sentences() {
		sentences = append(sentences, sent.Text)
	}
	return sentences
}
