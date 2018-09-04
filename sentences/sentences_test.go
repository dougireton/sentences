package sentences

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseText(t *testing.T) {
	f, err := ioutil.ReadFile("testdata/basic.txt")

	if err != nil {
		t.Fatal(err)
	}

	got := ParseText(string(f))
	want := []string{
		"This is a sentence with a period.",
		"Is it followed by a sentence with a question mark?",
		"Yes it is!",
		"Now we have a sentence with two spaces before it.",
		"New paragraph with \"quotes\".",
	}

	assert.Equal(t, got, want)
}
