package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

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
