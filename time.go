package null

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"time"
)

type Time struct {
	Time  time.Time
	Valid bool
}

// Scan implements the Scanner interface.
func (n *Time) Scan(value interface{}) error {
	var local sql.NullTime
	err := local.Scan(value)
	*n = Time{
		Time:  local.Time,
		Valid: local.Valid,
	}
	return err
}

// Value implements the driver Valuer interface.
func (n Time) Value() (driver.Value, error) {
	return sql.NullTime{
		Time:  n.Time,
		Valid: n.Valid,
	}.Value()
}

func (n Time) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return json.Marshal(nil)
	}

	return json.Marshal(n.Time)
}

func (n *Time) UnmarshalJSON(data []byte) error {
	var s *time.Time
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	*n = Time{Valid: s != nil}
	if n.Valid {
		n.Time = *s
	}
	return nil
}

func (n Time) Ptr() *time.Time {
	if !n.Valid {
		return nil
	}
	return &n.Time
}
