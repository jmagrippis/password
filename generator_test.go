package generator

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGenerator(t *testing.T) {
	var generator *Generator

	Convey("Given a new Generator with a set Dictionary and seed", t, func() {
		dictionary := &Dictionary{
			Adjectives: []string{"beautiful", "homely"},
			Subjects: []string{"mermaids", "unicorns"},
			Verbs: []string{"love", "fancy"},
			Adverbs: []string{"cuddling", "slapping"},
			Objects: []string{"teddy-bears", "diamonds"},
		}
		generator = NewGenerator(dictionary, 1)

		Convey("Given I just run it with no extra settings", func() {

			Convey("It returns secure passwords", func() {
					So(generator.generate(), ShouldEqual, "beautiful unicorns fancy cuddling teddy-bears")
				})

		})

		Convey("Given I change the number of returned words to x", func() {

			Convey("It returns a sentence of x words", nil)

		})

		Convey("Given I change the number of returned words to below 4", func() {

			Convey("It panics", nil)

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

	})

}