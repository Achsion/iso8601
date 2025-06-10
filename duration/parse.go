package duration

import (
	"fmt"
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

// decimalPointMultiplier stores the pre-computed multiplier used for decimal points calculation
var decimalPointMultiplier = [...]time.Duration{
	1e09, 1e08, 1e07, 1e06, 1e05, 1e04, 1e03, 1e02, 1e01, 1e00,
}

var invalidFormatErr = fmt.Errorf("invalid iso8601 duration format")

// ParseToDuration parses an ISO 8601 duration string into a time.Duration.
// It accepts negative durations but only by prepending a '-' like: "[-]P<duration>".
func ParseToDuration(durationString string) (time.Duration, error) {
	isNegative := false

	// consume [-]?
	if durationString != "" {
		firstChar := durationString[0]
		if firstChar == '-' {
			isNegative = true
			durationString = durationString[1:]
		}
	}

	if len(durationString) < 3 || durationString[0] != startDesignator {
		// duration string has to start with 'P' or '-P' and the shortest possible length is 3 (e.g. "P3D")
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

	return calculateDuration(durationParts, isNegative), nil
}

func calculateDuration(durationParts []string, isNegative bool) time.Duration {
	var resultDur time.Duration
	var nrVal int

	if durationParts[yearIdx-idxLookupShift] != "" {
		nrVal, _ = strconv.Atoi(durationParts[yearIdx-idxLookupShift])
		resultDur += TimeYear * time.Duration(nrVal)
	}
	if durationParts[monthIdx-idxLookupShift] != "" {
		nrVal, _ = strconv.Atoi(durationParts[monthIdx-idxLookupShift])
		resultDur += TimeMonth * time.Duration(nrVal)
	}
	if durationParts[dayIdx-idxLookupShift] != "" {
		nrVal, _ = strconv.Atoi(durationParts[dayIdx-idxLookupShift])
		resultDur += TimeDay * time.Duration(nrVal)
	}
	if durationParts[hourIdx-idxLookupShift] != "" {
		nrVal, _ = strconv.Atoi(durationParts[hourIdx-idxLookupShift])
		resultDur += time.Hour * time.Duration(nrVal)
	}
	if durationParts[minuteIdx-idxLookupShift] != "" {
		nrVal, _ = strconv.Atoi(durationParts[minuteIdx-idxLookupShift])
		resultDur += time.Minute * time.Duration(nrVal)
	}
	if durationParts[secondIdx-idxLookupShift] != "" {
		decimalSecsStr := durationParts[secondIdx-idxLookupShift]
		decimalPoints := len(decimalSecsStr)
		if decimalPoints > 9 {
			decimalSecsStr = decimalSecsStr[:9]
			decimalPoints = 9
		}
		nrVal, _ = strconv.Atoi(decimalSecsStr)

		if durationParts[secondSepIdx-idxLookupShift] != "" {
			sVal, _ := strconv.Atoi(durationParts[secondSepIdx-idxLookupShift])
			resultDur += time.Second*time.Duration(sVal) + calculateDecimalSecondsDuration(nrVal, decimalPoints)
		} else {
			resultDur += time.Second * time.Duration(nrVal)
		}
	}

	if isNegative {
		resultDur = -resultDur
	}

	return resultDur
}

func calculateDecimalSecondsDuration(decimalSeconds int, multiplyIdx int) time.Duration {
	if decimalSeconds == 0 {
		return 0
	}

	return time.Duration(decimalSeconds) * decimalPointMultiplier[multiplyIdx]
}
