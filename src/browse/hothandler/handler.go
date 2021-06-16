package hotwirehandler

import (
	"context"
	"errors"
	"net/http"
	"strings"
)

type RequestHandler interface {
	HandleRequest(http.ResponseWriter, *http.Request) (error, Model)
	RenderPage(context.Context, Model, http.ResponseWriter) error
	RenderStream(context.Context, Model, http.ResponseWriter) error
	CanHandleModel(string) bool
}

type TurboHandler struct {
	handler RequestHandler
}

func New(h RequestHandler) TurboHandler {
	return TurboHandler{handler: h}
}

func (h TurboHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err, m := h.handler.HandleRequest(w, r)

	if err != nil {
		panic(err)
	}

	if !h.handler.CanHandleModel(m.ModelName()) {
		panic(errors.New("Handler cannot handler model of type " + m.ModelName()))
	}

	ct := r.Header.Get("Accept")

	if strings.Contains(ct, "vnd.turbo-stream.html") {
		err = h.handler.RenderStream(r.Context(), m, w)
	} else {
		err = h.handler.RenderPage(r.Context(), m, w)
	}

	if err != nil {
		panic(err)
	}
}
