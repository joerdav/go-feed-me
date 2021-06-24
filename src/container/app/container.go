package app

import (
	"container/templates"
	"container/types"
	"log"
	"net/http"
)

type ContainerHandler struct {
	Config types.Config
}

func (h ContainerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")

	var url string
	if r.RequestURI == "/" || r.RequestURI == "" {
		url = "/apps" + h.Config.GetDefaultApp().Url
	} else {
		url = "/apps" + r.RequestURI
	}

	t := templates.LayoutTemplate(url, h.Config)
	err := t.Render(r.Context(), w)

	if err != nil {
		log.Fatal(err)
	}
}
