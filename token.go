package spacy

import (
	"fmt"
	"strings"
)

type SpacyReturnData struct {
	NumSentences int        `json:"num_sentences"`
	Processing_time int64   `json:"processing_time"`
	NumTokens int           `json:"num_tokens"`
	SentenceList []Sentence `json:"sentence_list"`
}

type Token struct {
	Index int               `json:"index"`
	AncestorList []int      `json:"list"`
	Tag   string            `json:"tag"`
	Text  string            `json:"text"`
	Dep   string            `json:"dep"`
	SynId int               `json:"synid"`
}

type Sentence []Token

// len for sort interface
func (s Sentence) Len() int {
	return len(s)
}

// less for sorting, sort by token index
func (s Sentence) Less(i, j int) bool {
	return s[i].Index < s[j].Index;
}

// sort interface
func (s Sentence) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (t Token) ToString() (string) {
	return fmt.Sprintf("%#v", t)
}

func (s Sentence) ToString() (string) {
	str := ""
	for _, token := range s {
		str += token.Text
		if strings.Contains(token.Dep, "sub") || strings.Contains(token.Dep, "obj") {
			str += "{" + token.Dep + "} " // display grammar rule
		} else {
			str += " "
		}
	}
	return strings.TrimSpace(str)
}

