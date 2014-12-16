package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/go-on/queue.v2"
	"gopkg.in/go-on/wrap.v2"
	"gopkg.in/go-on/wrap-contrib-testing.v2/wrapstesting"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type errorHandler struct {
	logger *log.Logger
}

func (eh errorHandler) log(r *http.Request, err error) {
	eh.logger.Printf("Error %T in %s %s: %s",
		err,
		r.Method,
		r.URL.String(),
		err.Error(),
	)
}

func (eh errorHandler) WriteError(w http.ResponseWriter, r *http.Request, err error) {
	eh.log(r, err)

	switch v := err.(type) {
	case wrapstesting.HTTPStatusError:
		// write all headers but the X- ones
		for k, _ := range v.Header {
			if !strings.HasPrefix(strings.ToLower(k), "x-") {
				w.Header().Set(k, v.Header.Get(k))
			}
		}
		switch v.Code {
		case 301, 302:
			w.WriteHeader(v.Code)
		case 404:
			w.WriteHeader(404)
			// TODO: switch content type
			w.Write([]byte("beautiful 404 message"))
		default:
			w.WriteHeader(v.Code)
			if v.Code < 500 {
				w.Write([]byte(err.Error()))
			}
		}
	case *strconv.NumError:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "not an int: %#v", v.Num)
	case wrapstesting.ValidationError:
		m, _ := json.Marshal(v.ValidationErrors())
		w.WriteHeader(http.StatusBadRequest)
		w.Write(m)
	default:
		w.WriteHeader(500)
		w.Write([]byte("Internal Server Error"))
	}
}

func HandleError(w http.ResponseWriter) *queue.Queue {
	return queue.New().OnError(w.(queue.ErrHandler))
}

func repeatInt(w http.ResponseWriter, r *http.Request) {
	HandleError(w).
		Add(strconv.Atoi, r.URL.Query().Get("int")).
		Add(fmt.Fprintf, w, "%d%d", queue.PIPE, queue.PIPE).
		Run()
}

func main() {
	stack := wrap.New(
		wrapstesting.NewErrorWrapper(errorHandler{
			log.New(os.Stdout, "repeatInt ", log.Ltime|log.Ldate|log.Lshortfile),
		}),
		wrap.HandlerFunc(repeatInt),
	)

	http.ListenAndServe(":8085", stack)
}
