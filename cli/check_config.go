package cli

import (
	"fmt"
	"os"
	"github.com/prometheus/alertmanager/config"
	"github.com/prometheus/alertmanager/template"
	"gopkg.in/alecthomas/kingpin.v2"
)

type checkConfigCmd struct{ files []string }

const checkConfigHelp = `Validate alertmanager config files

Will validate the syntax and schema for alertmanager config file
and associated templates. Non existing templates will not trigger
errors.
`

func configureCheckConfigCmd(app *kingpin.Application) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var (
		c		= &checkConfigCmd{}
		checkCmd	= app.Command("check-config", checkConfigHelp)
	)
	checkCmd.Arg("check-files", "Files to be validated").ExistingFilesVar(&c.files)
	checkCmd.Action(c.checkConfig)
}
func (c *checkConfigCmd) checkConfig(ctx *kingpin.ParseContext) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return CheckConfig(c.files)
}
func CheckConfig(args []string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(args) == 0 {
		stat, err := os.Stdin.Stat()
		if err != nil {
			kingpin.Fatalf("Failed to stat standard input: %v", err)
		}
		if (stat.Mode() & os.ModeCharDevice) != 0 {
			kingpin.Fatalf("Failed to read from standard input")
		}
		args = []string{os.Stdin.Name()}
	}
	failed := 0
	for _, arg := range args {
		fmt.Printf("Checking '%s'", arg)
		cfg, _, err := config.LoadFile(arg)
		if err != nil {
			fmt.Printf("  FAILED: %s\n", err)
			failed++
		} else {
			fmt.Printf("  SUCCESS\n")
		}
		if cfg != nil {
			fmt.Println("Found:")
			if cfg.Global != nil {
				fmt.Println(" - global config")
			}
			if cfg.Route != nil {
				fmt.Println(" - route")
			}
			fmt.Printf(" - %d inhibit rules\n", len(cfg.InhibitRules))
			fmt.Printf(" - %d receivers\n", len(cfg.Receivers))
			fmt.Printf(" - %d templates\n", len(cfg.Templates))
			if len(cfg.Templates) > 0 {
				_, err = template.FromGlobs(cfg.Templates...)
				if err != nil {
					fmt.Printf("  FAILED: %s\n", err)
					failed++
				} else {
					fmt.Printf("  SUCCESS\n")
				}
			}
		}
		fmt.Printf("\n")
	}
	if failed > 0 {
		return fmt.Errorf("failed to validate %d file(s)", failed)
	}
	return nil
}
