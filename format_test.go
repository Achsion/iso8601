package iso8601_test

import (
	"github.com/Achsion/iso8601"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	testCases := []struct {
		in       time.Duration
		expected string
	}{
		{
			in:       0,
			expected: "PT0S",
		},
		{
			in:       3 * time.Nanosecond,
			expected: "PT0.000000003S",
		},
		{
			in:       40 * time.Nanosecond,
			expected: "PT0.00000004S",
		},
		{
			in:       500 * time.Nanosecond,
			expected: "PT0.0000005S",
		},
		{
			in:       7 * time.Millisecond,
			expected: "PT0.007S",
		},
		{
			in:       345 * time.Millisecond,
			expected: "PT0.345S",
		},
		{
			in:       10 * time.Second,
			expected: "PT10S",
		},
		{
			in:       -3 * time.Second,
			expected: "-PT3S",
		},
		{
			in:       3 * time.Minute,
			expected: "PT3M0S",
		},
		{
			in:       4 * time.Hour,
			expected: "PT4H0M0S",
		},
		{
			in:       44 * time.Hour,
			expected: "PT44H0M0S",
		},
		{
			in:       3*time.Hour + 40*time.Minute,
			expected: "PT3H40M0S",
		},
		{
			in:       1*time.Hour + 2*time.Minute + 3*time.Second + 456*time.Millisecond,
			expected: "PT1H2M3.456S",
		},
		{
			in:       time.Duration(-1 << 63), // min duration
			expected: "-PT2562047H47M16.854775808S",
		},
		{
			in:       time.Duration(1<<63 - 1), // max duration
			expected: "PT2562047H47M16.854775807S",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.expected, func(t *testing.T) {
			actual := iso8601.Format(tc.in)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func BenchmarkFormat(b *testing.B) {
	cases := []struct {
		name string
		dur  time.Duration
	}{
		{name: "zero duration", dur: 0},
		{name: "one second", dur: 1 * time.Second},
		{name: "1 nanosecond", dur: time.Nanosecond},
		{name: "3h40m", dur: 3*time.Hour + 40*time.Minute},
		{name: "1h2m3.456s", dur: 1*time.Hour + 2*time.Minute + 3*time.Second + 456*time.Microsecond},
		{name: "large duration", dur: time.Duration(1<<63 - 1)},
	}

	b.ResetTimer()
	for _, benchCase := range cases {
		b.Run(benchCase.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = iso8601.Format(benchCase.dur)
			}
		})
	}
}
