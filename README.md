# go-jsd

[![Go Reference](https://pkg.go.dev/badge/github.com/mattn/go-jsd.svg)](https://pkg.go.dev/github.com/mattn/go-jsd)

Jaro-Winkler String Distance supporting unicode/multibyte strings.

## Usage

```go
distnace := jsd.StringDistance("こんにちわ世界", "こにゃにゃちわ世界")     // 0.0762
similarity := jsd.StringSimilarity("こんにちわ世界", "こにゃにゃちわ世界") // 0.9238
```

```go
distnace := jsd.Distance([]rune("accomodate"), []rune("accommodate")))     // 0.0181
similarity := jsd.Similarity([]rune("accomodate"), []rune("accommodate"))) // 0.9818
```

## Installation

```shellsession
$ go get github.com/mattn/go-jsd
```

## License

MIT

## Author

Yasuhiro Matsumoto (a.k.a. mattn)
