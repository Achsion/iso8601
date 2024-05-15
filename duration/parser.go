package duration

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

const (
	Day   = 24 * time.Hour
	Month = 30 * Day
	Year  = 365 * Day

	yearsPatternKey   = "year"
	monthsPatternKey  = "month"
	daysPatternKey    = "day"
	hoursPatternKey   = "hour"
	minutesPatternKey = "minute"
	secondsPatternKey = "second"
)

var isoRegex = regexp.MustCompile(
	fmt.Sprintf(
		`^P((?P<%s>\d+)Y)?((?P<%s>\d+)M)?((?P<%s>\d+)D)?(T((?P<%s>\d+)H)?((?P<%s>\d+)M)?((?P<%s>[\d.]+)S)?)?$`,
		yearsPatternKey, monthsPatternKey, daysPatternKey, hoursPatternKey, minutesPatternKey, secondsPatternKey,
	),
)

// ParseToDuration parses an ISO 8601 duration string into a time.Duration.
func ParseToDuration(durationString string) (time.Duration, error) {
	matches, err := findStringCaptureGroupMatches(isoRegex, durationString)
	if err != nil {
		return 0, err
	}

	var t time.Duration

	if yearStr, ok := matches[yearsPatternKey]; ok {
		t += calcSubDuration(yearStr, Year)
	}
	if monthStr, ok := matches[monthsPatternKey]; ok {
		t += calcSubDuration(monthStr, Month)
	}
	if dayStr, ok := matches[daysPatternKey]; ok {
		t += calcSubDuration(dayStr, Day)
	}
	if hourStr, ok := matches[hoursPatternKey]; ok {
		t += calcSubDuration(hourStr, time.Hour)
	}
	if minuteStr, ok := matches[minutesPatternKey]; ok {
		t += calcSubDuration(minuteStr, time.Minute)
	}
	if secondStr, ok := matches[secondsPatternKey]; ok {
		seconds, _ := strconv.ParseFloat(secondStr, 64)
		// ignore error, as the regex ensures that the string is a valid float
		t += time.Duration(seconds * float64(time.Second))
	}

	return t, nil
}

func calcSubDuration(subStr string, durationMultiplier time.Duration) time.Duration {
	val, err := strconv.Atoi(subStr)
	if err != nil {
		// ignore error, as the regex ensures that the string is a valid int
		return 0
	}

	return time.Duration(val) * durationMultiplier
}

func findStringCaptureGroupMatches(
	regex *regexp.Regexp,
	durationStr string,
) (map[string]string, error) {
	captureGroups := regex.SubexpNames()
	matches := regex.FindStringSubmatch(durationStr)

	if matches == nil {
		return nil, fmt.Errorf("could match duration string %q with regex", durationStr)
	}

	result := make(map[string]string, len(captureGroups))
	for i, name := range captureGroups {
		if i != 0 && name != "" && matches[i] != "" {
			result[name] = matches[i]
		}
	}

	return result, nil
}
