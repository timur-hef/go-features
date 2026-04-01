package clock

import (
	"math"
	"time"
)

// every clock has a centre of (150, 150)
// the hour hand is 50 long
// the minute hand is 80 long
// the second hand is 90 long.
const (
	hourRadius      = 50
	minuteRadius    = 80
	secondRadius    = 90
	RadianPerHour   = math.Pi / 6
	RadianPerMinute = math.Pi / 30
	RadianPerSecond = math.Pi / 30
	CenterX         = 150
	CenterY         = 150
)

// A Point represents a two-dimensional Cartesian coordinate
type Point struct {
	X float64
	Y float64
}

type Clock struct {
	hour   Point
	minute Point
	second Point
}

// SecondHand is the unit vector of the second hand of an analogue clock at time `t`
// represented as a Point.

func AnalogClock(t time.Time) Clock {
	minute := float64(t.Minute())
	second := float64(t.Second())
	hour := float64(t.Hour()%12) + minute/60

	hourHand := Point{
		CenterX + hourRadius*math.Sin(hour*RadianPerHour),
		CenterY - hourRadius*math.Cos(hour*RadianPerHour),
	}
	minuteHand := Point{
		CenterX + minuteRadius*math.Sin(minute*RadianPerMinute),
		CenterY - minuteRadius*math.Cos(minute*RadianPerMinute),
	}
	secondHand := Point{
		CenterX + secondRadius*math.Sin(second*RadianPerSecond),
		CenterY - secondRadius*math.Cos(second*RadianPerSecond),
	}

	return Clock{
		hour:   hourHand,
		minute: minuteHand,
		second: secondHand,
	}
}
