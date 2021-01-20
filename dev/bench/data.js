window.BENCHMARK_DATA = {
  "lastUpdate": 1611112555794,
  "repoUrl": "https://github.com/jimeh/rands",
  "entries": {
    "Benchmark": [
      {
        "commit": {
          "author": {
            "email": "contact@jimeh.me",
            "name": "Jim Myhrberg",
            "username": "jimeh"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "7379265233e699206dcb91ee0e70f82cb2726fef",
          "message": "Merge pull request #1 from jimeh/initial-implementation\n\nfeat(rands): initial implementation",
          "timestamp": "2021-01-20T03:09:37Z",
          "tree_id": "4b9935fa9c78eab6b4a9255850d3bb99635013d1",
          "url": "https://github.com/jimeh/rands/commit/7379265233e699206dcb91ee0e70f82cb2726fef"
        },
        "date": 1611112555209,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkString__greek",
            "value": 11.3,
            "unit": "ns/op",
            "extra": "100000000 times"
          },
          {
            "name": "BenchmarkString__chinese",
            "value": 12.8,
            "unit": "ns/op",
            "extra": "93307009 times"
          },
          {
            "name": "BenchmarkString__japanese",
            "value": 13,
            "unit": "ns/op",
            "extra": "93292501 times"
          },
          {
            "name": "BenchmarkString__uppercase",
            "value": 29487,
            "unit": "ns/op",
            "extra": "40597 times"
          },
          {
            "name": "BenchmarkString__lowercase",
            "value": 29350,
            "unit": "ns/op",
            "extra": "39716 times"
          },
          {
            "name": "BenchmarkUnicodeString__greek",
            "value": 42214,
            "unit": "ns/op",
            "extra": "28716 times"
          },
          {
            "name": "BenchmarkUnicodeString__chinese",
            "value": 64839,
            "unit": "ns/op",
            "extra": "18519 times"
          },
          {
            "name": "BenchmarkUnicodeString__japanese",
            "value": 65280,
            "unit": "ns/op",
            "extra": "18952 times"
          },
          {
            "name": "BenchmarkUnicodeString__uppercase",
            "value": 29965,
            "unit": "ns/op",
            "extra": "41234 times"
          },
          {
            "name": "BenchmarkUnicodeString__lowercase",
            "value": 30265,
            "unit": "ns/op",
            "extra": "39325 times"
          }
        ]
      }
    ]
  }
}