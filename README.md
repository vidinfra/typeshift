# TypeShift

A Go utility package providing type-safe pointer operations and null value handling with both generic and type-specific functions.

## 🎯 Why TypeShift?

Working with pointers and null values in Go can be verbose and error-prone. TypeShift simplifies these operations by providing:

- **Safe pointer dereferencing** with automatic zero-value fallbacks
- **Easy pointer creation** from values
- **Null type handling** for database operations
- **Modern generic functions** alongside backward-compatible type-specific functions
- **Zero runtime overhead** with compile-time type safety

## 📦 Installation

```bash
go get github.com/vidinfra/typeshift
```

## 🚀 Quick Start

```go
package main

import (
    "fmt"
    "github.com/vidinfra/typeshift"
)

func main() {
    // Create pointers from values
    intPtr := typeshift.Ptr(42)
    stringPtr := typeshift.Ptr("hello")
    
    // Safely dereference pointers
    value := typeshift.Deref(intPtr)    // 42
    text := typeshift.Deref(stringPtr)  // "hello"
    
    // Safe handling of nil pointers
    var nilPtr *int = nil
    safeValue := typeshift.Deref(nilPtr) // 0 (zero value)
    
    fmt.Printf("Value: %d, Text: %s, Safe: %d\n", value, text, safeValue)
}
```

## 📋 Features

### Generic Functions (Go 1.18+)

#### Pointer Creation
```go
// Create pointers from any value type
intPtr := typeshift.Ptr(42)           // *int
stringPtr := typeshift.Ptr("hello")  // *string
boolPtr := typeshift.Ptr(true)       // *bool
timePtr := typeshift.Ptr(time.Now()) // *time.Time
```

#### Safe Dereferencing
```go
// Safely dereference any pointer type
value := typeshift.Deref(intPtr)     // int: 42
text := typeshift.Deref(stringPtr)   // string: "hello"
flag := typeshift.Deref(boolPtr)     // bool: true

// Returns zero values for nil pointers
var nilInt *int = nil
safeInt := typeshift.Deref(nilInt)   // int: 0
```

### Type-Specific Functions

For backward compatibility and explicit type handling:

```go
// Pointer creation
intPtr := typeshift.Int(42)
stringPtr := typeshift.String("hello")
boolPtr := typeshift.Bool(true)

// Safe dereferencing
value := typeshift.DerefInt(intPtr)       // 42
text := typeshift.DerefString(stringPtr)  // "hello"
flag := typeshift.DerefBool(boolPtr)      // true
```

### Supported Types

**Numeric Types:**
- `int`, `int8`, `int16`, `int32`, `int64`
- `uint`, `uint8`, `uint16`, `uint32`, `uint64`
- `float32`, `float64`

**Other Types:**
- `string`, `bool`
- `time.Time`, `uuid.UUID`

### Enum Support

Handle string-based enums with type safety:

```go
type Status string
const (
    Active   Status = "active"
    Inactive Status = "inactive"
)

var statusPtr *Status = &Active
statusString := typeshift.DerefEnumToString(statusPtr) // "active"

// Safe nil handling
var nilStatus *Status = nil
safeString := typeshift.DerefEnumToString(nilStatus) // ""
```

### Database Null Types

Handle SQL null types seamlessly:

```go
import "database/sql"

// Handle nullable database fields
nullStr := sql.NullString{String: "hello", Valid: true}
value := typeshift.DerefNullString(nullStr) // "hello"

invalidStr := sql.NullString{Valid: false}
empty := typeshift.DerefNullString(invalidStr) // ""

// Supported null types:
// - sql.NullString
// - sql.NullInt64
// - sql.NullBool
// - sql.NullFloat64
// - sql.NullTime
```

### Map Operations

Safe map dereferencing with deep copying:

