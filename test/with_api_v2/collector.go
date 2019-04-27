package test

import (
	"encoding/json"
	"fmt"
	"sync"
	"testing"
	"time"
	"github.com/prometheus/alertmanager/test/with_api_v2/api_v2_client/models"
)

type Collector struct {
	t		*testing.T
	name		string
	opts		*AcceptanceOpts
	collected	map[float64][]models.GettableAlerts
	expected	map[Interval][]models.GettableAlerts
	mtx		sync.RWMutex
}

func (c *Collector) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.name
}
func (c *Collector) Collected() map[float64][]models.GettableAlerts {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.mtx.RLock()
	defer c.mtx.RUnlock()
	return c.collected
}
func batchesEqual(as, bs models.GettableAlerts, opts *AcceptanceOpts) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(as) != len(bs) {
		return false
	}
	for _, a := range as {
		found := false
		for _, b := range bs {
			if equalAlerts(a, b, opts) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
func (c *Collector) latest() float64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.mtx.RLock()
	defer c.mtx.RUnlock()
	var latest float64
	for iv := range c.expected {
		if iv.end > latest {
			latest = iv.end
		}
	}
	return latest
}
func (c *Collector) Want(iv Interval, alerts ...*TestAlert) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.mtx.Lock()
	defer c.mtx.Unlock()
	var nas models.GettableAlerts
	for _, a := range alerts {
		nas = append(nas, a.nativeAlert(c.opts))
	}
	c.expected[iv] = append(c.expected[iv], nas)
}
func (c *Collector) add(alerts ...*models.GettableAlert) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.mtx.Lock()
	defer c.mtx.Unlock()
	arrival := c.opts.relativeTime(time.Now())
	c.collected[arrival] = append(c.collected[arrival], models.GettableAlerts(alerts))
}
func (c *Collector) Check() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	report := fmt.Sprintf("\ncollector %q:\n\n", c)
	c.mtx.RLock()
	defer c.mtx.RUnlock()
	for iv, expected := range c.expected {
		report += fmt.Sprintf("interval %v\n", iv)
		var alerts []models.GettableAlerts
		for at, got := range c.collected {
			if iv.contains(at) {
				alerts = append(alerts, got...)
			}
		}
		for _, exp := range expected {
			found := len(exp) == 0 && len(alerts) == 0
			report += fmt.Sprintf("---\n")
			for _, e := range exp {
				report += fmt.Sprintf("- %v\n", c.opts.alertString(e))
			}
			for _, a := range alerts {
				if batchesEqual(exp, a, c.opts) {
					found = true
					break
				}
			}
			if found {
				report += fmt.Sprintf("  [ ✓ ]\n")
			} else {
				c.t.Fail()
				report += fmt.Sprintf("  [ ✗ ]\n")
			}
		}
	}
	var totalExp, totalAct int
	for _, exp := range c.expected {
		for _, e := range exp {
			totalExp += len(e)
		}
	}
	for _, act := range c.collected {
		for _, a := range act {
			if len(a) == 0 {
				c.t.Error("received empty notifications")
			}
			totalAct += len(a)
		}
	}
	if totalExp != totalAct {
		c.t.Fail()
		report += fmt.Sprintf("\nExpected total of %d alerts, got %d", totalExp, totalAct)
	}
	if c.t.Failed() {
		report += "\nreceived:\n"
		for at, col := range c.collected {
			for _, alerts := range col {
				report += fmt.Sprintf("@ %v\n", at)
				for _, a := range alerts {
					report += fmt.Sprintf("- %v\n", c.opts.alertString(a))
				}
			}
		}
	}
	return report
}
func alertsToString(as []*models.GettableAlert) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b, err := json.Marshal(as)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
func CompareCollectors(a, b *Collector, opts *AcceptanceOpts) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	f := func(collected map[float64][]models.GettableAlerts) []*models.GettableAlert {
		result := []*models.GettableAlert{}
		for _, batches := range collected {
			for _, batch := range batches {
				for _, alert := range batch {
					result = append(result, alert)
				}
			}
		}
		return result
	}
	aAlerts := f(a.Collected())
	bAlerts := f(b.Collected())
	if len(aAlerts) != len(bAlerts) {
		aAsString, err := alertsToString(aAlerts)
		if err != nil {
			return false, err
		}
		bAsString, err := alertsToString(bAlerts)
		if err != nil {
			return false, err
		}
		err = fmt.Errorf("first collector has %v alerts, second collector has %v alerts\n%v\n%v", len(aAlerts), len(bAlerts), aAsString, bAsString)
		return false, err
	}
	for _, aAlert := range aAlerts {
		found := false
		for _, bAlert := range bAlerts {
			if equalAlerts(aAlert, bAlert, opts) {
				found = true
				break
			}
		}
		if !found {
			aAsString, err := alertsToString([]*models.GettableAlert{aAlert})
			if err != nil {
				return false, err
			}
			bAsString, err := alertsToString(bAlerts)
			if err != nil {
				return false, err
			}
			err = fmt.Errorf("could not find matching alert for alert from first collector\n%v\nin alerts of second collector\n%v", aAsString, bAsString)
			return false, err
		}
	}
	return true, nil
}
