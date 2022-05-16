package git

import (
	"context"
	"fmt"
	"os"
	"testing"
)

/*
to run
export GITHUB_TOKEN=YOUR-token
(cd pkg/github && go test -run TestInfo)
*/

func githubToken() string {
	return os.Getenv("GITHUB_TOKEN")
}

func TestInfo(t *testing.T) {

	fmt.Println("GITHUB_TOKEN", githubToken())

	ctx := context.Background()
	client := NewClient(githubToken(), true)

	resp, err := client.GetRelease(ctx, NubeRubixService, TagLatest)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("GetRelease", resp.GetName())

}
