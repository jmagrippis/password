package password

import (
	"math/rand"
	"strings"
)

const (
	defaultWordCount uint8  = 5
	defaultSeparator string = " "
)

type Generator struct {
	dictionary *Dictionary
	seed       int64
	wordCount  uint8
	separator  string
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
	selected[3] = g.dictionary.Adjectives[rand.Intn(len(g.dictionary.Adjectives))]
	selected[4] = g.dictionary.Objects[rand.Intn(len(g.dictionary.Objects))]

	return strings.Join(selected, g.separator)
}

// NewGenerator seeds the RNG and returns a password Generator with the given Dictionary
// and the default settings
func NewGenerator(dictionary *Dictionary, seed int64) *Generator {
	rand.Seed(seed)
	return &Generator{
		dictionary: dictionary,
		seed:       seed,
		wordCount:  defaultWordCount,
		separator:  defaultSeparator,
	}
}
