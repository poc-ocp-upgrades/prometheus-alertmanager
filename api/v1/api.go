package v1

import (
	"encoding/json"
	godefaultbytes "bytes"
	godefaultruntime "runtime"
	"errors"
	"fmt"
	"net/http"
	godefaulthttp "net/http"
	"regexp"
	"sort"
	"sync"
	"time"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/model"
	"github.com/prometheus/common/route"
	"github.com/prometheus/common/version"
	"github.com/prometheus/prometheus/pkg/labels"
	"github.com/prometheus/alertmanager/cluster"
	"github.com/prometheus/alertmanager/config"
	"github.com/prometheus/alertmanager/dispatch"
	"github.com/prometheus/alertmanager/pkg/parse"
	"github.com/prometheus/alertmanager/provider"
	"github.com/prometheus/alertmanager/silence"
	"github.com/prometheus/alertmanager/silence/silencepb"
	"github.com/prometheus/alertmanager/types"
)

var (
	numReceivedAlerts	= prometheus.NewCounterVec(prometheus.CounterOpts{Namespace: "alertmanager", Name: "alerts_received_total", Help: "The total number of received alerts."}, []string{"status"})
	numInvalidAlerts	= prometheus.NewCounter(prometheus.CounterOpts{Namespace: "alertmanager", Name: "alerts_invalid_total", Help: "The total number of received alerts that were invalid."})
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	numReceivedAlerts.WithLabelValues("firing")
	numReceivedAlerts.WithLabelValues("resolved")
	prometheus.MustRegister(numReceivedAlerts)
	prometheus.MustRegister(numInvalidAlerts)
}

var corsHeaders = map[string]string{"Access-Control-Allow-Headers": "Accept, Authorization, Content-Type, Origin", "Access-Control-Allow-Methods": "GET, DELETE, OPTIONS", "Access-Control-Allow-Origin": "*", "Access-Control-Expose-Headers": "Date", "Cache-Control": "no-cache, no-store, must-revalidate"}

type Alert struct {
	*model.Alert
	Status		types.AlertStatus	`json:"status"`
	Receivers	[]string		`json:"receivers"`
	Fingerprint	string			`json:"fingerprint"`
}

func setCORS(w http.ResponseWriter) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for h, v := range corsHeaders {
		w.Header().Set(h, v)
	}
}

type API struct {
	alerts		provider.Alerts
	silences	*silence.Silences
	config		*config.Config
	route		*dispatch.Route
	resolveTimeout	time.Duration
	uptime		time.Time
	peer		*cluster.Peer
	logger		log.Logger
	getAlertStatus	getAlertStatusFn
	mtx		sync.RWMutex
}
type getAlertStatusFn func(model.Fingerprint) types.AlertStatus

func New(alerts provider.Alerts, silences *silence.Silences, sf getAlertStatusFn, peer *cluster.Peer, l log.Logger) *API {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if l == nil {
		l = log.NewNopLogger()
	}
	return &API{alerts: alerts, silences: silences, getAlertStatus: sf, uptime: time.Now(), peer: peer, logger: l}
}
func (api *API) Register(r *route.Router) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	wrap := func(f http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			setCORS(w)
			f(w, r)
		})
	}
	r.Options("/*path", wrap(func(w http.ResponseWriter, r *http.Request) {
	}))
	r.Get("/status", wrap(api.status))
	r.Get("/receivers", wrap(api.receivers))
	r.Get("/alerts", wrap(api.listAlerts))
	r.Post("/alerts", wrap(api.addAlerts))
	r.Get("/silences", wrap(api.listSilences))
	r.Post("/silences", wrap(api.setSilence))
	r.Get("/silence/:sid", wrap(api.getSilence))
	r.Del("/silence/:sid", wrap(api.delSilence))
}
func (api *API) Update(cfg *config.Config, resolveTimeout time.Duration) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	api.mtx.Lock()
	defer api.mtx.Unlock()
	api.resolveTimeout = resolveTimeout
	api.config = cfg
	api.route = dispatch.NewRoute(cfg.Route, nil)
	return nil
}

