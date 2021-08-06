# Concurrency: *goroutines* and *channels*

Most programming languages were designed for processors with a single core. Go was designed with parallel and concurrent processing in mind. Go enables two style of concurrent programming, *goroutines* and *channels* support **communicating sequential processes (CSP)**, a model of concurrency in which values are passed between independent activities (goroutines1) but variables are for the most part confined toa single activity. On the other hand, Go also supports the more traditional model of **shared memory multithreading**.

> The differences between threads and goroutines are essentially quantitative, not qualitative.

* Goroutines
* Channels
  * Unbuffered Channels
  * Pipelines
  * Undirirectional Channel Types
  * Buffered Channels

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

A *channel* has two principal operations, *send* and *receive*, collectively known as *communications*. A send statement transmits a value from one goroutine, through the channel, to another goroutine executing a corresponding receive operations. Both operations are written using the `<-` operator. In a send statement, the `<-` separates the channel and value operands. In a receive expression, `<-` precedes the channel operand.

```golang
ch <- x // a send statement
x = <-ch // a receive expression in an assignment statement
<- ch // a receive statement; result is discarded. This can be used for waiting for a goroutine to finish.
```

Channels support a third operation, *close*, which sets a flag indicating that no more values will ever be sent on this channel, subsequent attempts to send will panic. Receive operations on a closed channels yield the values that have been sent until no more values are left; any receive operations thereafter complete immediately and yield the zero value of the channel's element type.

```golang
close(ch)
```

A channel created with a simple call to make is an *unbuffered channel*, but *make* accepts an optional second argument, an integer called the channel's *capacity*. If the capacity is non-zero, *make* creates a *buffered* channel.

```golang
ch = make(chan int) // unbuffered channel
ch = make(chan int, 0) // unbuffered channel
ch = make(chan int, 3) // buffered channel with capacity 3
```

### Unbuffered Channels

A *send* operation on an unbuffered channel blocks the sending goroutine until another goroutine executes a corresponding *receive* on the same channel, at which point the value is transmitted and both goroutines may continue. Conversely, if the *receive* operation was attempted first, the receiving goroutine is blocked until another goroutine performs a send on the same channel.

Communication over an unbuffered channel causes the sneding and receiving goroutines to *synchronize*. Because of this, unbuffered channels are sometimes called *synchronous channels*. When a value is sent on an unbuffered channel, the receipt of the value happens before the reawakening of the sending goroutine.

> When `x` neither happens before `y` nor after `y`, we say that *x is concurrent with y*. This doesn't mean that they are necessarily simultaneous, merely that we cannot assume anything about their ordering.

### Pipelines

Channels can be used to connect goroutines so that the output of one is the input to another. This is called a *pipeline*.

If the sender knows that no further valeus will ever be sent on a channel, it's useful to communicate this fact to receiver goroutines so that they can stop waiting. This is accomplished by *closing* the channel with `close`.

After the closed channel has been *drained*, that is, after the last sent element has been received, all subsequent receive operations will proceed without blocking but will yield a zero value. There is no way to test directly whether a channel has been closed, but there is a variant of the *receive* operation that produces two results: the received channel element, plus a boolean value, conventionally called `ok`, which is `true` for a successful receive, and `false` for a receive on a closed and drained channel.

```go
go func() {
	for {
		x, ok := <-naturals
		if !ok {
			break // channel was closed and drained
		}
		squares <- x * x
	}
	close(squares)
}
```

Because the syntax above is clumsy and this pattern is common, the language lets us use a `range` loop to iterate over channels too, for receiving all the values sent on a channel and terminating the loop after the last one.

```golang
func main() {
	naturales := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// Squarer
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}

	// Printer (in main goroutine)
	for x:= range squares {
		fmt.Println(x)
	}
}
```

You needn't close every channel when you've finished with it. It's o0nly necessary to close a channel when it is important to tell the receiving goroutines that all data have been sent. A channel that the garbage collector determines to be unreachable will have its resources reclaimed whether or not is closed.

### Undirectional Channel Types

Go type system provides *undirectional channel* types that expose only sending or receiving operations. The type `chan<- int` a *send-only* channel of `int`, allows sends but not receives. Conversely, the type `<-chan int` a *receive-only* channel of `int`, allows receives but not sends. Violations of this discipline are detected at compile time.

```golang
func counter(out chan<- int) {
	for x:= 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	naturals := make(chan int)
	squares := makle(chan int)

	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}
```

### Buffered Channels

A *buffered channel* has a queue of elements. The queue's maximum size is determined when it is created, by the *capacity* argument to `make`.

A `send` operation on a buffered channel inserts an element at the back of the queue, and a `receive` operation removes an element from the front. If the channel is full, the `send` operation blocks its goroutine until space is made available by another goroutine's receive. Conversely, if the channel is empty, a `receive` operation blocks until a value is sent by another goroutine.

> Failure to allocate sufficient buffer capacity would cause the program to **deadlock**.
>
