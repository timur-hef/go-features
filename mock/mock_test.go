package mock

import (
	"bytes"
	"math"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	mockSleeper := &MockSleeper{}
	elapsed := RunAndMeasureTime(buffer, mockSleeper) // should NOT be any delays

	if math.Round(elapsed) != 0.0 {
		t.Error("the runtime of method is not correct, not equal to 0")
	}

	CompareStrings(t, buffer)
}

func TestRealCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	realSleeper := &RealSleeper{}
	elapsed := RunAndMeasureTime(buffer, realSleeper) // should be 1 second delays. in total 3 second

	if math.Round(elapsed) != 3.0 {
		t.Error("the runtime of method is not correct, not equal to 3")
	}

	CompareStrings(t, buffer)
}

func CompareStrings(t *testing.T, buffer *bytes.Buffer) {
	got := buffer.String()
	want := "3\n2\n1\nGo!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func RunAndMeasureTime(buffer *bytes.Buffer, sleeper Sleeper) float64 {
	start := time.Now()
	Countdown(buffer, sleeper)
	elapsed := time.Since(start)

	return elapsed.Seconds()
}
