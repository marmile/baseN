//Package baseN implements baseN encoding for integer numbers with given string to be used as base. 
package baseN

import (
	"fmt"
	_ "math"
	_ "math/big"
	"strconv"
	_ "strings"
)
var base int64 = 62

type Encoding struct {
	encode  string
	padding int
}

// Option sets a number of optional parameters on the encoding
func (e *Encoding) Option(opts ...option) *Encoding {
	for _, opt := range opts {
		opt(e)
	}

	// Return the encoding to allow chaining
	return e
}
//EncodeStd set by user chars 
var EncodeStd = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// NewEncoding returns a new Encoding defined by the given alphabet
func NewEncoding(encoder string) *Encoding {
	return &Encoding{
		encode: encoder,
	}
}

// NewStdEncoding returns an Encoding preconfigured with the standard base62 alphabet
func NewStdEncoding() *Encoding {
	return NewEncoding(EncodeStd)
}

// StdEncoding is the standard base62 encoding
var StdEncoding = NewStdEncoding()

// Configurable options for an Encoding

type option func(*Encoding)

// Padding sets the minimum string length returned when encoding
// strings shorter than this will be left padded with zeros
func Padding(n int) option {
	return func(e *Encoding) {
		e.padding = n
	}
}

/**
 * Encoder
 */

// EncodeInt64 returns the base62 encoding of n using the StdEncoding
func EncodeInt64(n int64, basechars string) string {
    EncodeStd = basechars
	return StdEncoding.EncodeInt64(n)
}

/**
TODO: modify EncodeBigInt to accept basechars 
// EncodeBigInt returns the base62 encoding of an arbitrary precision integer using the StdEncoding
func EncodeBigInt(n *big.Int) string {
	return StdEncoding.EncodeBigInt(n)
}
**/
// EncodeInt64 returns the baseN encoding of n
func (e *Encoding) EncodeInt64(n int64) string {
	var (
		c string
		rem int64
	)

	// Progressively divide by base, store remainder each time
	// Prepend as an additional character is the higher power
	base = int64(len(EncodeStd))
	for n > 0 {
		rem = n % base
		n = n / base
		c = string(EncodeStd[rem]) + c
	}
	s := c
	if e.padding > 0 {
		s = e.pad(s, e.padding)
	}

	return s
}
/**
TODO: modify rest of methods to accept basechars 
// EncodeBigInt returns the base62 encoding of an arbitrary precision integer
func (e *Encoding) EncodeBigInt(n *big.Int) string {
	var (
		b    = make([]byte, 0)
		rem  = new(big.Int)
		bse  = new(big.Int)
		zero = new(big.Int)
	)
	bse.SetInt64(base)
	zero.SetInt64(0)

	// Progressively divide by base, until we hit zero
	// store remainder each time
	// Prepend as an additional character is the higher power
	for n.Cmp(zero) == 1 {
		n, rem = n.DivMod(n, bse, rem)
		b = append([]byte{e.encode[rem.Int64()]}, b...)
	}

	s := string(b)
	if e.padding > 0 {
		s = e.pad(s, e.padding)
	}

	return s
}


 * Decoder
 

// DecodeToInt64 decodes a base62 encoded string using the StdEncoding
func DecodeToInt64(s string, basechars string) int64 {
	EncodeStd = basechars
	return StdEncoding.DecodeToInt64(s)
}


TODO: modify DecodeToBigInt to accept basechars  
// DecodeToBigInt returns an arbitrary precision integer from the base62
// encoded string using the StdEncoding
func DecodeToBigInt(s string) *big.Int {
	return StdEncoding.DecodeToBigInt(s)
}


// DecodeToInt64 decodes a base62 encoded string
func (e *Encoding) DecodeToInt64(s string) int64 {
	var (
		n     int64
		c     int64
		idx   int
		power int
	)

	for i, v := range s {
		idx = strings.IndexRune(e.encode, v)

		// Work downwards through powers of our base
		power = len(s) - (i + 1)

		// Calculate value at this position and add
		c = int64(idx) * int64(math.Pow(float64(base), float64(power)))
		n = n + c
	}

	return int64(n)
}

// DecodeToBigInt returns an arbitrary precision integer from the base62 encoded string
func (e *Encoding) DecodeToBigInt(s string) *big.Int {
	var (
		n = new(big.Int)

		c     = new(big.Int)
		idx   = new(big.Int)
		power = new(big.Int)
		exp   = new(big.Int)
		bse   = new(big.Int)
	)
	bse.SetInt64(base)

	// Run through each character to decode
	for i, v := range s {
		// Get index/position of the rune as a big int
		idx.SetInt64(int64(strings.IndexRune(e.encode, v)))

		// Work downwards through exponents
		exp.SetInt64(int64(len(s) - (i + 1)))

		// Calculate power for this exponent
		power.Exp(bse, exp, nil)

		// Multiplied by our index, gives us the value for this character
		c = c.Mul(idx, power)

		// Finally add to running total
		n.Add(n, c)
	}

	return n
}
**/
// pad a string to a minimum length with zero characters
func (e *Encoding) pad(s string, minlen int) string {
	if len(s) >= minlen {
		return s
	}

	format := fmt.Sprint(`%0`, strconv.Itoa(minlen), "s")
	return fmt.Sprintf(format, s)
}
