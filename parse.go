package spacy

import (
	"encoding/json"
	"net/http"
	"bytes"
	"strconv"
	"log"
)

// convert a json string to a parser result map
func jsonToParseResult(jsonStr string) []Sentence {
	res := SpacyReturnData{} // json string to parser response
	json.Unmarshal([]byte(jsonStr), &res)
	return res.SentenceList
}

// remove spaces from the parse results
func removeSpaces(sentence_list []Sentence) []Sentence {
	final_sentence_list := make([]Sentence,0)
	for _, sentence := range sentence_list {
		new_sentence := make(Sentence,0)
		for _, token := range sentence {
			if token.Text != " " {
				new_sentence = append(new_sentence, token)
			}
		}
		if len(new_sentence) > 0 {
			final_sentence_list = append(final_sentence_list, new_sentence)
		}
	}
	return final_sentence_list
}

// post a request to the parser server (parsey)
func postRequest(url string, text string) ([]Sentence) {
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(text))
	req.Header.Set("Content-Type", "text/plain")
	cl := strconv.Itoa(len(text))
	req.Header.Set("Content-Length", cl)
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer response.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	return removeSpaces(jsonToParseResult(buf.String()))
}

// parse a piece of text and return the list of sentences parsed by spaCy
// serviceAddress:  the spacy endpoint, e.g. "http://localhost:9000/parse"
// text: the text to parse
func ParseText(serviceAddress string, text string) []Sentence {
	return postRequest(serviceAddress, text)
}

