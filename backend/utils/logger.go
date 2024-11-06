package utils

import (
	"fmt"
	"net/http"
	"time"

	"github.com/fatih/color"
)

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrappedWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func datetimeFormat(t time.Time) string {
	year, month, day := t.Date()
	hour, minute, second := t.Clock()

	return fmt.Sprintf("%04d/%02d/%02d %02d:%02d:%02d", year, month, day, hour, minute, second)
}

func Logging(next http.Handler) http.Handler {
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
