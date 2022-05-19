package git

// AssetOptions are parameters to download an asset file.
type AssetOptions struct {
	Tag            string //latest or v.1.1.1
	Repo           string //my-repo
	Owner          string //NubeIO
	OS             string //linux, windows
	OSAlias        []string
	Arch           string //tag arch as in amd64 arm7
	ArchAlias      []string
	DestPath       string //where to unzip the download to--dest=bin
	Target         string //pass in a new target dir  --dest=bin --target=new  this will add the archive in bin/new/my-asset
	MatchOS        bool   //try and match the OS type from the name as in "myapp-linux-amd64", we would try and match the "linux"
	VersionDirName bool   // set this to true and the asset version number will be used in the naming of the target dir (eg: /bin/bios/rubix-0.5)
	DeleteZip      bool
	downloadOnly   bool
	manualInstall  manualInstall
}

type manualInstall struct {
	path  string //  /home/user/rubix-service-1.19.0-eb71da61.amd64.zip
	asset string // /rubix-service-1.19.0-eb71da61.amd64.zip
	//DestinationPath    string //this is used when the user does not want to download the build from the GitHub, as in the already have the zip, so they just want to install the build
	deleteZip bool //delete the zip after the installation is done

}
