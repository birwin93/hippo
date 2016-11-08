package users

import (
	"database/sql"

	"github.com/birwin93/familyfeud/core/models"
)

type FindByUsername struct {
	Username string
	User     models.User
}

func NewFindByUsername(username string) *FindByUsername {
	query := FindByUsername{}
	query.Username = username
	query.User = models.User{}
	return &query
}

func (q *FindByUsername) Query() string {
	return "SELECT id, username, password FROM users WHERE username = $1 LIMIT 1"
}

func (q *FindByUsername) QueryVals() []interface{} {
	return []interface{}{q.Username}
}

func (q *FindByUsername) Scan(row *sql.Row) error {
	return row.Scan(&q.User.Id, &q.User.Username, &q.User.Password)
}
