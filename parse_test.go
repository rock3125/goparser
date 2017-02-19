package spacy

import (
	"testing"
)

// test fn
func isTrue(t *testing.T, cond bool) {
	if !cond {
		t.Error("condition failed")
		panic("condition failed")
	}
}


// test we can parse
func TestParser1(t *testing.T) {

	parse_str := "Peter was here.  He then moved to London."
	sentence_list := ParseText("http://localhost:9000/parse", parse_str)
	isTrue(t, len(sentence_list) == 2)

	sentence_1 := sentence_list[0]
	isTrue(t, len(sentence_1) == 4)

	sentence_2 := sentence_list[1]
	isTrue(t, len(sentence_2) == 6)
}

