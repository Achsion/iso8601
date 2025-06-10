# iso8601 Duration parser and formatter

[![GoDoc](https://godoc.org/github.com/Achsion/iso8601?status.svg)](https://godoc.org/github.com/Achsion/iso8601)
[![GoReport](https://goreportcard.com/badge/github.com/Achsion/iso8601)](https://goreportcard.com/report/github.com/Achsion/iso8601) 

The current go native `time` pkg does not support any ISO 8601 duration strings.

This library provides the functionality to parse and format a native Go `time.Duration` to and from an ISO 8601 duration string.

## What ISO8601 duration strings are supported?

Currently, parsing only supports ISO 8601-1.  
However, formatting a Duration into a string already supports negative durations and the ISO 8601-2 extension.

## Should I use this?

Do you want to parse a full ISO8601 duration? If so, then probably not.  
Do you want to parse ISO8601 durations that hardly ever reach '1D' and never reaches a higher duration part (e.g. month)? Then yes, you may want to use this.

Some APIs (for example) provide ISO8601 durations in their response even though the duration never reaches a value higher than a few days.  
This project aims for those cases.

### Why should I not use this for other use cases?

How long is a month? How long is a year?

Both, a month and a year, don't really have a hard value as a duration. But this project needs one for them regardless.  
Currently a month is hard-coded as 30 days and a year received the duration of 365 days. This is not a perfect solution but will work for use cases where a duration hardly, if ever, reaches those values.

## Installation

```bash
go get github.com/Achsion/iso8601
```

will resolve and add the package to the current development module, along with its dependencies.

## Usage

```go
package main

import "github.com/Achsion/iso8601/duration"

func main() {
	duration, err := duration.ParseToDuration("P1Y1M1DT1H1M1.1S")
}

```
