package duration

import (
	"strconv"
	"time"
)

// TODO: add doc comment
func FormatQuick(duration time.Duration) string {
	// TODO: make faster, this is "just" a first solution because I need this functionality.
	// TODO: this is shit ugly and fucking sucks in how it creates the output string. fix this whole thing.
	//		it is okay for now, because i really quickly need this functionality without any regards of actual speed.
	//		all that matters for now is that this func exists and works as intended.
	secondStr := strconv.FormatInt(int64(duration/time.Second), 10)

	// Largest decimal section is .000000000
	var decimalBuf [10]byte
	var bufWriteIdx = len(decimalBuf)
	expectedIterations := 9
	startWriting := false

	if nsec := duration % time.Second; nsec > 0 {
		for nsec > 0 {
			expectedIterations--
			if !startWriting && nsec%10 == 0 {
				nsec /= 10
				continue
			}

			startWriting = true
			bufWriteIdx--
			decimalBuf[bufWriteIdx] = byte(nsec%10) + '0'
			nsec /= 10
		}

		for range expectedIterations {
			bufWriteIdx--
			decimalBuf[bufWriteIdx] = '0'
		}

		bufWriteIdx--
		decimalBuf[bufWriteIdx] = '.'
	}

	return "PT" + secondStr + string(decimalBuf[bufWriteIdx:]) + "S"
}
