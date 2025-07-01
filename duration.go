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

func DurationFromString(iso8601DurationStr string) (Duration, error) {
	// TODO: implement this

	return Duration{}, errors.New("not implemented")
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
// Go std time stuff /////////////////////////////////////////////////////////////////////

func (d Duration) AddToTime(stdTime time.Time) (time.Duration, error) {
	// TODO: implement this

	return 0, errors.New("not implemented")
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
