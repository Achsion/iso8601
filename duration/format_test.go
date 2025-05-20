package duration_test

import (
	"github.com/Achsion/iso8601/duration"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFormatQuick(t *testing.T) {
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
			in:       3 * time.Minute,
			expected: "PT180S",
		},
		{
			in:       4 * time.Hour,
			expected: "PT14400S",
		},
		{
			in:       44 * time.Hour,
			expected: "PT158400S",
		},
		{
			in:       1*time.Hour + 2*time.Minute + 3*time.Second + 456*time.Millisecond,
			expected: "PT3723.456S",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.expected, func(t *testing.T) {
			actual := duration.FormatQuick(tc.in)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

//func TestFormat(t *testing.T) { // TODO: format a bit more properly
//	testCases := []struct {
//		in       time.Duration
//		expected string
//	}{
//		{
//			in:       0,
//			expected: "PT0S",
//		},
//		{
//			in:       3 * time.Nanosecond,
//			expected: "PT0.000000003S",
//		},
//		{
//			in:       20 * time.Nanosecond,
//			expected: "PT0.00000002S",
//		},
//		{
//			in:       7 * time.Millisecond,
//			expected: "PT0.007S",
//		},
//		{
//			in:       345 * time.Millisecond,
//			expected: "PT0.345S",
//		},
//		{
//			in:       10 * time.Second,
//			expected: "PT10S",
//		},
//		{
//			in:       3 * time.Minute,
//			expected: "PT3M",
//		},
//		{
//			in:       4 * time.Hour,
//			expected: "PT4H",
//		},
//		{
//			in:       44 * time.Hour,
//			expected: "PT44H",
//		},
//		{
//			in:       1*time.Hour + 2*time.Minute + 3*time.Second + 456*time.Microsecond,
//			expected: "PT1H2M3.456S",
//		}, // TODO: also support negative duration values, as defined in ISO8601-2
//	}
//
//	for _, tc := range testCases {
//		t.Run(tc.in.String(), func(t *testing.T) {
//			actual := duration.Format(tc.in)
//			assert.Equal(t, tc.expected, actual)
//		})
//	}
//}

func BenchmarkFormatQuick(b *testing.B) {
	x := 1*time.Hour + 2*time.Minute + 3*time.Second + 456*time.Microsecond
	for i := 0; i < b.N; i++ {
		_ = duration.FormatQuick(x)
	}
}
