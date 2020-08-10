# baseN

`baseN` implements baseN encoding for integer numbers with given string to be used as base.
The code originate from mattheath/base62. 


## Usage

First `go get` the package:
```
go get github.com/marmile/baseN
```

This can then be used as below:
```go
package main

import (
    "fmt"
    "github.com/marmile/baseN"
)

func main() {

// Encoding 64 bit integers
var n int64 = 4815162342
var N string = "0123456789abcdefghijklmnop"
encoded := baseN.EncodeInt64(n,N)
fmt.Println(encoded)

```
