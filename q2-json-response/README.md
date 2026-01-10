# JSON Response Server

This is a Go server using the Echo framework that demonstrates different ways to return JSON responses and handle JSON input from API endpoints.

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

type UserSignupRequest struct {
    Name  string `json:"name,omitempty"`
    Email string `json:"email"`
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

    // POST /users: Accepts JSON request body, binds to struct, returns as JSON
    e.POST("/users", func(c echo.Context) error {
        var userReq UserSignupRequest

        // JSON → Go struct
        if err := c.Bind(&userReq); err != nil {
            return c.JSON(http.StatusBadRequest, echo.Map{
                "message": "invalid request payload",
            })
        }

        // Go struct → JSON
        return c.JSON(200, userReq)
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

### Handling JSON Input with Binding

```go
type UserSignupRequest struct {
    Name  string `json:"name,omitempty"`
    Email string `json:"email"`
}

e.POST("/users", func(c echo.Context) error {
    var userReq UserSignupRequest

    // JSON → Go struct
    if err := c.Bind(&userReq); err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{
            "message": "invalid request payload",
        })
    }

    // Go struct → JSON
    return c.JSON(200, userReq)
})
```

#### What is Binding?
- Binding is the process of converting incoming JSON data from an HTTP request body into Go data structures (usually structs).
- In Echo, `c.Bind(data)` automatically deserializes the JSON request body into the provided struct pointer.
- It uses Go's `encoding/json` package under the hood for unmarshaling.

#### Why Do We Need Binding?
- **Type Safety**: Ensures incoming data matches expected structure and types.
- **Validation**: Allows validation of required fields and data formats.
- **Convenience**: Automatically handles JSON parsing, reducing boilerplate code.
- **Error Handling**: Provides clear errors for malformed JSON or missing required fields.

#### How `c.Bind()` Works
1. Reads the request body.
2. Parses JSON into the target struct using reflection and JSON tags.
3. Returns an error if parsing fails (e.g., invalid JSON, type mismatches).
4. On success, the struct is populated with the parsed data.

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

### JSON Marshaling and Unmarshaling Basics

Go's `encoding/json` package provides core JSON functionality:

#### Marshaling (Go → JSON)
```go
import "encoding/json"

user := User{Id: 1, Name: "John"}
jsonBytes, err := json.Marshal(user)
// jsonBytes: {"id":1,"name":"John"}
```

- `json.Marshal(data)` converts Go data to JSON bytes.
- Works with structs, maps, slices, etc.

#### Unmarshaling (JSON → Go)
```go
jsonStr := `{"id":1,"name":"John"}`
var user User
err := json.Unmarshal([]byte(jsonStr), &user)
// user.Id = 1, user.Name = "John"
```

- `json.Unmarshal(jsonBytes, &target)` parses JSON into Go data.
- Target must be a pointer to the data structure.

### JSON Tags in Detail

JSON tags control serialization behavior:

```go
type User struct {
    ID       int    `json:"id"`           // Rename field
    Name     string `json:"name"`         // Rename field
    Password string `json:"-"`            // Ignore field
    Email    string `json:"email,omitempty"` // Omit if empty
}
```

- `json:"fieldName"`: Custom JSON key name.
- `json:"-"`: Skip field entirely.
- `json:",omitempty"`: Omit field if zero value (empty string, 0, false, nil).

### Omitempty and Other Tag Options

- `omitempty`: Excludes field from JSON if it's the zero value.
  - For strings: empty string `""`
  - For ints: `0`
  - For pointers/slices: `nil`

```go
type UserSignupRequest struct {
    Name  string `json:"name,omitempty"`  // Omitted if empty
    Email string `json:"email"`           // Always included
}
```

- Other options: `json:"name,string"` (encode as string), but rarely used.

### Manual JSON Handling

Instead of `c.Bind()`, you can manually handle JSON:

```go
e.POST("/users", func(c echo.Context) error {
    body, err := io.ReadAll(c.Request().Body)
    if err != nil {
        return err
    }

    var userReq UserSignupRequest
    if err := json.Unmarshal(body, &userReq); err != nil {
        return c.JSON(400, echo.Map{"error": "invalid JSON"})
    }

    return c.JSON(200, userReq)
})
```

- Gives more control but requires more code.
- Useful for custom validation or preprocessing.

### Alternatives to `c.Bind()`

Echo provides other binding methods:

- `c.Bind(i interface{})`: General binding (JSON, form, query).
- `c.BindJSON(i interface{})`: JSON-specific binding.
- `c.BindQuery(i interface{})`: Query parameters.
- `c.BindForm(i interface{})`: Form data.

For strict JSON only:
```go
if err := c.BindJSON(&userReq); err != nil {
    // handle error
}
```

### Error Handling in Binding

Common binding errors and handling:

```go
if err := c.Bind(&userReq); err != nil {
    switch err.(type) {
    case *json.SyntaxError:
        return c.JSON(400, echo.Map{"error": "invalid JSON syntax"})
    case *json.UnmarshalTypeError:
        return c.JSON(400, echo.Map{"error": "invalid data type"})
    default:
        return c.JSON(400, echo.Map{"error": "bad request"})
    }
}
```

- `*json.SyntaxError`: Malformed JSON.
- `*json.UnmarshalTypeError`: Type mismatch (e.g., string in int field).

### Custom JSON Marshaling

Implement interfaces for advanced control:

```go
type User struct {
    Id   int
    Name string
}

// Custom marshaling
func (u User) MarshalJSON() ([]byte, error) {
    return json.Marshal(map[string]interface{}{
        "user_id": u.Id,
        "full_name": u.Name,
        "timestamp": time.Now().Unix(),
    })
}

// Custom unmarshaling
func (u *User) UnmarshalJSON(data []byte) error {
    var temp map[string]interface{}
    if err := json.Unmarshal(data, &temp); err != nil {
        return err
    }
    u.Id = int(temp["user_id"].(float64))
    u.Name = temp["full_name"].(string)
    return nil
}
```

- `MarshalJSON()`: Customize Go → JSON conversion.
- `UnmarshalJSON()`: Customize JSON → Go conversion.

### Different JSON Serialization Methods

| Method | Pros | Cons |
|--------|------|------|
| `map[string]interface{}` | Flexible, no struct needed | No type safety, prone to errors |
| `echo.Map` | Same as map, Echo convenience | Same as map |
| Struct with tags | Type-safe, clear structure | Requires struct definition |
| Custom marshaling | Full control over output | More complex to implement |

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

- **JSON Tags**: Control serialization field names, omitempty, ignore fields.
- **Marshaling**: Converting Go data to JSON (`json.Marshal`).
- **Unmarshaling**: Converting JSON to Go data (`json.Unmarshal`).
- **Binding**: Automatic JSON → Go struct conversion in Echo.
- **c.JSON()**: Echo's method for JSON responses.
- **Maps vs Structs**: Flexibility vs type safety in JSON handling.
- **Omitempty**: Exclude zero values from JSON output.
- **Custom Marshaling**: Implement interfaces for advanced JSON control.
- **Error Handling**: Handle binding errors for robust APIs.

This setup demonstrates comprehensive JSON handling in Go web APIs using Echo, covering both input and output operations.
