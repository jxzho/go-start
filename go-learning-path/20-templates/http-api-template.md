# HTTP API Starter 模板

目录结构：

```
my-api/
  cmd/
    server/
      main.go
  internal/
    transport/
      http.go
  go.mod
```

`cmd/server/main.go`
```go
package main

import "my-api/internal/transport"

func main() {
	transport.StartHTTP()
}
```

`internal/transport/http.go`
```go
package transport

import (
	"encoding/json"
	"net/http"
)

func StartHTTP() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})
	_ = http.ListenAndServe(":8080", nil)
}
```
