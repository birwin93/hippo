package sessions

import (
	"github.com/birwin93/hippo"
	"github.com/birwin93/hippo/demo/filters"
	"github.com/birwin93/hippo/demo/transactions/api_keys"
)

type LogoutHandler struct {
	hippo.Handler
}

func (l *LogoutHandler) Handle(ct *hippo.Context) *hippo.Error {
	destroyApiKey := api_keys.NewDestroy(ct.Params.Get("token"))
	err := ct.DB.Destroy(destroyApiKey)
	if err != nil {
		return hippo.NewError(500, "Error signing out")
	} else {
		ct.ResponseWriter.WriteHeader(200)
		return nil
	}
}

func (h *LogoutHandler) Filters() []hippo.FilterHandler {
	return []hippo.FilterHandler{&filters.Auth{}}
}
