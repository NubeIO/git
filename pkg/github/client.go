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

const NubeIO = "NubeIO"
const NubeFlow = "NubeIO/flow-framework"
const NubeBios = "NubeIO/rubix-bios"
const NubeRubixService = "NubeIO/rubix-service"
const NubeWires = "NubeIO/rubix-wires"
const NubeWiresBuild = "NubeIO/wires-builds"
const NubeRubixIO = "NubeIO/nubeio-rubix-app-pi-gpio-go"
const TagLatest = "latest"

// Client is a GitHub oauth2 client.
type Client struct {
	client *github.Client
}

// NewClient creates github client.
func NewClient(accessToken string) *Client {
	ctx := context.Background()
	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	return &Client{
		client: github.NewClient(oauth2.NewClient(ctx, tokenSource)),
	}
}

// ListReleases get release list.
func (c *Client) ListReleases(ctx context.Context, repo Repository, opt *ListOptions) ([]*RepositoryRelease, error) {
	if err := repo.valid(); err != nil {
		return nil, err
	}

	releases, _, err := c.client.Repositories.ListReleases(ctx, repo.Owner(), repo.Name(), opt)
	return releases, err
}

// GetRelease gets release info.
func (c *Client) GetRelease(ctx context.Context, repo Repository, tag string) (*RepositoryRelease, error) {
	if err := repo.valid(); err != nil {
		return nil, err
	}

	if tag == "latest" {
		release, _, err := c.client.Repositories.GetLatestRelease(ctx, repo.Owner(), repo.Name())
		return release, err
	}

	release, _, err := c.client.Repositories.GetReleaseByTag(ctx, repo.Owner(), repo.Name(), tag)
	return release, err
}

// DownloadReleaseAsset downloads a release asset file.
// first returns release asset info.
// second returns download progress info or error info use a stream.
// third returns initialize error info.
func (c *Client) DownloadReleaseAsset(ctx context.Context, repo Repository, opt *AssetOptions) (*ReleaseAsset, error) {
	if err := repo.valid(); err != nil {
		return nil, err
	}

	release, err := c.GetRelease(ctx, repo, opt.Tag)
	if err != nil {
		return nil, err
	}

	asset := c.findReleaseAsset(release, opt)
	if asset == nil {
		err := fmt.Errorf("not found asset: [name: %s, os: %s, arch: %s]", opt.Name, opt.OS, opt.Arch)
		return nil, err
	}
	log.Infof("found asset: [name: %s, os: %s, arch: %s]", opt.Name, opt.OS, opt.Arch)
	err = c.downloadAsset(asset, opt)
	return asset, err
}

func (c *Client) findReleaseAsset(release *RepositoryRelease, opt *AssetOptions) *ReleaseAsset {
	for _, asset := range release.Assets {
		name := strings.ToLower(asset.GetName())
		matchedName := strings.Contains(name, strings.ToLower(opt.Name))
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

func (c *Client) downloadAsset(asset *ReleaseAsset, opt *AssetOptions) error {
	url := asset.GetBrowserDownloadURL()
	log.Infof("release dl url:%s", url)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	filename := path.Base(url)
	destination := filepath.Join(opt.DestPath, filename)
	tempExt := ".rubix-downloads"

	if err := os.MkdirAll(filepath.Dir(destination), os.ModePerm); err != nil {
		return err
	}
	file, err := os.Create(destination + tempExt)
	log.Infof("make tmp dir:%s", destination+tempExt)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err = io.Copy(file, resp.Body); err != nil {
		return err
	}

	if err := os.Rename(destination+tempExt, destination); err != nil {
		return err
	}
	log.Infof("rename tmp old: %s new:%s", destination+tempExt, destination)
	defer func() {
		log.Infof("delete:%s", destination+tempExt)
		_ = os.Remove(destination + tempExt)
		log.Infof("delete:%s", destination)
		_ = os.Remove(destination)
	}()

	if !archive.Support(filename) {
		if opt.Target != "" {
			newDestination := filepath.Join(opt.DestPath, opt.Target)
			if err := os.Rename(destination, newDestination); err != nil {
				return err
			}
		}
		return nil
	}

	newDestination := filepath.Join(opt.DestPath, opt.Target)
	log.Infof("new destination of zip/tar:%s", newDestination)
	if err := archive.UnArchive(destination, newDestination); err != nil {
		return err
	}
	return nil

}
