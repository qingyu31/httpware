package httpware

import "net/http"

//NewRecoverWare create a new midware process user defined function after catch panic.
func NewRecoverWare(function func(http.ResponseWriter, *http.Request, interface{})) midware {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if rec := recover(); rec != nil {
					function(w, r, rec)
				}
			}()
		})

	}
}
