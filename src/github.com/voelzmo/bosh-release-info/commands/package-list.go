package commands

import (
	"fmt"
	"os"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshcmd "github.com/cloudfoundry/bosh-utils/fileutil"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	boshsys "github.com/cloudfoundry/bosh-utils/system"
	"github.com/codegangsta/cli"
	"github.com/voelzmo/bosh-release-info/release"
)

func PackageListCommand(fs boshsys.FileSystem, comp boshcmd.Compressor, logger boshlog.Logger) cli.Command {
	command := cli.Command{
		Name:  "package-list",
		Usage: "lists all packages in this release",
		Action: func(c *cli.Context) {
			if len(c.Args()) != 1 {
				err := bosherr.Error("that's not how it works. provide the name to the release tarball!")
				fail(err, logger)
			}
			releasePath := c.Args()[0]
			fmt.Printf("Info for release: %s\n", releasePath)
			fmt.Printf("\n")

			reader := release.NewReader(releasePath, "/Users/d058546/.bosh-release-info-tmp", fs, comp)

			manifest, err := reader.Read()
			if err != nil {
				fail(bosherr.WrapError(err, "Failed reading release:"), logger)
			}
			fmt.Printf("Release name: %s\n", manifest.Name)
			fmt.Printf("\n")

			for _, pkg := range manifest.Packages {
				pkgSpecFiles, err := reader.ReadPackageSpecs(pkg.Name)
				if err != nil {
					fail(bosherr.WrapError(err, "Failed reading package spec files:"), logger)
				}

				fmt.Printf("Files for package '%s': %s", pkg.Name, pkgSpecFiles)

				fmt.Printf("\n")

			}
		},
	}
	return command
}

func fail(err error, logger boshlog.Logger) {
	logger.Error("main", err.Error())
	os.Exit(1)
}
