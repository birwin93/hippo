package api_keys

import (
	"database/sql"

	"github.com/birwin93/hippo/demo/models"
)

type Create struct {
	Key models.ApiKey
}

func NewCreate(userId int64, token string) *Create {
	query := Create{}
	query.Key = models.ApiKey{UserId: userId, Token: token}
	return &query
}

func (s *Create) Query() string {
	return "INSERT INTO api_keys (user_id, token) VALUES ($1, $2)"
}

func (s *Create) QueryVals() []interface{} {
	return []interface{}{s.Key.UserId, s.Key.Token}
}

func (s *Create) Scan(row *sql.Row) error {
	return nil
}
