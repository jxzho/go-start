# CLI Starter 模板

目录结构：

```
my-cli/
  cmd/
    main.go
  internal/
    app/
      app.go
  go.mod
```

`cmd/main.go`
```go
package main

import "my-cli/internal/app"

func main() {
	app.Run()
}
```

`internal/app/app.go`
```go
package app

import (
	"fmt"
	"os"
)

func Run() {
	name := "World"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	fmt.Printf("Hello, %s\n", name)
}
```
