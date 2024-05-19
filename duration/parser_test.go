package duration_test

import (
	"github.com/Achsion/iso8601/duration"
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
			expected: 1 * duration.TimeYear,
		},
		{
			isoStr:   "P1M",
			expected: 1 * duration.TimeMonth,
		},
		{
			isoStr:   "P1D",
			expected: 1 * duration.TimeDay,
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
			expected: 1*duration.TimeYear + 2*duration.TimeMonth + 3*duration.TimeDay + 4*time.Hour + 5*time.Minute + 6*time.Second,
		},
		{
			isoStr:   "P1Y2M3DT4H5M6.7S",
			expected: 1*duration.TimeYear + 2*duration.TimeMonth + 3*duration.TimeDay + 4*time.Hour + 5*time.Minute + 6*time.Second + 700*time.Millisecond,
		},
		{
			isoStr:   "P12Y32M153DT7H15M6.7023S",
			expected: 12*duration.TimeYear + 32*duration.TimeMonth + 153*duration.TimeDay + 7*time.Hour + 15*time.Minute + 6*time.Second + 702*time.Millisecond + 300*time.Microsecond,
		},
		{
			isoStr:   "P7Y3M4D",
			expected: 7*duration.TimeYear + 3*duration.TimeMonth + 4*duration.TimeDay,
		},
		{
			isoStr:   "PT40H5M1.0103S",
			expected: 40*time.Hour + 5*time.Minute + 1*time.Second + 10*time.Millisecond + 300*time.Microsecond,
		},
		{
			isoStr:   "P7Y6DT5M",
			expected: 7*duration.TimeYear + 6*duration.TimeDay + 5*time.Minute,
		},
	}

	for _, test := range testCases {
		t.Run(test.isoStr, func(t *testing.T) {
			dur, err := duration.ParseToDuration(test.isoStr)
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
			dur, err := duration.ParseToDuration(test.isoStr)
			require.Empty(t, dur)
			assert.Error(t, err, dur)
		})
	}
}
