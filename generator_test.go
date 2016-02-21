package password_test

import (
	"github.com/jmagrippis/password"
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestGenerator(t *testing.T) {
	var generator *password.Generator

	Convey("Given a new Generator with a set Dictionary and seed", t, func() {
		dictionary := &password.Dictionary{
			Adverbs:    []string{"cuddling", "slapping", "shouting", "jumping", "ducking", "mocking", "trotting", "galloping"},
			Subjects:   []string{"mermaids", "unicorns", "lions", "piranhas", "cuttlefish", "llamas", "dragons"},
			Verbs:      []string{"love", "fancy", "eat", "bring", "fear", "aggravate", "detest", "adore", "belittle", "ravish"},
			Adjectives: []string{"beautiful", "homely", "magical", "posh", "excellent", "portly", "lovely"},
			Objects:    []string{"teddy-bears", "diamonds", "buckets", "boxes", "dishes", "ornaments"},
		}
		generator = password.NewGenerator(dictionary, 2)

		Convey("Given I just run it with no extra settings", func() {

			Convey("It returns secure passwords", func() {
				So(generator.Generate(), ShouldEqual, "shouting unicorns eat posh buckets")
				So(generator.Generate(), ShouldEqual, "trotting lions fear portly teddy-bears")
				So(generator.Generate(), ShouldEqual, "trotting mermaids aggravate portly buckets")
			})

			Convey("Consisting of 5 words each", func() {
				So(len(strings.Split(generator.Generate(), " ")), ShouldEqual, 5)
				So(len(strings.Split(generator.Generate(), " ")), ShouldEqual, 5)
				So(len(strings.Split(generator.Generate(), " ")), ShouldEqual, 5)
			})
		})

		Convey("Given I change the number of returned words to x", func() {

			Convey("It returns a sentence of x words", func() {
				var x uint8 = 8
				generator.SetWordCount(x)
				So(len(strings.Split(generator.Generate(), " ")), ShouldEqual, x)
				x = 15
				generator.SetWordCount(x)
				So(len(strings.Split(generator.Generate(), " ")), ShouldEqual, x)
				x = 12
				generator.SetWordCount(x)
				So(len(strings.Split(generator.Generate(), " ")), ShouldEqual, x)
			})
		})

		Convey("When I try to change the number of returned words to below 4", func() {

			Convey("It returns an error", func() {
				So(generator.SetWordCount(3), ShouldNotEqual, nil)
				So(generator.SetWordCount(1), ShouldNotEqual, nil)
				So(generator.SetWordCount(4), ShouldEqual, nil)
			})
		})

		Convey("Given I set titleCase to true", func() {
			generator.SetTitleCase(true)
			Convey("It returns a sentence with each first letter capitalized", func() {
				So(generator.Generate(), ShouldEqual, "Shouting Unicorns Eat Posh Buckets")
				So(generator.Generate(), ShouldEqual, "Trotting Lions Fear Portly Teddy-Bears")
				So(generator.Generate(), ShouldEqual, "Trotting Mermaids Aggravate Portly Buckets")
			})
		})

		Convey("Given I change the delimiter to x", func() {

			Convey("It returns the sentence delimited by x instead of \" \"", func() {
				var x string = "-"
				generator.SetDelimiter(x)
				password := generator.Generate()
				So(password, ShouldEqual, "shouting-unicorns-eat-posh-buckets")
				So(len(strings.Split(password, x)), ShouldEqual, 5)

				x = "/"
				generator.SetDelimiter(x)
				So(len(strings.Split(generator.Generate(), x)), ShouldEqual, 5)

				x = "1"
				generator.SetDelimiter(x)
				So(len(strings.Split(generator.Generate(), x)), ShouldEqual, 5)
			})

		})

		Convey("Given I change the prefix to x", func() {

			Convey("It prefaces the sentence prefixed by x", func() {
				var x string = "!"
				generator.SetPrefix(x)
				password := generator.Generate()
				So(password, ShouldEqual, "!shouting unicorns eat posh buckets")
				first, _ := utf8.DecodeRuneInString(password)
				So(string(first), ShouldEqual, x)

				x = "爱"
				generator.SetPrefix(x)
				password = generator.Generate()
				first, _ = utf8.DecodeRuneInString(password)
				So(string(first), ShouldEqual, x)

				x = "xoxo"
				prefixCharacterCount := utf8.RuneCountInString(x)
				generator.SetPrefix(x)
				password = generator.Generate()
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

		Convey("Given I change the suffix to x", func() {

			Convey("It ends the sentence with x", func() {
				var x string = "!"
				generator.SetSuffix(x)
				password := generator.Generate()
				So(password, ShouldEqual, "shouting unicorns eat posh buckets!")
				last, _ := utf8.DecodeLastRuneInString(password)
				So(string(last), ShouldEqual, x)

				x = "爱"
				generator.SetSuffix(x)
				password = generator.Generate()
				last, _ = utf8.DecodeLastRuneInString(password)
				So(string(last), ShouldEqual, x)

				x = "xoxo"
				suffixCharacterCount := utf8.RuneCountInString(x)
				generator.SetSuffix(x)
				password = generator.Generate()
				var endingCharacters = string(password[len(password)-suffixCharacterCount:])
				So(endingCharacters, ShouldEqual, x)
			})

		})

		Convey("Given I reseed", func() {
			generator.Reseed(4)
			Convey("I get a different set of results", func() {
				So(generator.Generate(), ShouldNotEqual, "shouting unicorns eat posh buckets")
				So(generator.Generate(), ShouldNotEqual, "trotting lions fear portly teddy-bears")
				So(generator.Generate(), ShouldNotEqual, "trotting mermaids aggravate portly buckets")
			})

		})

	})

}
