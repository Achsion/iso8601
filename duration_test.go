package iso8601_test

import (
	"github.com/Achsion/iso8601"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDuration_String(t *testing.T) {
	testCases := []struct {
		name            string
		iso8601Duration iso8601.Duration
		expected        string
	}{
		{
			name:            "zero",
			iso8601Duration: iso8601.NewDuration(true, 0, 0, 0, 0, 0, 0, 0),
			expected:        "PT0S",
		},
		{
			name:            "3 nanoseconds",
			iso8601Duration: iso8601.NewDuration(true, 0, 0, 0, 0, 0, 0, 0.000000003),
			expected:        "PT0.000000003S",
		},
		{
			name:            "40 nanoseconds",
			iso8601Duration: iso8601.NewDuration(true, 0, 0, 0, 0, 0, 0, 0.00000004),
			expected:        "PT0.00000004S",
		},
		{
			name:            "500 nanoseconds",
			iso8601Duration: iso8601.NewDuration(true, 0, 0, 0, 0, 0, 0, 0.0000005),
			expected:        "PT0.0000005S",
		},
		{
			name:            "7 milliseconds",
			iso8601Duration: iso8601.NewDuration(true, 0, 0, 0, 0, 0, 0, 0.007),
			expected:        "PT0.007S",
		},
		{
			name:            "345 milliseconds",
			iso8601Duration: iso8601.NewDuration(true, 0, 0, 0, 0, 0, 0, 0.345),
			expected:        "PT0.345S",
		},
		{
			name:            "seconds",
			iso8601Duration: iso8601.NewDuration(true, 0, 0, 0, 0, 0, 0, 3),
			expected:        "PT3S",
		},
		{
			name:            "negative duration",
			iso8601Duration: iso8601.NewDuration(false, 0, 0, 0, 0, 0, 0, 3),
			expected:        "-PT3S",
		},
		{
			name:            "minutes",
			iso8601Duration: iso8601.NewDuration(true, 0, 0, 0, 0, 0, 3, 0),
			expected:        "PT3M",
		},
		{
			name:            "hours",
			iso8601Duration: iso8601.NewDuration(true, 0, 0, 0, 0, 44, 0, 0),
			expected:        "PT44H",
		},
		{
			name:            "3 hours 40 minutes",
			iso8601Duration: iso8601.NewDuration(true, 0, 0, 0, 0, 3, 40, 0),
			expected:        "PT3H40M",
		},
		{
			name:            "days",
			iso8601Duration: iso8601.NewDuration(true, 0, 0, 0, 6, 0, 0, 0),
			expected:        "P6D",
		},
		{
			name:            "weeks",
			iso8601Duration: iso8601.NewDuration(true, 0, 0, 80, 0, 0, 0, 0),
			expected:        "P80W",
		},
		{
			name:            "months",
			iso8601Duration: iso8601.NewDuration(true, 0, 981, 0, 0, 0, 0, 0),
			expected:        "P981M",
		},
		{
			name:            "years",
			iso8601Duration: iso8601.NewDuration(true, 9, 0, 0, 0, 0, 0, 0),
			expected:        "P9Y",
		},
		{
			name:            "one of everything",
			iso8601Duration: iso8601.NewDuration(true, 1, 2, 3, 4, 5, 6, 7.89),
			expected:        "P1Y2M3W4DT5H6M7.89S",
		},
		{
			name:            "one of everything, negative",
			iso8601Duration: iso8601.NewDuration(false, 1, 2, 3, 4, 5, 6, 7.89),
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

func BenchmarkDuration_String(b *testing.B) {
	cases := []struct {
		name string
		dur  iso8601.Duration
	}{
		{name: "zero duration", dur: iso8601.NewDuration(true, 0, 0, 0, 0, 0, 0, 0)},
		{name: "one second", dur: iso8601.NewDuration(true, 0, 0, 0, 0, 0, 0, 1)},
		{name: "1 nanosecond", dur: iso8601.NewDuration(true, 0, 0, 0, 0, 0, 0, 0.000000001)},
		{name: "3h40m", dur: iso8601.NewDuration(true, 0, 0, 0, 0, 3, 40, 0)},
		{name: "-3h40m", dur: iso8601.NewDuration(false, 0, 0, 0, 0, 3, 40, 0)},
		{name: "1h2m3.456s", dur: iso8601.NewDuration(true, 0, 0, 0, 0, 1, 2, 3.456)},
		{name: "one of everything", dur: iso8601.NewDuration(true, 1, 2, 3, 4, 5, 6, 7.89)},
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
