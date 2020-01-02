package null

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
)

type String struct {
	String string
	Valid  bool
}

// Scan implements the Scanner interface.
func (n *String) Scan(value interface{}) error {
	var local sql.NullString
	err := local.Scan(value)
	*n = String{
		String: local.String,
		Valid:  local.Valid,
	}
	return err
}

// Value implements the driver Valuer interface.
func (n String) Value() (driver.Value, error) {
	return sql.NullString{
		String: n.String,
		Valid:  n.Valid,
	}.Value()
}

func (n String) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return json.Marshal(nil)
	}

	return json.Marshal(n.String)
}

func (n *String) UnmarshalJSON(data []byte) error {
	var s *string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	*n = String{Valid: s != nil}
	if n.Valid {
		n.String = *s
	}
	return nil
}

func (n String) Ptr() *string {
	if !n.Valid {
		return nil
	}
	return &n.String
}
