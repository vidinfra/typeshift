package typeshift

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

func DerefInt(p *int) int {
	if p == nil {
		return 0
	}
	return *p
}

func DerefInt64(p *int64) int64 {
	if p == nil {
		return 0
	}
	return *p
}

func DerefUint(p *uint) uint {
	if p == nil {
		return 0
	}
	return *p
}

func DerefFloat64(p *float64) float64 {
	if p == nil {
		return 0.0
	}
	return *p
}

func DerefString(p *string) string {
	if p == nil {
		return ""
	}
	return *p
}

func DerefBool(p *bool) bool {
	if p == nil {
		return false
	}
	return *p
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
