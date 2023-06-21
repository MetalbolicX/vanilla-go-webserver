package middlewares

import (
	"log"
	"net/http"
	"time"

	"github.com/MetalbolicX/vanilla-go-webserver/types"
)

// The Logging logs the duration of each HTTP request.
// It captures the start time before invoking the handler
// function and calculates the duration after the handler
// function completes. The URL path and duration are
// then logged using the log.Println function.
func Logging() types.Middleware {
	return func(handlerLogic http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() {
				log.Println(r.URL.Path, time.Since(start))
			}()
			handlerLogic(w, r)
		}
	}
}
