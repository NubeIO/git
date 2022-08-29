package git

import (
	"context"
	"errors"

	"fmt"
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
	RepositoryRelease *RepositoryRelease
	AssetName         string
}

type DownloadOptions struct {
	DownloadDestination string `json:"download_destination"`
	AssetName           string `json:"asset_name"`
	MatchName           bool   `json:"match_name"`
	MatchArch           bool   `json:"match_arch"`
	MatchOS             bool   `json:"match_os"`
	DownloadFirst       bool   `json:"download_first"`
}

func (inst *Client) downloadReleaseAsset(owner, repo string, id int64) (rc io.ReadCloser, redirectURL string, err error) {
	return inst.hub.Repositories.DownloadReleaseAsset(inst.CTX, owner, repo, id, nil)
}

func (inst *Client) DownloadRelease(owner, repo, destination string, id int64) error {
	if destination == "" {
		return errors.New("destination can not be empty")
	}
	_, url, err := inst.downloadReleaseAsset(owner, repo, id)
	if err != nil {
		return err
	}
	return downloadFile(destination, url)
}

func downloadFile(filepath string, url string) (err error) {
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

type AssetInfo struct {
	RepositoryRelease *github.ReleaseAsset `json:"repository_release"`
	Url               string               `json:"url"`
}

// MatchAssetInfo get info from a release asset file.
func (inst *Client) MatchAssetInfo(options DownloadOptions) (*AssetInfo, error) {
	var assetName = options.AssetName
	if assetName == "" {
		return nil, errors.New("asset name can not be empty (the asset name my not always be the repo name), try flow-framework")
	}
	var destination = options.DownloadDestination
	if destination == "" {
		return nil, errors.New("destination can not be empty try, /data/store/apps")
	}
	var matchName = options.MatchName
	var matchArch = options.MatchArch
	var matchOS = options.MatchOS
	var downloadFirst = options.DownloadFirst
	opt := inst.Opts
	release, err := inst.GetRelease()
	if err != nil {
		return nil, err
	}
	url := ""
	var asset *github.ReleaseAsset
	if len(release.Assets) == 0 {
		url = *release.ZipballURL
	} else {
		asset = inst.findReleaseAsset(release, assetName, matchName, matchArch, matchOS, downloadFirst)
		if asset == nil {
			err := fmt.Errorf("not found asset: [name: %s, os: %s, arch: %s]", opt.Repo, opt.OS, opt.Arch)
			return nil, err
		}
		url = asset.GetBrowserDownloadURL()
	}
	resp := &AssetInfo{
		RepositoryRelease: asset,
		Url:               url,
	}
	return resp, err
}

// Download downloads a release asset file.
func (inst *Client) Download(options DownloadOptions) (*DownloadResponse, error) {
	var assetName = options.AssetName
	if assetName == "" {
		return nil, errors.New("asset name can not be empty (the asset name my not always be the repo name), try flow-framework")
	}
	var destination = options.DownloadDestination
	if destination == "" {
		return nil, errors.New("destination can not be empty try, /data/store/apps")
	}
	var matchName = options.MatchName
	var matchArch = options.MatchArch
	var matchOS = options.MatchOS
	var downloadFirst = options.DownloadFirst
	opt := inst.Opts
	release, err := inst.GetRelease()
	if err != nil {
		return nil, err
	}
	url := ""
	if len(release.Assets) == 0 {
		url = *release.ZipballURL
	} else {
		asset := inst.findReleaseAsset(release, assetName, matchName, matchArch, matchOS, downloadFirst)
		if asset == nil {
			err := fmt.Errorf("not found asset: [name: %s, os: %s, arch: %s]", opt.Repo, opt.OS, opt.Arch)
			return nil, err
		}
		url = asset.GetBrowserDownloadURL()
	}
	log.Infof("found asset: [name: %s, os: %s, arch: %s]", opt.Repo, opt.OS, opt.Arch)
	log.Infof("release dl url:%s", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	filename := path.Base(url)
	destination = filepath.Join(destination, filename)
	out, err := os.Create(destination)
	if err != nil {
		return nil, err
	}
	defer out.Close()
	defer resp.Body.Close()
	log.Infof("put asset at destination: %s", destination)
	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return nil, err
	}
	res := &DownloadResponse{}
	res.AssetName = filename
	res.RepositoryRelease = release
	return res, err
}

func (inst *Client) findReleaseAsset(release *RepositoryRelease, assetName string, matchName, matchArch, matchOS, downloadFirst bool) *ReleaseAsset {
	if downloadFirst {
		if len(release.Assets) > 0 {
			return release.Assets[0]
		} else {
			return nil
		}
	}
	opt := inst.Opts
	for _, asset := range release.Assets {
		name := strings.ToLower(asset.GetName())
		if assetName == "" {
			assetName = opt.Repo
		}
		//log.Infof("matched: [name: %s]", name)
		matchedName := strings.Contains(name, strings.ToLower(assetName))
		//fmt.Println(matchedName, name, assetName)
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
		//log.Infof("matched: [name: %s, os: %t, arch: %t]", name, matchedOS, matchedArch)
		if matchArch && matchName {
			if matchedName && matchedArch {
				return asset
			}
		}
		if matchOS && matchName {
			if matchedName && matchedOS {
				return asset
			}
		}

	}
	return nil
}
