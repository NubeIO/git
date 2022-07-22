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

	//_, raw, err := client.DownloadContents("NubeIO", "releases", "flow/v0.6.1.json", 0, &github.RepositoryContentGetOptions{})
	//fmt.Println(err)
	//pprint.Log(raw)
	//var r Release
	//json.Unmarshal(raw, &r)
	//pprint.Log(r)

}
