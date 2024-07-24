package utils

/*
description: feeds the value into a channel
arguments:
  - channel: the channel to feed the value
  - val: the value to feed to the channel

return: no return
*/
func Feed[T any](channel chan T, val T) {
	if channel != nil {
		channel <- val
	}
}
