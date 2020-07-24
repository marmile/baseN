# baseN

`baseN` implements baseN encoding for integer numbers of arbitrary precision with given string to be used as base.
The code originate from mattheath/base62. 

Currently both `int64` and `*big.Int` are supported


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
    "math/big"
    "github.com/marmile/baseN"
)

func main() {
	// Encoding 64 bit integers
	var n int64 = 4815162342
  var N string = "0123456789abcdefghijklmnop"
    encoded := baseN.EncodeInt64(n,N)
    fmt.Println(encoded)

    // Arbitrary precision integers can be specified using the math/big pkg
    var b *big.Int = new(big.Int)
    b.SetString("340282366920938463463374607431768211455") // 128bit unsigned int
    bigEncoded := baseN.EncodeBigInt(b, N)
    fmt.Println(bigEncoded) // prints 7n42DGM5Tflk9n8mt7Fhc7

    // Padding can be specified using an option on an Encoding
    // eg. to pad strings to a minimum length of 15 chars:
    e := baseN.NewStdEncoding().Option(Padding(15))
    encoded = e.EncodeInt64(n)
    fmt.Println(encoded) 
}
```
