// Package event is a simple event system.
package event

import (
	"errors"
	"reflect"
	"sync"
)

// Event implementation
type event struct {
	sync.RWMutex

	events map[string]interface{}
}

// Dispatcher event interface
type Dispatcher interface {
	On(name string, fn interface{}) error
	Do(name string, params ...interface{}) error
	Has(name string) bool
	Remove(names string)
}

// New returns a new event.Dispatcher
func New() Dispatcher {
	return &event{
		events: make(map[string]interface{}),
	}
}

// On set new listener
func (e *event) On(name string, fn interface{}) error {
	e.Lock()
	defer e.Unlock()

	if fn == nil {
		return errors.New("fn is nil")
	}

	t := reflect.TypeOf(fn)
	if t.Kind() != reflect.Func {
		return errors.New("fn is not a function")
	}
	if t.NumOut() != 1 {
		return errors.New("fn must have one return value")
	}
	if t.Out(0) != reflect.TypeOf((*error)(nil)).Elem() {
		return errors.New("fn must return an error message")
	}
	if f, ok := e.events[name]; ok {
		tt := reflect.TypeOf(f)
		if tt.NumIn() != t.NumIn() {
			return errors.New("fn signature is not equal")
		}
		for i := 0; i < tt.NumIn(); i++ {
			if tt.In(i) != t.In(i) {
				return errors.New("fn signature is not equal")
			}
		}
	}

	e.events[name] = fn
	return nil
}

// Do firing an event
func (e *event) Do(name string, params ...interface{}) error {
	e.RLock()
	defer e.RUnlock()

	fn := e.events[name]

	err := e.call(fn, params...)
	if err != nil {
		return err
	}

	return nil
}

func (e *event) call(fn interface{}, params ...interface{}) (err error) {
	var (
		f     = reflect.ValueOf(fn)
		t     = f.Type()
		numIn = t.NumIn()
		in    = make([]reflect.Value, 0, numIn)
	)

	if t.IsVariadic() {
		n := numIn - 1
		if len(params) < n {
			return errors.New("parameters mismatched")
		}
		for _, param := range params[:n] {
			in = append(in, reflect.ValueOf(param))
		}
		s := reflect.MakeSlice(t.In(n), 0, len(params[n:]))
		for _, param := range params[n:] {
			s = reflect.Append(s, reflect.ValueOf(param))
		}
		in = append(in, s)

		err, _ = f.CallSlice(in)[0].Interface().(error)
		return err
	}

	if len(params) != numIn {
		return errors.New("parameters mismatched")
	}
	for _, param := range params {
		in = append(in, reflect.ValueOf(param))
	}

	err, _ = f.Call(in)[0].Interface().(error)
	return err
}

// Has returns true if a event exists
func (e *event) Has(name string) bool {
	e.RLock()
	defer e.RUnlock()
	_, ok := e.events[name]
	return ok
}

// Remove delete events from the event list
func (e *event) Remove(name string) {
	e.Lock()
	defer e.Unlock()

	delete(e.events, name)
}
