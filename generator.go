package password

import (
	"errors"
	"math/rand"
	"strings"
)

const (
	minWordCount     uint8  = 4
	defaultWordCount uint8  = 5
	defaultDelimiter string = " "
)

type Generator struct {
	dictionary *Dictionary
	seed       int64
	wordCount  uint8
	delimiter  string
	prefix     string
}

type Dictionary struct {
	Adjectives []string
	Subjects   []string
	Verbs      []string
	Adverbs    []string
	Objects    []string
}

// Generates a secure password based on the given dictionary and settings
func (g Generator) generate() string {
	selected := make([]string, g.wordCount)

	selected[0] = g.dictionary.Adverbs[rand.Intn(len(g.dictionary.Adverbs))]
	selected[1] = g.dictionary.Subjects[rand.Intn(len(g.dictionary.Subjects))]
	selected[2] = g.dictionary.Verbs[rand.Intn(len(g.dictionary.Verbs))]
	var key uint8 = 3
	for ; key < g.wordCount-1; key++ {
		selected[key] = g.dictionary.Adjectives[rand.Intn(len(g.dictionary.Adjectives))]
	}
	selected[key] = g.dictionary.Objects[rand.Intn(len(g.dictionary.Objects))]

	result := strings.Join(selected, g.delimiter)

	if g.prefix != "" {
		result = strings.Join([]string{g.prefix, result}, "")
	}

	return result
}

// setWordCount sets the number of words returned by the generate function
func (g *Generator) setWordCount(wordCount uint8) error {
	if wordCount < minWordCount {
		return errors.New("Cannot return so few words!")
	}
	g.wordCount = wordCount
	return nil
}

// setDelimeter sets the delimiter used to break up the words of the sentence
func (g *Generator) setDelimiter(delimiter string) {
	g.delimiter = delimiter
}

// setPrefix sets the prefix used to start the sentence
func (g *Generator) setPrefix(prefix string) {
	g.prefix = prefix
}

// NewGenerator seeds the RNG and returns a password Generator with the given Dictionary
// and the default settings
func NewGenerator(dictionary *Dictionary, seed int64) *Generator {
	rand.Seed(seed)
	return &Generator{
		dictionary: dictionary,
		seed:       seed,
		wordCount:  defaultWordCount,
		delimiter:  defaultDelimiter,
	}
}
