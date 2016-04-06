package hippo

import (
	"database/sql"
	_ "github.com/lib/pq"
	"net/http"
	"net/url"
)

type Context struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	Params         url.Values
	DB             *sql.DB
	Config         *Config
}
