# Concurrency: *goroutines* and *channels*

Most programming languages were designed for processors with a single core. Go was designed with parallel and concurrent processing in mind.

## Goroutine

Go has a feature called **goroutine**, a function that can be run concurrently to the main program or other goroutines. Sometimes dubbed *lightweight threads*, goroutines are managed by the Go runtime, where they're mapped and moved to the appropriate operating system thread and garbage-collected when no longer needed. When multiple processor cores are available, the goroutines can be run in parallel because various threads are running on different processing cores. But from the developer's point of view, creating a goroutine is as easy as wiriting a function.

### Goroutine Example: Printing concurrently

```go
// Function to execute as a goroutine
func count() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 1)
	}
}

func main() {
	go count() // Starts goroutine
	time.Sleep(time.Millisecond * 2)
	fmt.Println("Hello World")
	time.Sleep(time.Millisecond * 5)
}
```

Expected output:

```
0
1
Hello World
2
3
4
```

## Channel

*Channels* provide a way for two goroutines to communicate with each other. By default, they block execution, allowing goroutines to synchronize. They can be one-directional or bidirectional.
