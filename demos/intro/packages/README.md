# Packages

Go comes with an extensive standard library of useful packages, and the Go community has created and shared many more. Before you embark on any new program, it's a good idea to see if packages already exist that might help you get your job done more easily.

You can find an index of the *standard library package* at `https://golang.org/pkg` and the packages contributed by the community at `https://godoc.org`. The `go doc` tool makes these documents easily accessible from the command line:

```
$ go doc http.ListenAndServe

package http // import "net/http"

func ListenAndServe(addr string, handler Handler) error
    ListenAndServe listens on the TCP network address addr and then calls Serve
    with handler to handle requests on incoming connections. Accepted
    connections are configured to enable TCP keep-alives.

    The handler is typically nil, in which case the DefaultServeMux is used.

    ListenAndServe always returns a non-nil error.
```
