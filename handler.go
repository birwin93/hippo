package hippo

import (
	"encoding/json"
)

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

func (h *Handler) WriteJSON(ct *Context, data interface{}) *Error {
	err := json.NewEncoder(ct.ResponseWriter).Encode(data)
	if err != nil {
		return NewError(300, "Error packing json")
	}
	return nil
}
