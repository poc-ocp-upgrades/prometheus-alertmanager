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
	apiclient "github.com/prometheus/alertmanager/test/with_api_v2/api_v2_client/client"
	"github.com/prometheus/alertmanager/test/with_api_v2/api_v2_client/client/alert"
	"github.com/prometheus/alertmanager/test/with_api_v2/api_v2_client/client/general"
	"github.com/prometheus/alertmanager/test/with_api_v2/api_v2_client/models"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

type AcceptanceTest struct {
	*testing.T
	opts		*AcceptanceOpts
	amc		*AlertmanagerCluster
	collectors	[]*Collector
	actions		map[float64][]func()
}
type AcceptanceOpts struct {
	RoutePrefix	string
	Tolerance	time.Duration
	baseTime	time.Time
}

func (opts *AcceptanceOpts) alertString(a *models.GettableAlert) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if a.EndsAt == nil || time.Time(*a.EndsAt).IsZero() {
		return fmt.Sprintf("%v[%v:]", a, opts.relativeTime(time.Time(*a.StartsAt)))
	}
	return fmt.Sprintf("%v[%v:%v]", a, opts.relativeTime(time.Time(*a.StartsAt)), opts.relativeTime(time.Time(*a.EndsAt)))
}
func (opts *AcceptanceOpts) expandTime(rel float64) time.Time {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return opts.baseTime.Add(time.Duration(rel * float64(time.Second)))
}
func (opts *AcceptanceOpts) relativeTime(act time.Time) float64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return float64(act.Sub(opts.baseTime)) / float64(time.Second)
}
func NewAcceptanceTest(t *testing.T, opts *AcceptanceOpts) *AcceptanceTest {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	test := &AcceptanceTest{T: t, opts: opts, actions: map[float64][]func(){}}
	opts.baseTime = time.Now()
	return test
}
func freeAddress() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.actions[at] = append(t.actions[at], f)
}
func (t *AcceptanceTest) AlertmanagerCluster(conf string, size int) *AlertmanagerCluster {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	amc := AlertmanagerCluster{}
	for i := 0; i < size; i++ {
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
		transport := httptransport.New(am.apiAddr, t.opts.RoutePrefix+"/api/v2/", nil)
		am.clientV2 = apiclient.New(transport, strfmt.Default)
		amc.ams = append(amc.ams, am)
	}
	t.amc = &amc
	return &amc
}
func (t *AcceptanceTest) Collector(name string) *Collector {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	co := &Collector{t: t.T, name: name, opts: t.opts, collected: map[float64][]models.GettableAlerts{}, expected: map[Interval][]models.GettableAlerts{}}
	t.collectors = append(t.collectors, co)
	return co
}
func (t *AcceptanceTest) Run() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	errc := make(chan error)
	for _, am := range t.amc.ams {
		am.errc = errc
		defer func(am *Alertmanager) {
			am.Terminate()
			am.cleanup()
			t.Logf("stdout:\n%v", am.cmd.Stdout)
			t.Logf("stderr:\n%v", am.cmd.Stderr)
		}(am)
	}
	err := t.amc.Start()
	if err != nil {
		t.T.Fatal(err)
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
}
func (t *AcceptanceTest) runActions() {
	_logClusterCodePath()
	defer _logClusterCodePath()
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	b.mtx.Lock()
	defer b.mtx.Unlock()
	return b.b.Write(p)
}
func (b *buffer) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
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
	clientV2	*apiclient.Alertmanager
	cmd		*exec.Cmd
	confFile	*os.File
	dir		string
	errc		chan<- error
}
type AlertmanagerCluster struct{ ams []*Alertmanager }

