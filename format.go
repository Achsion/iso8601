package iso8601

import "time"

// Format returns a string representing the duration in the ISO 8601 format, but only
// with the hours being the highest time element, e.g. "PT44H7M3.15s". Leading zero units are omitted.
// The result counts as a valid ISO8601 duration.
// It supports negative durations, as detailed in the extension ISO 8601-2.
func Format(duration time.Duration) string {
	// This is inlinable to take advantage of "function outlining".
	// Thus, the caller can decide whether a string must be heap allocated.
	var arr [27]byte
	bufWriteIdx := format(duration, &arr)

	return string(arr[bufWriteIdx:])
}

// format formats the ISO8601 string representation of duration into the end of outBuf and
// returns the offset of the first character.
func format(duration time.Duration, outBuf *[27]byte) int {
	// Largest possible string: '-PT2562047H47M16.854775808S' -> 27 chars
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

	bufWriteIdx = fmtInt(outBuf[:bufWriteIdx], durVal%60)
	durVal /= 60 // durVal is now integer minutes

	if durVal > 0 {
		bufWriteIdx--
		outBuf[bufWriteIdx] = 'M'
		bufWriteIdx = fmtInt(outBuf[:bufWriteIdx], durVal%60)
		durVal /= 60 // durVal is now integer hours

		if durVal > 0 {
			bufWriteIdx--
			outBuf[bufWriteIdx] = 'H'
			bufWriteIdx = fmtInt(outBuf[:bufWriteIdx], durVal)
			// stop at hours because days can be different lengths
		}
	}

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
