package mock

import (
	"bytes"
	"math"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}

	start := time.Now()
	mockSleeper := &MockSleeper{}
	Countdown(buffer, mockSleeper) // should NOT be any delays
	elapsed := time.Since(start)

	got := buffer.String()
	want := "3\n2\n1\nGo!"

	if math.Round(elapsed.Seconds()) != 0.0 {
		t.Error("the runtime of method is not correct, not equal to 0")
	}

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestRealCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}

	start := time.Now()
	realSleeper := &RealSleeper{} // should be 1 second delays. in total 3 second
	Countdown(buffer, realSleeper)
	elapsed := time.Since(start)

	if math.Round(elapsed.Seconds()) != 3.0 {
		t.Error("the runtime of method is not correct, not equal to 3")
	}

	got := buffer.String()
	want := "3\n2\n1\nGo!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
