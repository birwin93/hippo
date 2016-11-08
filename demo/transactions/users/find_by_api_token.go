package users

import (
	"database/sql"

	"github.com/birwin93/hippo/demo/models"
)

type FindByApiToken struct {
	Token string
	User  models.User
}

func NewFindByApiToken(token string) *FindByApiToken {
	query := FindByApiToken{}
	query.Token = token
	query.User = models.User{}
	return &query
}

func (q *FindByApiToken) Query() string {
	return "SELECT u.id, u.username " +
		"FROM users u, api_keys ak " +
		"WHERE ak.user_id = u.id AND ak.token = $1 " +
		"LIMIT 1"
}

func (q *FindByApiToken) QueryVals() []interface{} {
	return []interface{}{q.Token}
}

func (q *FindByApiToken) Scan(row *sql.Row) error {
	return row.Scan(&q.User.Id, &q.User.Username)
}
