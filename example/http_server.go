package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/wjiec/greuseport"
)

func main() {
	addr := ":10898"
	if len(os.Args) > 1 {
		addr = os.Args[1]
	}

	listener, _ := greuseport.Listen("tcp", addr)
	defer func() {
		_ = listener.Close()
	}()

	server := &http.Server{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "gid: %d, pid: %d\n", os.Getgid(), os.Getpid())
	})

	fmt.Printf("gid: %d, pid: %d\n", os.Getgid(), os.Getpid())
	log.Printf("server running at %s\n", addr)
	log.Fatal(server.Serve(listener))
}
