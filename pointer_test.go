package typeshift

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestPointerFunctions(t *testing.T) {
	// Int types
	assert.Equal(t, 10, *Int(10))
	assert.Equal(t, int8(8), *Int8(8))
	assert.Equal(t, int16(16), *Int16(16))
	assert.Equal(t, int32(32), *Int32(32))
	assert.Equal(t, int64(64), *Int64(64))

	// Uint types
	assert.Equal(t, uint(10), *Uint(10))
	assert.Equal(t, uint8(8), *Uint8(8))
	assert.Equal(t, uint16(16), *Uint16(16))
	assert.Equal(t, uint32(32), *Uint32(32))
	assert.Equal(t, uint64(64), *Uint64(64))

	// Float types
	assert.Equal(t, float32(3.14), *Float32(3.14))
	assert.Equal(t, 3.1415, *Float64(3.1415))

	// Other types
	assert.Equal(t, "hello", *String("hello"))
	assert.Equal(t, true, *Bool(true))

	// Time & UUID
	now := time.Now()
	assert.Equal(t, now, *Time(now))

	id := uuid.New()
	assert.Equal(t, id, *UUID(id))
}
