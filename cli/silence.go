package cli

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

func configureSilenceCmd(app *kingpin.Application) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	silenceCmd := app.Command("silence", "Add, expire or view silences. For more information and additional flags see query help").PreAction(requireAlertManagerURL)
	configureSilenceAddCmd(silenceCmd)
	configureSilenceExpireCmd(silenceCmd)
	configureSilenceImportCmd(silenceCmd)
	configureSilenceQueryCmd(silenceCmd)
	configureSilenceUpdateCmd(silenceCmd)
}
