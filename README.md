# Golang Memorable Password Generator

This simple library generates a memorable but secure* password, given a Dictionary of words.

It returns sentences in the form of:

```
[Adverb] [Subject] [Verb] [Adjective(s)] [Object]
```

The bigger the dictionary you feed it, the more secure the password! Of course, using a password manager is still the way to go with these things, but why not use something for the unique per-account passwords you can actually easily remember for a couple of minutes and type out if need arises?

Inspired by the [popular xkcd strip][].

## Sample usage

```go
generator = NewGenerator(dictionary, time.Now().UnixNano())
pass := generator.generate()
// pass = "shouting unicorns eat posh buckets"
pass = generator.generate()
// pass = "trotting lions fear ravishing teddy-bears"
```

## Testing
Library built using BDD with the excellent [GoConvey][]

[popular xkcd strip]: https://xkcd.com/936/ "Password Strength"
[GoConvey]: http://goconvey.co/ "Satisfying Behaviour Driven Development"