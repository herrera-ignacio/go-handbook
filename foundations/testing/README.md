# Testing

## Overview

Each file must inmport the testing package `import "testing"`. Test functions have the following signature:

```go
// test functions
func TestName(t *testing.T) {
  // ...
}

// benchmark functions
func BenchmarkName(b* *testing.B) {
  for i := 0; i < b.N; i++ {
    test()
  }
}

func benchmark(b *testing.B, size int) { /* ... */ }
func Benchmark10(b *testing.B) { benchmark(b, 10) }
func Benchmark1000(b *testing.B) { benchmark(b, 1000) }

//e xample functions
func ExampleIsPalindrome() {
  fmt.Println(IsPalindrome("A man, a plan, a canal: Panama"))
  fmt.Println(IsPalindrome("palindrome"))
  // Output:
  // true
  // false
}
```

The `t` parameter provides methods for reporting test failures and logging additional information.

The `go test` tool scans the `*_test.go` files for these special functions, generates a temporary `main` package that calls them all in the proper way, builds and runs it, reports the results and then cleans up. The `-run` flag, whose argument is a regular expression, causes `go test` to run only those tests whose function name matches the pattern.

## Types

* A *test function*, which is a function whose name begins with `Test`, exercises some program logic for correct behavior.

* A *benchmark function* has a name beginning with `Benchmark` and measures the performance of some operation; `go test` reports the mean execution time of the operation. `go test -bench=.`

* An *example function*, whose name starts with `Example`, provides machine-checked documentation.

## Coverage

```
go tool -cover [-html=c.out] [-coverprofile=c.out] [-covermode=count]
go test -cpuprofile=cpu.out -blockprofile=block.out -memprofile=mem.out
```
