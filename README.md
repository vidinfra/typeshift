# typeshift

`typeshift` is a simple Go package that provides helper functions for working with pointers and nullable types. It's useful when working with database models, APIs, or any context where pointers are required or dereferencing needs to be safe and concise.

---

## ✨ Features

- Easily create pointers to primitive types
- Safely dereference pointers with default fallback values
- Handle common `sql.Null*` types
- Includes helpers for:
  - `int`, `int8`, `int16`, `int32`, `int64`
  - `uint`, `uint8`, `uint16`, `uint32`, `uint64`
  - `float32`, `float64`
  - `string`, `bool`
  - `time.Time`, `uuid.UUID`

---

## 📦 Installation

```bash
go get github.com/vidinfra/typeshift
```
