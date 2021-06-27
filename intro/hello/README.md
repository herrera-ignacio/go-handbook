# Hello World example

Go is a compiled language. The Go toolchain converts a source program and the things it depends on into instructions in the native machine language of a computer.

These tools are accessed through the `go` command, that has a number of subcommands. The simplest of these subcommands is `run`, which compiles the source code from one or more source files whose names end in `.go`, links it with libraries, then runs the resulting executable file.

```bash
go run ./demos/intro/helloworld/helloworld.go
```

You can compile it and save the compiled result for later use.

```bash
go build ./demos/intro/helloworld/helloworld.go
./demos/intro/helloworld/helloworld
```
