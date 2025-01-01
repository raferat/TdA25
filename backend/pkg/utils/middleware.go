package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/andybalholm/brotli"
	"github.com/fatih/color"
)

//========================================================  Logging  ========================================================

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrappedWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapped := &wrappedWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}
		next.ServeHTTP(wrapped, r)

		var status *color.Color
		if wrapped.statusCode < 400 {
			status = color.New(color.FgGreen, color.Bold)
		} else if wrapped.statusCode < 500 {
			status = color.New(color.FgMagenta, color.Bold)
		} else {
			status = color.New(color.FgHiRed, color.Bold)
		}

		fmt.Printf("%s | %s | %s | %s | %s\n",
			datetimeFormat(time.Now()),
			status.Sprint(wrapped.statusCode),
			r.Method,
			r.URL.Path,
			time.Since(start),
		)
	})
}

//======================================================= JSON Encode =======================================================

type JSONReturnFunction func(r *http.Request) (code int, object any)

func JSONEncodeMiddleware(next JSONReturnFunction) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		code, obj := next(r)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(code)
		if obj != nil {
			enc := json.NewEncoder(w)
			enc.Encode(obj)
		}
	})
}

//======================================================= Compression =======================================================

type redirectedWriter struct {
	w http.ResponseWriter
	r io.Writer
}

func (r *redirectedWriter) Header() http.Header {
	return r.w.Header()
}

func (r *redirectedWriter) WriteHeader(code int) {
	r.w.WriteHeader(code)
}

func (r *redirectedWriter) Write(p []byte) (int, error) {
	n, err := r.r.Write(p)
	return n, err
}

func NewRedirectedWriter(parent http.ResponseWriter, redirect io.Writer) http.ResponseWriter {
	return &redirectedWriter{
		w: parent,
		r: redirect,
	}
}

func CompressionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !canCompress(r) {
			next.ServeHTTP(w, r)
			return
		}

		w.Header().Add("Content-Encoding", "br")
		encoder := brotli.NewWriterLevel(w, 5)
		rw := NewRedirectedWriter(w, encoder)
		next.ServeHTTP(rw, r)
		encoder.Flush()
		encoder.Close()
	})
}

func canCompress(r *http.Request) bool {
	if val, ok := r.Header["Accept-Encoding"]; ok {
		for _, v := range val {
			if strings.Contains(v, "br") {
				return true
			}
		}
		fmt.Printf("%#v\n", val[0])
	}
	return false
}
