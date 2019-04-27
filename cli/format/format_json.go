package format

import (
	"encoding/json"
	"io"
	"os"
	"github.com/prometheus/alertmanager/client"
	"github.com/prometheus/alertmanager/types"
)

type JSONFormatter struct{ writer io.Writer }

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	Formatters["json"] = &JSONFormatter{writer: os.Stdout}
}
func (formatter *JSONFormatter) SetOutput(writer io.Writer) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	formatter.writer = writer
}
func (formatter *JSONFormatter) FormatSilences(silences []types.Silence) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	enc := json.NewEncoder(formatter.writer)
	return enc.Encode(silences)
}
func (formatter *JSONFormatter) FormatAlerts(alerts []*client.ExtendedAlert) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	enc := json.NewEncoder(formatter.writer)
	return enc.Encode(alerts)
}
func (formatter *JSONFormatter) FormatConfig(status *client.ServerStatus) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	enc := json.NewEncoder(formatter.writer)
	return enc.Encode(status)
}