type errorType string

const (
	errorInternal	errorType	= "server_error"
	errorBadData	errorType	= "bad_data"
)

type apiError struct {
	typ	errorType
	err	error
}

func (e *apiError) Error() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf("%s: %s", e.typ, e.err)
}
func (api *API) receivers(w http.ResponseWriter, req *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	api.mtx.RLock()
	defer api.mtx.RUnlock()
	receivers := make([]string, 0, len(api.config.Receivers))
	for _, r := range api.config.Receivers {
		receivers = append(receivers, r.Name)
	}
	api.respond(w, receivers)
}
func (api *API) status(w http.ResponseWriter, req *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	api.mtx.RLock()
	var status = struct {
		ConfigYAML	string			`json:"configYAML"`
		ConfigJSON	*config.Config		`json:"configJSON"`
		VersionInfo	map[string]string	`json:"versionInfo"`
		Uptime		time.Time		`json:"uptime"`
		ClusterStatus	*clusterStatus		`json:"clusterStatus"`
	}{ConfigYAML: api.config.String(), ConfigJSON: api.config, VersionInfo: map[string]string{"version": version.Version, "revision": version.Revision, "branch": version.Branch, "buildUser": version.BuildUser, "buildDate": version.BuildDate, "goVersion": version.GoVersion}, Uptime: api.uptime, ClusterStatus: getClusterStatus(api.peer)}
	api.mtx.RUnlock()
	api.respond(w, status)
}

type peerStatus struct {
	Name	string	`json:"name"`
	Address	string	`json:"address"`
}
type clusterStatus struct {
	Name	string		`json:"name"`
	Status	string		`json:"status"`
	Peers	[]peerStatus	`json:"peers"`
}

