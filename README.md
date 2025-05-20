# iso8601 Duration parser

[![GoDoc](https://godoc.org/github.com/Achsion/iso8601?status.svg)](https://godoc.org/github.com/Achsion/iso8601)
[![GoReport](https://goreportcard.com/badge/github.com/Achsion/iso8601)](https://goreportcard.com/report/github.com/Achsion/iso8601) 

The current go native `time.ParseDuration()` does not support any ISO8601 duration strings.

This library parses any ISO8601-1 duration string into a native Go `time.Duration` object.

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

## Benchmark

Benchmarks created with
```bash
go test -bench=. -benchmem ./...
```

### `v1.1.0`
```text
goos: linux
goarch: amd64
pkg: github.com/Achsion/iso8601/duration
cpu: AMD Ryzen 7 PRO 5875U with Radeon Graphics     
BenchmarkFormatQuick-16        	17331862	        69.89 ns/op	      20 B/op	       2 allocs/op
BenchmarkParseToDuration-16    	12591764	        87.91 ns/op	       0 B/op	       0 allocs/op
```

<details>
  <summary>Older versions</summary>

### `v1.0.1`
```text
goos: linux
goarch: amd64
pkg: github.com/Achsion/iso8601/duration
cpu: AMD Ryzen 7 PRO 5875U with Radeon Graphics     
BenchmarkParseToDuration-16    	13721776	        86.28 ns/op	       0 B/op	       0 allocs/op
```

### `v1.0.0`
```text
goos: linux
goarch: amd64
pkg: github.com/Achsion/iso8601/duration
cpu: AMD Ryzen 7 PRO 5875U with Radeon Graphics     
BenchmarkParseToDuration-16    	13843465	        85.54 ns/op	       0 B/op	       0 allocs/op
```

### `v0.1.0`
```text
goos: linux
goarch: amd64
pkg: github.com/Achsion/iso8601/duration
cpu: AMD Ryzen 7 PRO 5875U with Radeon Graphics     
BenchmarkParseToDuration-16    	  878901	      1154 ns/op	    1655 B/op	       4 allocs/op
```
</details>
