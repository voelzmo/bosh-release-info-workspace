package release

import (
	"fmt"
	"path"

	"gopkg.in/yaml.v2"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshcmd "github.com/cloudfoundry/bosh-utils/fileutil"
	boshsys "github.com/cloudfoundry/bosh-utils/system"
	"github.com/voelzmo/bosh-release-info/fileutil"
)

type Manifest struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`

	CommitHash         string `yaml:"commit_hash"`
	UncommittedChanges bool   `yaml:"uncommitted_changes"`

	Jobs     []JobRef     `yaml:"jobs"`
	Packages []PackageRef `yaml:"packages"`
}

type JobRef struct {
	Name        string `yaml:"name"`
	Fingerprint string `yaml:"fingerprint"`
	SHA1        string `yaml:"sha1"`
}

type PackageRef struct {
	Name         string   `yaml:"name"`
	Fingerprint  string   `yaml:"fingerprint"`
	SHA1         string   `yaml:"sha1"`
	Dependencies []string `yaml:"dependencies"`
}

type reader struct {
	tarFilePath          string
	extractedReleasePath string
	fs                   boshsys.FileSystem
	extractor            boshcmd.Compressor
	lister               fileutil.Lister
}

type PackageSpec struct {
	Name  string   `yaml:"name"`
	Files []string `yaml:"files"`
}

type Reader interface {
	ReadManifest() (Manifest, error)
	ReadPackageSpecs(string) (string, error)
}

func NewReader(
	tarFilePath string,
	extractedReleasePath string,
	fs boshsys.FileSystem,
	runner boshsys.CmdRunner,
) Reader {
	tarballExtractor := boshcmd.NewTarballCompressor(runner, fs)
	tarballLister := fileutil.NewTarballLister(runner, fs)

	return &reader{
		tarFilePath:          tarFilePath,
		extractedReleasePath: extractedReleasePath,
		fs:                   fs,
		extractor:            tarballExtractor,
		lister:               tarballLister,
	}
}

func (r *reader) ReadManifest() (Manifest, error) {
	err := r.extractor.DecompressFileToDir(r.tarFilePath, r.extractedReleasePath, boshcmd.CompressorOptions{})
	if err != nil {
		return Manifest{}, bosherr.WrapError(err, "Extracting release")
	}

	releaseManifestPath := path.Join(r.extractedReleasePath, "release.MF")
	releaseManifestBytes, err := r.fs.ReadFile(releaseManifestPath)
	if err != nil {
		return Manifest{}, bosherr.WrapErrorf(err, "Reading release manifest '%s'", releaseManifestPath)
	}

	var manifest Manifest
	err = yaml.Unmarshal(releaseManifestBytes, &manifest)
	if err != nil {
		return Manifest{}, bosherr.WrapError(err, "Parsing release manifest")
	}

	return manifest, nil
}

func (r *reader) ReadPackageSpecs(pkgName string) (string, error) {
	var allFiles string
	allFiles, _ = r.lister.ListFilesInArchive(path.Join(r.extractedReleasePath, fmt.Sprintf("packages/%s.tgz", pkgName)))

	return allFiles, nil
}