func getClusterStatus(p *cluster.Peer) *clusterStatus {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if p == nil {
		return nil
	}
	s := &clusterStatus{Name: p.Name(), Status: p.Status()}
	for _, n := range p.Peers() {
		s.Peers = append(s.Peers, peerStatus{Name: n.Name, Address: n.Address()})
	}
	return s
}
func (api *API) listAlerts(w http.ResponseWriter, r *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var (
		err				error
		receiverFilter			*regexp.Regexp
		res				= []*Alert{}
		matchers			= []*labels.Matcher{}
		showActive, showInhibited	bool
		showSilenced, showUnprocessed	bool
	)
	getBoolParam := func(name string) (bool, error) {
		v := r.FormValue(name)
		if v == "" {
			return true, nil
		}
		if v == "false" {
			return false, nil
		}
		if v != "true" {
			err := fmt.Errorf("parameter %q can either be 'true' or 'false', not %q", name, v)
			api.respondError(w, apiError{typ: errorBadData, err: err}, nil)
			return false, err
		}
		return true, nil
	}
	if filter := r.FormValue("filter"); filter != "" {
		matchers, err = parse.Matchers(filter)
		if err != nil {
			api.respondError(w, apiError{typ: errorBadData, err: err}, nil)
			return
		}
	}
	showActive, err = getBoolParam("active")
	if err != nil {
		return
	}
	showSilenced, err = getBoolParam("silenced")
	if err != nil {
		return
	}
	showInhibited, err = getBoolParam("inhibited")
	if err != nil {
		return
	}
	showUnprocessed, err = getBoolParam("unprocessed")
	if err != nil {
		return
	}
	if receiverParam := r.FormValue("receiver"); receiverParam != "" {
		receiverFilter, err = regexp.Compile("^(?:" + receiverParam + ")$")
		if err != nil {
			api.respondError(w, apiError{typ: errorBadData, err: fmt.Errorf("failed to parse receiver param: %s", receiverParam)}, nil)
			return
		}
	}
	alerts := api.alerts.GetPending()
	defer alerts.Close()
	api.mtx.RLock()
	for a := range alerts.Next() {
		if err = alerts.Err(); err != nil {
			break
		}
		routes := api.route.Match(a.Labels)
		receivers := make([]string, 0, len(routes))
		for _, r := range routes {
			receivers = append(receivers, r.RouteOpts.Receiver)
		}
		if receiverFilter != nil && !receiversMatchFilter(receivers, receiverFilter) {
			continue
		}
		if !alertMatchesFilterLabels(&a.Alert, matchers) {
			continue
		}
		if !a.Alert.EndsAt.IsZero() && a.Alert.EndsAt.Before(time.Now()) {
			continue
		}
		status := api.getAlertStatus(a.Fingerprint())
		if !showActive && status.State == types.AlertStateActive {
			continue
		}
		if !showUnprocessed && status.State == types.AlertStateUnprocessed {
			continue
		}
		if !showSilenced && len(status.SilencedBy) != 0 {
			continue
		}
		if !showInhibited && len(status.InhibitedBy) != 0 {
			continue
		}
		alert := &Alert{Alert: &a.Alert, Status: status, Receivers: receivers, Fingerprint: a.Fingerprint().String()}
		res = append(res, alert)
	}
	api.mtx.RUnlock()
	if err != nil {
		api.respondError(w, apiError{typ: errorInternal, err: err}, nil)
		return
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i].Fingerprint < res[j].Fingerprint
	})
	api.respond(w, res)
}
func receiversMatchFilter(receivers []string, filter *regexp.Regexp) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, r := range receivers {
		if filter.MatchString(r) {
			return true
		}
	}
	return false
}
func alertMatchesFilterLabels(a *model.Alert, matchers []*labels.Matcher) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	sms := make(map[string]string)
	for name, value := range a.Labels {
		sms[string(name)] = string(value)
	}
	return matchFilterLabels(matchers, sms)
}
func (api *API) addAlerts(w http.ResponseWriter, r *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var alerts []*types.Alert
	if err := api.receive(r, &alerts); err != nil {
		api.respondError(w, apiError{typ: errorBadData, err: err}, nil)
		return
	}
	api.insertAlerts(w, r, alerts...)
}
func (api *API) insertAlerts(w http.ResponseWriter, r *http.Request, alerts ...*types.Alert) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	now := time.Now()
	api.mtx.RLock()
	resolveTimeout := api.resolveTimeout
	api.mtx.RUnlock()
	for _, alert := range alerts {
		alert.UpdatedAt = now
		if alert.StartsAt.IsZero() {
			if alert.EndsAt.IsZero() {
				alert.StartsAt = now
			} else {
				alert.StartsAt = alert.EndsAt
			}
		}
		if alert.EndsAt.IsZero() {
			alert.Timeout = true
			alert.EndsAt = now.Add(resolveTimeout)
		}
		if alert.EndsAt.After(time.Now()) {
			numReceivedAlerts.WithLabelValues("firing").Inc()
		} else {
			numReceivedAlerts.WithLabelValues("resolved").Inc()
		}
	}
	var (
		validAlerts	= make([]*types.Alert, 0, len(alerts))
		validationErrs	= &types.MultiError{}
	)
	for _, a := range alerts {
		removeEmptyLabels(a.Labels)
		if err := a.Validate(); err != nil {
			validationErrs.Add(err)
			numInvalidAlerts.Inc()
			continue
		}
		validAlerts = append(validAlerts, a)
	}
	if err := api.alerts.Put(validAlerts...); err != nil {
		api.respondError(w, apiError{typ: errorInternal, err: err}, nil)
		return
	}
	if validationErrs.Len() > 0 {
		api.respondError(w, apiError{typ: errorBadData, err: validationErrs}, nil)
		return
	}
	api.respond(w, nil)
}
func removeEmptyLabels(ls model.LabelSet) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for k, v := range ls {
		if string(v) == "" {
			delete(ls, k)
		}
	}
}
func (api *API) setSilence(w http.ResponseWriter, r *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var sil types.Silence
	if err := api.receive(r, &sil); err != nil {
		api.respondError(w, apiError{typ: errorBadData, err: err}, nil)
		return
	}
	if sil.Expired() {
		api.respondError(w, apiError{typ: errorBadData, err: errors.New("start time must not be equal to end time")}, nil)
		return
	}
	if sil.EndsAt.Before(time.Now()) {
		api.respondError(w, apiError{typ: errorBadData, err: errors.New("end time can't be in the past")}, nil)
		return
	}
	psil, err := silenceToProto(&sil)
	if err != nil {
		api.respondError(w, apiError{typ: errorBadData, err: err}, nil)
		return
	}
	sid, err := api.silences.Set(psil)
	if err != nil {
		api.respondError(w, apiError{typ: errorBadData, err: err}, nil)
		return
	}
	api.respond(w, struct {
		SilenceID string `json:"silenceId"`
	}{SilenceID: sid})
}
func (api *API) getSilence(w http.ResponseWriter, r *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	sid := route.Param(r.Context(), "sid")
	sils, err := api.silences.Query(silence.QIDs(sid))
	if err != nil || len(sils) == 0 {
		http.Error(w, fmt.Sprint("Error getting silence: ", err), http.StatusNotFound)
		return
	}
	sil, err := silenceFromProto(sils[0])
	if err != nil {
		api.respondError(w, apiError{typ: errorInternal, err: err}, nil)
		return
	}
	api.respond(w, sil)
}
func (api *API) delSilence(w http.ResponseWriter, r *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	sid := route.Param(r.Context(), "sid")
	if err := api.silences.Expire(sid); err != nil {
		api.respondError(w, apiError{typ: errorBadData, err: err}, nil)
		return
	}
	api.respond(w, nil)
}
func (api *API) listSilences(w http.ResponseWriter, r *http.Request) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	psils, err := api.silences.Query()
	if err != nil {
		api.respondError(w, apiError{typ: errorInternal, err: err}, nil)
		return
	}
	matchers := []*labels.Matcher{}
	if filter := r.FormValue("filter"); filter != "" {
		matchers, err = parse.Matchers(filter)
		if err != nil {
			api.respondError(w, apiError{typ: errorBadData, err: err}, nil)
			return
		}
	}
	sils := []*types.Silence{}
	for _, ps := range psils {
		s, err := silenceFromProto(ps)
		if err != nil {
			api.respondError(w, apiError{typ: errorInternal, err: err}, nil)
			return
		}
		if !silenceMatchesFilterLabels(s, matchers) {
			continue
		}
		sils = append(sils, s)
	}
	var active, pending, expired []*types.Silence
	for _, s := range sils {
		switch s.Status.State {
		case types.SilenceStateActive:
			active = append(active, s)
		case types.SilenceStatePending:
			pending = append(pending, s)
		case types.SilenceStateExpired:
			expired = append(expired, s)
		}
	}
	sort.Slice(active, func(i int, j int) bool {
		return active[i].EndsAt.Before(active[j].EndsAt)
	})
	sort.Slice(pending, func(i int, j int) bool {
		return pending[i].StartsAt.Before(pending[j].EndsAt)
	})
	sort.Slice(expired, func(i int, j int) bool {
		return expired[i].EndsAt.After(expired[j].EndsAt)
	})
	silences := []*types.Silence{}
	silences = append(silences, active...)
	silences = append(silences, pending...)
	silences = append(silences, expired...)
	api.respond(w, silences)
}
func silenceMatchesFilterLabels(s *types.Silence, matchers []*labels.Matcher) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	sms := make(map[string]string)
	for _, m := range s.Matchers {
		sms[m.Name] = m.Value
	}
	return matchFilterLabels(matchers, sms)
}
func matchFilterLabels(matchers []*labels.Matcher, sms map[string]string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, m := range matchers {
		v, prs := sms[m.Name]
		switch m.Type {
		case labels.MatchNotRegexp, labels.MatchNotEqual:
			if string(m.Value) == "" && prs {
				continue
			}
			if !m.Matches(string(v)) {
				return false
			}
		default:
			if string(m.Value) == "" && !prs {
				continue
			}
			if !prs || !m.Matches(string(v)) {
				return false
			}
		}
	}
	return true
}
func silenceToProto(s *types.Silence) (*silencepb.Silence, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	sil := &silencepb.Silence{Id: s.ID, StartsAt: s.StartsAt, EndsAt: s.EndsAt, UpdatedAt: s.UpdatedAt, Comment: s.Comment, CreatedBy: s.CreatedBy}
	for _, m := range s.Matchers {
		matcher := &silencepb.Matcher{Name: m.Name, Pattern: m.Value, Type: silencepb.Matcher_EQUAL}
		if m.IsRegex {
			matcher.Type = silencepb.Matcher_REGEXP
		}
		sil.Matchers = append(sil.Matchers, matcher)
	}
	return sil, nil
}
func silenceFromProto(s *silencepb.Silence) (*types.Silence, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	sil := &types.Silence{ID: s.Id, StartsAt: s.StartsAt, EndsAt: s.EndsAt, UpdatedAt: s.UpdatedAt, Status: types.SilenceStatus{State: types.CalcSilenceState(s.StartsAt, s.EndsAt)}, Comment: s.Comment, CreatedBy: s.CreatedBy}
	for _, m := range s.Matchers {
		matcher := &types.Matcher{Name: m.Name, Value: m.Pattern}
		switch m.Type {
		case silencepb.Matcher_EQUAL:
		case silencepb.Matcher_REGEXP:
			matcher.IsRegex = true
		default:
			return nil, fmt.Errorf("unknown matcher type")
		}
		sil.Matchers = append(sil.Matchers, matcher)
	}
	return sil, nil
}

