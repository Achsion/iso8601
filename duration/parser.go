package duration

import (
	"fmt"
	"math"
	"strconv"
	"time"
	"unicode"
)

// time values for missing time values
const (
	TimeDay   = 24 * time.Hour
	TimeMonth = 30 * TimeDay
	TimeYear  = 365 * TimeDay
)

const (
	durationStringShift = 1
	idxLookupShift      = 1

	startDesignator      = 'P'
	timeSwitchDesignator = 'T'
)

const (
	yearIdx = iota + 1
	monthIdx
	dayIdx
	hourIdx
	minuteIdx
	secondSepIdx
	secondIdx
)

var dateLookup = map[int32]int{
	'Y': yearIdx,
	'M': monthIdx,
	'D': dayIdx,
}
var timeLookup = map[int32]int{
	'H': hourIdx,
	'M': minuteIdx,
	'.': secondSepIdx,
	'S': secondIdx,
}

var invalidFormatErr = fmt.Errorf("invalid iso8601 duration format")

// ParseToDuration parses an ISO 8601 duration string into a time.Duration.
func ParseToDuration(durationString string) (time.Duration, error) {
	if durationString[0] != startDesignator || len(durationString) < 3 {
		// duration string has to start with 'P' and the shortest possible length is 3 (e.g. "P3D")
		return 0, invalidFormatErr
	}

	var stringPart string
	var idx int
	lastIdx := -1
	interpretDate := true

	// duration string split in its parts
	durationParts := make([]string, 7)

	// separating duration string into parts
	numberStartIndex := 0
	for charIndex, nextChar := range durationString[durationStringShift:] {
		if nextChar == timeSwitchDesignator {
			interpretDate = false
			numberStartIndex = charIndex + 1
			continue
		}

		if !unicode.IsNumber(nextChar) {
			stringPart = durationString[numberStartIndex+durationStringShift : charIndex+durationStringShift]

			if interpretDate {
				idx = dateLookup[nextChar]
			} else {
				idx = timeLookup[nextChar]
			}

			if idx == 0 || lastIdx >= idx {
				// the designators are in the wrong order / there is a duplicate designator
				return 0, invalidFormatErr
			}

			durationParts[idx-idxLookupShift] = stringPart
			numberStartIndex = charIndex + 1
			lastIdx = idx
		}
	}

	if numberStartIndex+1 < len(durationString) {
		// there are still some characters 'left' in the string that should not be there
		return 0, invalidFormatErr
	}

	return calculateDuration(durationParts), nil
}

func calculateDuration(durationParts []string) time.Duration {
	var t time.Duration
	var nrVal int

	if durationParts[yearIdx-idxLookupShift] != "" {
		nrVal, _ = strconv.Atoi(durationParts[yearIdx-idxLookupShift])
		t += TimeYear * time.Duration(nrVal)
	}
	if durationParts[monthIdx-idxLookupShift] != "" {
		nrVal, _ = strconv.Atoi(durationParts[monthIdx-idxLookupShift])
		t += TimeMonth * time.Duration(nrVal)
	}
	if durationParts[dayIdx-idxLookupShift] != "" {
		nrVal, _ = strconv.Atoi(durationParts[dayIdx-idxLookupShift])
		t += TimeDay * time.Duration(nrVal)
	}
	if durationParts[hourIdx-idxLookupShift] != "" {
		nrVal, _ = strconv.Atoi(durationParts[hourIdx-idxLookupShift])
		t += time.Hour * time.Duration(nrVal)
	}
	if durationParts[minuteIdx-idxLookupShift] != "" {
		nrVal, _ = strconv.Atoi(durationParts[minuteIdx-idxLookupShift])
		t += time.Minute * time.Duration(nrVal)
	}
	if durationParts[secondIdx-idxLookupShift] != "" {
		lesserSecsStr := durationParts[secondIdx-idxLookupShift]
		decimalPoints := len(lesserSecsStr)
		if decimalPoints > 9 {
			lesserSecsStr = lesserSecsStr[:9]
			decimalPoints = 9
		}
		nrVal, _ = strconv.Atoi(lesserSecsStr)

		if durationParts[secondSepIdx-idxLookupShift] != "" {
			sVal, _ := strconv.Atoi(durationParts[secondSepIdx-idxLookupShift])
			t += time.Second*time.Duration(sVal) + calculateLesserSecondsDuration(nrVal, decimalPoints)
		} else {
			t += time.Second * time.Duration(nrVal)
		}
	}

	return t
}

func calculateLesserSecondsDuration(lesserSeconds int, originalLength int) time.Duration {
	if lesserSeconds == 0 {
		return 0
	}

	multiplier := time.Second / time.Duration(math.Pow10(originalLength))

	return time.Duration(lesserSeconds) * multiplier
}
