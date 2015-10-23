//
// Sets up the URL routing and starts the web server
//
// Copyright (c) 2015 Jon Carlson.  All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.
//
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Taken from http://blog.golang.org/error-handling-and-go
// errorHandler adds a ServeHttp method to every errorHandler function
type errorHandler func(http.ResponseWriter, *http.Request) error

// Adds a ServeHttp method to every errorHandler function
func (fn errorHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if err := fn(w, req); err != nil {
		log.Println("Error handling request", err)
		http.Error(w, "Internal Server Error.  Check logs for actual error", 500)
	}
}

// errorableHandler converts an http.Handler into an errorHandler function
//func errorableHandler(handler http.Handler) func(w http.ResponseWriter, req *http.Request) error {
func errorableHandler(handler http.Handler) errorHandler {
	return func(w http.ResponseWriter, req *http.Request) error {
		handler.ServeHTTP(w, req)
		return nil
	}
}

// authBasic wraps a request handler (that returns an error: AKA errorHandler)
// with another one that requires BASIC HTTP authentication
func authBasic(handler errorHandler) errorHandler {
	return func(w http.ResponseWriter, req *http.Request) error {

		if insecure {
			return handler(w, req)
		}

		//
		// Ensure request has an "Authorization" header
		// (needed for "Basic" authentication)
		//
		username, _, ok := req.BasicAuth()
		if !ok {
			// Request the "Authorization" header
			w.Header().Set("WWW-Authenticate", `Basic realm="go-example-web"`)
			http.Error(w, "Access Denied", http.StatusUnauthorized)
			return nil
		}

		if username != secretUsername {
			// User authentication failed
			w.Header().Set("WWW-Authenticate", `Basic realm="Enter token for username"`)
			http.Error(w, "Access Denied", http.StatusUnauthorized)
			return nil
		}

		//
		// The credentials match, so run the wrapped handler
		//
		return handler(w, req)
	}
}

// redirectToHttps returns a function that redirects anything on the http
// port over to the https port. We have to wrap the function in a function
// so we can dynamically provide the http and https ports.
func redirectToHttps(httpPort int, httpsPort int) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		newHost := strings.Replace(req.Host, strconv.Itoa(httpPort), strconv.Itoa(httpsPort), 1)
		newUrl := fmt.Sprintf("https://%s/%s", newHost, req.RequestURI)
		http.Redirect(w, req, newUrl, http.StatusMovedPermanently)
	}
}
