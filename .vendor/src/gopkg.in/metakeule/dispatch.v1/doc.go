// Copyright 2013 Marc René Arns. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package dispatch provides an flexible way to structure dispatching based on types.

Advantages over using a type switch:

	- fallback methods for unhandled types
	- add and remove types from the outside
	- add, remove and overwrite type handling functions from the outside
	- add fallback functions from the outside

Example

Let's say you have a struct with a method that uses a type switch

	type MyStruct {}

	func (ø *MyStruct) DoStuff(i interface{}){
		switch t := i.(type) {
		case int:
			ø.doIntStuff(v)
		// ...add some more
		default:
			ø.fallback(t)
		}
	}

	func (ø *MyStruct) doIntStuff(i int){}
	func (ø *MyStruct) fallback(i interface{}){}

If you want to handle a new type
  - you will need to change your DoStuff method.
  - If it is from another package, you will need to fork the code and change it.
  - your changes will affect all other uses of the DoStuff method


Using dispatch you have more flexibility:

	import (
		"fmt"
		"github.com/metakeule/dispatch"
	)

	type MyStruct { *Dispatcher}

	func NewMyStruct() *MyStruct {
		ø := &MyStruct{dispatcher.New()}
		ø.AddType(0)
		ø.SetHandler("int", doIntStuff)
		ø.AddFallback(fallback)
	}

	func (ø *MyStruct) DoStuff(i interface{}, out interface{}) (err error) {
		fmt.Printf("before dispatching: %v", i)
		err = ø.Dispatch(i,"")
		return
	}

	func doIntStuff(value interface{}, out interface{}) (err error)
		// you have to typecast, but you can be sure the  type cast will work
		i := value.(int)
		fmt.Printf("is an int: %v\n", i)
		return
	}

	func fallback(value interface{}, out interface{}) (didHandle bool, err error){
		didHandle = true // return didHandle = true if the fallback did handle the value
		fmt.Printf("in the fallback: %v\n", value)
		return
	}


Now if MyStruct is used, new types, handlers and fallbacks can be added. No need to change anything inside MyStruct.

	type Special int

	func specialHandler(value interface{}, out interface{}) (didHandle bool, err error) {
		s := value.(Special) // no need for error handling
		fmt.Printf("is a special int: %v\n", s)
		return
	}

	// overwrite int handling
	func myIntHandler(value interface{}, out interface{}) (didHandle bool, err error) {
		i := value.(int)
		fmt.Printf("is an int in the overwritten handler:: %v\n", i)
		return
	}

	func main() {
		my := NewMyStruct()
		my.AddType(Special(0))
		my.SetHandler("Special",specialHandler)
		my.SetHandler("int", myStringHandler)
		my.DoStuff(2)
		my.DoStuff(Special(2))
	}

If you use more than one fallback, be aware, that they are called in the reverse order of registration (last in first out).
Why is that?

If you want to add a fallback from the outside, you might want to intercept the fallback chain and decide,
if the other fallbacks are called by returning true or false.

Why does Dispatch() return an error instead of a centralized error handling?

Well if you need a centralized error handling,
simply wrap the call to Dispatch() into another function that handles the error. But you also could do the error handling
when calling the Dispatch() where you have more information about the value, you pass to Dispatch and where you can make
better decisions about what to do best.
*/
package dispatch
