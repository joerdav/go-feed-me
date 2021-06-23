package app

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gorilla/mux"
)

var paths = []struct {
	path  string
	proxy string
}{
	{path: "/content", proxy: "http://content"},
	{path: "/apps/browse", proxy: "http://browse"},
	{path: "/apps/basket", proxy: "http://basket"},
	{path: "/apps/random", proxy: "http://random"},
	{path: "/apps/details", proxy: "http://details"},
}

func getDestination(path string) string {
	for _, p := range paths {
		if p.path == path || strings.HasPrefix(path, p.path+"/") {
			return p.proxy
		}
	}

	return "http://container"
}

func handleRequestAndRedirect(res http.ResponseWriter, req *http.Request) {
	url, _ := url.Parse(getDestination(req.RequestURI))

	proxy := httputil.NewSingleHostReverseProxy(url)

	req.URL.Host = url.Host
	req.URL.Scheme = url.Scheme
	req.Host = url.Host
	proxy.ServeHTTP(res, req)
}

func Run(env string) {
	r := mux.NewRouter()

	r.
		PathPrefix("/").
		HandlerFunc(handleRequestAndRedirect)

	fmt.Println("Listening: :80")

	err := http.ListenAndServe(":80", r)

	log.Fatal(err)
}
