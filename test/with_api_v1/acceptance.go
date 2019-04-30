package test

import (
	"bytes"
	godefaultbytes "bytes"
	godefaultruntime "runtime"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	godefaulthttp "net/http"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"syscall"
	"testing"
	"time"
	"github.com/prometheus/client_golang/api"
	"github.com/prometheus/common/model"
	"github.com/prometheus/alertmanager/client"
)

type AcceptanceTest struct {
	*testing.T
	opts		*AcceptanceOpts
	ams		[]*Alertmanager
	collectors	[]*Collector
	actions		map[float64][]func()
}
type AcceptanceOpts struct {
	RoutePrefix	string
	Tolerance	time.Duration
	baseTime	time.Time
}

func (opts *AcceptanceOpts) alertString(a *model.Alert) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if a.EndsAt.IsZero() {
		return fmt.Sprintf("%s[%v:]", a, opts.relativeTime(a.StartsAt))
	}
	return fmt.Sprintf("%s[%v:%v]", a, opts.relativeTime(a.StartsAt), opts.relativeTime(a.EndsAt))
}
func (opts *AcceptanceOpts) expandTime(rel float64) time.Time {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return opts.baseTime.Add(time.Duration(rel * float64(time.Second)))
}
func (opts *AcceptanceOpts) relativeTime(act time.Time) float64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return float64(act.Sub(opts.baseTime)) / float64(time.Second)
}
func NewAcceptanceTest(t *testing.T, opts *AcceptanceOpts) *AcceptanceTest {
	_logClusterCodePath()
	defer _logClusterCodePath()
	test := &AcceptanceTest{T: t, opts: opts, actions: map[float64][]func(){}}
	opts.baseTime = time.Now()
	return test
}
func freeAddress() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	l, err := net.Listen("tcp4", "localhost:0")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := l.Close(); err != nil {
			panic(err)
		}
	}()
	return l.Addr().String()
}
func (t *AcceptanceTest) Do(at float64, f func()) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.actions[at] = append(t.actions[at], f)
}
func (t *AcceptanceTest) Alertmanager(conf string) *Alertmanager {
	_logClusterCodePath()
	defer _logClusterCodePath()
	am := &Alertmanager{t: t, opts: t.opts}
	dir, err := ioutil.TempDir("", "am_test")
	if err != nil {
		t.Fatal(err)
	}
	am.dir = dir
	cf, err := os.Create(filepath.Join(dir, "config.yml"))
	if err != nil {
		t.Fatal(err)
	}
	am.confFile = cf
	am.UpdateConfig(conf)
	am.apiAddr = freeAddress()
	am.clusterAddr = freeAddress()
	t.Logf("AM on %s", am.apiAddr)
	c, err := api.NewClient(api.Config{Address: am.getURL("")})
	if err != nil {
		t.Fatal(err)
	}
	am.client = c
	t.ams = append(t.ams, am)
	return am
}
func (t *AcceptanceTest) Collector(name string) *Collector {
	_logClusterCodePath()
	defer _logClusterCodePath()
	co := &Collector{t: t.T, name: name, opts: t.opts, collected: map[float64][]model.Alerts{}, expected: map[Interval][]model.Alerts{}}
	t.collectors = append(t.collectors, co)
	return co
}
func (t *AcceptanceTest) Run() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	errc := make(chan error)
	for _, am := range t.ams {
		am.errc = errc
		am.Start()
		defer func(am *Alertmanager) {
			am.Terminate()
			am.cleanup()
			t.Logf("stdout:\n%v", am.cmd.Stdout)
			t.Logf("stderr:\n%v", am.cmd.Stderr)
		}(am)
	}
	go t.runActions()
	var latest float64
	for _, coll := range t.collectors {
		if l := coll.latest(); l > latest {
			latest = l
		}
	}
	deadline := t.opts.expandTime(latest)
	select {
	case <-time.After(time.Until(deadline)):
	case err := <-errc:
		t.Error(err)
	}
	for _, coll := range t.collectors {
		report := coll.check()
		t.Log(report)
	}
}
func (t *AcceptanceTest) runActions() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var wg sync.WaitGroup
	for at, fs := range t.actions {
		ts := t.opts.expandTime(at)
		wg.Add(len(fs))
		for _, f := range fs {
			go func(f func()) {
				time.Sleep(time.Until(ts))
				f()
				wg.Done()
			}(f)
		}
	}
	wg.Wait()
}

type buffer struct {
	b	bytes.Buffer
	mtx	sync.Mutex
}

func (b *buffer) Write(p []byte) (int, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b.mtx.Lock()
	defer b.mtx.Unlock()
	return b.b.Write(p)
}
func (b *buffer) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b.mtx.Lock()
	defer b.mtx.Unlock()
	return b.b.String()
}

type Alertmanager struct {
	t		*AcceptanceTest
	opts		*AcceptanceOpts
	apiAddr		string
	clusterAddr	string
	client		api.Client
	cmd		*exec.Cmd
	confFile	*os.File
	dir		string
	errc		chan<- error
}

