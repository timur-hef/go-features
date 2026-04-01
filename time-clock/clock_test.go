package clock

import (
	"math"
	"testing"
	"time"
)

func approxEq(a, b float64) bool {
	const eps = 1e-9
	return math.Abs(a-b) <= eps
}

func clockApproxEq(a, b Clock) bool {
	return approxEq(a.hour.X, b.hour.X) && approxEq(a.hour.Y, b.hour.Y) &&
		approxEq(a.minute.X, b.minute.X) && approxEq(a.minute.Y, b.minute.Y) &&
		approxEq(a.second.X, b.second.X) && approxEq(a.second.Y, b.second.Y)
}

func TestAnalogClockCases(t *testing.T) {
	tests := []struct {
		name string
		tm   time.Time
		want Clock
	}{
		{
			name: "полночь",
			tm:   time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC),
			want: Clock{
				hour:   Point{X: 150, Y: 100},
				minute: Point{X: 150, Y: 70},
				second: Point{X: 150, Y: 60},
			},
		},
		{
			name: "03:30",
			tm:   time.Date(1337, time.January, 1, 3, 30, 0, 0, time.UTC),
			want: Clock{
				hour:   Point{X: 198.29629131445341, Y: 162.94095225512604},
				minute: Point{X: 150, Y: 230},
				second: Point{X: 150, Y: 60},
			},
		},
		{
			name: "21:30 — 9:30 вечера",
			tm:   time.Date(1337, time.January, 1, 21, 30, 0, 0, time.UTC),
			want: Clock{
				hour:   Point{X: 101.70370868554659, Y: 137.05904774487396},
				minute: Point{X: 150, Y: 230},
				second: Point{X: 150, Y: 60},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnalogClock(tt.tm)
			if !clockApproxEq(got, tt.want) {
				t.Errorf("AnalogClock(%v)\ngot  %#v\nwant %#v", tt.tm, got, tt.want)
			}
		})
	}
}
