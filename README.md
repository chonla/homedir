HomeDir
======

![](https://github.com/chonla/homedir/workflows/Go/badge.svg)

Golang package that help you to create, get items in user's home directory.

# Example

```go
package main

import (
    "fmt"
    "log"
)

func main() {
    h, e := homedir.NewHomeDir("user")

    if e != nil {
        log.Fatalf("some error: %v", e)
    }

    fmt.Printf("user home directory is %s", h.Path())
    fmt.Printf("note.txt full path in home directory is %s", h.With("note.txt))
}
```

# License

This project is licensed under the terms of the MIT license.