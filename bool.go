package null

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
)

type Bool struct {
	Bool  bool
	Valid bool
}

// Scan implements the Scanner interface.
func (n *Bool) Scan(value interface{}) error {
	var local sql.NullBool
	err := local.Scan(value)
	*n = Bool{
		Bool:  local.Bool,
		Valid: local.Valid,
	}
	return err
}

// Value implements the driver Valuer interface.
func (n Bool) Value() (driver.Value, error) {
	return sql.NullBool{
		Bool:  n.Bool,
		Valid: n.Valid,
	}.Value()
}

func (n Bool) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return json.Marshal(nil)
	}

	return json.Marshal(n.Bool)
}

func (n *Bool) UnmarshalJSON(data []byte) error {
	var s *bool
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	*n = Bool{Valid: s != nil}
	if n.Valid {
		n.Bool = *s
	}
	return nil
}

func (n Bool) Ptr() *bool {
	if !n.Valid {
		return nil
	}
	return &n.Bool
}