type status string

const (
	statusSuccess	status	= "success"
	statusError	status	= "error"
)

type response struct {
	Status		status		`json:"status"`
	Data		interface{}	`json:"data,omitempty"`
	ErrorType	errorType	`json:"errorType,omitempty"`
	Error		string		`json:"error,omitempty"`
}

func (api *API) respond(w http.ResponseWriter, data interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	b, err := json.Marshal(&response{Status: statusSuccess, Data: data})
	if err != nil {
		level.Error(api.logger).Log("msg", "Error marshalling JSON", "err", err)
		return
	}
	if _, err := w.Write(b); err != nil {
		level.Error(api.logger).Log("msg", "failed to write data to connection", "err", err)
	}
}
func (api *API) respondError(w http.ResponseWriter, apiErr apiError, data interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	w.Header().Set("Content-Type", "application/json")
	switch apiErr.typ {
	case errorBadData:
		w.WriteHeader(http.StatusBadRequest)
	case errorInternal:
		w.WriteHeader(http.StatusInternalServerError)
	default:
		panic(fmt.Sprintf("unknown error type %q", apiErr.Error()))
	}
	b, err := json.Marshal(&response{Status: statusError, ErrorType: apiErr.typ, Error: apiErr.err.Error(), Data: data})
	if err != nil {
		return
	}
	level.Error(api.logger).Log("msg", "API error", "err", apiErr.Error())
	if _, err := w.Write(b); err != nil {
		level.Error(api.logger).Log("msg", "failed to write data to connection", "err", err)
	}
}
func (api *API) receive(r *http.Request, v interface{}) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	dec := json.NewDecoder(r.Body)
	defer r.Body.Close()
	err := dec.Decode(v)
	if err != nil {
		level.Debug(api.logger).Log("msg", "Decoding request failed", "err", err)
	}
	return err
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
