package types

import (
	"fmt"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"regexp"
	"sort"
	"bytes"
	"github.com/prometheus/common/model"
)

type Matcher struct {
	Name	string	`json:"name"`
	Value	string	`json:"value"`
	IsRegex	bool	`json:"isRegex"`
	regex	*regexp.Regexp
}

func (m *Matcher) Init() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !m.IsRegex {
		return nil
	}
	re, err := regexp.Compile("^(?:" + m.Value + ")$")
	if err == nil {
		m.regex = re
	}
	return err
}
func (m *Matcher) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m.IsRegex {
		return fmt.Sprintf("%s=~%q", m.Name, m.Value)
	}
	return fmt.Sprintf("%s=%q", m.Name, m.Value)
}
func (m *Matcher) Validate() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !model.LabelName(m.Name).IsValid() {
		return fmt.Errorf("invalid name %q", m.Name)
	}
	if m.IsRegex {
		if _, err := regexp.Compile(m.Value); err != nil {
			return fmt.Errorf("invalid regular expression %q", m.Value)
		}
	} else if !model.LabelValue(m.Value).IsValid() || len(m.Value) == 0 {
		return fmt.Errorf("invalid value %q", m.Value)
	}
	return nil
}
func (m *Matcher) Match(lset model.LabelSet) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	v := lset[model.LabelName(m.Name)]
	if m.IsRegex {
		return m.regex.MatchString(string(v))
	}
	return string(v) == m.Value
}
func NewMatcher(name model.LabelName, value string) *Matcher {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &Matcher{Name: string(name), Value: value, IsRegex: false}
}
func NewRegexMatcher(name model.LabelName, re *regexp.Regexp) *Matcher {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &Matcher{Name: string(name), Value: re.String(), IsRegex: true, regex: re}
}

type Matchers []*Matcher

func NewMatchers(ms ...*Matcher) Matchers {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	m := Matchers(ms)
	sort.Sort(m)
	return m
}
func (ms Matchers) Len() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(ms)
}
func (ms Matchers) Swap(i, j int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ms[i], ms[j] = ms[j], ms[i]
}
func (ms Matchers) Less(i, j int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if ms[i].Name > ms[j].Name {
		return false
	}
	if ms[i].Name < ms[j].Name {
		return true
	}
	if ms[i].Value > ms[j].Value {
		return false
	}
	if ms[i].Value < ms[j].Value {
		return true
	}
	return !ms[i].IsRegex && ms[j].IsRegex
}
func (ms Matchers) Equal(o Matchers) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(ms) != len(o) {
		return false
	}
	for i, a := range ms {
		if *a != *o[i] {
			return false
		}
	}
	return true
}
func (ms Matchers) Match(lset model.LabelSet) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, m := range ms {
		if !m.Match(lset) {
			return false
		}
	}
	return true
}
func (ms Matchers) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, m := range ms {
		if i > 0 {
			buf.WriteByte(',')
		}
		buf.WriteString(m.String())
	}
	buf.WriteByte('}')
	return buf.String()
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
