# Simple Echo Server

This is a basic HTTP server using the Echo framework in Go. It responds with "Hello world" on the root path.

## Code Breakdown

```go
package main

import "github.com/labstack/echo/v4"

func main() {
    e := echo.New()

    e.GET("/", func(c echo.Context) error {
        return c.String(200, "Hello world")
    })

    e.Logger.Fatal(e.Start(":8080"))
}
```

## Breaking Down `e.Logger.Fatal(e.Start(":8080"))`

### `e`

- `e` is the Echo instance created with `echo.New()`.
- It provides methods for routing, middleware, starting the server, logging, and configuration.

### `e.Start(":8080")`

- Starts the HTTP server on port 8080, listening on all interfaces (0.0.0.0).
- Returns an error if the server cannot start.
- Blocks until the server is stopped or encounters a fatal error.

### `e.Logger.Fatal(...)`

- `e.Logger` is the Echo logger.
- `Fatal` logs the message at FATAL level and exits the program with a non-zero status.
- Wraps `e.Start(":8080")` to log any startup errors and terminate if necessary.

## Alternative Ways to Start the Server

### Option 1: Custom Host Binding

```go
e.Logger.Fatal(e.Start("127.0.0.1:9090")) // Bind only to localhost
```

### Option 2: HTTPS with StartTLS

```go
e.Logger.Fatal(e.StartTLS(":8443", "cert.pem", "key.pem"))
```

- Starts an HTTPS server using SSL certificate and key.

### Option 3: Advanced Control with Server Struct

```go
import "net/http"
import "time"

s := &http.Server{
    Addr:         ":8081",
    ReadTimeout:  10 * time.Second,
    WriteTimeout: 10 * time.Second,
    IdleTimeout:  60 * time.Second,
}

e.Logger.Fatal(e.StartServer(s))
```

- Provides fine-grained control over timeouts and host binding.
- Preferred for production use.

### Option 4: Manual Error Handling

```go
import "fmt"
import "os"

err := e.Start(":8080")
if err != nil {
    fmt.Println("Server failed:", err)
    os.Exit(1)
}
```

- Handles errors manually instead of using Fatal.

## Summary

| Method                                       | Description                           |
| -------------------------------------------- | ------------------------------------- |
| `e.Start(":8080")`                           | Start simple HTTP server on port 8080 |
| `e.Logger.Fatal(e.Start(":8080"))`           | Start server, log fatal errors        |
| `e.StartTLS(":8443", "cert.pem", "key.pem")` | Start HTTPS server                    |
| `e.StartServer(&http.Server{...})`           | Advanced control with custom settings |
| Manual error handling                        | Start and handle errors yourself      |

## What is Fatal in Go (and Echo)

In Echo, `e.Logger.Fatal(err)`:

- Logs the message at FATAL level (critical error).
- Exits the program immediately with a non-zero exit code.

This ensures that if the server fails to start, the program stops and reports the issue.
