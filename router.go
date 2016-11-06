package hippo

import (
	"net/url"
	"regexp"
	"strings"
)

type Route struct {
	regex   *regexp.Regexp
	params  map[int]string
	handler HandlerInterface
}

type Router struct {
	routes     []*Route
	middleware *Middleware
}

func NewRouter() *Router {
	return &Router{routes: make([]*Route, 0)}
}

func (r *Router) UseMiddleware(middleware *Middleware) {
	r.middleware = middleware
}

func (p *Router) Add(pattern string, h HandlerInterface) {
	parts := strings.Split(pattern, "/")

	j := 0
	params := make(map[int]string)
	for i, part := range parts {
		if strings.HasPrefix(part, ":") {
			expr := "([^/]+)"

			//a user may choose to override the defult expression
			// similar to expressjs: ‘/user/:id([0-9]+)’

			if index := strings.Index(part, "("); index != -1 {
				expr = part[index:]
				part = part[:index]
			}
			params[j] = part
			parts[i] = expr
			j++
		}
	}

	//recreate the url pattern, with parameters replaced
	//by regular expressions. then compile the regex

	pattern = strings.Join(parts, "/")
	regex, regexErr := regexp.Compile(pattern)
	if regexErr != nil {

		//TODO add error handling here to avoid panic
		panic(regexErr)
	}

	//now create the Route
	route := &Route{}
	route.regex = regex
	route.params = params
	route.handler = h

	p.routes = append(p.routes, route)

}

func (p *Router) Handle(context *Context) *Error {

	requestPath := context.Request.URL.Path

	//find a matching Route
	for _, route := range p.routes {

		//check if Route pattern matches url
		if !route.regex.MatchString(requestPath) {
			continue
		}

		//get submatches (params)
		matches := route.regex.FindStringSubmatch(requestPath)

		//double check that the Route matches the URL pattern.
		if len(matches[0]) != len(requestPath) {
			continue
		}

		params := context.Request.URL.Query()
		if len(route.params) > 0 {
			//add url parameters to the query param map
			for i, match := range matches[1:] {
				params.Add(route.params[i], match)
			}
			//reassemble query params and add to RawQuery
			context.Request.URL.RawQuery = url.Values(params).Encode() + "&" + context.Request.URL.RawQuery
		}

		context.Params = params

		//Invoke the request handler
		handler := route.handler
		middleware := NewMiddleware()
		middleware.UseMultiple(p.middleware.Handlers)
		middleware.UseMultiple(handler.Filters())
		middleware.UseHandler(handler)
		return middleware.Handle(context)
	}

	// Route not found
	return NewError(404, "Route could not be found")

}
