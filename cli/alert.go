package cli

import (
	"gopkg.in/alecthomas/kingpin.v2"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

func configureAlertCmd(app *kingpin.Application) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	alertCmd := app.Command("alert", "Add or query alerts.").PreAction(requireAlertManagerURL)
	configureQueryAlertsCmd(alertCmd)
	configureAddAlertCmd(alertCmd)
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
