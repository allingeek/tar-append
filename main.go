// Copyright 2017 Jeff Nickoloff. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.
package main

import (
	"archive/tar"
	"bufio"
	"io"
	"log"
	"os"
	path "path/filepath"
	"strings"
)

func main() {
	// buffered read tar from stdin into buffer
	// until reached EOF write buffer to stdout
	// write header for target file to stdout
	// read, buffer, and write target file to stdout

	log.SetOutput(os.Stderr)

	if len(os.Args) != 2 {
		log.Fatal(`requires one argument - relative name of the file to be included in the tar stream`)
	}
	fn := os.Args[1]

	// Learn something about the file
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	if path.IsAbs(fn) {
		fn, err = path.Rel(cwd, fn)
		if err != nil {
			log.Fatal(err)
		}
	}
	if strings.Index(fn, `..`) == 0 {
		log.Fatal(`The target file cannot be added from outside the context of the CWD`)
	}

	fi, err := os.Stat(fn)
	if err != nil {
		log.Fatal(err)
	}

	// Start pipeing the tar

	r := bufio.NewReader(os.Stdin)
	tin := tar.NewReader(r)
	tw := tar.NewWriter(os.Stdout)

	for {
		h, err := tin.Next()
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			log.Fatalln(err)
		}

		// Just a file, write it to stdout
		if err := tw.WriteHeader(h); err != nil {
			log.Fatalln(err)
		}
		if _, err := io.Copy(tw, tin); err != nil {
			log.Fatal(err)
		}
	}
	// Reached the end of the tar stream

	// Build the tar header
	head := &tar.Header{
		Name: fn,
		Mode: int64(fi.Mode()),
		Size: fi.Size(),
	}

	// Write the header
	if err = tw.WriteHeader(head); err != nil {
		log.Fatal(err)
	}

	// Open the target file, write a header, write the file, close everything
	fh, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
	}
	defer fh.Close()

	if _, err := io.Copy(tw, fh); err != nil {
		log.Fatal(err)
	}
	if err := tw.Close(); err != nil {
		log.Fatalln(err)
	}
}

