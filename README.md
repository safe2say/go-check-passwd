# go-check-passwd

A tool to see if password is commonly used or not

## Install

go get github.com/ninja-software/go-check-passwd

## Example

```go
package main

import (
    "fmt"
    "time"

    "github.com/ninja-software/go-check-passwd"
)

func main() {
    pass := "123456"

    if gocheckpasswd.IsCommon(pass) {}
        fmt.Printf("%s is a commonly used password\n", pass)
    else {
        fmt.Printf("%s is not a commonly used password\n", pass)
    }
}
```
