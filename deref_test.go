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
