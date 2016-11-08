package users

import (
	"log"
	"strconv"

	"github.com/birwin93/hippo"
	"github.com/birwin93/hippo/demo/transactions/users"
)

type FindByIdHandler struct {
	hippo.Handler
}

func (h *FindByIdHandler) Handle(ct *hippo.Context) *hippo.Error {
	gameId, _ := strconv.ParseInt(ct.Params.Get("id"), 10, 64)
	findById := users.NewFindById(gameId)
	err := ct.DB.Select(findById)
	if err != nil {
		log.Println(err)
		return hippo.NewError(500, "Could not retrieve game with id")
	}
	return h.WriteJSON(ct, findById.User)
}
