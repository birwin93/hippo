package hippo

import (
	"net/http"
	"net/url"
)

type Context struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	Params         url.Values
	DB             *Database
	Config         *Config
	UserId         int64
}
