package git

import (
	"github.com/google/go-github/v32/github"
)

// Repository is github repository. (owner/name)
//type Repository string
//
//// Owner is a repository owner part.
//func (r Repository) Owner() string {
//	parts := strings.Split(string(r), "/")
//	if len(parts) == 2 {
//		return parts[0]
//	}
//	return ""
//}
//
//// Name is a repository name part.
//func (r Repository) Name() string {
//	parts := strings.Split(string(r), "/")
//	if len(parts) == 2 {
//		return parts[1]
//	}
//	return ""
//}
//
//func (r Repository) valid() error {
//	parts := strings.Split(string(r), "/")
//	if len(parts) != 2 {
//		return fmt.Errorf("malformed github repository: %s", r)
//	}
//	return nil
//}

// ListOptions specifies the optional parameters to various List methods that support pagination.
type ListOptions = github.ListOptions

// RepositoryRelease represents a GitHub release in a repository.
type RepositoryRelease = github.RepositoryRelease

// ReleaseAsset represents a GitHub release asset in a repository.
type ReleaseAsset = github.ReleaseAsset
