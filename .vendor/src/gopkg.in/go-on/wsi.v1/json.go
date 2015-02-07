package wsi

import (
	"bytes"
	"encoding/json"
	"net/http"
	"reflect"

	"gopkg.in/go-on/lib.v3/internal/meta"
)

// JSONStreamer streams a json array to an http.ResponseWriter
type JSONStreamer struct {
	w     http.ResponseWriter
	enc   *json.Encoder
	first bool
}

// MapViaJSON transforms an object to its map representation via
// json marshalling.
// MapViaJSON is not fast (doing a json encoding and decoding each time),
// but it is convenient and universal somehow
func MapViaJSON(v interface{}) (m map[string]interface{}, err error) {
	var bf bytes.Buffer
	err = json.NewEncoder(&bf).Encode(v)
	if err != nil {
		return nil, err
	}
	m = map[string]interface{}{}
	err = json.NewDecoder(&bf).Decode(&m)
	return
}

func MustMapViaJSON(v interface{}) map[string]interface{} {
	m, err := MapViaJSON(v)
	if err != nil {
		panic(err.Error())
	}
	return m
}

// MapSQL similar to MapViaJSON (but should be faster), but for sql tags. structPtr must be a pointer to a struct
func MapSQL(structPtr interface{}) (map[string]interface{}, error) {
	s, err := meta.StructByValue(reflect.ValueOf(structPtr))
	if err != nil {
		return nil, err
	}
	return s.ToMap("sql"), nil
}

func MustMapSQL(structPtr interface{}) map[string]interface{} {
	m, err := MapSQL(structPtr)
	if err != nil {
		panic(err.Error())
	}
	return m
}

// NewJSONStreamer returns a JSONStreamer for the given ResponseWriter and starts writing to it.
// The json content type is set and the opening bracket of the json array is written.
// The next step should be to call the Encode method for every json object that should be written
// and to call the Finish method at the end to write the closing bracket of the array.
func NewJSONStreamer(w http.ResponseWriter) (StreamEncoder, error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write([]byte("["))
	return &JSONStreamer{w, json.NewEncoder(w), true}, nil
}

// Encode writes an json object for the given value to the underlying ResponseWriter.
// Don't forget to call the Finish() method at the end.
func (j *JSONStreamer) Encode(v interface{}) error {
	if !j.first {
		j.w.Write([]byte(","))
	}
	j.first = false
	return j.enc.Encode(v)
}

// Finish writes the closing bracket of the array to the underlying ResponseWriter.
// Don't write to the underlying ResponseWriter after Finish has been run.
func (j *JSONStreamer) Finish() {
	j.w.Write([]byte("]"))
}

func ServeJSON(i interface{}, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(i)
}

type jsonDecoder struct{}

func (j jsonDecoder) Decode(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

var JSONDecoder RequestDecoder = jsonDecoder{}
