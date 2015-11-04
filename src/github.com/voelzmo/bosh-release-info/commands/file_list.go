package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshcmd "github.com/cloudfoundry/bosh-utils/fileutil"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	boshsys "github.com/cloudfoundry/bosh-utils/system"
	"github.com/codegangsta/cli"
	"github.com/voelzmo/bosh-release-info/output"
	"github.com/voelzmo/bosh-release-info/release"
)

func FileListCommand(fs boshsys.FileSystem, comp boshcmd.Compressor, logger boshlog.Logger) cli.Command {
	command := cli.Command{
		Name:  "file-list",
		Usage: "lists all files in all packages in this release",
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "sort-by-type",
				Usage: "sort the files according to their filetype",
			},
		},
		Action: func(c *cli.Context) {
			if len(c.Args()) != 1 {
				err := bosherr.Error("that's not how it works. provide the name to the release tarball!")
				output.Fail(err, logger)
			}
			releasePath := c.Args()[0]
			sortByType := c.Bool("sort-by-type")

			// relasePathSplit := strings.Split(releasePath, "/")
			// releaseName := relasePathSplit[len(relasePathSplit)-1]
			releaseName := path.Base(releasePath)
			fmt.Printf("Info for release: %s\n", releaseName)
			fmt.Printf("\n")

			tmpDir, err := ioutil.TempDir("", releaseName)
			if err != nil {
				output.Fail(bosherr.WrapError(err, "Failed creating temporary directory:"), logger)
			}
			defer os.RemoveAll(tmpDir)

			reader := release.NewReader(releasePath, tmpDir, fs, comp)

			manifest, err := reader.Read()
			if err != nil {
				output.Fail(bosherr.WrapError(err, "Failed reading release:"), logger)
			}
			fmt.Printf("Release name: %s\n", manifest.Name)
			fmt.Printf("\n")

			for _, pkg := range manifest.Packages {
				pkgSpecFiles, err := reader.ReadPackageSpecs(pkg.Name)
				if err != nil {
					output.Fail(bosherr.WrapError(err, "Failed reading package spec files:"), logger)
				}

				if sortByType {
					filesInPackage := strings.Split(pkgSpecFiles, "\n")
					filesByType := make(map[string][]string)
					for _, fileWithPath := range filesInPackage {
						file := path.Base(fileWithPath)
						re, _ := regexp.Compile(".*(LICENSE|license|License)[\\.md]??")
						if re.MatchString(fileWithPath) {
							fileSplit := strings.Split(file, ".")
							fileType := fileSplit[len(fileSplit)-1]
							filesByType[fileType] = append(filesByType[fileType], fmt.Sprintf("%s:%s", manifest.Name, fileWithPath))
						} else {

						}

					}
					for typeName, files := range filesByType {
						fmt.Printf("Files for type '%s'\n: %s", typeName, strings.Join(files, "\n"))
						fmt.Printf("\n")
					}

				} else {
					fmt.Printf("Files for package '%s': %s", pkg.Name, pkgSpecFiles)
					fmt.Printf("\n")
				}
			}
		},
	}
	return command
}
