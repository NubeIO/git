package git

import (
	"context"
	"fmt"
	pprint "github.com/NubeIO/git/pkg/helpers/print"
	"github.com/google/go-github/v32/github"
	"io"
	"os"
	"testing"
)

type Release struct {
	Name    string `json:"name"`
	Repo    string `json:"repo"`
	Release string `json:"release"`
	Apps    []struct {
		Name             string   `json:"name"`
		Repo             string   `json:"repo"`
		Description      string   `json:"description"`
		Products         []string `json:"products"`
		Versions         []string `json:"versions"`
		FlowDependency   bool     `json:"flow_dependency"`
		PluginDependency string   `json:"plugin_dependency"`
	} `json:"apps"`
}

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

	asset, _, err := client.DownloadReleaseAsset(opts.Owner, opts.Repo, 75784608)
	fmt.Println(err)
	if err != nil {
		return
	}

	outFile, err := os.Create("./aap.zip")
	// handle err
	defer outFile.Close()

	_, err = io.Copy(outFile, asset)
	fmt.Println(err)
	if err != nil {

	}

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

	//asset, s, err := inst.hub.Repositories.DownloadReleaseAsset(inst.CTX, "NubeIO", "flow-framework")
	//if err != nil {
	//	return nil, err
	//}
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
	//pprint.PrintJOSN(contents)
	pprint.PrintJOSN(dir)

	//content, err := contents.GetContent()
	//if err != nil {
	//	return
	//}
	//fmt.Println(content)
	//var r *Release
	//json.Unmarshal([]byte(content), &r)
	//fmt.Println()
	//pprint.Print(r)

}
