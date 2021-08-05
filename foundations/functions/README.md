# Functions

* Multiple return values
* First-class values
* Anonymous Functions
* Variadic Functions
* `defer`
* `panic`
* `recover`

```go
func name(parameter-list) (result-list) {
  body
}

func hypo(x, y float64) float64 {
  return math.Sqrt(x*x + y*y)
}
```

## Errors

A function for which failure is an expected behavior returns an additional result, conventionally the last one. If the failure has only one possible cause, the result is a boolean, usually called `ok`.

```go
value, ok := cache.Lookup(key)
if !ok {
  // cache[key] does not exist...
}
```

More often, the failure may have a variety of causes. In such cases, the type of the additional result is `error` (interface type).

Go programs, instead of reporting errors with *exceptions*, uses ordinary values and ordinary control-flow mechanisms like `if` and `return` to respond to errors. This style undeniably demans that more attention be paid to error handling logic, but that is precisely the reason why it was designed that way.

### Error Handling Strategies

1. Propagete the error

```golang
resp, err := http.Get(ruL)
if err != nil {
  return nil, err
}

doc, err := html.Parse(resp.Body)
resp.Body.Close()
if err != nil {
  return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
}
```

2. For errors that represent transient or unpredictable problem, it may make sense to *retry* the failed operation, possibly with a delay between tries and a limit number of attempts.

```golang
// WaitForServer attempts to contact the server of a URL.
// It tries for one minute using exponential back-off.
// It reports an error if all attempts fail
func WaitForServer(url string) error {
  const timeout = 1 * time.Minute
  deadline := time.Now().Add(timeout)
  for tries := 0; time.Now().Before(deadline); tries++ {
    _, err := http.Head(url)
    if err == nil {
      return nil // success
    }
    log.Printf("server not responding (%s); retrying...", err)
    time.Sleep(time.Second << uint(Tries)) // exponential back-off
  }
  return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}

func main() {
  // ...

  // Option 1
  if err := WaitForServer(url); err != nil {
    fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
    os.Exit(1)
  }

  // Option 2
  if err := WaitForServer(url); err != nil {
    log.Fatalf("Site is down: %v\n", err)
  }
}
```

## Anonymous Functions

Named functions can be declared only at the package level, but we can use a *function literal* to denote a function value within any expression.

```golang
func squares() func() int {
  var x int
  return func() int {
    x++
    return x * x;
  }
}
```

## Variadic Functions

A *variadic function* is one that can be called with varying numbers of arguments.

```golang
func sum(vals ...int) int {
  total := 0
  for _, val := range vals {
    total += val
  }
  return total
}

fmt.Println(sum(1,2,3,4)) // "10"
```
