package typeshift

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// Deref safely dereferences a pointer, returning the zero value if nil
func Deref[T any](p *T) T {
	if p == nil {
		var zero T
		return zero
	}
	return *p
}

// Legacy functions for backward compatibility
func DerefInt(p *int) int {
	return Deref(p)
}

func DerefInt8(ptr *int8) int8 {
	return Deref(ptr)
}

func DerefInt16(ptr *int16) int16 {
	return Deref(ptr)
}

func DerefInt32(ptr *int32) int32 {
	return Deref(ptr)
}

func DerefInt64(p *int64) int64 {
	return Deref(p)
}

func DerefUint(p *uint) uint {
	return Deref(p)
}

func DerefUint8(p *uint8) uint8 {
	return Deref(p)
}

func DerefUint16(p *uint16) uint16 {
	return Deref(p)
}

func DerefUint32(p *uint32) uint32 {
	return Deref(p)
}

func DerefUint64(p *uint64) uint64 {
	return Deref(p)
}

func DerefFloat32(p *float32) float32 {
	return Deref(p)
}

func DerefFloat64(p *float64) float64 {
	return Deref(p)
}

func DerefString(p *string) string {
	return Deref(p)
}

func DerefBool(p *bool) bool {
	return Deref(p)
}

func DerefTime(p *time.Time) time.Time {
	if p == nil {
		return time.Time{}
	}
	return *p
}

func DerefUUID(p *uuid.UUID) uuid.UUID {
	if p == nil {
		return uuid.Nil
	}
	return *p
}

func DerefNullString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}

func DerefNullInt64(ns sql.NullInt64) int64 {
	if ns.Valid {
		return ns.Int64
	}
	return 0
}

func DerefNullBool(nb sql.NullBool) bool {
	if nb.Valid {
		return nb.Bool
	}
	return false
}

func DerefNullFloat64(nf sql.NullFloat64) float64 {
	if nf.Valid {
		return nf.Float64
	}
	return 0.0
}

func DerefNullTime(nt sql.NullTime) time.Time {
	if nt.Valid {
		return nt.Time
	}
	return time.Time{}
}

func DerefEnumToString[T ~string](value *T) string {
	if value == nil {
		return ""
	}
	return string(*value)
}

func DerefMapStringString(ptr *map[string]string) map[string]string {
	if ptr == nil {
		return nil
	}
	clone := make(map[string]string, len(*ptr))
	for k, v := range *ptr {
		clone[k] = v
	}
	return clone
}
