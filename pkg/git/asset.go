package git

// AssetOptions are parameters to download an asset file.
type AssetOptions struct {
	Tag   string // latest or v.1.1.1
	Repo  string // repository
	Owner string // NubeIO
	Arch  string // tag arch as in amd64 arm7
}
