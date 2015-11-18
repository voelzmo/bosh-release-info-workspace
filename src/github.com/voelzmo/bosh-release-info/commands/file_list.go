package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	boshsys "github.com/cloudfoundry/bosh-utils/system"

	"github.com/codegangsta/cli"
	"github.com/voelzmo/bosh-release-info/output"
	"github.com/voelzmo/bosh-release-info/release"
)

func FileListCommand(fs boshsys.FileSystem, runner boshsys.CmdRunner, logger boshlog.Logger) cli.Command {
	command := cli.Command{
		Name:  "file-list",
		Usage: "lists all files in all packages in this release",
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "by-type",
				Usage: "group files by type. Supported filetypes: LICENSE files, rubygems, archives",
			},
		},

		Action: func(c *cli.Context) {
			sortByType := c.Bool("by-type")

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

			filesByType := make(map[string][]string)

			for _, pkg := range manifest.Packages {
				pkgSpecFiles, err := reader.ReadPackageSpecs(pkg.Name)
				if err != nil {
					output.Fail(bosherr.WrapError(err, "Failed reading package spec files:"), logger)
				}

				if sortByType {
					filesInPackage := strings.Split(pkgSpecFiles, "\n")
					for _, fileWithPath := range filesInPackage {
						licenseRegex, _ := regexp.Compile(".*(LICENSE|license|License)[\\..*]??")
						gemRegex, _ := regexp.Compile(".*\\.gem$")
						archiveRegex, _ := regexp.Compile(".*\\.(gz|tgz|zip)$")
						if licenseRegex.MatchString(fileWithPath) {
							filesByType["License"] = append(filesByType["License"], fmt.Sprintf("%s:%s", pkg.Name, fileWithPath))
						} else if gemRegex.MatchString(fileWithPath) {
							filesByType["Rubygem"] = append(filesByType["Rubygem"], fmt.Sprintf("%s:%s", pkg.Name, fileWithPath))
						} else if archiveRegex.MatchString(fileWithPath) {
							filesByType["Archive"] = append(filesByType["Archive"], fmt.Sprintf("%s:%s", pkg.Name, fileWithPath))
						}
					}
				} else {
					fmt.Printf("Files for package '%s': %s\n", pkg.Name, pkgSpecFiles)
				}
			}

			if sortByType {
				for typeName, files := range filesByType {
					fmt.Printf("=== Files for type '%s' (%v) ===\n%s", typeName, len(files), strings.Join(files, "\n"))
					fmt.Printf("\n\n")
				}
			}
		},
	}
	return command
}
