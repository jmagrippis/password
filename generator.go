package generator

type Generator struct {
	dictionary *Dictionary
	seed int64
}

type Dictionary struct {
	Adjectives []string
	Subjects []string
	Verbs []string
	Adverbs []string
	Objects []string
}

func (g Generator) generate() string {
	return "hello, world!"
}

func NewGenerator(dictionary *Dictionary, seed int64) *Generator {
	return &Generator{dictionary: dictionary, seed: seed}
}