// Fork from gin
// https://github.com/gin-gonic/gin/blob/a889c58de78711cb9b53de6cfcc9272c8518c729/response_writer.go

package middleware

import (
	"net/http"
)

const noWrittenStatus = -1

// ResponseWriter ...
type ResponseWriter interface {
	http.ResponseWriter
	// Status returns the HTTP response status code of the current request.
	Status() int

	// Size returns the number of bytes already written into the response http body.
	// See Written()
	Size() int
}

type responseWriter struct {
	http.ResponseWriter
	size   int
	status int
}

func newResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{
		ResponseWriter: w,
		size:           0,
		status:         noWrittenStatus,
	}
}

var _ ResponseWriter = (*responseWriter)(nil)

func (w *responseWriter) WriteHeader(code int) {
	w.ResponseWriter.WriteHeader(code)

	if w.status == noWrittenStatus {
		w.status = code
	}
}

func (w *responseWriter) Write(data []byte) (n int, err error) {
	n, err = w.ResponseWriter.Write(data)
	w.size += n
	return
}

func (w *responseWriter) Status() int {
	return w.status
}

func (w *responseWriter) Size() int {
	return w.size
}
