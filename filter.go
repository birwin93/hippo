package hippo

type FilterHandler interface {
	Handle(context *Context, next HandlerInterface) *Error
}

type FilterHandlerFunc func(context *Context, next HandlerInterface) *Error

func (f FilterHandlerFunc) Handle(context *Context, next HandlerInterface) *Error {
	return f(context, next)
}

func (f FilterHandlerFunc) Filters() []FilterHandler {
	return make([]FilterHandler, 0)
}

type Filter struct {
	handler FilterHandler
	next    *Filter
}

func (f *Filter) Handle(context *Context) *Error {
	return f.handler.Handle(context, f.next)
}

func (f *Filter) Filters() []FilterHandler {
	return make([]FilterHandler, 0)
}
