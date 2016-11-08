package sessions

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/birwin93/hippo"
	"github.com/birwin93/hippo/demo/transactions/api_keys"
	"github.com/birwin93/hippo/demo/transactions/users"
)

type LoginHandler struct {
	hippo.Handler
}

func (l *LoginHandler) Handle(ct *hippo.Context) *hippo.Error {
	findUserByUsername := users.NewFindByUsername(ct.Params.Get("username"))
	err := ct.DB.Select(findUserByUsername)
	if err != nil {
		return hippo.NewError(300, "Could not find user with username")
	}

	if findUserByUsername.User.Password == ct.Params.Get("password") {
		token, _ := generateToken()
		createApiKey := api_keys.NewCreate(findUserByUsername.User.Id, token)
		err = ct.DB.Insert(createApiKey)
		if err != nil {
			return hippo.NewError(300, "Error generating api key")
		}
		return l.WriteJSON(ct, createApiKey.Key)
	} else {
		return hippo.NewError(300, "Username/password combo did not match")
	}
}

func generateRandomBytes() ([]byte, error) {
	b := make([]byte, 24)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func generateToken() (string, error) {
	bytes, err := generateRandomBytes()
	return base64.URLEncoding.EncodeToString(bytes), err
}
