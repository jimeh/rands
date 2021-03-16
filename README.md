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
s, _ := rands.Base64(16)       // => CYxqEdUB1Rzno3SyZu2g/g==
s, _ := rands.Base64URL(16)    // => zlqw9aFqcFggbk2asn3_aQ
s, _ := rands.Hex(16)          // => 956e2ec9e7f19ddd58bb935826926531
s, _ := rands.Alphanumeric(16) // => Fvk1PkrmG5crgOjT
s, _ := rands.Alphabetic(16)   // => XEJIzcZufHkuUmRM
s, _ := rands.Upper(16)        // => UMAGAFPPNDRGLUPZ
s, _ := rands.UpperNumeric(16) // => DF0CQS0TK9CPUO3E
s, _ := rands.Lower(16)        // => ocsmggykzrxzfwgt
s, _ := rands.LowerNumeric(16) // => rwlv7a1p7klqffs5
s, _ := rands.Numeric(16)      // => 9403373143598295

s, _ := rands.String(16, "abcdefABCDEF")                                // => adCDCaDEdeffeDeb
s, _ := rands.UnicodeString(16, []rune("九七二人入八力十下三千上口土夕大")) // => 下下口九力下土夕下土八上二夕大三

s, _ := rands.DNSLabel(16) // => z0ij9o8qkbs0ru-h
s, _ := rands.UUID()       // => a62b8712-f238-43ba-a47e-333f5fffe785

n, _ := rands.Int(2147483647)                   // => 1334400235
n, _ := rands.Int64(int64(9223372036854775807)) // => 8256935979116161233

b, _ := rands.Bytes(8) // => [0 220 137 243 135 204 34 63]
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
