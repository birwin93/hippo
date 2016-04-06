package hippo

import (
	"database/sql"
	"errors"
	"flag"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type App struct {
	router     *Router
	middleware *Middleware
	db         *sql.DB
	config     *Config
}

func InitApp() App {

	config := createConfig()

	db := NewDB(config)

	log.Println(config)

	middleware := NewMiddleware()
	middleware.Use(NewLogger())

	router := NewRouter()
	router.UseMiddleware(middleware)

	app := App{router, middleware, db, config}
	return app
}

func (a App) Add(pattern string, h HandlerInterface) {
	a.router.Add(pattern, h)
}

func (a App) Use(filter FilterHandler) {
	a.middleware.Use(filter)
}

func (a App) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	context := &Context{}
	context.Request = r
	context.ResponseWriter = rw
	context.DB = a.db
	context.Config = a.config
	err := a.router.Handle(context)
	if err != nil {
		http.Error(rw, err.Message, err.Code)
	}
}

func createConfig() *Config {
	var env string
	var configPath string
	flag.StringVar(&env, "env", "development", "env=development/test/production")
	flag.StringVar(&configPath, "config", "", "path to config.yml file")
	flag.Parse()

	if env != "development" &&
		env != "test" &&
		env != "production" {
		err := errors.New("incorrect env: valid values are development, test, production")
		panic(err)
	}

	if configPath == "" {
		err := errors.New("please provide a -config path")
		panic(err)
	}

	return NewConfig(env, configPath)
}
