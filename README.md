# iso8601 Duration parser and formatter

[![GoDoc](https://godoc.org/github.com/Achsion/iso8601?status.svg)](https://godoc.org/github.com/Achsion/iso8601)
[![GoReport](https://goreportcard.com/badge/github.com/Achsion/iso8601)](https://goreportcard.com/report/github.com/Achsion/iso8601) 

The current go native `time` pkg does not support any ISO 8601 duration strings.

This library provides the functionality to parse and format a native Go `time.Duration` to and from an ISO 8601-2 duration string.

## What is still missing?

Currently, fast formatting (`Format`) does not support duration parts smaller than one second.

## Installation

```bash
go get github.com/Achsion/iso8601/v2
```

will resolve and add the package to the current development module, along with its dependencies.

## Usage

```go
package main

import (
	"github.com/Achsion/iso8601"
)

func main() {
	// Quick parsing to go duration:
	duration, err := iso8601.ParseToDuration("P1Y1M1DT1H1M1.1S")

	// Slower, but more complete parsing to custom duration struct:
	isoDuration, err := iso8601.DurationFromString("P1Y1M1DT1H1M1.1S")
}

```
