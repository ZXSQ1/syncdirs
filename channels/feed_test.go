package channels

import (
	"fmt"
	"testing"
)

func TestFeed(t *testing.T) {
	intChan := make(chan int)

	go Feed(intChan, 90)

	fmt.Println(<-intChan)

	go Feed(nil, 90)
}
