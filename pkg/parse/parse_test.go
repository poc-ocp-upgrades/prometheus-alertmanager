package parse

import (
	"reflect"
	"testing"
	"github.com/prometheus/prometheus/pkg/labels"
)

func TestMatchers(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testCases := []struct {
		input	string
		want	[]*labels.Matcher
		err	error
	}{{input: `{foo="bar"}`, want: func() []*labels.Matcher {
		ms := []*labels.Matcher{}
		m, _ := labels.NewMatcher(labels.MatchEqual, "foo", "bar")
		return append(ms, m)
	}()}, {input: `{foo=~"bar.*"}`, want: func() []*labels.Matcher {
		ms := []*labels.Matcher{}
		m, _ := labels.NewMatcher(labels.MatchRegexp, "foo", "bar.*")
		return append(ms, m)
	}()}, {input: `{foo!="bar"}`, want: func() []*labels.Matcher {
		ms := []*labels.Matcher{}
		m, _ := labels.NewMatcher(labels.MatchNotEqual, "foo", "bar")
		return append(ms, m)
	}()}, {input: `{foo!~"bar.*"}`, want: func() []*labels.Matcher {
		ms := []*labels.Matcher{}
		m, _ := labels.NewMatcher(labels.MatchNotRegexp, "foo", "bar.*")
		return append(ms, m)
	}()}, {input: `{foo="bar", baz!="quux"}`, want: func() []*labels.Matcher {
		ms := []*labels.Matcher{}
		m, _ := labels.NewMatcher(labels.MatchEqual, "foo", "bar")
		m2, _ := labels.NewMatcher(labels.MatchNotEqual, "baz", "quux")
		return append(ms, m, m2)
	}()}, {input: `{foo="bar", baz!~"quux.*"}`, want: func() []*labels.Matcher {
		ms := []*labels.Matcher{}
		m, _ := labels.NewMatcher(labels.MatchEqual, "foo", "bar")
		m2, _ := labels.NewMatcher(labels.MatchNotRegexp, "baz", "quux.*")
		return append(ms, m, m2)
	}()}, {input: `{foo="bar",baz!~".*quux", derp="wat"}`, want: func() []*labels.Matcher {
		ms := []*labels.Matcher{}
		m, _ := labels.NewMatcher(labels.MatchEqual, "foo", "bar")
		m2, _ := labels.NewMatcher(labels.MatchNotRegexp, "baz", ".*quux")
		m3, _ := labels.NewMatcher(labels.MatchEqual, "derp", "wat")
		return append(ms, m, m2, m3)
	}()}, {input: `{foo="bar", baz!="quux", derp="wat"}`, want: func() []*labels.Matcher {
		ms := []*labels.Matcher{}
		m, _ := labels.NewMatcher(labels.MatchEqual, "foo", "bar")
		m2, _ := labels.NewMatcher(labels.MatchNotEqual, "baz", "quux")
		m3, _ := labels.NewMatcher(labels.MatchEqual, "derp", "wat")
		return append(ms, m, m2, m3)
	}()}, {input: `{foo="bar", baz!~".*quux.*", derp="wat"}`, want: func() []*labels.Matcher {
		ms := []*labels.Matcher{}
		m, _ := labels.NewMatcher(labels.MatchEqual, "foo", "bar")
		m2, _ := labels.NewMatcher(labels.MatchNotRegexp, "baz", ".*quux.*")
		m3, _ := labels.NewMatcher(labels.MatchEqual, "derp", "wat")
		return append(ms, m, m2, m3)
	}()}, {input: `{foo="bar", instance=~"some-api.*"}`, want: func() []*labels.Matcher {
		ms := []*labels.Matcher{}
		m, _ := labels.NewMatcher(labels.MatchEqual, "foo", "bar")
		m2, _ := labels.NewMatcher(labels.MatchRegexp, "instance", "some-api.*")
		return append(ms, m, m2)
	}()}, {input: `{foo=""}`, want: func() []*labels.Matcher {
		ms := []*labels.Matcher{}
		m, _ := labels.NewMatcher(labels.MatchEqual, "foo", "")
		return append(ms, m)
	}()}, {input: `{foo="bar,quux", job="job1"}`, want: func() []*labels.Matcher {
		ms := []*labels.Matcher{}
		m, _ := labels.NewMatcher(labels.MatchEqual, "foo", "bar,quux")
		m2, _ := labels.NewMatcher(labels.MatchEqual, "job", "job1")
		return append(ms, m, m2)
	}()}}
	for i, tc := range testCases {
		got, err := Matchers(tc.input)
		if tc.err != err {
			t.Fatalf("error not equal (i=%d):\ngot  %v\nwant %v", i, err, tc.err)
		}
		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("labels not equal (i=%d):\ngot  %v\nwant %v", i, got, tc.want)
		}
	}
}
