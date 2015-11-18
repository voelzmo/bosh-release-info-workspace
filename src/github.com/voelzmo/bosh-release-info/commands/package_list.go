package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	boshsys "github.com/cloudfoundry/bosh-utils/system"
	"github.com/codegangsta/cli"
	"github.com/voelzmo/bosh-release-info/output"
	"github.com/voelzmo/bosh-release-info/release"
)

func PackageListCommand(fs boshsys.FileSystem, runner boshsys.CmdRunner, logger boshlog.Logger) cli.Command {
	command := cli.Command{
		Name:  "package-list",
		Usage: "lists all packages in this release",
		Action: func(c *cli.Context) {
			releasePath, err := validateAndGetArgument(c.Args(), 1)
			if err != nil {
				output.Fail(bosherr.WrapError(err, "Failed to get path to release tarball:"), logger)
			}

			releaseName := path.Base(releasePath)
			tmpDir, err := ioutil.TempDir("", releaseName)
			if err != nil {
				output.Fail(bosherr.WrapError(err, "Failed creating temporary directory:"), logger)
			}
			defer os.RemoveAll(tmpDir)

			reader := release.NewReader(releasePath, tmpDir, fs, runner)

			manifest, err := reader.ReadManifest()
			if err != nil {
				output.Fail(bosherr.WrapError(err, "Failed reading release:"), logger)
			}
			fmt.Printf("Info for release '%s', version '%s'\n\n", manifest.Name, manifest.Version)

			var packageNames []string
			for _, pkg := range manifest.Packages {
				if err != nil {
					output.Fail(bosherr.WrapError(err, "Failed reading package spec files:"), logger)
				}

				packageNames = append(packageNames, pkg.Name)
			}
			fmt.Printf("Packages: %s\n", strings.Join(packageNames, ", "))
		},
	}
	return command
}
