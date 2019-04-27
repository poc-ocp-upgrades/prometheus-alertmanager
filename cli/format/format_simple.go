package format

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"text/tabwriter"
	"github.com/prometheus/alertmanager/client"
	"github.com/prometheus/alertmanager/types"
)

type SimpleFormatter struct{ writer io.Writer }

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	Formatters["simple"] = &SimpleFormatter{writer: os.Stdout}
}
func (formatter *SimpleFormatter) SetOutput(writer io.Writer) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	formatter.writer = writer
}
func (formatter *SimpleFormatter) FormatSilences(silences []types.Silence) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	w := tabwriter.NewWriter(formatter.writer, 0, 0, 2, ' ', 0)
	sort.Sort(ByEndAt(silences))
	fmt.Fprintln(w, "ID\tMatchers\tEnds At\tCreated By\tComment\t")
	for _, silence := range silences {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t\n", silence.ID, simpleFormatMatchers(silence.Matchers), FormatDate(silence.EndsAt), silence.CreatedBy, silence.Comment)
	}
	return w.Flush()
}
func (formatter *SimpleFormatter) FormatAlerts(alerts []*client.ExtendedAlert) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	w := tabwriter.NewWriter(formatter.writer, 0, 0, 2, ' ', 0)
	sort.Sort(ByStartsAt(alerts))
	fmt.Fprintln(w, "Alertname\tStarts At\tSummary\t")
	for _, alert := range alerts {
		fmt.Fprintf(w, "%s\t%s\t%s\t\n", alert.Labels["alertname"], FormatDate(alert.StartsAt), alert.Annotations["summary"])
	}
	return w.Flush()
}
func (formatter *SimpleFormatter) FormatConfig(status *client.ServerStatus) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fmt.Fprintln(formatter.writer, status.ConfigYAML)
	return nil
}
func simpleFormatMatchers(matchers types.Matchers) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	output := []string{}
	for _, matcher := range matchers {
		output = append(output, simpleFormatMatcher(*matcher))
	}
	return strings.Join(output, " ")
}
func simpleFormatMatcher(matcher types.Matcher) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if matcher.IsRegex {
		return fmt.Sprintf("%s=~%s", matcher.Name, matcher.Value)
	}
	return fmt.Sprintf("%s=%s", matcher.Name, matcher.Value)
}
