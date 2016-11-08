package api_keys

import "database/sql"

type Destroy struct {
	Token string
}

func NewDestroy(token string) *Destroy {
	return &Destroy{token}
}

func (s *Destroy) Query() string {
	return "DELETE FROM api_keys WHERE token = $1"
}

func (s *Destroy) QueryVals() []interface{} {
	return []interface{}{s.Token}
}

func (s *Destroy) Scan(row *sql.Row) error {
	return nil
}
