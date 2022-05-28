package git

import (
	"context"

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
}

// Download downloads a release asset file.
func (inst *Client) Download(destination string) (*DownloadResponse, error) {
	opt := inst.Opts
	release, err := inst.GetRelease()
	if err != nil {
		return nil, err
	}
	url := ""
	if len(release.Assets) == 0 {
		url = *release.ZipballURL
	} else {
		asset := inst.findReleaseAsset(release)
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

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return nil, err
	}
	res := &DownloadResponse{}
	res.RepositoryRelease = release
	return res, err
}

func (inst *Client) findReleaseAsset(release *RepositoryRelease) *ReleaseAsset {
	opt := inst.Opts
	for _, asset := range release.Assets {
		name := strings.ToLower(asset.GetName())
		log.Infof("matched: [name: %s]", name)
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
		log.Infof("matched: [name: %s, os: %t, arch: %t]", name, matchedOS, matchedArch)
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
