package hippo

import (
	"net/http"
	"net/url"

	"github.com/birwin93/db"
)

type Context struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	Params         url.Values
	DB             *db.Database
	Config         *Config
	UserId         int64
}