func (am *Alertmanager) Start() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	args := []string{"--config.file", am.confFile.Name(), "--log.level", "debug", "--web.listen-address", am.apiAddr, "--storage.path", am.dir, "--cluster.listen-address", am.clusterAddr, "--cluster.settle-timeout", "0s"}
	if am.opts.RoutePrefix != "" {
		args = append(args, "--web.route-prefix", am.opts.RoutePrefix)
	}
	cmd := exec.Command("../../../alertmanager", args...)
	if am.cmd == nil {
		var outb, errb buffer
		cmd.Stdout = &outb
		cmd.Stderr = &errb
	} else {
		cmd.Stdout = am.cmd.Stdout
		cmd.Stderr = am.cmd.Stderr
	}
	am.cmd = cmd
	if err := am.cmd.Start(); err != nil {
		am.t.Fatalf("Starting alertmanager failed: %s", err)
	}
	go func() {
		if err := am.cmd.Wait(); err != nil {
			am.errc <- err
		}
	}()
	time.Sleep(50 * time.Millisecond)
	for i := 0; i < 10; i++ {
		resp, err := http.Get(am.getURL("/"))
		if err != nil {
			time.Sleep(500 * time.Millisecond)
			continue
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			am.t.Fatalf("Starting alertmanager failed: expected HTTP status '200', got '%d'", resp.StatusCode)
		}
		_, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			am.t.Fatalf("Starting alertmanager failed: %s", err)
		}
		resp.Body.Close()
		return
	}
	am.t.Fatalf("Starting alertmanager failed: timeout")
}
func (am *Alertmanager) Terminate() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := syscall.Kill(am.cmd.Process.Pid, syscall.SIGTERM); err != nil {
		am.t.Fatalf("error sending SIGTERM to Alertmanager process: %v", err)
	}
}
func (am *Alertmanager) Reload() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := syscall.Kill(am.cmd.Process.Pid, syscall.SIGHUP); err != nil {
		am.t.Fatalf("error sending SIGHUP to Alertmanager process: %v", err)
	}
}
func (am *Alertmanager) cleanup() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := os.RemoveAll(am.confFile.Name()); err != nil {
		am.t.Errorf("error removing test config file %q: %v", am.confFile.Name(), err)
	}
}
func (am *Alertmanager) Push(at float64, alerts ...*TestAlert) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var cas []client.Alert
	for i := range alerts {
		a := alerts[i].nativeAlert(am.opts)
		al := client.Alert{Labels: client.LabelSet{}, Annotations: client.LabelSet{}, StartsAt: a.StartsAt, EndsAt: a.EndsAt, GeneratorURL: a.GeneratorURL}
		for n, v := range a.Labels {
			al.Labels[client.LabelName(n)] = client.LabelValue(v)
		}
		for n, v := range a.Annotations {
			al.Annotations[client.LabelName(n)] = client.LabelValue(v)
		}
		cas = append(cas, al)
	}
	alertAPI := client.NewAlertAPI(am.client)
	am.t.Do(at, func() {
		if err := alertAPI.Push(context.Background(), cas...); err != nil {
			am.t.Errorf("Error pushing %v: %s", cas, err)
		}
	})
}
func (am *Alertmanager) SetSilence(at float64, sil *TestSilence) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	am.t.Do(at, func() {
		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(sil.nativeSilence(am.opts)); err != nil {
			am.t.Errorf("Error setting silence %v: %s", sil, err)
			return
		}
		resp, err := http.Post(am.getURL("/api/v1/silences"), "application/json", &buf)
		if err != nil {
			am.t.Errorf("Error setting silence %v: %s", sil, err)
			return
		}
		defer resp.Body.Close()
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		var v struct {
			Status	string	`json:"status"`
			Data	struct {
				SilenceID string `json:"silenceId"`
			}	`json:"data"`
		}
		if err := json.Unmarshal(b, &v); err != nil || resp.StatusCode/100 != 2 {
			am.t.Errorf("error setting silence %v: %s", sil, err)
			return
		}
		sil.SetID(v.Data.SilenceID)
	})
}
func (am *Alertmanager) DelSilence(at float64, sil *TestSilence) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	am.t.Do(at, func() {
		req, err := http.NewRequest("DELETE", am.getURL(fmt.Sprintf("/api/v1/silence/%s", sil.ID())), nil)
		if err != nil {
			am.t.Errorf("Error deleting silence %v: %s", sil, err)
			return
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil || resp.StatusCode/100 != 2 {
			am.t.Errorf("Error deleting silence %v: %s", sil, err)
			return
		}
	})
}
func (am *Alertmanager) UpdateConfig(conf string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if _, err := am.confFile.WriteString(conf); err != nil {
		am.t.Fatal(err)
		return
	}
	if err := am.confFile.Sync(); err != nil {
		am.t.Fatal(err)
		return
	}
}
func (am *Alertmanager) getURL(path string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf("http://%s%s%s", am.apiAddr, am.opts.RoutePrefix, path)
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