```go
originalMap := map[string]string{
    "key1": "value1",
    "key2": "value2",
}

// Deep copy to prevent shared state
clonedMap := typeshift.DerefMapStringString(&originalMap)

// Modifications don't affect original
clonedMap["key1"] = "modified"
// originalMap["key1"] is still "value1"

// Safe nil handling
var nilMap *map[string]string = nil
safeMap := typeshift.DerefMapStringString(nilMap) // nil
```

## � Examples

### API Response Handling

```go
type User struct {
    ID    *int64  `json:"id"`
    Name  *string `json:"name"`
    Email *string `json:"email"`
    Age   *int    `json:"age"`
}

func processUser(user User) {
    // Safe dereferencing with zero values for missing fields
    id := typeshift.Deref(user.ID)       // 0 if nil
    name := typeshift.Deref(user.Name)   // "" if nil
    email := typeshift.Deref(user.Email) // "" if nil
    age := typeshift.Deref(user.Age)     // 0 if nil
    
    fmt.Printf("User: %d, %s <%s>, age %d\n", id, name, email, age)
}
```

### Database Operations

```go
func handleDatabaseRow(rows *sql.Rows) error {
    var name sql.NullString
    var age sql.NullInt64
    var createdAt sql.NullTime
    
    if err := rows.Scan(&name, &age, &createdAt); err != nil {
        return err
    }
    
    // Safe extraction with fallbacks
    safeName := typeshift.DerefNullString(name)         // "" if NULL
    safeAge := typeshift.DerefNullInt64(age)            // 0 if NULL
    safeTime := typeshift.DerefNullTime(createdAt)      // time.Time{} if NULL
    
    fmt.Printf("Record: %s, %d, %v\n", safeName, safeAge, safeTime)
    return nil
}
```

### Configuration with Optional Values

```go
type Config struct {
    Host    *string `yaml:"host"`
    Port    *int    `yaml:"port"`
    Timeout *int    `yaml:"timeout"`
    Debug   *bool   `yaml:"debug"`
}

func applyConfig(cfg Config) {
    // Apply with sensible defaults
    host := typeshift.Deref(cfg.Host)       // "" if not specified
    port := typeshift.Deref(cfg.Port)       // 0 if not specified
    timeout := typeshift.Deref(cfg.Timeout) // 0 if not specified
    debug := typeshift.Deref(cfg.Debug)     // false if not specified
    
    // Use default values where needed
    if host == "" {
        host = "localhost"
    }
    if port == 0 {
        port = 8080
    }
    
    fmt.Printf("Config: %s:%d, timeout=%d, debug=%v\n", host, port, timeout, debug)
}
```

## 🔄 Migration Guide

### From Manual Pointer Handling

**Before:**
```go
func handleOptionalValue(ptr *string) string {
    if ptr == nil {
        return ""
    }
    return *ptr
}

func createPointer(value string) *string {
    return &value
}
```

**After:**
```go
func handleOptionalValue(ptr *string) string {
    return typeshift.Deref(ptr)
}

func createPointer(value string) *string {
    return typeshift.Ptr(value)
}
```

### Gradual Adoption

You can adopt TypeShift gradually:

1. **Start with new code** using generic functions
2. **Replace repetitive patterns** in existing code
3. **Keep existing type-specific calls** - they continue to work unchanged

## 🎯 Design Principles

- **Type Safety**: Leverage Go's type system for compile-time guarantees
- **Zero Runtime Overhead**: Generic functions compile to the same code as manual implementations
- **Backward Compatibility**: Existing APIs remain unchanged
- **Simplicity**: Reduce boilerplate without hiding behavior
- **Consistency**: Uniform behavior across all supported types

## 📊 Performance

TypeShift has zero runtime overhead:

- Generic functions are specialized at compile time
- No reflection or runtime type checking
- Identical performance to manual pointer operations
- Memory efficient with no additional allocations

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🏷️ Version History

- **v1.1.0**: Added generic functions and enhanced type support
- **v1.0.0**: Initial release with type-specific functions

---

**TypeShift** - Making pointer operations in Go simple, safe, and elegant. ✨
