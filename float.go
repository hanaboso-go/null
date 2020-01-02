package null

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
)

type Float64 struct {
	Float64 float64
	Valid   bool
}

// Scan implements the Scanner interface.
func (n *Float64) Scan(value interface{}) error {
	var local sql.NullFloat64
	err := local.Scan(value)
	*n = Float64{
		Float64: local.Float64,
		Valid:   local.Valid,
	}
	return err
}

// Value implements the driver Valuer interface.
func (n Float64) Value() (driver.Value, error) {
	return sql.NullFloat64{
		Float64: n.Float64,
		Valid:   n.Valid,
	}.Value()
}

func (n Float64) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return json.Marshal(nil)
	}

	return json.Marshal(n.Float64)
}

func (n *Float64) UnmarshalJSON(data []byte) error {
	var s *float64
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	*n = Float64{Valid: s != nil}
	if n.Valid {
		n.Float64 = *s
	}
	return nil
}

func (n Float64) Ptr() *float64 {
	if !n.Valid {
		return nil
	}
	return &n.Float64
}
