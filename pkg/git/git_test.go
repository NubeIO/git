package git

import (
	"context"
	"fmt"
	pprint "github.com/NubeIO/git/pkg/helpers/print"
	"github.com/google/go-github/v32/github"
	"testing"
)

const token = "Z2hwX3pIdklCZFZPWmd5N1M2YXFtcHBWMHRkcndIbUk5eTNEMnlQMg=="

func TestDownloadReleaseAsset(t *testing.T) {
	opts := &AssetOptions{
		Owner: "NubeIO",
		Repo:  "nubeio-rubix-app-lora-serial-py",
		Tag:   "latest",
		Arch:  "armv7",
	}
	ctx := context.Background()
	client := NewClient(DecodeToken(token), opts, ctx)
	err := client.DownloadReleaseAsset(opts.Owner, opts.Repo, "./test.zip", 75784608)
	fmt.Println(err)
}

func TestGetAssetInfo(t *testing.T) {
	opts := &AssetOptions{
		Owner: "NubeIO",
		Repo:  "rubix-io-fw",
		Tag:   "v3.4",
		Arch:  "",
	}

	ctx := context.Background()
	client := NewClient(DecodeToken(token), opts, ctx)
	releaseAsset, err := client.GetReleaseAsset(DownloadOptions{
		DownloadDestination: ".",
		AssetName:           "r-io-modbus-v3.4",
		MatchName:           false,
		MatchArch:           false,
		NameContains:        true,
	})
	fmt.Println(err, releaseAsset)
	pprint.PrintJOSN(releaseAsset)
}

func TestList(t *testing.T) {
	opts := &AssetOptions{
		Owner: "NubeIO",
		Repo:  "rubix-io-fw",
		Tag:   "latest",
		Arch:  "",
	}

	ctx := context.Background()
	client := NewClient(DecodeToken(token), opts, ctx)

	releases, err := client.ListReleases(&ListOptions{
		Page:    0,
		PerPage: 0,
	})
	pprint.PrintJOSN(releases)
	if err != nil {
		return
	}
}

func TestDownloadZipball(t *testing.T) {
	opts := &AssetOptions{
		Owner: "NubeIO",
		Repo:  "wires-builds",
		Tag:   "latest",
	}

	ctx := context.Background()
	client := NewClient(DecodeToken(token), opts, ctx)
	download, err := client.DownloadZipball(DownloadOptions{
		DownloadDestination: ".",
		AssetName:           "wires-builds",
	})
	fmt.Println(err)
	if err != nil {
		return
	}
	pprint.Print(download)
}

func TestInfo(t *testing.T) {
	opts := &AssetOptions{
		Owner: "NubeIO",
		Repo:  "releases",
		Tag:   "latest",
	}

	ctx := context.Background()
	client := NewClient(DecodeToken(token), opts, ctx)
	_, dir, _, err := client.GetContents("NubeIO", "releases", "flow", &github.RepositoryContentGetOptions{})
	if err != nil {
		return
	}
	pprint.PrintJOSN(dir)
}
