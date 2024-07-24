package channels

import "testing"

func TestUnfeed(t *testing.T) {
	intChan := make(chan int)

	go func() {
		intChan <- 90
	}()

	if Unfeed(intChan) != 90 {
		t.Fail()
	}
}
