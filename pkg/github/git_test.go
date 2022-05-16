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
func githubRepo() string {
	return os.Getenv("GITHUB_REPO")
}
func githubTag() string {
	return os.Getenv("GITHUB_TAG")
}

func TestInfo(t *testing.T) {

	fmt.Println("GITHUB_TOKEN", githubToken(), "REPO", Repository(githubRepo()), "TAG", githubTag())

	ctx := context.Background()
	client := NewClient(githubToken())

	resp, err := client.GetRelease(ctx, Repository(githubRepo()), githubTag())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("GetRelease", resp.GetName())

}
