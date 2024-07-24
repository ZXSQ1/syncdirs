package channels

import "testing"

func TestClose(t *testing.T) {
	chanInt := make(chan int)

	go func() {
		chanInt <- 90
	}()

	go func() {
		print(<-chanInt)
	}()

	Close(chanInt)

	chanInt = nil

	Close(chanInt)
}
