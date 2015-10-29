package main

import (
	"os"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshcmd "github.com/cloudfoundry/bosh-utils/fileutil"
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
	compressor := boshcmd.NewTarballCompressor(cmdRunner, fileSystem)

	app.Commands = []cli.Command{
		commands.PackageListCommand(fileSystem, compressor, logger),
		commands.FileListCommand(fileSystem, compressor, logger),
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
