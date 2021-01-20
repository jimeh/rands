window.BENCHMARK_DATA = {
  "lastUpdate": 1611113059471,
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
      },
      {
        "commit": {
          "author": {
            "email": "contact@jimeh.me",
            "name": "Jim Myhrberg",
            "username": "jimeh"
          },
          "committer": {
            "email": "contact@jimeh.me",
            "name": "Jim Myhrberg",
            "username": "jimeh"
          },
          "distinct": true,
          "id": "f13952d55ef8c66c0b217ad0099418a595c2f5b2",
          "message": "docs(readme): fix formatting issue",
          "timestamp": "2021-01-20T03:14:20Z",
          "tree_id": "73e3ae19340a5552f0b232dee25c2a69b2096276",
          "url": "https://github.com/jimeh/rands/commit/f13952d55ef8c66c0b217ad0099418a595c2f5b2"
        },
        "date": 1611113056445,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkString__greek",
            "value": 11,
            "unit": "ns/op",
            "extra": "100000000 times"
          },
          {
            "name": "BenchmarkString__chinese",
            "value": 12.6,
            "unit": "ns/op",
            "extra": "93399517 times"
          },
          {
            "name": "BenchmarkString__japanese",
            "value": 12.7,
            "unit": "ns/op",
            "extra": "94113018 times"
          },
          {
            "name": "BenchmarkString__uppercase",
            "value": 28680,
            "unit": "ns/op",
            "extra": "41763 times"
          },
          {
            "name": "BenchmarkString__lowercase",
            "value": 29038,
            "unit": "ns/op",
            "extra": "41145 times"
          },
          {
            "name": "BenchmarkUnicodeString__greek",
            "value": 41576,
            "unit": "ns/op",
            "extra": "29212 times"
          },
          {
            "name": "BenchmarkUnicodeString__chinese",
            "value": 62802,
            "unit": "ns/op",
            "extra": "19082 times"
          },
          {
            "name": "BenchmarkUnicodeString__japanese",
            "value": 61330,
            "unit": "ns/op",
            "extra": "19392 times"
          },
          {
            "name": "BenchmarkUnicodeString__uppercase",
            "value": 29148,
            "unit": "ns/op",
            "extra": "42187 times"
          },
          {
            "name": "BenchmarkUnicodeString__lowercase",
            "value": 29328,
            "unit": "ns/op",
            "extra": "41301 times"
          }
        ]
      }
    ]
  }
}