package password

import (
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestGenerator(t *testing.T) {
	var generator *Generator

	Convey("Given a new Generator with a set Dictionary and seed", t, func() {
		dictionary := &Dictionary{
			Adverbs:    []string{"cuddling", "slapping", "shouting", "jumping", "ducking", "mocking", "trotting", "galloping"},
			Subjects:   []string{"mermaids", "unicorns", "lions", "piranhas", "cuttlefish", "llamas", "dragons"},
			Verbs:      []string{"love", "fancy", "eat", "bring", "fear", "aggravate", "detest", "adore", "belittle", "ravish"},
			Adjectives: []string{"beautiful", "homely", "magical", "posh", "excellent", "portly", "lovely"},
			Objects:    []string{"teddy-bears", "diamonds", "buckets", "boxes", "dishes", "ornaments"},
		}
		generator = NewGenerator(dictionary, 2)

		Convey("Given I just run it with no extra settings", func() {

			Convey("It returns secure passwords", func() {
				So(generator.generate(), ShouldEqual, "shouting unicorns eat posh buckets")
				So(generator.generate(), ShouldEqual, "trotting lions fear portly teddy-bears")
				So(generator.generate(), ShouldEqual, "trotting mermaids aggravate portly buckets")
			})

			Convey("Consisting of 5 words each", func() {
				So(len(strings.Split(generator.generate(), " ")), ShouldEqual, 5)
				So(len(strings.Split(generator.generate(), " ")), ShouldEqual, 5)
				So(len(strings.Split(generator.generate(), " ")), ShouldEqual, 5)
			})
		})

		Convey("Given I change the number of returned words to x", func() {

			Convey("It returns a sentence of x words", func() {
				var x uint8 = 8
				generator.setWordCount(x)
				So(len(strings.Split(generator.generate(), " ")), ShouldEqual, x)
				x = 15
				generator.setWordCount(x)
				So(len(strings.Split(generator.generate(), " ")), ShouldEqual, x)
				x = 12
				generator.setWordCount(x)
				So(len(strings.Split(generator.generate(), " ")), ShouldEqual, x)
			})
		})

		Convey("When I try to change the number of returned words to below 4", func() {

			Convey("It returns an error", func() {
				So(generator.setWordCount(3), ShouldNotEqual, nil)
				So(generator.setWordCount(1), ShouldNotEqual, nil)
				So(generator.setWordCount(4), ShouldEqual, nil)
			})
		})

		Convey("Given I change the delimiter to x", func() {

			Convey("It returns the sentence delimited by x instead of \" \"", func() {
				var x string = "-"
				generator.setDelimiter(x)
				password := generator.generate()
				So(password, ShouldEqual, "shouting-unicorns-eat-posh-buckets")
				So(len(strings.Split(password, x)), ShouldEqual, 5)

				x = "/"
				generator.setDelimiter(x)
				So(len(strings.Split(generator.generate(), x)), ShouldEqual, 5)

				x = "1"
				generator.setDelimiter(x)
				So(len(strings.Split(generator.generate(), x)), ShouldEqual, 5)
			})

		})

		Convey("Given I change the prefix to x", func() {

			Convey("It prefaces the sentence prefixed by x", func() {
				var x string = "!"
				generator.setPrefix(x)
				password := generator.generate()
				So(password, ShouldEqual, "!shouting unicorns eat posh buckets")
				first, _ := utf8.DecodeRuneInString(password)
				So(string(first), ShouldEqual, x)

				x = "爱"
				generator.setPrefix(x)
				password = generator.generate()
				first, _ = utf8.DecodeRuneInString(password)
				So(string(first), ShouldEqual, x)

				x = "xoxo"
				prefixCharacterCount := utf8.RuneCountInString(x)
				generator.setPrefix(x)
				password = generator.generate()
				var startingCharacters = make([]string, prefixCharacterCount)
				for i, character := range password {
					startingCharacters[i] = string(character)
					if i == prefixCharacterCount-1 {
						break
					}
				}
				So(strings.Join(startingCharacters, ""), ShouldEqual, x)
			})

		})

		Convey("Given I change to suffix", func() {

			Convey("It ends the sentence with the given suffix", nil)

		})

		Convey("Given I reseed", func() {

			Convey("I get a different set of results", nil)

		})

	})

}
