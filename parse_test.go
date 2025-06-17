package iso8601_test

import (
	"github.com/Achsion/iso8601"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestParseToDuration(t *testing.T) {
	testCases := []struct {
		isoStr   string
		expected time.Duration
	}{
		{
			isoStr:   "P1Y",
			expected: 1 * iso8601.TimeYear,
		},
		{
			isoStr:   "P1M",
			expected: 1 * iso8601.TimeMonth,
		},
		{
			isoStr:   "P1D",
			expected: 1 * iso8601.TimeDay,
		},
		{
			isoStr:   "PT3H40M0S",
			expected: 3*time.Hour + 40*time.Minute,
		},
		{
			isoStr:   "-PT3H40M0S", // negative duration, as detailed in the extension ISO8601-2
			expected: -3*time.Hour - 40*time.Minute,
		},
		{
			isoStr:   "PT1H",
			expected: 1 * time.Hour,
		},
		{
			isoStr:   "PT1M",
			expected: 1 * time.Minute,
		},
		{
			isoStr:   "PT1S",
			expected: 1 * time.Second,
		},
		{
			isoStr:   "PT0.1S",
			expected: 100 * time.Millisecond,
		},
		{
			isoStr:   "PT0.01S",
			expected: 10 * time.Millisecond,
		},
		{
			isoStr:   "PT0.001S",
			expected: 1 * time.Millisecond,
		},
		{
			isoStr:   "PT1.0023S",
			expected: 1*time.Second + 2*time.Millisecond + 300*time.Microsecond,
		},
		{
			isoStr:   "P1Y2M3DT4H5M6S",
			expected: 1*iso8601.TimeYear + 2*iso8601.TimeMonth + 3*iso8601.TimeDay + 4*time.Hour + 5*time.Minute + 6*time.Second,
		},
		{
			isoStr:   "P1Y2M3DT4H5M6.7S",
			expected: 1*iso8601.TimeYear + 2*iso8601.TimeMonth + 3*iso8601.TimeDay + 4*time.Hour + 5*time.Minute + 6*time.Second + 700*time.Millisecond,
		},
		{
			isoStr:   "P12Y32M153DT7H15M6.7023S",
			expected: 12*iso8601.TimeYear + 32*iso8601.TimeMonth + 153*iso8601.TimeDay + 7*time.Hour + 15*time.Minute + 6*time.Second + 702*time.Millisecond + 300*time.Microsecond,
		},
		{
			isoStr:   "P7Y3M4D",
			expected: 7*iso8601.TimeYear + 3*iso8601.TimeMonth + 4*iso8601.TimeDay,
		},
		{
			isoStr:   "PT40H5M1.0103S",
			expected: 40*time.Hour + 5*time.Minute + 1*time.Second + 10*time.Millisecond + 300*time.Microsecond,
		},
		{
			isoStr:   "PT1.23456789123S", // Too many decimal points, the last '23' will be cut/removed/ignored.
			expected: 1*time.Second + 234*time.Millisecond + 567*time.Microsecond + 891*time.Nanosecond,
		},
		{
			isoStr:   "P7Y6DT5M",
			expected: 7*iso8601.TimeYear + 6*iso8601.TimeDay + 5*time.Minute,
		},
	}

	for _, test := range testCases {
		t.Run(test.isoStr, func(t *testing.T) {
			dur, err := iso8601.ParseToDuration(test.isoStr)
			require.NoError(t, err)
			assert.Equal(t, test.expected, dur)
		})
	}
}

func TestParseToDurationError(t *testing.T) {
	testCases := []struct {
		name   string
		isoStr string
	}{
		{
			name:   "missing 'P' prefix",
			isoStr: "1Y2M3DT4H5M6S",
		},
		{
			name:   "not a duration",
			isoStr: "abc",
		},
		{
			name:   "invalid designator 'G' as last designator",
			isoStr: "P1Y2M4D1G",
		},
		{
			name:   "invalid designator 'G' in the middle",
			isoStr: "P1Y2M40G1D",
		},
		{
			name:   "invalid designator 'G' as first designator",
			isoStr: "P40G1D",
		},
		{
			name:   "wrong order of designators",
			isoStr: "PT5M4H6S",
		},
		{
			name:   "string with prefix",
			isoStr: " P7Y3M4D",
		},
		{
			name:   "string with suffix",
			isoStr: "P7Y3M4D ",
		},
		{
			name:   "string with number prefix",
			isoStr: "1P7Y3M4D",
		},
		{
			name:   "string with number suffix",
			isoStr: "P7Y3M4D1",
		},
		{
			name:   "double designator",
			isoStr: "P1Y2M3DT4H3H5M6S",
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			dur, err := iso8601.ParseToDuration(test.isoStr)
			require.Empty(t, dur)
			assert.Error(t, err, dur)
		})
	}
}

func BenchmarkParseToDuration(b *testing.B) {
	x := "P12Y32M153DT7H15M6.7023S"
	for i := 0; i < b.N; i++ {
		_, _ = iso8601.ParseToDuration(x)
	}
}
