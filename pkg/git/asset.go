package git

// AssetOptions are parameters to download an asset file.
type AssetOptions struct {
	Tag       string //latest or v.1.1.1
	Repo      string //my-repo
	Owner     string //NubeIO
	OS        string //linux, windows
	OSAlias   []string
	Arch      string //tag arch as in amd64 arm7
	ArchAlias []string
	MatchOS   bool //try and match the OS type from the name as in "myapp-linux-amd64", we would try and match the "linux"
}
