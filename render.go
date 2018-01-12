package gorender

import (
	"encoding/json"
	"log"
	"net/http"
)

const (
	Charset                        = "charset=UTF-8"
	HeaderContentType              = "Content-Type"
	MIMEApplicationJSON            = "application/json"
	MIMEApplicationJSONCharsetUTF8 = MIMEApplicationJSON + "; " + Charset
)

type Render struct {
	opts *Options
}

type Options struct{}

func New(opts *Options) *Render {
	r := &Render{
		opts: opts,
	}

	return r
}

func (r *Render) JSON(w http.ResponseWriter, code int, data interface{}) {
	b, err := json.Marshal(data)
	if err != nil {
		r.ServerError(w, err)
		return
	}
	r.JSONBlob(w, code, b)
	return
}

func (r *Render) JSONBlob(w http.ResponseWriter, code int, b []byte) {
	r.Blob(w, code, MIMEApplicationJSONCharsetUTF8, b)
}

func (r *Render) Blob(w http.ResponseWriter, code int, contentType string, b []byte) {
	w.Header().Set(HeaderContentType, contentType)
	w.WriteHeader(code)
	w.Write(b)
}

func (r *Render) ServerError(w http.ResponseWriter, err error) {
	log.Printf("Error occured while encoding to JSON: %v\n", err)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
