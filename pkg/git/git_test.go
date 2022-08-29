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
		//Repo:  "flow-framework",
		Repo: "nubeio-rubix-app-lora-serial-py",
		Tag:  "latest",
		//Arch:  "armv7",  amd64
		Arch: "armv7",
	}
	token := "Z2hwX2pDU0tteWxrVjkzN1Z5NmFFUHlPVFpObEhoTEdITjBYemxkSA=="
	ctx := context.Background()
	client := NewClient(DecodeToken(token), opts, ctx)
	err := client.DownloadRelease(opts.Owner, opts.Repo, "./test.zip", 75784608)
	fmt.Println(err)

}

func TestGetAssetInfo(t *testing.T) {

	opts := &AssetOptions{
		Owner: "NubeIO",
		//Repo:  "flow-framework",
		Repo: "flow-framework",
		Tag:  "latest",
		//Arch:  "armv7",  amd64
		Arch: "amd64",
	}

	token := "Z2hwX2pDU0tteWxrVjkzN1Z5NmFFUHlPVFpObEhoTEdITjBYemxkSA=="

	ctx := context.Background()
	client := NewClient(DecodeToken(token), opts, ctx)
	download, err := client.MatchAssetInfo(DownloadOptions{
		DownloadDestination: ".",
		AssetName:           "bacnetmaster",
		MatchName:           true,
		MatchArch:           true,
		MatchOS:             false,
		DownloadFirst:       false,
	})
	fmt.Println(err, download.RepositoryRelease.ID)
	pprint.PrintJOSN(download)

}

func TestList(t *testing.T) {

	opts := &AssetOptions{
		Owner: "NubeIO",
		//Repo:  "flow-framework",
		Repo: "nubeio-rubix-app-lora-serial-py",
		Tag:  "latest",
		//Arch:  "armv7",  amd64
		Arch: "armv7",
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

func TestDownload(t *testing.T) {
	opts := &AssetOptions{
		Owner: "NubeIO",
		//Repo:  "flow-framework",
		Repo: "nubeio-rubix-app-lora-serial-py",
		Tag:  "latest",
		//Arch:  "armv7",  amd64
		Arch: "armv7",
	}

	token := "Z2hwX2pDU0tteWxrVjkzN1Z5NmFFUHlPVFpObEhoTEdITjBYemxkSA=="

	ctx := context.Background()
	client := NewClient(DecodeToken(token), opts, ctx)
	download, err := client.Download(DownloadOptions{
		DownloadDestination: ".",
		AssetName:           "nubeio-rubix-app-lora-serial-py",
		MatchName:           true,
		MatchArch:           true,
		MatchOS:             false,
		DownloadFirst:       false,
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
		Arch:  "",
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
