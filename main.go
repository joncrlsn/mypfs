//
// Copyright (c) 2015 Jon Carlson.  All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.
//

package main

//
// Runs an HTTP static file server and file upload receiver from/to
// the directory that this is executed from.
//

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var (
	action         string = "help"
	port           int    = 8080
	timeoutMinutes int64  = 10
	version        string = "0.9"
)

// uploadHandler returns an HTML upload form
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintf(w, `<html>
<head>
  <title>GoLang HTTP Fileserver</title>
</head>

<body>

<h4>Choose a file to upload</h4>

<form action="/fs-receive" method="post" enctype="multipart/form-data">
  <input type="file" name="file" id="file">
  <br> <br>
  <input type="submit" name="submit" value="Submit">
</form>

</body>
</html>`)
	}
}

// receiveHandler accepts the file and saves it to the current working directory
func receiveHandler(w http.ResponseWriter, r *http.Request) {

	// the FormFile function takes in the POST input id file
	file, header, err := r.FormFile("file")

	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	defer file.Close()

	out, err := os.Create(header.Filename)
	if err != nil {
		fmt.Fprintf(w, "Unable to create the file for writing. Check your write access privilege")
		return
	}

	defer out.Close()

	// write the content from POST to the file
	_, err = io.Copy(out, file)
	if err != nil {
		fmt.Fprintln(w, err)
	}

	fmt.Fprintf(w, `<html>
File uploaded successfully: %s 
<p><a href="/">Back</a></p>
<html>`, header.Filename)
}

func main() {
	err := myPfsCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

	var portStr = ":" + strconv.Itoa(port)

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("err=", err)
		os.Exit(1)
	}

	if action == "help" {
		os.Exit(0)
	} else if action == "version" {
		os.Exit(0)
	} else if action == "up" {
		log.Printf("Allowing uploads to the current directory for %v minutes on port %v\n", timeoutMinutes, port)

		// Show the upload form
		http.HandleFunc("/", uploadHandler)
		// Handle the incoming file
		http.HandleFunc("/fs-receive", receiveHandler)

	} else if action == "down" {
		log.Printf("Allowing downloads from the current directory for %v minutes on port %v\n", timeoutMinutes, port)

		// Show the download page using the standard http FileServer
		http.Handle("/", http.FileServer(http.Dir(dir)))

	} else if action == "up/down" {
		log.Printf("Allowing downloads from (and uploads to) the current directory for %v minutes on port %v\n", timeoutMinutes, port)

		// Display the upload form
		http.HandleFunc("/fs-upload", uploadHandler)
		// Handle the incoming file
		http.HandleFunc("/fs-receive", receiveHandler)

		// Show the download page using a customized FileServer
		// copied from net/http/fs.go. This version adds a header
		// to the top when we list a directory (in dirList() func)
		http.Handle("/", FileServer(Dir(dir)))
	}

	go func() {
		time.Sleep(time.Duration(timeoutMinutes) * time.Minute)
		log.Println("Fileserver timed out.  Exiting.")
		os.Exit(0)
	}()

	log.Fatal(http.ListenAndServe(portStr, nil))
}
