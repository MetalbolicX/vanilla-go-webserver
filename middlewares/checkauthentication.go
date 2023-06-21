package middlewares

import (
	"fmt"
	"net/http"

	"github.com/MetalbolicX/vanilla-go-webserver/types"
)

//	The CheckAuth performs authentication checks.If
//
// the authentication check fails, it aborts the
// request and returns an "Unauthorized" response.
// If the authentication check succeeds,
// it allows the execution to continue to the
// subsequent handler function.
func CheckAuth() types.Middleware {
	return func(handlerLogic http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			flag := true
			if !flag {
				fmt.Println("Unauthorized")
				return
			}
			handlerLogic(w, r)
		}
	}
}
