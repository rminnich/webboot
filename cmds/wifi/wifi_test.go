package main

import (
	ui "github.com/gizak/termui/v3"
)

// TODO: figure out how to run tests for network.

func pressKey(ch chan ui.Event, input []string) {
	var key ui.Event
	for _, id := range input {
		key = ui.Event{
			Type: ui.KeyboardEvent,
			ID:   id,
		}
		ch <- key
	}
}

func stringToKeypress(str string) []string {
	var keyPresses []string
	for i := 0; i < len(str); i++ {
		keyPresses = append(keyPresses, str[i:i+1])
	}
	return keyPresses
}
