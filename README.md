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
  <a href="https://pkg.go.dev/github.com/jimeh/rands">
    <img src="https://img.shields.io/badge/%E2%80%8B-reference-387b97.svg?logo=go&logoColor=white"
  alt="Go Reference">
  </a>
  <a href="https://github.com/jimeh/rands/releases">
    <img src="https://img.shields.io/github/v/tag/jimeh/rands?label=release" alt="GitHub tag (latest SemVer)">
  </a>
  <a href="https://github.com/jimeh/rands/actions">
    <img src="https://img.shields.io/github/workflow/status/jimeh/rands/CI.svg?logo=github" alt="Actions Status">
  </a>
  <a href="https://codeclimate.com/github/jimeh/rands">
    <img src="https://img.shields.io/codeclimate/coverage/jimeh/rands.svg?logo=code%20climate" alt="Coverage">
  </a>
  <a href="https://github.com/jimeh/rands/issues">
    <img src="https://img.shields.io/github/issues-raw/jimeh/rands.svg?style=flat&logo=github&logoColor=white"
alt="GitHub issues">
  </a>
  <a href="https://github.com/jimeh/rands/pulls">
    <img src="https://img.shields.io/github/issues-pr-raw/jimeh/rands.svg?style=flat&logo=github&logoColor=white" alt="GitHub pull requests">
  </a>
  <a href="https://github.com/jimeh/rands/blob/master/LICENSE">
    <img src="https://img.shields.io/github/license/jimeh/rands.svg?style=flat" alt="License Status">
  </a>
</p>

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

n, err := rands.Int(2147483647)                   // => 1334400235
n, err := rands.Int64(int64(9223372036854775807)) // => 8256935979116161233

b, err := rands.Bytes(8) // => [0 220 137 243 135 204 34 63]
```

## Import

```
import "github.com/jimeh/rands"
```

## Documentation

Please see the
[Go Reference](https://pkg.go.dev/github.com/jimeh/rands#section-documentation)
for documentation and examples.

## Benchmarks

Benchmark reports and graphs are available here:
https://jimeh.me/rands/dev/bench/

## License

[MIT](https://github.com/jimeh/rands/blob/master/LICENSE)
