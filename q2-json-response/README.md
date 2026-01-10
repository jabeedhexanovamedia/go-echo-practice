# JSON Response Server

This is a Go server using the Echo framework that demonstrates different ways to return JSON responses from API endpoints.

## Code Breakdown

```go
package main

import (
    "net/http"
    "github.com/labstack/echo/v4"
)

type User struct {
    Id   int    `json:"id"`
    Name string `json:"name"`
}

func main() {
    e := echo.New()

    e.GET("/users", func(c echo.Context) error {
        return c.JSON(http.StatusOK, map[string]interface{}{
            "id":   1,
            "name": "john",
        })
    })

    // using echo.Map
    e.GET("/user", func(c echo.Context) error {
        return c.JSON(http.StatusOK, echo.Map{
            "id":   1,
            "name": "John",
        })
    })

    // using struct
    e.GET("/users2", func(C echo.Context) error {
        return C.JSON(200, User{
            Id:   2,
            Name: "john 2",
        })
    })

    e.Logger.Fatal(e.Start("127.0.0.1:3000"))
}
```

## Breaking Down Key Concepts

### Struct Definition with JSON Tags

```go
type User struct {
    Id   int    `json:"id"`
    Name string `json:"name"`
}
```

- Defines a `User` struct with fields `Id` (int) and `Name` (string).
- JSON tags (`json:"id"`, `json:"name"`) control how fields are serialized to JSON.
- Without tags, Go uses field names as-is; tags allow custom naming (e.g., lowercase keys).

### Returning JSON with `c.JSON()`

Echo's `c.JSON(statusCode, data)` method:

- Serializes `data` to JSON.
- Sets `Content-Type: application/json`.
- Returns the specified HTTP status code.

### Method 1: Using `map[string]interface{}`

```go
e.GET("/users", func(c echo.Context) error {
    return c.JSON(http.StatusOK, map[string]interface{}{
        "id":   1,
        "name": "john",
    })
})
```

- `map[string]interface{}` is a flexible map for dynamic data.
- Keys are strings, values can be any type.
- Useful for ad-hoc JSON without defining structs.

### Method 2: Using `echo.Map`

```go
e.GET("/user", func(c echo.Context) error {
    return c.JSON(http.StatusOK, echo.Map{
        "id":   1,
        "name": "John",
    })
})
```

- `echo.Map` is an alias for `map[string]interface{}`.
- Provided by Echo for convenience.
- Same functionality as standard map, but Echo-specific.

### Method 3: Using Structs

```go
e.GET("/users2", func(C echo.Context) error {
    return C.JSON(200, User{
        Id:   2,
        Name: "john 2",
    })
})
```

- Passes a struct instance directly to `c.JSON()`.
- Echo uses reflection to serialize based on JSON tags.
- Preferred for typed data: ensures structure, enables validation, and improves maintainability.

### Server Startup

```go
e.Logger.Fatal(e.Start("127.0.0.1:3000"))
```

- Starts the server on localhost (127.0.0.1) port 3000.
- `e.Logger.Fatal()` logs errors and exits on failure.
- Binds only to localhost, not accessible externally.

## Alternatives and Best Practices

### Different JSON Serialization Methods

| Method                   | Pros                          | Cons                            |
| ------------------------ | ----------------------------- | ------------------------------- |
| `map[string]interface{}` | Flexible, no struct needed    | No type safety, prone to errors |
| `echo.Map`               | Same as map, Echo convenience | Same as map                     |
| Struct with tags         | Type-safe, clear structure    | Requires struct definition      |

### Handling JSON Input

To read JSON from requests:

```go
e.POST("/user", func(c echo.Context) error {
    u := new(User)
    if err := c.Bind(u); err != nil {
        return err
    }
    return c.JSON(200, u)
})
```

- `c.Bind(u)` deserializes JSON body into struct.
- Validates against struct fields.

### Custom JSON Marshaling

For advanced control, implement `json.Marshaler`:

```go
func (u User) MarshalJSON() ([]byte, error) {
    return json.Marshal(map[string]interface{}{
        "user_id": u.Id,
        "full_name": u.Name,
    })
}
```

- Allows custom JSON output logic.

### Server Configuration Alternatives

Similar to the simple server:

- `e.Start(":3000")` - Listen on all interfaces.
- `e.StartServer(&http.Server{...})` - Custom server with timeouts.

## Core Concepts Refresher

- **JSON Tags**: Control serialization field names.
- **c.JSON()**: Echo's method for JSON responses.
- **Maps vs Structs**: Flexibility vs type safety.
- **Binding**: Reading JSON into Go structs.
- **Server Binding**: Host/port configuration for security.

This setup demonstrates fundamental JSON handling in Go web APIs using Echo.
