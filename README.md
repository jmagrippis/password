# Golang Memorable Password Generator

This simple library generates a memorable but secure<sup>[citation needed]</sup> password, given a Dictionary of words.

It returns sentences in the form of:

```
[Adverb] [Subject] [Verb] [Adjective(s)] [Object]
```

The bigger the dictionary you feed it, the more secure the password! Of course, using a password manager is still the way to go with these things, but why not use something for the unique per-account passwords you can actually easily remember for a couple of minutes, and type out if need arises?

Inspired by the [popular xkcd strip][].

## Sample usage

```go
package example

import "github.com/jmagrippis/password"

dictionary := &password.Dictionary{
    Adverbs:    []string{"cuddling", "slapping", "shouting", "jumping"},
    Subjects:   []string{"mermaids", "unicorns", "lions", "piranhas"},
    Verbs:      []string{"love", "fancy", "eat", "bring", "fear", "aggravate"},
    Adjectives: []string{"beautiful", "homely", "magical", "posh", "excellent"},
    Objects:    []string{"teddy-bears", "diamonds", "buckets", "boxes"},
}

generator := password.NewGenerator(dictionary, time.Now().UnixNano())
pass := generator.Generate()
// pass = "shouting unicorns eat posh buckets"
pass = generator.Generate()
// pass = "jumping lions fear homely teddy-bears"
```

### Additional Settings

You may define strings to prefix or suffix the returned passwords, change the delimiter, change the amount of words in the returned password to any number from 4 to 255, or have the password returned In Title Case.

```go
// continuing from the code above...

generator.SetDelimiter("-❤-")
pass = generator.Generate()
// pass = "slapping-❤-mermaids-❤-fear-❤-excellent-❤-diamonds"

generator.SetDelimiter("|")
generator.setPrefix("77")
generator.setSuffix("42")
pass = generator.Generate()
// pass = "77cuddling|lions|eat|excellent|boxes42"

generator.Reseed(time.Now().UnixNano())
generator.SetTitleCase(true)
generator.SetDelimiter(" ")
generator.setPrefix("¡")
generator.setSuffix("!")
pass = generator.Generate()
// pass = "¡Cuddling Piranhas Fear Beautiful Diamonds!"
```

## Testing
Library built using BDD with the excellent [GoConvey][].

[popular xkcd strip]: https://xkcd.com/936/ "Password Strength"
[GoConvey]: http://goconvey.co/ "Satisfying Behaviour Driven Development"