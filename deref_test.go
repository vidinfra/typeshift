package typeshift

import (
	"database/sql"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestDerefFunctions(t *testing.T) {
	assert.Equal(t, 42, DerefInt(Int(42)))
	assert.Equal(t, 0, DerefInt(nil))

	assert.Equal(t, int64(123456), DerefInt64(Int64(123456)))
	assert.Equal(t, int64(0), DerefInt64(nil))

	assert.Equal(t, uint(100), DerefUint(Uint(100)))
	assert.Equal(t, uint(0), DerefUint(nil))

	assert.Equal(t, 3.14, DerefFloat64(Float64(3.14)))
	assert.Equal(t, 0.0, DerefFloat64(nil))

	assert.Equal(t, "Go", DerefString(String("Go")))
	assert.Equal(t, "", DerefString(nil))

	assert.Equal(t, true, DerefBool(Bool(true)))
	assert.Equal(t, false, DerefBool(nil))

	now := time.Now()
	assert.Equal(t, now, DerefTime(Time(now)))
	assert.Equal(t, time.Time{}, DerefTime(nil))

	id := uuid.New()
	assert.Equal(t, id, DerefUUID(UUID(id)))
	assert.Equal(t, uuid.Nil, DerefUUID(nil))
}

func TestDerefNullTypes(t *testing.T) {
	assert.Equal(t, "abc", DerefNullString(sql.NullString{String: "abc", Valid: true}))
	assert.Equal(t, "", DerefNullString(sql.NullString{Valid: false}))

	assert.Equal(t, int64(99), DerefNullInt64(sql.NullInt64{Int64: 99, Valid: true}))
	assert.Equal(t, int64(0), DerefNullInt64(sql.NullInt64{Valid: false}))

	assert.Equal(t, true, DerefNullBool(sql.NullBool{Bool: true, Valid: true}))
	assert.Equal(t, false, DerefNullBool(sql.NullBool{Valid: false}))

	assert.Equal(t, 5.55, DerefNullFloat64(sql.NullFloat64{Float64: 5.55, Valid: true}))
	assert.Equal(t, 0.0, DerefNullFloat64(sql.NullFloat64{Valid: false}))

	now := time.Now()
	assert.Equal(t, now, DerefNullTime(sql.NullTime{Time: now, Valid: true}))
	assert.Equal(t, time.Time{}, DerefNullTime(sql.NullTime{Valid: false}))
}

func TestGenericDeref(t *testing.T) {
	// Test generic Deref with various types
	assert.Equal(t, 42, Deref(Int(42)))
	assert.Equal(t, 0, Deref[int](nil))

	assert.Equal(t, "hello", Deref(String("hello")))
	assert.Equal(t, "", Deref[string](nil))

	assert.Equal(t, true, Deref(Bool(true)))
	assert.Equal(t, false, Deref[bool](nil))

	assert.Equal(t, int64(123), Deref(Int64(123)))
	assert.Equal(t, int64(0), Deref[int64](nil))

	assert.Equal(t, 3.14, Deref(Float64(3.14)))
	assert.Equal(t, 0.0, Deref[float64](nil))

	// Test with time and uuid
	now := time.Now()
	assert.Equal(t, now, Deref(Time(now)))
	assert.Equal(t, time.Time{}, Deref[time.Time](nil))

	id := uuid.New()
	assert.Equal(t, id, Deref(UUID(id)))
	assert.Equal(t, uuid.UUID{}, Deref[uuid.UUID](nil))
}

func TestDerefNewNumericTypes(t *testing.T) {
	// Test int8
	assert.Equal(t, int8(127), DerefInt8(Int8(127)))
	assert.Equal(t, int8(0), DerefInt8(nil))

	// Test int16
	assert.Equal(t, int16(32767), DerefInt16(Int16(32767)))
	assert.Equal(t, int16(0), DerefInt16(nil))

	// Test int32
	assert.Equal(t, int32(2147483647), DerefInt32(Int32(2147483647)))
	assert.Equal(t, int32(0), DerefInt32(nil))

	// Test uint8
	assert.Equal(t, uint8(255), DerefUint8(Uint8(255)))
	assert.Equal(t, uint8(0), DerefUint8(nil))

	// Test uint16
	assert.Equal(t, uint16(65535), DerefUint16(Uint16(65535)))
	assert.Equal(t, uint16(0), DerefUint16(nil))

	// Test uint32
	assert.Equal(t, uint32(4294967295), DerefUint32(Uint32(4294967295)))
	assert.Equal(t, uint32(0), DerefUint32(nil))

	// Test uint64
	assert.Equal(t, uint64(18446744073709551615), DerefUint64(Uint64(18446744073709551615)))
	assert.Equal(t, uint64(0), DerefUint64(nil))

	// Test float32
	assert.Equal(t, float32(3.14), DerefFloat32(Float32(3.14)))
	assert.Equal(t, float32(0.0), DerefFloat32(nil))
}

func TestDerefEnumToString(t *testing.T) {
	type Status string
	const (
		Active   Status = "active"
		Inactive Status = "inactive"
	)

	type UserRole string
	const (
		Admin UserRole = "admin"
		User  UserRole = "user"
	)

	// Test with Status enum
	activeStatus := Active
	assert.Equal(t, "active", DerefEnumToString(&activeStatus))
	assert.Equal(t, "", DerefEnumToString[Status](nil))

	inactiveStatus := Inactive
	assert.Equal(t, "inactive", DerefEnumToString(&inactiveStatus))

	// Test with UserRole enum
	adminRole := Admin
	assert.Equal(t, "admin", DerefEnumToString(&adminRole))
	assert.Equal(t, "", DerefEnumToString[UserRole](nil))

	userRole := User
	assert.Equal(t, "user", DerefEnumToString(&userRole))
}

func TestDerefMapStringString(t *testing.T) {
	// Test with nil map
	assert.Nil(t, DerefMapStringString(nil))

	// Test with valid map
	originalMap := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}

	clonedMap := DerefMapStringString(&originalMap)

	// Assert maps are equal
	assert.Equal(t, originalMap, clonedMap)

	// Assert they are different objects (deep copy)
	clonedMap["key1"] = "modified"
	assert.Equal(t, "value1", originalMap["key1"]) // Original should be unchanged
	assert.Equal(t, "modified", clonedMap["key1"]) // Clone should be modified

	// Test with empty map
	emptyMap := map[string]string{}
	clonedEmpty := DerefMapStringString(&emptyMap)
	assert.Equal(t, emptyMap, clonedEmpty)
	assert.NotNil(t, clonedEmpty) // Should not be nil, should be empty map
}
