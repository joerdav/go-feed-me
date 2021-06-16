package app

import (
	"container/templates"
	"container/types"
	"context"
	"net/http"

	"github.com/joe-davidson1802/hotwirehandler"
)

type ContainerModel struct {
	url string
}

func (m ContainerModel) ModelName() string { return "ContainerModel" }

type HomeHandler struct {
	Config types.Config
}

func (h HomeHandler) CanHandleModel(m string) bool {
	return m == ContainerModel{}.ModelName()
}

func (h HomeHandler) HandleRequest(w http.ResponseWriter, r *http.Request) (error, hotwirehandler.Model) {
	return nil, ContainerModel{
		url: r.RequestURI,
	}
}

func (h HomeHandler) RenderPage(ctx context.Context, m hotwirehandler.Model, w http.ResponseWriter) error {
	cm := m.(ContainerModel)

	w.Header().Add("Content-Type", "text/html")

	var url string
	if cm.url == "/" || cm.url == "" {
		url = "/apps" + h.Config.GetDefaultApp().Url
	} else {
		url = "/apps" + cm.url
	}

	t := templates.LayoutTemplate(url, h.Config)
	err := t.Render(ctx, w)
	return err
}

func (h HomeHandler) RenderStream(ctx context.Context, m hotwirehandler.Model, w http.ResponseWriter) error {
	err := templates.NavTemplate(h.Config.Apps).Render(ctx, w)
	return err
}
