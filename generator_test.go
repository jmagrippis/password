package password

import (
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
)

func TestGenerator(t *testing.T) {
	var generator *Generator

	Convey("Given a new Generator with a set Dictionary and seed", t, func() {
		dictionary := &Dictionary{
			Adverbs:    []string{"cuddling", "slapping", "shouting", "jumping", "ducking", "mocking", "trotting", "galloping"},
			Subjects:   []string{"mermaids", "unicorns", "lions", "piranhas", "cuttlefish", "llamas", "dragons"},
			Verbs:      []string{"love", "fancy", "eat", "bring", "fear", "aggravate", "detest", "adore", "belittle", "ravish"},
			Adjectives: []string{"beautiful", "homely", "magical", "posh", "excellent", "ravishing", "lovely"},
			Objects:    []string{"teddy-bears", "diamonds", "buckets", "boxes", "dishes", "ornaments"},
		}
		generator = NewGenerator(dictionary, 2)

		Convey("Given I just run it with no extra settings", func() {

			Convey("It returns secure passwords", func() {
				So(generator.generate(), ShouldEqual, "shouting unicorns eat posh buckets")
				So(generator.generate(), ShouldEqual, "trotting lions fear ravishing teddy-bears")
				So(generator.generate(), ShouldEqual, "trotting mermaids aggravate ravishing buckets")
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
			Convey("It panics", nil)
			So(generator.setWordCount(3), ShouldNotEqual, nil)
			So(generator.setWordCount(1), ShouldNotEqual, nil)
			So(generator.setWordCount(4), ShouldEqual, nil)
		})

		Convey("Given I change the delimiter to x", func() {

			Convey("It returns the sentence delimited by x instead of \" \"", nil)

		})

		Convey("Given I change the prefix", func() {

			Convey("It prefaces the sentence with the given prefix", nil)

		})

		Convey("Given I change to suffix", func() {

			Convey("It ends the sentence with the given suffix", nil)

		})

		Convey("Given I reseed", func() {

			Convey("I get a different set of results", nil)

		})

	})

}
