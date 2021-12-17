package randsmust

var testCases = []struct {
	name string
	n    int
}{
	{name: "n=0", n: 0},
	{name: "n=1", n: 1},
	{name: "n=2", n: 2},
	{name: "n=7", n: 7},
	{name: "n=8", n: 8},
	{name: "n=16", n: 16},
	{name: "n=32", n: 32},
	{name: "n=128", n: 128},
	{name: "n=1024", n: 1024},
	{name: "n=409600", n: 409600},
	{name: "n=2394345", n: 2394345},
}

func recoverPanic(f func()) (p interface{}) {
	defer func() {
		if r := recover(); r != nil {
			p = r
		}
	}()

	f()

	return
}