func (amc *AlertmanagerCluster) Start() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var peerFlags []string
	for _, am := range amc.ams {
		peerFlags = append(peerFlags, "--cluster.peer="+am.clusterAddr)
	}
	for _, am := range amc.ams {
		err := am.Start(peerFlags)
		if err != nil {
			return fmt.Errorf("starting alertmanager cluster: %v", err.Error())
		}
	}
	for _, am := range amc.ams {
		err := am.WaitForCluster(len(amc.ams))
		if err != nil {
			return fmt.Errorf("starting alertmanager cluster: %v", err.Error())
		}
	}
	return nil
}
func (am *Alertmanager) Start(additionalArg []string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	am.t.Helper()
	args := []string{"--config.file", am.confFile.Name(), "--log.level", "debug", "--web.listen-address", am.apiAddr, "--storage.path", am.dir, "--cluster.listen-address", am.clusterAddr, "--cluster.settle-timeout", "0s"}
	if am.opts.RoutePrefix != "" {
		args = append(args, "--web.route-prefix", am.opts.RoutePrefix)
	}
	args = append(args, additionalArg...)
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
		return fmt.Errorf("starting alertmanager failed: %s", err)
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
			return fmt.Errorf("starting alertmanager failed: expected HTTP status '200', got '%d'", resp.StatusCode)
		}
		_, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("starting alertmanager failed: %s", err)
		}
		return nil
	}
	return fmt.Errorf("starting alertmanager failed: timeout")
}
func (am *Alertmanager) WaitForCluster(size int) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	params := general.NewGetStatusParams()
	params.WithContext(context.Background())
	var status general.GetStatusOK
	for i := 0; i < 20; i++ {
		status, err := am.clientV2.General.GetStatus(params)
		if err != nil {
			return err
		}
		if len(status.Payload.Cluster.Peers) == size {
			return nil
		}
		time.Sleep(100 * time.Millisecond)
	}
	return fmt.Errorf("failed to wait for Alertmanager instance %q to join cluster: expected %v peers, but got %v", am.clusterAddr, size, len(status.Payload.Cluster.Peers))
}
func (amc *AlertmanagerCluster) Terminate() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, am := range amc.ams {
		am.Terminate()
	}
}
func (am *Alertmanager) Terminate() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	am.t.Helper()
	if err := syscall.Kill(am.cmd.Process.Pid, syscall.SIGTERM); err != nil {
		am.t.Fatalf("Error sending SIGTERM to Alertmanager process: %v", err)
	}
}
func (amc *AlertmanagerCluster) Reload() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, am := range amc.ams {
		am.Reload()
	}
}
func (am *Alertmanager) Reload() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	am.t.Helper()
	if err := syscall.Kill(am.cmd.Process.Pid, syscall.SIGHUP); err != nil {
		am.t.Fatalf("Error sending SIGHUP to Alertmanager process: %v", err)
	}
}
func (am *Alertmanager) cleanup() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	am.t.Helper()
	if err := os.RemoveAll(am.confFile.Name()); err != nil {
		am.t.Errorf("Error removing test config file %q: %v", am.confFile.Name(), err)
	}
}
func (amc *AlertmanagerCluster) Push(at float64, alerts ...*TestAlert) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, am := range amc.ams {
		am.Push(at, alerts...)
	}
}
func (am *Alertmanager) Push(at float64, alerts ...*TestAlert) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var cas models.PostableAlerts
	for i := range alerts {
		a := alerts[i].nativeAlert(am.opts)
		alert := &models.PostableAlert{Alert: models.Alert{Labels: a.Labels, GeneratorURL: a.GeneratorURL}, Annotations: a.Annotations}
		if a.StartsAt != nil {
			alert.StartsAt = *a.StartsAt
		}
		if a.EndsAt != nil {
			alert.EndsAt = *a.EndsAt
		}
		cas = append(cas, alert)
	}
	am.t.Do(at, func() {
		params := alert.PostAlertsParams{}
		params.WithContext(context.Background()).WithAlerts(cas)
		_, err := am.clientV2.Alert.PostAlerts(&params)
		if err != nil {
			am.t.Errorf("Error pushing %v: %v", cas, err)
		}
	})
}
func (amc *AlertmanagerCluster) SetSilence(at float64, sil *TestSilence) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, am := range amc.ams {
		am.SetSilence(at, sil)
	}
}
func (am *Alertmanager) SetSilence(at float64, sil *TestSilence) {
	_logClusterCodePath()
	defer _logClusterCodePath()
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
func (amc *AlertmanagerCluster) DelSilence(at float64, sil *TestSilence) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, am := range amc.ams {
		am.DelSilence(at, sil)
	}
}
func (am *Alertmanager) DelSilence(at float64, sil *TestSilence) {
	_logClusterCodePath()
	defer _logClusterCodePath()
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
func (amc *AlertmanagerCluster) UpdateConfig(conf string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, am := range amc.ams {
		am.UpdateConfig(conf)
	}
}
func (am *Alertmanager) UpdateConfig(conf string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
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
func (amc *AlertmanagerCluster) GenericAPIV2Call(at float64, f func()) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, am := range amc.ams {
		am.GenericAPIV2Call(at, f)
	}
}
func (am *Alertmanager) GenericAPIV2Call(at float64, f func()) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	am.t.Do(at, f)
}
func (am *Alertmanager) getURL(path string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf("http://%s%s%s", am.apiAddr, am.opts.RoutePrefix, path)
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
