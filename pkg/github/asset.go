package git

// AssetOptions are parameters to download an asset file.
type AssetOptions struct {
	Tag           string //latest or v.1.1.1
	Name          string //NubeIO/my-repo
	OS            string //linux, windows
	OSAlias       []string
	Arch          string //tag arch as in amd64 arm7
	ArchAlias     []string
	DestPath      string //where to unzip the download to--dest=bin
	Target        string //pass in a new target dir  --dest=bin --target=new  this will add the archive in bin/new/my-asset
	MatchOS       bool   //try and match the OS type from the name as in "myapp-linux-amd64", we would try and match the "linux"
	ManualInstall ManualInstall
}

type ManualInstall struct {
	Path  string //  /home/user/rubix-service-1.19.0-eb71da61.amd64.zip
	Asset string // /rubix-service-1.19.0-eb71da61.amd64.zip
	//DestinationPath    string //this is used when the user does not want to download the build from the GitHub, as in the already have the zip, so they just want to install the build
	DeleteAsset bool //delete the zip after the installation is done

}
