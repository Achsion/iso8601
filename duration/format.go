package duration

import (
	"time"
)

// FormatQuick
// Deprecated: use FormatSeconds instead.
func FormatQuick(duration time.Duration) string {
	return FormatSeconds(duration)
}

// FormatSeconds returns a string representing the duration in the ISO8601 format, but only
// in seconds, e.g. "PT345.15s". Leading zero units are omitted.
// The result counts as a valid ISO8601 duration. It is faster to format than using the 'full'
// duration format, just not as readable.
// It supports negative durations, as detailed in the extension ISO8601-2.
func FormatSeconds(duration time.Duration) string {
	// This is inlinable to take advantage of "function outlining".
	// Thus, the caller can decide whether a string must be heap allocated.
	var arr [24]byte
	n := format(duration, &arr)
	return string(arr[n:])
}

// format formats the representation of d into the end of buf and
// returns the offset of the first character.
func format(duration time.Duration, outBuf *[24]byte) int {
	// Largest possible string: '-PT9223372036.854775808S'
	bufWriteIdx := len(outBuf)
	isNegative := duration < 0

	durVal := uint64(duration)
	if isNegative {
		durVal = -durVal
	}

	bufWriteIdx--
	outBuf[bufWriteIdx] = 'S'

	bufWriteIdx, durVal = fmtFraction(outBuf[:bufWriteIdx], durVal, 9)

	// durVal is now integer seconds
	bufWriteIdx = fmtInt(outBuf[:bufWriteIdx], durVal)

	bufWriteIdx--
	outBuf[bufWriteIdx] = 'T'
	bufWriteIdx--
	outBuf[bufWriteIdx] = 'P'
	if isNegative {
		bufWriteIdx--
		outBuf[bufWriteIdx] = '-'
	}

	return bufWriteIdx
}

// fmtFrac formats the fraction of in/10**precision (e.g., ".12345") into the
// tail of buf, omitting trailing zeros. It omits the decimal
// point too when the fraction is 0. It returns the index where the
// output bytes begin and the value v/10**prec.
func fmtFraction(buf []byte, value uint64, precision int) (idx int, newValue uint64) {
	if value == 0 {
		return len(buf), value
	}

	// Omit trailing zeros up to and including decimal point.
	bufWriteIdx := len(buf)
	isPrinting := false
	var digit uint64

	for i := 0; i < precision; i++ {
		digit = value % 10

		isPrinting = isPrinting || digit != 0
		if isPrinting {
			bufWriteIdx--
			buf[bufWriteIdx] = byte(digit) + '0'
		}

		value /= 10
	}

	if isPrinting {
		bufWriteIdx--
		buf[bufWriteIdx] = '.'
	}

	return bufWriteIdx, value
}

// fmtInt formats value into the tail of buf.
// It returns the index where the output begins.
func fmtInt(buf []byte, value uint64) int {
	bufWriteIdx := len(buf)
	if value == 0 {
		bufWriteIdx--
		buf[bufWriteIdx] = '0'

		return bufWriteIdx
	}

	for value > 0 {
		bufWriteIdx--
		buf[bufWriteIdx] = byte(value%10) + '0'
		value /= 10
	}

	return bufWriteIdx
}
