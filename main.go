package main

import (
	"flag"
	"net/http"
	"strconv"
)

// file  server port
var port = flag.Int64("port", 1234, "file server port")
var file_path = flag.String("path", "./", "file server path")

func init() {
	flag.Parse()
}

func main() {
	http.Handle("/", http.FileServer(http.Dir(*file_path)))

	addr := ":" + strconv.FormatInt(*port, 10)

	println(http.ListenAndServe(addr, nil))

}
