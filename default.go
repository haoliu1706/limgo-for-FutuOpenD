package limgo

import (
	"fmt"
	"limgo/event"
)

// handlersID map
var handlersID = make(map[uint32]string)

// Default event instance
var handlers = event.New()

// On set new listener
func On(name string, fn interface{}) error {
	return handlers.On(name, fn)
}

// Do firing an event
func Do(name string, params ...interface{}) error {
	if !HasHandler(name) {
		fmt.Println(name, " handler func is not exist")
	}

	return handlers.Do(name, params...)
}

// HasHandler returns true if a event exists
func HasHandler(name string) bool {
	return handlers.Has(name)
}

// SetHandlerID to handlersID
func SetHandlerID(id uint32, name string) {
	handlersID[id] = name
}

// TransHandlerID from handlersID
func TransHandlerID(id uint32) (string, bool) {
	str, ok := handlersID[id]

	return str, ok
}
