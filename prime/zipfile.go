package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	rfile = flag.String("rfile", "", "read zip file path")
	wfile = flag.String("wfile", "", "write zip file path")
	url   = flag.String("url", "http://primes.utm.edu/lists/small/millions/primes1.zip", "get a zip file from this url")
)

func zipread(file string) {
	// Open a zip archive for reading.
	r, err := zip.OpenReader(file)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	// Iterate through the files in the archive,
	// printing some of their contents.
	for _, f := range r.File {
		fmt.Printf("Contents of %s:\n", f.Name)
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.CopyN(os.Stdout, rc, 68)
		if err != nil {
			log.Fatal(err)
		}
		rc.Close()
		fmt.Println()
	}
}

func zipget(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Get url error", err)
		return nil, err
	}
	defer resp.Body.Close()

	log.Println(resp)

	zipfile, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Read file error")
		return nil, err
	}

	return zipfile, nil
}

func zipsave(url string, saveto string) (err error) {
	file, err := zipget(url)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(saveto, file, os.ModePerm)
	return err
}

func main() {
	flag.Parse()

	i := strings.LastIndex(*url, "/")
	if i == -1 {
		return
	}
	file := fmt.Sprintf("%s%s", ".", string([]byte(*url)[i:]))
	zipsave(*url, file)
	zipread(file)
}
