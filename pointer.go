package typeshift

import (
	"time"

	"github.com/google/uuid"
)

// Ptr functions to create pointers from values
func Ptr[T any](v T) *T { return &v }

// Int types
func Int(v int) *int       { return &v }
func Int8(v int8) *int8    { return &v }
func Int16(v int16) *int16 { return &v }
func Int32(v int32) *int32 { return &v }
func Int64(v int64) *int64 { return &v }

// Uint types
func Uint(v uint) *uint       { return &v }
func Uint8(v uint8) *uint8    { return &v }
func Uint16(v uint16) *uint16 { return &v }
func Uint32(v uint32) *uint32 { return &v }
func Uint64(v uint64) *uint64 { return &v }

// Float types
func Float32(v float32) *float32 { return &v }
func Float64(v float64) *float64 { return &v }

// Other basic types
func String(v string) *string { return &v }
func Bool(v bool) *bool       { return &v }

// Special types
func Time(v time.Time) *time.Time { return &v }
func UUID(v uuid.UUID) *uuid.UUID { return &v }
