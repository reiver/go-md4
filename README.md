# go-md4

Package **md4** implements the MD4 cryptographic hash algorithm as per [IETF RFC-1320](https://datatracker.ietf.org/doc/html/rfc1320), for the Go programming language.

Note that MD4 is no longer considered _secure_.
However, it may still be useful for some non-secure (and historical) use-cases.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-md4

[![GoDoc](https://godoc.org/github.com/reiver/go-md4?status.svg)](https://godoc.org/github.com/reiver/go-md4)

## Examples

Here is a simple example of calculating the MD4 digest of data contained in a `[]byte`:

```go
import "github.com/reiver/go-md4"

// ...

var data []byte = []byte("Hello world!")

digest := md4.Sum(data)
```

Here is a more complex example of putting data into the an MD4 hasher in multiple pieces.

```go
import "github.com/reiver/go-md4"

// ...

hasher := md4.New()

// This is equivalent to writing:
//
//	[]byte("once+twice-thrice_fource")
hasher.Write([]byte("once"))
hasher.Write([]byte("+"))
hasher.Write([]byte("twice"))
hasher.Write([]byte("-"))
hasher.Write([]byte("thrice"))
hasher.Write([]byte("_"))
hasher.Write([]byte("fource"))

digest := hasher.Sum(nil)
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
