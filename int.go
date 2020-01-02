package null

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
)

type Int64 struct {
	Int64 int64
	Valid bool
}

// Scan implements the Scanner interface.
func (n *Int64) Scan(value interface{}) error {
	var local sql.NullInt64
	err := local.Scan(value)
	*n = Int64{
		Int64: local.Int64,
		Valid: local.Valid,
	}
	return err
}

// Value implements the driver Valuer interface.
func (n Int64) Value() (driver.Value, error) {
	return sql.NullInt64{
		Int64: n.Int64,
		Valid: n.Valid,
	}.Value()
}

func (n Int64) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return json.Marshal(nil)
	}

	return json.Marshal(n.Int64)
}

func (n *Int64) UnmarshalJSON(data []byte) error {
	var s *int64
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	*n = Int64{Valid: s != nil}
	if n.Valid {
		n.Int64 = *s
	}
	return nil
}

func (n Int64) Ptr() *int64 {
	if !n.Valid {
		return nil
	}
	return &n.Int64
}

type Int32 struct {
	Int32 int32
	Valid bool
}

// Scan implements the Scanner interface.
func (n *Int32) Scan(value interface{}) error {
	var local sql.NullInt32
	err := local.Scan(value)
	*n = Int32{
		Int32: local.Int32,
		Valid: local.Valid,
	}
	return err
}

// Value implements the driver Valuer interface.
func (n Int32) Value() (driver.Value, error) {
	return sql.NullInt32{
		Int32: n.Int32,
		Valid: n.Valid,
	}.Value()
}

func (n Int32) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return json.Marshal(nil)
	}

	return json.Marshal(n.Int32)
}

func (n *Int32) UnmarshalJSON(data []byte) error {
	var s *int32
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	*n = Int32{Valid: s != nil}
	if n.Valid {
		n.Int32 = *s
	}
	return nil
}

func (n Int32) Ptr() *int32 {
	if !n.Valid {
		return nil
	}
	return &n.Int32
}
