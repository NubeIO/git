package git

import (
	"context"
	pprint "github.com/NubeIO/git/pkg/helpers/print"
	"github.com/google/go-github/v32/github"
	"testing"
)

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

	_, i, _, err := client.GetContents("NubeIO", "releases", "flow", &github.RepositoryContentGetOptions{})
	if err != nil {
		return
	}
	pprint.PrintJOSN(i)

}
