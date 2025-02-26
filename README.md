<h1 align="center">
  rands
</h1>

<p align="center">
  <strong>
    Go package providing a suite of functions that use <code>crypto/rand</code>
    to generate cryptographically secure random strings in various formats, as
    well as ints and bytes.
  </strong>
</p>

<p align="center">
  <a href="https://pkg.go.dev/github.com/jimeh/rands"><img src="https://img.shields.io/badge/%E2%80%8B-reference-387b97.svg?logo=go&logoColor=white" alt="Go Reference"></a>
  <a href="https://github.com/jimeh/rands/releases"><img src="https://img.shields.io/github/v/tag/jimeh/rands?label=release" alt="GitHub tag (latest SemVer)"></a>
  <a href="https://github.com/jimeh/rands/actions"><img src="https://img.shields.io/github/actions/workflow/status/jimeh/rands/ci.yml?branch=main&logo=github" alt="Actions Status"></a>
  <a href="https://codeclimate.com/github/jimeh/rands"><img src="https://img.shields.io/codeclimate/coverage/jimeh/rands.svg?logo=code%20climate" alt="Coverage"></a>
  <a href="https://github.com/jimeh/rands/issues"><img src="https://img.shields.io/github/issues-raw/jimeh/rands.svg?style=flat&logo=github&logoColor=white" alt="GitHub issues"></a>
  <a href="https://github.com/jimeh/rands/pulls"><img src="https://img.shields.io/github/issues-pr-raw/jimeh/rands.svg?style=flat&logo=github&logoColor=white" alt="GitHub pull requests"></a>
  <a href="https://github.com/jimeh/rands/blob/master/LICENSE"><img src="https://img.shields.io/github/license/jimeh/rands.svg?style=flat" alt="License Status"></a>
</p>

## [`rands`](https://pkg.go.dev/github.com/jimeh/rands) package

`rands` is intended for use in production code where random data generation is
required. All functions have a error return value which should be checked.

For tests there is the `randsmust` package, which has all the same functions but
with single return values, and they panic in the event of an error.

### Import

```
import "github.com/jimeh/rands"
```

### Usage

```go
s, err := rands.Base64(16)       // => CYxqEdUB1Rzno3SyZu2g/g==
s, err := rands.Base64URL(16)    // => zlqw9aFqcFggbk2asn3_aQ
s, err := rands.Hex(16)          // => 956e2ec9e7f19ddd58bb935826926531
s, err := rands.Alphanumeric(16) // => Fvk1PkrmG5crgOjT
s, err := rands.Alphabetic(16)   // => XEJIzcZufHkuUmRM
s, err := rands.Upper(16)        // => UMAGAFPPNDRGLUPZ
s, err := rands.UpperNumeric(16) // => DF0CQS0TK9CPUO3E
s, err := rands.Lower(16)        // => ocsmggykzrxzfwgt
s, err := rands.LowerNumeric(16) // => rwlv7a1p7klqffs5
s, err := rands.Numeric(16)      // => 9403373143598295

s, err := rands.String(16, "abcdefABCDEF")                               // => adCDCaDEdeffeDeb
s, err := rands.UnicodeString(16, []rune("九七二人入八力十下三千上口土夕大")) // => 下下口九力下土夕下土八上二夕大三

s, err := rands.DNSLabel(16) // => z0ij9o8qkbs0ru-h
s, err := rands.UUID()       // => a62b8712-f238-43ba-a47e-333f5fffe785
s, err := rands.UUIDv7()     // => 01954a31-867f-7ffb-876e-b818f960ec3b

n, err := rands.Int(2147483647)                   // => 1334400235
n, err := rands.Int64(int64(9223372036854775807)) // => 8256935979116161233

b, err := rands.Bytes(8) // => [0 220 137 243 135 204 34 63]
```

## [`randsmust`](https://pkg.go.dev/github.com/jimeh/rands/randsmust) package

`randsmust` is specifically intended as an alternative to `rands` for use in
tests. All functions return a single value, and panic in the event of an error.
This makes them easy to use when building structs in test cases that need random
data.

For production code, make sure to use the `rands` package and check returned
errors.

### Import

```
import "github.com/jimeh/rands/randsmust"
```

### Usage

```go
s := randsmust.Base64(16)       // => d1wm/wS6AQGduO3uaey1Cg==
s := randsmust.Base64URL(16)    // => 4pHWVcddXsL_45vhOfCdng
s := randsmust.Hex(16)          // => b5552558bc009264d129c422a666fe56
s := randsmust.Alphanumeric(16) // => j5WkpNKmW8K701XF
s := randsmust.Alphabetic(16)   // => OXxsqfFjNLvmZqDb
s := randsmust.Upper(16)        // => AOTLYQRCVNMEPRCX
s := randsmust.UpperNumeric(16) // => 1NTY6KATDVAXBTY2
s := randsmust.Lower(16)        // => xmftrwvurrritqfu
s := randsmust.LowerNumeric(16) // => yszg56fzeql7pjpl
s := randsmust.Numeric(16)      // => 0761782105447226

s := randsmust.String(16, "abcdefABCDEF")                               // => dfAbBfaDDdDFDaEa
s := randsmust.UnicodeString(16, []rune("九七二人入八力十下三千上口土夕大")) // => 十十千口三十十下九上千口七夕土口

s := randsmust.DNSLabel(16) // => pu31o0gqyk76x35f
s := randsmust.UUID()       // => d616c873-f3dd-4690-bcd6-ed307eec1105
s := randsmust.UUIDv7()     // => 01954a30-add2-7590-8238-6cf6b2790c1e

n := randsmust.Int(2147483647)                   // => 1293388115
n := randsmust.Int64(int64(9223372036854775807)) // => 6168113630900161239

b := randsmust.Bytes(8) // => [205 128 54 95 0 95 53 51]
```

## Documentation

Please see the Go Reference for documentation and examples:

- [`rands`](https://pkg.go.dev/github.com/jimeh/rands)
- [`randsmust`](https://pkg.go.dev/github.com/jimeh/rands/randsmust)

## Benchmarks

Benchmark reports and graphs are available here:
https://jimeh.me/rands/dev/bench/

## License

[MIT](https://github.com/jimeh/rands/blob/main/LICENSE)
