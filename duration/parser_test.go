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
			expected: 1 * duration.Year,
		},
		{
			isoStr:   "P1M",
			expected: 1 * duration.Month,
		},
		{
			isoStr:   "P1D",
			expected: 1 * duration.Day,
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
			isoStr:   "P12Y32M153DT7H15M6.7023S",
			expected: 12*duration.Year + 32*duration.Month + 153*duration.Day + 7*time.Hour + 15*time.Minute + 6*time.Second + 702*time.Millisecond + 300*time.Microsecond,
		},
		{
			isoStr:   "P7Y3M4D",
			expected: 7*duration.Year + 3*duration.Month + 4*duration.Day,
		},
		{
			isoStr:   "PT40H5M1.0103S",
			expected: 40*time.Hour + 5*time.Minute + 1*time.Second + 10*time.Millisecond + 300*time.Microsecond,
		},
		{
			isoStr:   "P7Y6DT5M",
			expected: 7*duration.Year + 6*duration.Day + 5*time.Minute,
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

//TODO: invalid input cases
