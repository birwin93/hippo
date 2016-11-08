package filters

import (
	"github.com/birwin93/hippo"
	"github.com/birwin93/hippo/demo/transactions/users"
)

type Auth struct{}

func (l *Auth) Handle(ct *hippo.Context, next hippo.HandlerInterface) *hippo.Error {
	token := ct.Params.Get("token")
	findByToken := users.NewFindByApiToken(token)
	err := ct.DB.Select(findByToken)
	if err != nil {
		return hippo.NewError(403, "Unauthorized")
	}
	ct.UserId = findByToken.User.Id
	return next.Handle(ct)
}
