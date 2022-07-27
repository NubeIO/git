package git

import (
	"context"
	"encoding/json"
	"fmt"
	pprint "github.com/NubeIO/git/pkg/helpers/print"
	"github.com/google/go-github/v32/github"
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

func TestList(t *testing.T) {

	//word1 := "nubeio-rubix-app-bbb-rest-py-0.0.2-0f2f6c5d.amd64.zip"
	//
	//word2 := "bbb-rest-py"
	//res, err := edlib.StringsSimilarity(word1, word2, edlib.Levenshtein)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Printf("Similarity: %f", res)
	//}
}

func TestDownload(t *testing.T) {
	opts := &AssetOptions{
		Owner: "NubeIO",
		//Repo:  "flow-framework",
		Repo: "wires-builds",
		Tag:  "latest",
		//Arch:  "armv7",
		Arch: "amd64",
	}

	token := "Z2hwX2pDU0tteWxrVjkzN1Z5NmFFUHlPVFpObEhoTEdITjBYemxkSA=="

	ctx := context.Background()
	client := NewClient(DecodeToken(token), opts, ctx)
	download, err := client.Download(DownloadOptions{
		DownloadDestination: "./",
		AssetName:           "flow-framework",
		MatchName:           true,
		MatchArch:           true,
		MatchOS:             false,
		DownloadFirst:       true,
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

	contents, _, _, err := client.GetContents("NubeIO", "releases", "flow/v0.6.1.json", &github.RepositoryContentGetOptions{})
	if err != nil {
		return
	}

	content, err := contents.GetContent()
	if err != nil {
		return
	}
	var r *Release
	json.Unmarshal([]byte(content), &r)
	fmt.Println()
	pprint.Print(r.Release)

}
