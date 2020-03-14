# greuseport
SO_REUSEADDR and SO_REUSEPORT in golang


### Installation
```shell
go get -u -v github.com/wjiec/greuseport
```

or use go module
```shell
require github.com/wjiec/greuseport latest
```

### Quick Start
```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/wjiec/greuseport"
)

func main() {
	listener, _ := greuseport.Listen("tcp", "8898")
	defer func() {
		_ = listener.Close()
	}()

	server := &http.Server{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "gid: %d, pid: %d\n", os.Getgid(), os.Getpid())
	})

	log.Fatal(server.Serve(listener))
}

```
