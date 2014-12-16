package dispatch

import (
	"fmt"
	"reflect"
)

type TypeHandler func(in interface{}, out interface{}) error

type Dispatcher struct {
	handlers  map[string]TypeHandler
	Fallbacks []func(in interface{}, out interface{}) (didHandle bool, err error)
}

type NotHandled struct {
	Value interface{}
	Type  string
}

func (ø NotHandled) Error() string {
	return fmt.Sprintf("not handled: value %#v type %s", ø.Value, ø.Type)
}

func New() (ø *Dispatcher) {
	return &Dispatcher{
		handlers:  map[string]TypeHandler{},
		Fallbacks: []func(in interface{}, out interface{}) (didHandle bool, err error){}}
}

/*
	A fallback function is called by Dispatch(), if no type handler could be found for a certain type of the given value.
	A fallback functions is expected to return a boolean to indicate, if it did handle the given value and/or an error if some error occured.
	A the fallback returns false, the next (LIFO) fallback function will be called from Dispatch() (see there).
*/
func (ø *Dispatcher) AddFallback(fb func(in interface{}, out interface{}) (didHandle bool, err error)) {
	ø.Fallbacks = append(ø.Fallbacks, fb)
}

/*
Set the handler for a specific type

Only one handler may be specified for a certain type.

You may check if a handler already exists with HasHandler()

You may also get the handler with GetHandler()

If the type is unknown for the registry, an error is returned
*/
func (ø *Dispatcher) SetHandler(i interface{}, f func(interface{}, interface{}) error) {
	ø.handlers[ø.Type(i)] = TypeHandler(f)
}

// removes all fallback functions
func (ø *Dispatcher) HasHandler(i interface{}) bool                  { return ø.handlers[ø.Type(i)] != nil }
func (ø *Dispatcher) GetHandler(i interface{}) (handler TypeHandler) { return ø.handlers[ø.Type(i)] }
func (ø *Dispatcher) RemoveHandler(i interface{})                    { delete(ø.handlers, ø.Type(i)) }
func (ø *Dispatcher) Type(i interface{}) string                      { return reflect.TypeOf(i).String() }

/*
Dispatch() takes any value, looks if it can find a handler function to handle its type and calls that handler with the value.

If there is no specific handler for the type of value, the fallback functions are called in the reverse order of their registration (LIFO) until one of them
either returns an error or a boolean with true that means, that the function did handle the value.

Dispatch() returns an error if one of the following conditions are met:

	- the handler function returns an error. the error is passed through
	- there is no handler function for the type of the value and there also is
	  no fallback function that could/did handle the value. fix it with SetHandler() or AddFallback()
	- a fallback function returned an error. the error is passed through
*/
func (ø *Dispatcher) Dispatch(in interface{}, out interface{}) error {
	tt := ø.Type(in)
	m := ø.handlers[tt]

	if m == nil {
		didHandle := false
		lenfb := len(ø.Fallbacks)
		for i := lenfb - 1; i > -1; i-- {
			didHandle, e := ø.Fallbacks[i](in, out)
			if e != nil {
				return e
			}
			if didHandle {
				return nil
			}
		}
		if !didHandle {
			return NotHandled{in, tt}
		}
	}
	return m(in, out)
}
