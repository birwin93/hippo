package hippo

type HandlerInterface interface {
	Handle(context *Context) *Error //process request
	Filters() []FilterHandler       //filters for request eg. Auth
}

type Handler struct{}

func (h *Handler) Handle(context *Context) *Error {
	return NewError(405, "Method not allowed")
}

func (h *Handler) Filters() []FilterHandler {
	return make([]FilterHandler, 0)
}
