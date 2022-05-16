# git

## example

set github token

```
export GITHUB_TAG=latest
export GITHUB_REPO=NubeIO/rubix-service
export GITHUB_TOKEN=YOUR-token
(cd pkg/github && go test -run TestInfo)
```

```go
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
```

## command docs

[CLI](docs/cli.md)
