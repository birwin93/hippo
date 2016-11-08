package users

import (
	"database/sql"

	"github.com/birwin93/hippo/demo/models"
)

type FindById struct {
	User models.User
}

func NewFindById(id int64) *FindById {
	query := FindById{}
	query.User = models.User{Id: id}
	return &query
}

func (q *FindById) Query() string {
	return "SELECT id, username FROM users WHERE id = $1 LIMIT 1"
}

func (q *FindById) QueryVals() []interface{} {
	return []interface{}{q.User.Id}
}

func (q *FindById) Scan(row *sql.Row) error {
	return row.Scan(&q.User.Id, &q.User.Username)
}
