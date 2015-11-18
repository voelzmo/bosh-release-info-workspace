package fileutil

import (
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshsys "github.com/cloudfoundry/bosh-utils/system"
)

type Lister interface {
	ListFilesInArchive(tarballPath string) (string, error)
}

type tarballLister struct {
	cmdRunner boshsys.CmdRunner
	fs        boshsys.FileSystem
}

func NewTarballLister(
	cmdRunner boshsys.CmdRunner,
	fs boshsys.FileSystem) Lister {
	return tarballLister{cmdRunner, fs}
}

func (t tarballLister) ListFilesInArchive(tarballPath string) (string, error) {

	stdout, _, _, err := t.cmdRunner.RunCommand("tar", "-tzf", tarballPath)
	if err != nil {
		return "", bosherr.WrapError(err, "Shelling out to tar")
	}

	return stdout, nil
}
