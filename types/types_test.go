package types

import (
	"reflect"
	"sort"
	"strconv"
	"testing"
	"time"
	"github.com/prometheus/common/model"
	"github.com/stretchr/testify/require"
)

func TestAlertMerge(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	now := time.Now()
	pairs := []struct{ A, B, Res *Alert }{{A: &Alert{Alert: model.Alert{StartsAt: now.Add(-2 * time.Minute), EndsAt: now.Add(2 * time.Minute)}, UpdatedAt: now, Timeout: true}, B: &Alert{Alert: model.Alert{StartsAt: now.Add(-time.Minute), EndsAt: now.Add(3 * time.Minute)}, UpdatedAt: now.Add(time.Minute), Timeout: true}, Res: &Alert{Alert: model.Alert{StartsAt: now.Add(-2 * time.Minute), EndsAt: now.Add(3 * time.Minute)}, UpdatedAt: now.Add(time.Minute), Timeout: true}}, {A: &Alert{Alert: model.Alert{StartsAt: now.Add(-time.Minute), EndsAt: now.Add(3 * time.Minute)}, UpdatedAt: now, Timeout: true}, B: &Alert{Alert: model.Alert{StartsAt: now, EndsAt: now.Add(2 * time.Minute)}, UpdatedAt: now.Add(time.Minute)}, Res: &Alert{Alert: model.Alert{StartsAt: now.Add(-time.Minute), EndsAt: now.Add(2 * time.Minute)}, UpdatedAt: now.Add(time.Minute)}}, {A: &Alert{Alert: model.Alert{StartsAt: now.Add(-time.Minute), EndsAt: now.Add(3 * time.Minute)}, UpdatedAt: now}, B: &Alert{Alert: model.Alert{StartsAt: now, EndsAt: now.Add(2 * time.Minute)}, UpdatedAt: now.Add(time.Minute), Timeout: true}, Res: &Alert{Alert: model.Alert{StartsAt: now.Add(-time.Minute), EndsAt: now.Add(3 * time.Minute)}, UpdatedAt: now.Add(time.Minute), Timeout: true}}, {A: &Alert{Alert: model.Alert{StartsAt: now.Add(-time.Minute), EndsAt: now.Add(3 * time.Minute)}, UpdatedAt: now}, B: &Alert{Alert: model.Alert{StartsAt: now, EndsAt: now.Add(2 * time.Minute)}, UpdatedAt: now.Add(time.Minute)}, Res: &Alert{Alert: model.Alert{StartsAt: now.Add(-time.Minute), EndsAt: now.Add(3 * time.Minute)}, UpdatedAt: now.Add(time.Minute)}}, {A: &Alert{Alert: model.Alert{StartsAt: now.Add(-time.Minute), EndsAt: now.Add(3 * time.Minute)}, UpdatedAt: now}, B: &Alert{Alert: model.Alert{StartsAt: now.Add(-time.Minute), EndsAt: now.Add(4 * time.Minute)}, UpdatedAt: now.Add(time.Minute)}, Res: &Alert{Alert: model.Alert{StartsAt: now.Add(-time.Minute), EndsAt: now.Add(4 * time.Minute)}, UpdatedAt: now.Add(time.Minute)}}, {A: &Alert{Alert: model.Alert{StartsAt: now.Add(-3 * time.Minute), EndsAt: now.Add(-time.Minute)}, UpdatedAt: now}, B: &Alert{Alert: model.Alert{StartsAt: now.Add(-2 * time.Minute), EndsAt: now.Add(time.Minute)}, UpdatedAt: now.Add(time.Minute)}, Res: &Alert{Alert: model.Alert{StartsAt: now.Add(-3 * time.Minute), EndsAt: now.Add(time.Minute)}, UpdatedAt: now.Add(time.Minute)}}, {A: &Alert{Alert: model.Alert{StartsAt: now.Add(-2 * time.Minute), EndsAt: now.Add(3 * time.Minute)}, UpdatedAt: now}, B: &Alert{Alert: model.Alert{StartsAt: now.Add(-2 * time.Minute), EndsAt: now}, UpdatedAt: now.Add(time.Minute)}, Res: &Alert{Alert: model.Alert{StartsAt: now.Add(-2 * time.Minute), EndsAt: now}, UpdatedAt: now.Add(time.Minute)}}, {A: &Alert{Alert: model.Alert{StartsAt: now.Add(-3 * time.Minute), EndsAt: now.Add(-time.Minute)}, UpdatedAt: now.Add(-time.Minute)}, B: &Alert{Alert: model.Alert{StartsAt: now.Add(-4 * time.Minute), EndsAt: now.Add(-2 * time.Minute)}, UpdatedAt: now.Add(time.Minute)}, Res: &Alert{Alert: model.Alert{StartsAt: now.Add(-4 * time.Minute), EndsAt: now.Add(-1 * time.Minute)}, UpdatedAt: now.Add(time.Minute)}}}
	for i, p := range pairs {
		p := p
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if res := p.A.Merge(p.B); !reflect.DeepEqual(p.Res, res) {
				t.Errorf("unexpected merged alert %#v", res)
			}
			if res := p.B.Merge(p.A); !reflect.DeepEqual(p.Res, res) {
				t.Errorf("unexpected merged alert %#v", res)
			}
		})
	}
}
func TestCalcSilenceState(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var (
		pastStartTime	= time.Now()
		pastEndTime	= time.Now()
		futureStartTime	= time.Now().Add(time.Hour)
		futureEndTime	= time.Now().Add(time.Hour)
	)
	expected := CalcSilenceState(futureStartTime, futureEndTime)
	require.Equal(t, SilenceStatePending, expected)
	expected = CalcSilenceState(pastStartTime, futureEndTime)
	require.Equal(t, SilenceStateActive, expected)
	expected = CalcSilenceState(pastStartTime, pastEndTime)
	require.Equal(t, SilenceStateExpired, expected)
}
func TestSilenceExpired(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	now := time.Now()
	silence := Silence{StartsAt: now, EndsAt: now}
	require.True(t, silence.Expired())
	silence = Silence{StartsAt: now.Add(time.Hour), EndsAt: now.Add(time.Hour)}
	require.True(t, silence.Expired())
	silence = Silence{StartsAt: now, EndsAt: now.Add(time.Hour)}
	require.False(t, silence.Expired())
}
func TestAlertSliceSort(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var (
		a1	= &Alert{Alert: model.Alert{Labels: model.LabelSet{"job": "j1", "instance": "i1", "alertname": "an1"}}}
		a2	= &Alert{Alert: model.Alert{Labels: model.LabelSet{"job": "j1", "instance": "i1", "alertname": "an2"}}}
		a3	= &Alert{Alert: model.Alert{Labels: model.LabelSet{"job": "j2", "instance": "i1", "alertname": "an1"}}}
		a4	= &Alert{Alert: model.Alert{Labels: model.LabelSet{"alertname": "an1"}}}
		a5	= &Alert{Alert: model.Alert{Labels: model.LabelSet{"alertname": "an2"}}}
	)
	cases := []struct {
		alerts	AlertSlice
		exp	AlertSlice
	}{{alerts: AlertSlice{a2, a1}, exp: AlertSlice{a1, a2}}, {alerts: AlertSlice{a3, a2, a1}, exp: AlertSlice{a1, a2, a3}}, {alerts: AlertSlice{a4, a2, a4}, exp: AlertSlice{a2, a4, a4}}, {alerts: AlertSlice{a5, a4}, exp: AlertSlice{a4, a5}}}
	for _, tc := range cases {
		sort.Stable(tc.alerts)
		if !reflect.DeepEqual(tc.alerts, tc.exp) {
			t.Fatalf("expected %v but got %v", tc.exp, tc.alerts)
		}
	}
}
