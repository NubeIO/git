package git

import (
	"context"
	"fmt"
	"github.com/NubeIO/git/pkg/archive"
	"github.com/google/go-github/v32/github"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// Client is a GitHub oauth2 hub.
type Client struct {
	hub  *github.Client //github
	Opts *AssetOptions
	CTX  context.Context
}

// NewClient creates GitHub hub.
func NewClient(accessToken string, Opts *AssetOptions, ctx context.Context) *Client {
	c := context.Background()
	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	return &Client{
		hub:  github.NewClient(oauth2.NewClient(c, tokenSource)),
		Opts: Opts,
		CTX:  ctx,
	}
}

// ListReleases get release list.
func (inst *Client) ListReleases(opt *ListOptions) ([]*RepositoryRelease, error) {
	releases, _, err := inst.hub.Repositories.ListReleases(inst.CTX, inst.Opts.Owner, inst.Opts.Repo, opt)
	return releases, err
}

// GetRelease gets release info.
func (inst *Client) GetRelease() (*RepositoryRelease, error) {

	if inst.Opts.Tag == "latest" {
		release, _, err := inst.hub.Repositories.GetLatestRelease(inst.CTX, inst.Opts.Owner, inst.Opts.Repo)
		return release, err
	}

	release, _, err := inst.hub.Repositories.GetReleaseByTag(inst.CTX, inst.Opts.Owner, inst.Opts.Repo, inst.Opts.Tag)
	return release, err
}

type DownloadResponse struct {
	ReleaseAsset     *ReleaseAsset
	Destination      string //the final destination of the build / ..bin/rubix-bios/v0.1.1
	DestinationFull  string //same as above but with a `pwd` /home/user/bin/rubix-bios/v0.1.1
	ExtractedVersion string //version number of the asset eg:v0.1.1
}

// DownloadInstall downloads a release asset file.
// first returns release asset info.
// third returns initialize error info.
func (inst *Client) DownloadInstall() (*DownloadResponse, error) {
	opt := inst.Opts
	release, err := inst.GetRelease()
	if err != nil {
		return nil, err
	}
	asset := inst.findReleaseAsset(release)
	if asset == nil {
		err := fmt.Errorf("not found asset: [name: %s, os: %s, arch: %s]", opt.Repo, opt.OS, opt.Arch)
		return nil, err
	}
	log.Infof("found asset: [name: %s, os: %s, arch: %s]", opt.Repo, opt.OS, opt.Arch)
	url := asset.GetBrowserDownloadURL()
	log.Infof("release dl url:%s", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	filename := path.Base(url)
	res, err := inst.unPacAsset(filename, resp.Body, opt)
	if err != nil {
		return nil, err
	}
	res.ReleaseAsset = asset
	return res, err
}

//DownloadOnly just download and unzip
func (inst *Client) DownloadOnly() (*DownloadResponse, error) {
	inst.Opts.DeleteZip = false
	inst.Opts.downloadOnly = true
	res, err := inst.DownloadInstall()
	return res, err
}

//InstallFromZip just install a pre downloaded  unzip
func (inst *Client) InstallFromZip(unzipPath, assetName string, deleteZip bool) (*DownloadResponse, error) {
	inst.Opts.DeleteZip = deleteZip
	full := fmt.Sprintf("%s/%s", unzipPath, assetName)
	log.Infof("do a manual install destination path: %s", full)
	zip, err := os.Open(full)
	if err != nil {
		return nil, err
	}
	defer zip.Close()
	res, err := inst.unPacAsset(full, zip, inst.Opts)
	return res, err
}

func (inst *Client) findReleaseAsset(release *RepositoryRelease) *ReleaseAsset {
	opt := inst.Opts
	for _, asset := range release.Assets {
		name := strings.ToLower(asset.GetName())
		matchedName := strings.Contains(name, strings.ToLower(opt.Repo))
		matchedOS := strings.Contains(name, strings.ToLower(opt.OS))
		if !matchedOS {
			for _, v := range opt.OSAlias {
				if matchedOS = strings.Contains(name, v); matchedOS {
					break
				}
			}
		}
		matchedArch := strings.Contains(name, strings.ToLower(opt.Arch))
		if !matchedArch {
			for _, v := range opt.ArchAlias {
				if matchedArch = strings.Contains(name, v); matchedArch {
					break
				}
			}
		}
		log.Infof("matched: [name: %t, os: %t, arch: %t]", matchedName, matchedOS, matchedArch)
		if opt.MatchOS {
			if matchedName && matchedArch && matchedOS {
				return asset
			}
		} else {
			if matchedName && matchedArch {
				return asset
			}
		}
	}
	return nil
}

func (inst *Client) InstallAsset(asset *ReleaseAsset) (*DownloadResponse, error) {
	opt := inst.Opts
	url := asset.GetBrowserDownloadURL()
	log.Infof("release dl url:%s", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	filename := path.Base(url)
	res, err := inst.unPacAsset(filename, resp.Body, opt)
	return res, err

}

func (inst *Client) unPacAsset(filename string, body io.ReadCloser, opt *AssetOptions) (*DownloadResponse, error) {
	destination := filepath.Join(opt.DestPath, filename)
	tempExt := ".rubix-downloads"

	if err := os.MkdirAll(filepath.Dir(destination), os.ModePerm); err != nil {
		return nil, err
	}
	file, err := os.Create(destination + tempExt)
	log.Infof("make tmp dir:%s", destination+tempExt)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	if _, err = io.Copy(file, body); err != nil {
		return nil, err
	}

	if err := os.Rename(destination+tempExt, destination); err != nil {
		return nil, err
	}
	log.Infof("rename tmp old: %s new:%s", destination+tempExt, destination)
	defer func() {
		if opt.manualInstall.deleteZip == true { //dont delete when doing manual install
			log.Infof("delete:%s", destination+tempExt)
			_ = os.Remove(destination + tempExt)
			log.Infof("delete:%s", destination)
			_ = os.Remove(destination)
			if opt.manualInstall.path != "" {
				manualPath := opt.manualInstall.path
				manualAsset := opt.manualInstall.asset
				deleteManual := fmt.Sprintf("%s/%s", manualPath, manualAsset)
				log.Infof("delete-manual:%s", deleteManual)
				_ = os.Remove(deleteManual)
			}
		}
	}()

	if !archive.Support(filename) {
		if opt.Target != "" {
			newDestination := filepath.Join(opt.DestPath, opt.Target)
			if err := os.Rename(destination, newDestination); err != nil {
				return nil, err
			}
		}
		return nil, err
	}
	newDestination := filepath.Join(opt.DestPath, opt.Target)
	if inst.Opts.VersionDirName {
		newDestination = fmt.Sprintf("%s/v%s", newDestination, getAssetVersion(filename))
	}

	pwd, _ := os.Getwd()
	res := &DownloadResponse{
		Destination:      newDestination,
		DestinationFull:  cleanPath(fmt.Sprintf("%s/%s", pwd, newDestination)),
		ExtractedVersion: getAssetVersion(filename),
	}

	log.Infof("new destination :%s", newDestination)
	if !inst.Opts.downloadOnly {
		if err := archive.UnArchive(destination, newDestination); err != nil {
			return nil, err
		}
	}

	return res, nil

}
