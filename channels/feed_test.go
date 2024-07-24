package channels

import (
	"testing"
)

func TestFeed(t *testing.T) {
	intChan := make(chan int)

	Feed(intChan, 90)

	intChan = nil

	Feed(intChan, 90)
}
