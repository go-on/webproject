package dispatch

import (
	"fmt"
	"testing"
)

func err(t *testing.T, msg string, is interface{}, shouldbe interface{}) {
	t.Errorf(msg+": is %s, should be %s\n", is, shouldbe)
}

type known int
type notRegistered int

var lastUnhandledContent interface{} = nil
var lastString = ""
var d = New()

func fallback(in interface{}, out interface{}) (didHandle bool, err error) {
	didHandle = true
	lastUnhandledContent = in
	return
}

func handleString(in interface{}, out interface{}) error {
	lastString = in.(string)
	return nil
}

func returnError(in interface{}, out interface{}) (didHandle bool, err error) {
	err = fmt.Errorf("test error")
	return
}

func TestDispatch(t *testing.T) {
	d.AddFallback(fallback)
	d.SetHandler("", handleString)
	d.Dispatch("hello world", "")

	if lastString != "hello world" {
		err(t, "Dispatch string", lastString, "hello world")
	}

	lastError := d.Dispatch(known(1), "")

	if lastError != nil {
		err(t, "Dispatch error with known", true, false)
	}

	if lastUnhandledContent != known(1) {
		err(t, "Dispatch error with unhandledContent", lastUnhandledContent, known(1))
	}
}

func TestFallbacks(t *testing.T) {
	res := []string{}
	fb1 := func(in interface{}, out interface{}) (didHandle bool, err error) {
		res = append(res, "fb1"+in.(string))
		return
	}
	fb2 := func(in interface{}, out interface{}) (didHandle bool, err error) {
		res = append(res, "fb2"+in.(string))
		didHandle = true
		return
	}
	fb3 := func(in interface{}, out interface{}) (didHandle bool, err error) {
		res = append(res, "fb3"+in.(string))
		didHandle = true
		return
	}

	disp := New()
	disp.AddFallback(fb3)
	disp.AddFallback(fb2)
	disp.AddFallback(fb1)

	disp.Dispatch("hiho", "")

	if res[0] != "fb1hiho" {
		err(t, "Fallback order 0", res[0], "fb1hiho")
	}

	if res[1] != "fb2hiho" {
		err(t, "Fallback order 1", res[1], "fb2hiho")
	}

	if len(res) > 2 {
		err(t, "Fallback order 2", res[2], "")
	}
}

func TestHasHandler(t *testing.T) {
	d1 := New()
	d1.SetHandler("", handleString)

	if !d1.HasHandler("") {
		err(t, "HasHandler", false, true)
	}

	if h := d1.GetHandler(""); h == nil {
		err(t, "HasHandler", h, handleString)
	}

	d1.RemoveHandler("")

	if d1.HasHandler("") {
		err(t, "RemoveHandler", true, false)
	}

}

func TestNoHandler(t *testing.T) {
	d1 := New()
	out := ""
	e := d1.Dispatch("", &out)
	if e == nil {
		err(t, "no handler error", false, true)
	}

	if e.Error() == "" {
		err(t, "no error message", false, true)
	}
}

func TestFallbackErr(t *testing.T) {
	d1 := New()
	d1.AddFallback(returnError)
	out := ""
	if e := d1.Dispatch("", &out); e == nil {
		err(t, "error of handler", false, true)
	}
}
