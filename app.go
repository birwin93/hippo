package hippo

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/birwin93/db"

	_ "github.com/lib/pq"
)

type App struct {
	router     *Router
	middleware *Middleware
	db         *db.Database
	config     *Config
}

func NewApp() App {

	config := createConfig()

	dbConfig := db.Config{
		Name:      config.Database.Name,
		User:      config.Database.User,
		EnableSSL: false,
	}
	db, err := db.NewDB(dbConfig)
	if err != nil {
		panic(err)
	}

	middleware := NewMiddleware()
	middleware.Use(NewLogger())

	router := NewRouter()
	router.UseMiddleware(middleware)

	app := App{router, middleware, db, config}
	return app
}

func (a App) Start(host string, port string) {
	log.Println("Starting up server")
	address := fmt.Sprintf("%s:%s", host, port)
	http.ListenAndServe(address, a)
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
