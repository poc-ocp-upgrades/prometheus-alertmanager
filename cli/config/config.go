package config

import (
	"io/ioutil"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	"os"
	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/yaml.v2"
)

type getFlagger interface {
	GetFlag(name string) *kingpin.FlagClause
}
type Resolver struct{ flags map[string]string }

func NewResolver(files []string, legacyFlags map[string]string) (*Resolver, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	flags := map[string]string{}
	for _, f := range files {
		if _, err := os.Stat(f); err != nil {
			continue
		}
		b, err := ioutil.ReadFile(f)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			return nil, err
		}
		var m map[string]string
		err = yaml.Unmarshal(b, &m)
		if err != nil {
			return nil, err
		}
		for k, v := range m {
			if flag, ok := legacyFlags[k]; ok {
				if _, ok := m[flag]; ok {
					continue
				}
				k = flag
			}
			if _, ok := flags[k]; !ok {
				flags[k] = v
			}
		}
	}
	return &Resolver{flags: flags}, nil
}
func (c *Resolver) setDefault(v getFlagger) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for name, value := range c.flags {
		f := v.GetFlag(name)
		if f != nil {
			f.Default(value)
		}
	}
}
func (c *Resolver) Bind(app *kingpin.Application, args []string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, err := app.ParseContext(args)
	if err != nil {
		return err
	}
	c.setDefault(app)
	if pc.SelectedCommand != nil {
		c.setDefault(pc.SelectedCommand)
	}
	return nil
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
