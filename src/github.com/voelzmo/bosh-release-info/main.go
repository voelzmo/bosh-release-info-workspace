package main

import (
	"os"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	boshsys "github.com/cloudfoundry/bosh-utils/system"
	"github.com/codegangsta/cli"
	"github.com/voelzmo/bosh-release-info/commands"
	"github.com/voelzmo/bosh-release-info/output"
)

func main() {
	app := cli.NewApp()
	app.Name = "bosh-release-info"
	app.Usage = "lists info about a boshrelease"
	app.Version = "0.1"

	logger := newLogger()
	fileSystem := boshsys.NewOsFileSystemWithStrictTempRoot(logger)
	cmdRunner := boshsys.NewExecCmdRunner(logger)

	app.Commands = []cli.Command{
		commands.PackageListCommand(fileSystem, cmdRunner, logger),
		commands.FileListCommand(fileSystem, cmdRunner, logger),
	}

	app.Run(os.Args)
}

func newLogger() boshlog.Logger {
	logLevelString := "error"
	level, err := boshlog.Levelify(logLevelString)
	if err != nil {
		err = bosherr.WrapError(err, "Invalid log level value")
		logger := boshlog.NewLogger(boshlog.LevelError)
		output.Fail(err, logger)
	}

	return boshlog.NewLogger(level)

}
