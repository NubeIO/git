package git

import (
	"github.com/google/go-github/v32/github"
	"io"
)

// GetContents gets release info.
func (inst *Client) GetContents(owner, repo, path string, opts *github.RepositoryContentGetOptions) (fileContent *github.RepositoryContent, directoryContent []*github.RepositoryContent, resp *github.Response, err error) {
	return inst.hub.Repositories.GetContents(inst.CTX, owner, repo, path, opts)
}

// DownloadContents download a file
func (inst *Client) DownloadContents(owner, repo, path string, byteSize int, opts *github.RepositoryContentGetOptions) (io.ReadCloser, []byte, error) {
	contents, err := inst.hub.Repositories.DownloadContents(inst.CTX, owner, repo, path, opts)
	if err != nil {
		return nil, nil, err
	}
	if byteSize == 0 {
		byteSize = 1026
	}
	b := make([]byte, byteSize)
	_, err = contents.Read(b)
	if err != nil {
		return nil, nil, err
	}
	return contents, b, err
}
