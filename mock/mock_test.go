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
	RunAndMeasureTime(t, buffer, mockSleeper, 0.0) // should NOT be any delays
	CompareStrings(t, buffer)
}

func TestRealCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	realSleeper := &RealSleeper{}
	RunAndMeasureTime(t, buffer, realSleeper, 3.0) // should be 1 second delays. in total 3 second

	CompareStrings(t, buffer)
}

func CompareStrings(t *testing.T, buffer *bytes.Buffer) {
	got := buffer.String()
	want := "3\n2\n1\nGo!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func RunAndMeasureTime(t *testing.T, buffer *bytes.Buffer, sleeper Sleeper, expectedRuntime float64) {
	start := time.Now()
	Countdown(buffer, sleeper)
	elapsed := time.Since(start)

	if math.Round(elapsed.Seconds()) != expectedRuntime {
		t.Error("the runtime of method is not correct, not equal to 3")
	}
}
