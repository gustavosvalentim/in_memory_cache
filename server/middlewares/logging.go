package middlewares

import (
	"fmt"
	"bytes"
	"io/ioutil"
	"net/http"
)

// Logging : 
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)

		if (err != nil) {
			fmt.Println(err.Error())
		}

		fmt.Println(r.Method, r.URL.Path)
		fmt.Println(string(body))

		/*
		 *	bytes.NewReader converts []byte to io.Reader
		 *	ioutil.NopCloser will conver io.Reader to io.ReaderCloser as required for r.Body
		 *	
		 *	This is needed because ioutil.ReadAll reads r.Body and return the bytes,
		 *	after reading, r.Body closes and can't be read anymore.
		 */
		r.Body = ioutil.NopCloser(bytes.NewReader(body))

		next.ServeHTTP(w, r)
	})
}