package git

import (
	"context"
	"fmt"
	pprint "github.com/NubeIO/git/pkg/helpers/print"
	"github.com/google/go-github/v32/github"
	"testing"
)

func TestDownloadReleaseAsset(t *testing.T) {
	opts := &AssetOptions{
		Owner: "NubeIO",
		Repo:  "nubeio-rubix-app-lora-serial-py",
		Tag:   "latest",
		Arch:  "armv7",
	}
	token := "Z2hwX2pDU0tteWxrVjkzN1Z5NmFFUHlPVFpObEhoTEdITjBYemxkSA=="
	ctx := context.Background()
	client := NewClient(DecodeToken(token), opts, ctx)
	err := client.DownloadReleaseAsset(opts.Owner, opts.Repo, "./test.zip", 75784608)
	fmt.Println(err)
}

func TestGetAssetInfo(t *testing.T) {
	opts := &AssetOptions{
		Owner: "NubeIO",
		Repo:  "flow-framework",
		Tag:   "latest",
		Arch:  "amd64",
	}

	token := "Z2hwX2pDU0tteWxrVjkzN1Z5NmFFUHlPVFpObEhoTEdITjBYemxkSA=="

	ctx := context.Background()
	client := NewClient(DecodeToken(token), opts, ctx)
	releaseAsset, err := client.GetReleaseAsset(DownloadOptions{
		DownloadDestination: ".",
		AssetName:           "bacnetmaster",
		MatchName:           true,
		MatchArch:           true,
		MatchOS:             false,
	})
	fmt.Println(err, releaseAsset.ID)
	pprint.PrintJOSN(releaseAsset)

}

func TestList(t *testing.T) {
	opts := &AssetOptions{
		Owner: "NubeIO",
		Repo:  "nubeio-rubix-app-lora-serial-py",
		Tag:   "latest",
		Arch:  "armv7",
	}

	token := "Z2hwX2pDU0tteWxrVjkzN1Z5NmFFUHlPVFpObEhoTEdITjBYemxkSA=="

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

	token := "Z2hwX2pDU0tteWxrVjkzN1Z5NmFFUHlPVFpObEhoTEdITjBYemxkSA=="

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
	token := "Z2hwX2pDU0tteWxrVjkzN1Z5NmFFUHlPVFpObEhoTEdITjBYemxkSA=="
	ctx := context.Background()
	client := NewClient(DecodeToken(token), opts, ctx)
	_, dir, _, err := client.GetContents("NubeIO", "releases", "flow", &github.RepositoryContentGetOptions{})
	if err != nil {
		return
	}
	pprint.PrintJOSN(dir)
}
