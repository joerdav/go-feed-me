package app

import (
	"container/hothandler"
	"container/templates"
	"container/types"
	"context"
	"net/http"
)

type EmptyModel struct {
}

func (m EmptyModel) ModelName() string { return "EmptyModel" }

type HomeHandler struct {
	Config types.Config
}

func (h HomeHandler) CanHandleModel(m string) bool {
	return m == EmptyModel{}.ModelName()
}

func (h HomeHandler) HandleRequest(w http.ResponseWriter, r *http.Request) (error, hothandler.Model) {
	return nil, EmptyModel{}
}

func (h HomeHandler) RenderPage(ctx context.Context, m hothandler.Model, w http.ResponseWriter) error {
	w.Header().Add("Content-Type", "text/html")
	t := templates.LayoutTemplate(h.Config)
	err := t.Render(ctx, w)
	return err
}

func (h HomeHandler) RenderStream(ctx context.Context, m hothandler.Model, w http.ResponseWriter) error {
	err := templates.NavTemplate(h.Config.Apps).Render(ctx, w)
	return err
}
