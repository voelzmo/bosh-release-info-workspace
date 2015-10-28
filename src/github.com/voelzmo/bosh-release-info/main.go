package main

import (
	"os"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshcmd "github.com/cloudfoundry/bosh-utils/fileutil"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	boshsys "github.com/cloudfoundry/bosh-utils/system"
	"github.com/codegangsta/cli"
	"github.com/voelzmo/bosh-release-info/commands"
)

func main() {
	app := cli.NewApp()
	app.Name = "bosh-release-info"
	app.Usage = "lists info about a boshrelease"
	app.Version = "0.1"

	logger := newLogger()
	fileSystem := boshsys.NewOsFileSystemWithStrictTempRoot(logger)
	cmdRunner := boshsys.NewExecCmdRunner(logger)
	compressor := boshcmd.NewTarballCompressor(cmdRunner, fileSystem)

	app.Commands = []cli.Command{
		commands.PackageListCommand(fileSystem, compressor, logger),
	}

	app.Run(os.Args)
}

func newLogger() boshlog.Logger {
	logLevelString := "none"
	level, err := boshlog.Levelify(logLevelString)
	if err != nil {
		err = bosherr.WrapError(err, "Invalid log level value")
		logger := boshlog.NewLogger(boshlog.LevelError)
		fail(err, logger)
	}

	return boshlog.NewLogger(level)
}

func fail(err error, logger boshlog.Logger) {
	logger.Error("main", err.Error())
	os.Exit(1)
}
