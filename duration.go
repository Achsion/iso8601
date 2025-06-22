package iso8601

import (
	"errors"
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

func NewDuration(isPositive bool, years, months, weeks, days, hours, minutes, seconds float64) Duration {
	return Duration{
		isPositive: isPositive,
		years:      years,
		months:     months,
		weeks:      weeks,
		days:       days,
		hours:      hours,
		minutes:    minutes,
		seconds:    seconds,
	}
}

func DurationFromString(durationStr string) (Duration, error) {
	// TODO: implement this

	return Duration{}, errors.New("not implemented")
}

func DurationFromTimeDuration(in time.Duration) (Duration, error) {
	return Duration{}, errors.New("not implemented")
}

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
