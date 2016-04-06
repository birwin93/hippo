package hippo

type Middleware struct {
	filter   Filter
	Handlers []FilterHandler
}

func NewMiddleware() *Middleware {
	return &Middleware{
		Handlers: make([]FilterHandler, 0),
	}
}

func (m *Middleware) Handle(context *Context) *Error {
	return m.filter.Handle(context)
}

func (m *Middleware) UseMultiple(handlers []FilterHandler) {
	m.Handlers = append(m.Handlers, handlers...)
	m.filter = build(m.Handlers)
}

func (m *Middleware) Use(handler FilterHandler) {
	m.Handlers = append(m.Handlers, handler)
	m.filter = build(m.Handlers)
}

func (m *Middleware) UseHandler(handler HandlerInterface) {
	filterHandler := FilterHandlerFunc(func(context *Context, next HandlerInterface) *Error {
		err := handler.Handle(context)
		if err != nil {
			return err
		}
		return next.Handle(context)
	})
	m.Use(filterHandler)
}

func build(handlers []FilterHandler) Filter {
	var next Filter
	if len(handlers) == 0 {
		return voidFilter()
	} else if len(handlers) > 1 {
		next = build(handlers[1:])
	} else {
		next = voidFilter()
	}
	return Filter{handler: handlers[0], next: &next}
}

func voidFilter() Filter {
	return Filter{
		FilterHandlerFunc(func(context *Context, next HandlerInterface) *Error { return nil }),
		&Filter{},
	}
}
