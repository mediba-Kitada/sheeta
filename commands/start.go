package commands

import (
	"flag"
	"os"
	"path/filepath"

	"context"

	c "github.com/future-architect/vuls/config"
	"github.com/google/subcommands"

	s "github.com/mediba-Kitada/Sheeta/start"
)

// StartCmd is Subcommand
type StartCmd struct {
	configPath string
	debug      bool
}

// Name return subcommand name
func (*StartCmd) Name() string { return "start" }

// Synopsis return synopsis
func (*StartCmd) Synopsis() string { return "Start vuls target instances" }

// Usage return usage
func (*StartCmd) Usage() string {
	return `start:
	start
		[-config=/path/to/config.toml]
		[-debug]
	`
}

// SetFlags set flags
func (p *StartCmd) SetFlags(f *flag.FlagSet) {
	wd, _ := os.Getwd()
	defaultConfPath := filepath.Join(wd, "../", "config.toml")
	f.StringVar(&p.configPath, "config", defaultConfPath, "/path/to/toml")
}

// Execute execute
func (p *StartCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	// Setup Logger
	c.Conf.Debug = p.debug

	err := c.Load(p.configPath, "")
	if err != nil {
		return subcommands.ExitUsageError
	}

	for servername, info := range c.Conf.Servers {
		for _, v := range info.Optional {
			if v[0] == "ami" {
				// todo goroutineを発行して、対象のAMI IDを基にEC2インスタンスを起動
				go s.Start(v[1], info)
			}
		}
	}

	// todo 全てのgoroutineの処理が正常に完了したら0を返却
	return subcommands.ExitSuccess
}
