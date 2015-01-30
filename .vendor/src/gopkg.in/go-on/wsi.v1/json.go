package wsi

import (
	"encoding/json"
	"net/http"
)

// JSONStreamer streams a json array to an http.ResponseWriter
type JSONStreamer struct {
	w     http.ResponseWriter
	enc   *json.Encoder
	first bool
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
