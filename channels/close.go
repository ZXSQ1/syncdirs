package channels

/*
description: closes the channel
arguments:
  - channel: the channel to close

return: no return
*/
func Close[T any](channel chan T) {
	if channel != nil {
		close(channel)
	}
}
