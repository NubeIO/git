package git

import "github.com/google/go-github/v32/github"

// GetContents gets release info.
func (inst *Client) GetContents(owner, repo, path string, opts *github.RepositoryContentGetOptions) (fileContent *github.RepositoryContent, directoryContent []*github.RepositoryContent, resp *github.Response, err error) {
	return inst.hub.Repositories.GetContents(inst.CTX, owner, repo, path, opts)
}
