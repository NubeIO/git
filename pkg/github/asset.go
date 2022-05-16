package git

// AssetOptions are parameters to download an asset file.
type AssetOptions struct {
	Tag       string
	Name      string
	OS        string
	OSAlias   []string
	Arch      string
	ArchAlias []string
	DestPath  string
	Target    string
	MatchOS   bool //try and match the OS type from the name as in "myapp-linux-amd64", we would try and match the "linux"
}
