package mock

import (
	"fmt"
	"io"
	"time"
)

type Sleeper interface {
	Sleep()
}

type RealSleeper struct{}

func (d *RealSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

type MockSleeper struct {
	Calls int
}

func (s *MockSleeper) Sleep() {
	s.Calls++
}

// we accept interface, instead of concrete object, so we can mock in tests
func Countdown(out io.Writer, sleeper Sleeper) {

	for i := 3; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep() // need to mock in tests, otherwise the tests will be too long!
	}

	fmt.Fprintf(out, "Go!")
}
