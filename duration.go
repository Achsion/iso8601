package iso8601

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"time"
)

type Duration struct {
	isPositive bool

	years   float64
	months  float64
	weeks   float64
	days    float64
	hours   float64
	minutes float64
	seconds float64
}

// TODO: add docu in this file

func NewDuration(isPositive bool, years, months, weeks, days, hours, minutes, seconds float64) (Duration, error) {
	if years < 0 || months < 0 || weeks < 0 || days < 0 || hours < 0 || minutes < 0 || seconds < 0 {
		return Duration{}, errors.New("all unit values must be greater than or equal to zero")
	}

	return Duration{
		isPositive: isPositive,
		years:      years,
		months:     months,
		weeks:      weeks,
		days:       days,
		hours:      hours,
		minutes:    minutes,
		seconds:    seconds,
	}, nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Parsing ///////////////////////////////////////////////////////////////////////////////

const (
	negativePatternKey = "negative"
	yearsPatternKey    = "year"
	monthsPatternKey   = "month"
	weeksPatternKey    = "week"
	daysPatternKey     = "day"
	hoursPatternKey    = "hour"
	minutesPatternKey  = "minute"
	secondsPatternKey  = "second"
)

var durationRegex = regexp.MustCompile(
	fmt.Sprintf(
		`^((?P<%s>\-))?P((?P<%s>[0-9.]+)Y)?((?P<%s>[0-9.]+)M)?((?P<%s>[0-9.]+)W)?((?P<%s>[0-9.]+)D)?(T((?P<%s>[0-9.]+)H)?((?P<%s>[0-9.]+)M)?((?P<%s>[\d.]+)S)?)?$`,
		negativePatternKey, yearsPatternKey, monthsPatternKey, weeksPatternKey, daysPatternKey, hoursPatternKey, minutesPatternKey, secondsPatternKey,
	),
)

func DurationFromString(iso8601DurationStr string) (Duration, error) {
	// Implemented with regex for now, as speed should not be of major importance when using this func.
	// If speed is crucial, `iso8601.ParseToDuration` should be used.
	// This implementation will probably be changed to something faster but regex should suffice for now.

	matches, err := findStringCaptureGroupMatches(durationRegex, iso8601DurationStr)
	if err != nil {
		return Duration{}, err
	}

	out := Duration{}

	_, isNegative := matches[negativePatternKey]
	out.isPositive = !isNegative

	if yearStr, ok := matches[yearsPatternKey]; ok {
		out.years = mustStringToFloat64(yearStr)
	}
	if monthStr, ok := matches[monthsPatternKey]; ok {
		out.months = mustStringToFloat64(monthStr)
	}
	if weekStr, ok := matches[weeksPatternKey]; ok {
		out.weeks = mustStringToFloat64(weekStr)
	}
	if dayStr, ok := matches[daysPatternKey]; ok {
		out.days = mustStringToFloat64(dayStr)
	}
	if hourStr, ok := matches[hoursPatternKey]; ok {
		out.hours = mustStringToFloat64(hourStr)
	}
	if minuteStr, ok := matches[minutesPatternKey]; ok {
		out.minutes = mustStringToFloat64(minuteStr)
	}
	if secondStr, ok := matches[secondsPatternKey]; ok {
		out.seconds = mustStringToFloat64(secondStr)
	}

	return out, nil
}

func mustStringToFloat64(in string) float64 {
	// ignore error, as the regex ensures that the string is a valid number
	out, _ := strconv.ParseFloat(in, 64)

	return out
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

func DurationFromTimeDuration(in time.Duration) Duration {
	durVal := in.Abs()

	hours := float64(durVal / time.Hour)
	durVal = durVal % time.Hour

	minutes := float64(durVal / time.Minute)
	durVal = durVal % time.Minute

	seconds := durVal.Seconds()

	return Duration{
		isPositive: in >= 0,
		hours:      hours,
		minutes:    minutes,
		seconds:    seconds,
	}
}

//////////////////////////////////////////////////////////////////////////////////////////
// Getter ////////////////////////////////////////////////////////////////////////////////

func (d Duration) IsPositive() bool {
	return d.isPositive
}

func (d Duration) Years() float64 {
	return d.years
}

func (d Duration) Months() float64 {
	return d.months
}

func (d Duration) Weeks() float64 {
	return d.weeks
}

func (d Duration) Days() float64 {
	return d.days
}

func (d Duration) Hours() float64 {
	return d.hours
}

func (d Duration) Minutes() float64 {
	return d.minutes
}

func (d Duration) Seconds() float64 {
	return d.seconds
}

//////////////////////////////////////////////////////////////////////////////////////////
// Determination funcs ///////////////////////////////////////////////////////////////////

func (d Duration) IsZero() bool {
	return d.years == 0 &&
		d.months == 0 &&
		d.weeks == 0 &&
		d.days == 0 &&
		d.hours == 0 &&
		d.minutes == 0 &&
		d.seconds == 0
}

//////////////////////////////////////////////////////////////////////////////////////////
// Go std time stuff /////////////////////////////////////////////////////////////////////

func (d Duration) AddToTime(stdTime time.Time) (time.Time, error) {
	if d.IsZero() {
		return stdTime, nil
	}

	multiplier := 1
	if !d.isPositive {
		multiplier = -1
	}

	yearAdd, err := float64ToInt(d.years)
	if err != nil {
		return time.Time{}, errors.New("could not convert year to int")
	}
	monthAdd, err := float64ToInt(d.months)
	if err != nil {
		return time.Time{}, errors.New("could not convert month to int")
	}

	weeks, decimalWeeks := math.Modf(d.weeks)
	weekAdd, err := float64ToInt(weeks)
	if err != nil {
		return time.Time{}, errors.New("could not convert week to int")
	}

	days, decimalDays := math.Modf(d.days + 7*decimalWeeks)
	dayAdd, err := float64ToInt(days)
	if err != nil {
		return time.Time{}, errors.New("could not convert day + remaining week to int")
	}

	hours, decimalHours := math.Modf(d.hours + 24*decimalDays)
	hourAdd, err := float64ToInt(hours)
	if err != nil {
		return time.Time{}, errors.New("could not convert hour + remaining day to int")
	}

	minutes, decimalMinutes := math.Modf(d.minutes + 60*decimalHours)
	minuteAdd, err := float64ToInt(minutes)
	if err != nil {
		return time.Time{}, errors.New("could not convert minute + remaining hour to int")
	}

	seconds, nanoSeconds := math.Modf(d.seconds + 60*decimalMinutes)
	secondAdd, err := float64ToInt(seconds)
	if err != nil {
		return time.Time{}, errors.New("could not convert second + remaining minute to int")
	}

	nanoSecondAdd, err := float64ToInt(nanoSeconds * float64(time.Second))
	if err != nil {
		return time.Time{}, errors.New("could not convert nanosecond to int")
	}

	out := stdTime.AddDate(
		multiplier*yearAdd,
		multiplier*monthAdd,
		multiplier*(weekAdd*7+dayAdd),
	).Add(
		time.Hour*time.Duration(hourAdd) +
			time.Minute*time.Duration(minuteAdd) +
			time.Second*time.Duration(secondAdd) +
			time.Nanosecond*time.Duration(nanoSecondAdd),
	)

	return out, nil
}

func float64ToInt(in float64) (int, error) {
	intVal, fracVal := math.Modf(in)
	if fracVal != 0.0 {
		return 0, errors.New("float containing decimals not supported")
	}

	out := int(intVal)

	if (out < 0) != (intVal < 0) {
		return 0, errors.New("float int val exceeds integer capacity")
	}

	return out, nil
}

//////////////////////////////////////////////////////////////////////////////////////////
// Formatting ////////////////////////////////////////////////////////////////////////////

func (d Duration) String() string {
	hasDate := false
	hasTime := false

	out := "P"
	if !d.isPositive {
		out = "-P"
	}

	if d.years != 0 {
		appendDurationPart(&out, d.years, yearDesignator)
		hasDate = true
	}
	if d.months != 0 {
		appendDurationPart(&out, d.months, monthDesignator)
		hasDate = true
	}
	if d.weeks != 0 {
		appendDurationPart(&out, d.weeks, weekDesignator)
		hasDate = true
	}
	if d.days != 0 {
		appendDurationPart(&out, d.days, dayDesignator)
		hasDate = true
	}

	if d.hours != 0 {
		out += "T"
		appendDurationPart(&out, d.hours, hourDesignator)
		hasTime = true
	}
	if d.minutes != 0 {
		if !hasTime {
			out += "T"
		}
		appendDurationPart(&out, d.minutes, minuteDesignator)
		hasTime = true
	}
	if d.seconds != 0 {
		if !hasTime {
			out += "T"
		}
		appendDurationPart(&out, d.seconds, secondDesignator)
		hasTime = true
	}

	if !hasDate && !hasTime {
		out += "T0S"
	}

	return out
}

func appendDurationPart(buf *string, value float64, suffix rune) {
	*buf += strconv.FormatFloat(value, 'f', -1, 64) + string(suffix)
}
