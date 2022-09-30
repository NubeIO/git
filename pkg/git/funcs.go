package git

import (
	"github.com/google/go-github/v32/github"
	"regexp"
	"strings"
)

// takes in flow-framework-0.5.0-340c0ad8.amd64.zip
// return 0.5.0
func getAssetVersion(asset string) string {
	parts := strings.Split(asset, "-")
	for _, p := range parts {
		match, _ := regexp.MatchString(`^(\d+\.)?(\d+\.)?(\*|\d+)$`, p)
		if match {
			return p
		}
	}
	return ""
}

// takes in /home/user/../bin/rubix-bios-app/v1.5.2
// return /home/user/bin/rubix-bios-app/v1.5.2
func cleanPath(asset string) string {
	return strings.Replace(asset, "/..", "", -1)
}

// ListOptions specifies the optional parameters to various List methods that support pagination.
type ListOptions = github.ListOptions

// RepositoryRelease represents a GitHub release in a repository.
type RepositoryRelease = github.RepositoryRelease

// ReleaseAsset represents a GitHub release asset in a repository.
type ReleaseAsset = github.ReleaseAsset
