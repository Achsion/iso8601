package iso8601_test

import (
	"github.com/Achsion/iso8601"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func newDurationT(
	t *testing.T,
	isPositive bool, years, months, weeks, days, hours, minutes, seconds float64,
) iso8601.Duration {
	out, err := iso8601.NewDuration(isPositive, years, months, weeks, days, hours, minutes, seconds)
	require.NoError(t, err)

	return out
}

func newDurationB(
	b *testing.B,
	isPositive bool, years, months, weeks, days, hours, minutes, seconds float64,
) iso8601.Duration {
	out, err := iso8601.NewDuration(isPositive, years, months, weeks, days, hours, minutes, seconds)
	require.NoError(b, err)

	return out
}

func TestNewDuration_Error(t *testing.T) {
	t.Run("negative year", func(t *testing.T) {
		_, err := iso8601.NewDuration(true, -1, 0, 0, 0, 0, 0, 0)
		assert.Error(t, err)
	})
	t.Run("negative month", func(t *testing.T) {
		_, err := iso8601.NewDuration(true, 0, -1, 0, 0, 0, 0, 0)
		assert.Error(t, err)
	})
	t.Run("negative week", func(t *testing.T) {
		_, err := iso8601.NewDuration(true, 0, 0, -1, 0, 0, 0, 0)
		assert.Error(t, err)
	})
	t.Run("negative day", func(t *testing.T) {
		_, err := iso8601.NewDuration(true, 0, 0, 0, -1, 0, 0, 0)
		assert.Error(t, err)
	})
	t.Run("negative hour", func(t *testing.T) {
		_, err := iso8601.NewDuration(true, 0, 0, 0, 0, -1, 0, 0)
		assert.Error(t, err)
	})
	t.Run("negative minute", func(t *testing.T) {
		_, err := iso8601.NewDuration(true, 0, 0, 0, 0, 0, -1, 0)
		assert.Error(t, err)
	})
	t.Run("negative second", func(t *testing.T) {
		_, err := iso8601.NewDuration(true, 0, 0, 0, 0, 0, 0, -1)
		assert.Error(t, err)
	})
}

func TestDuration_String(t *testing.T) {
	testCases := []struct {
		name            string
		iso8601Duration iso8601.Duration
		expected        string
	}{
		{
			name:            "zero",
			iso8601Duration: newDurationT(t, true, 0, 0, 0, 0, 0, 0, 0),
			expected:        "PT0S",
		},
		{
			name:            "3 nanoseconds",
			iso8601Duration: newDurationT(t, true, 0, 0, 0, 0, 0, 0, 0.000000003),
			expected:        "PT0.000000003S",
		},
		{
			name:            "40 nanoseconds",
			iso8601Duration: newDurationT(t, true, 0, 0, 0, 0, 0, 0, 0.00000004),
			expected:        "PT0.00000004S",
		},
		{
			name:            "500 nanoseconds",
			iso8601Duration: newDurationT(t, true, 0, 0, 0, 0, 0, 0, 0.0000005),
			expected:        "PT0.0000005S",
		},
		{
			name:            "7 milliseconds",
			iso8601Duration: newDurationT(t, true, 0, 0, 0, 0, 0, 0, 0.007),
			expected:        "PT0.007S",
		},
		{
			name:            "345 milliseconds",
			iso8601Duration: newDurationT(t, true, 0, 0, 0, 0, 0, 0, 0.345),
			expected:        "PT0.345S",
		},
		{
			name:            "seconds",
			iso8601Duration: newDurationT(t, true, 0, 0, 0, 0, 0, 0, 3),
			expected:        "PT3S",
		},
		{
			name:            "negative duration",
			iso8601Duration: newDurationT(t, false, 0, 0, 0, 0, 0, 0, 3),
			expected:        "-PT3S",
		},
		{
			name:            "minutes",
			iso8601Duration: newDurationT(t, true, 0, 0, 0, 0, 0, 3, 0),
			expected:        "PT3M",
		},
		{
			name:            "floating point minutes",
			iso8601Duration: newDurationT(t, true, 0, 0, 0, 0, 0, 3.765, 0),
			expected:        "PT3.765M",
		},
		{
			name:            "hours",
			iso8601Duration: newDurationT(t, true, 0, 0, 0, 0, 44, 0, 0),
			expected:        "PT44H",
		},
		{
			name:            "floating point hours",
			iso8601Duration: newDurationT(t, true, 0, 0, 0, 0, 44.44, 0, 0),
			expected:        "PT44.44H",
		},
		{
			name:            "3 hours 40 minutes",
			iso8601Duration: newDurationT(t, true, 0, 0, 0, 0, 3, 40, 0),
			expected:        "PT3H40M",
		},
		{
			name:            "days",
			iso8601Duration: newDurationT(t, true, 0, 0, 0, 6, 0, 0, 0),
			expected:        "P6D",
		},
		{
			name:            "floating point days",
			iso8601Duration: newDurationT(t, true, 0, 0, 0, 6.33, 0, 0, 0),
			expected:        "P6.33D",
		},
		{
			name:            "weeks",
			iso8601Duration: newDurationT(t, true, 0, 0, 80, 0, 0, 0, 0),
			expected:        "P80W",
		},
		{
			name:            "floating point weeks",
			iso8601Duration: newDurationT(t, true, 0, 0, 80.99, 0, 0, 0, 0),
			expected:        "P80.99W",
		},
		{
			name:            "months",
			iso8601Duration: newDurationT(t, true, 0, 981, 0, 0, 0, 0, 0),
			expected:        "P981M",
		},
		{
			name:            "floating point months",
			iso8601Duration: newDurationT(t, true, 0, 981.004, 0, 0, 0, 0, 0),
			expected:        "P981.004M",
		},
		{
			name:            "years",
			iso8601Duration: newDurationT(t, true, 9, 0, 0, 0, 0, 0, 0),
			expected:        "P9Y",
		},
		{
			name:            "floating point years",
			iso8601Duration: newDurationT(t, true, 9.12, 0, 0, 0, 0, 0, 0),
			expected:        "P9.12Y",
		},
		{
			name:            "one of everything",
			iso8601Duration: newDurationT(t, true, 1, 2, 3, 4, 5, 6, 7.89),
			expected:        "P1Y2M3W4DT5H6M7.89S",
		},
		{
			name:            "one of everything, negative",
			iso8601Duration: newDurationT(t, false, 1, 2, 3, 4, 5, 6, 7.89),
			expected:        "-P1Y2M3W4DT5H6M7.89S",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.iso8601Duration.String()
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestDurationFromTimeDuration(t *testing.T) {
	testCases := []struct {
		name     string
		in       time.Duration
		expected iso8601.Duration
	}{
		{
			name:     "zero",
			in:       0,
			expected: newDurationT(t, true, 0, 0, 0, 0, 0, 0, 0),
		},
		{
			name:     "1ns",
			in:       1 * time.Nanosecond,
			expected: newDurationT(t, true, 0, 0, 0, 0, 0, 0, 0.000000001),
		},
		{
			name:     "1s",
			in:       1 * time.Second,
			expected: newDurationT(t, true, 0, 0, 0, 0, 0, 0, 1),
		},
		{
			name:     "48h10m7s",
			in:       48*time.Hour + 10*time.Minute + 7*time.Second,
			expected: newDurationT(t, true, 0, 0, 0, 0, 48, 10, 7),
		},
		{
			name:     "-48h10m7s",
			in:       -48*time.Hour - 10*time.Minute - 7*time.Second,
			expected: newDurationT(t, false, 0, 0, 0, 0, 48, 10, 7),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := iso8601.DurationFromTimeDuration(tc.in)

			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestDuration_AddToTime_Success(t *testing.T) {
	testCases := []struct {
		name     string
		dur      iso8601.Duration
		stdTime  time.Time
		expected time.Time
	}{
		{ //TODO: test more
			name:     "first test",
			dur:      newDurationT(t, true, 1, 1, 0, 1, 3, 3, 3),
			stdTime:  time.Date(2003, 3, 3, 15, 15, 15, 0, time.UTC),
			expected: time.Date(2004, 4, 4, 18, 18, 18, 0, time.UTC),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, actualErr := tc.dur.AddToTime(tc.stdTime)
			require.NoError(t, actualErr)

			assert.Equal(t, tc.expected, actual)
		})
	}
}

func BenchmarkDuration_String(b *testing.B) {
	cases := []struct {
		name string
		dur  iso8601.Duration
	}{
		{name: "zero duration", dur: newDurationB(b, true, 0, 0, 0, 0, 0, 0, 0)},
		{name: "one second", dur: newDurationB(b, true, 0, 0, 0, 0, 0, 0, 1)},
		{name: "1.55 second", dur: newDurationB(b, true, 0, 0, 0, 0, 0, 0, 1.55)},
		{name: "1 nanosecond", dur: newDurationB(b, true, 0, 0, 0, 0, 0, 0, 0.000000001)},
		{name: "3h40m", dur: newDurationB(b, true, 0, 0, 0, 0, 3, 40, 0)},
		{name: "-3h40m", dur: newDurationB(b, false, 0, 0, 0, 0, 3, 40, 0)},
		{name: "1h2m3.456s", dur: newDurationB(b, true, 0, 0, 0, 0, 1, 2, 3.456)},
		{name: "one of everything", dur: newDurationB(b, true, 1, 2, 3, 4, 5, 6, 7.89)},
	}

	b.ResetTimer()
	for _, benchCase := range cases {
		b.Run(benchCase.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = benchCase.dur.String()
			}
		})
	}
}
