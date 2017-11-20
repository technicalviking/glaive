package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

type handler func(fName string) []byte

func (fn handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body := fn(r.URL.Path)

	w.Write(body)
}

func handle(fName string) []byte {

	fmt.Printf("wheee! %s  ", fName)
	joiner := string(os.PathSeparator)

	path := strings.Join([]string{"client", fName}, joiner)

	file, openErr := os.Open(path)

	if openErr != nil {
		return []byte(openErr.Error())
	}

	stat, statErr := file.Stat()

	if statErr != nil {
		return []byte(statErr.Error())
	}

	contents := make([]byte, stat.Size())

	_, readErr := file.Read(contents)

	if readErr != nil {
		return []byte(readErr.Error())
	}

	return contents
}

func main() {

	// *ServeMux
	http.Handle("/", handler(handle))
	http.ListenAndServe(":9001", nil)
}
