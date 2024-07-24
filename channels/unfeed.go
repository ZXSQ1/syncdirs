package channels

/*
description: unfeeds (gets) the value in the channel
arguments:
  - channel: the channel to get the value out of

return: the any type containing the value
*/
func Unfeed[T any](channel chan T) any {
	if channel != nil {
		if val, ok := <-channel; ok {
			return val
		}
	}

	return nil
}
