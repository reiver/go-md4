# go-md4

Package **md4** implements the MD4 cryptographic hash algorithm as per [IETF RFC-1320](https://datatracker.ietf.org/doc/html/rfc1320), for the Go programming language.

Note that MD4 is no longer considered _secure_.
However, it may still be useful for some non-secure (and historical) use-cases.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-md4

[![GoDoc](https://godoc.org/github.com/reiver/go-md4?status.svg)](https://godoc.org/github.com/reiver/go-md4)

## Examples

Here is an example of using package md4:

```go
import "github.com/reiver/go-md4"

//


var data []byte = []byte("Hello world!")

digest := md4.Sum(data)
```

## Import

To import package **md4** use `import` code like the following:
```
import "github.com/reiver/go-md4"
```

## Installation

To install package **md4** do the following:
```
GOPROXY=direct go get github.com/reiver/go-md4
```

## Author

Package **md4** was written by [Charles Iliya Krempeaux](http://reiver.link)
