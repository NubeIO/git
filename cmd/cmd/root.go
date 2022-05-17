package cmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/NubeIO/git/pkg/git"
	"github.com/fatih/color"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"runtime"
	"strings"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:           "github-dl",
	Short:         "Download a github repository release asset.",
	Long:          fmt.Sprintf(`Download a github repository release asset`),
	SilenceErrors: true,
	SilenceUsage:  true,
	Run:           runRoot,
}

func initClient() (*git.Client, error) {
	opt, err := makeAssetOptions()
	ctx := context.Background()
	client := git.NewClient(githubToken(), opt, ctx)
	return client, err

}

func runRoot(cmd *cobra.Command, args []string) {
	client, err := initClient()

	if err != nil {
		color.Magenta(err.Error())
		fmt.Println(cmd.UsageString())
		os.Exit(1)
	}
	ass, err := client.DownloadReleaseAsset()
	if err != nil {
		log.Errorln(err)
		return
	}
	log.Infoln("download completed GetName", ass.ReleaseAsset.GetName())
	log.Infoln("download completed Destination:", ass.Destination)
	log.Infoln("download completed DestinationFull:", ass.DestinationFull)
	log.Infoln("download completed ExtractedVersion:", ass.ExtractedVersion)
}

var (
	tokenEnv string
	token    string
	repo     string
)

var (
	asset           string
	owner           string
	tag             = "latest"
	osName          = runtime.GOOS
	osAlias         = "darwin:macos,osx;windows:win"
	arch            = runtime.GOARCH
	archAlias       = "amd64:x86_64"
	dest, _         = os.Getwd()
	target          string
	manualDeleteZip bool //delete the zip after the installation
	manualPath      string
	manualAsset     string
	versionDirName  bool
)

func init() {
	pFlagSet := rootCmd.PersistentFlags()
	pFlagSet.StringVar(&tokenEnv, "token-env", "GITHUB_TOKEN", "github oauth2 token environment name")
	pFlagSet.StringVar(&token, "token", token, "github oauth2 token value (optional)")
	pFlagSet.StringVarP(&owner, "owner", "", "NubeIO", "github repository (OWNER/name)")
	pFlagSet.StringVarP(&repo, "repo", "", "NubeIO/rubix-bios", "github repository (owner/NAME)")
	pFlagSet.StringVar(&dest, "dest", dest, "destination path")
	pFlagSet.StringVar(&target, "target", target, "rename destination file (optional)")
	pFlagSet.BoolVarP(&versionDirName, "version-in-target", "", false, "set this to true and the asset version number will be used in the naming of the target dir (eg: /bin/bios/rubix-0.5)")

	flagSet := rootCmd.Flags()
	flagSet.StringVar(&asset, "asset", asset, "asset name keyword")
	flagSet.StringVar(&tag, "tag", tag, "release tag")
	flagSet.StringVar(&osName, "os", osName, "os keyword")
	flagSet.StringVar(&osAlias, "os-alias", osAlias, "os keyword alias")
	flagSet.StringVar(&arch, "arch", arch, "arch keyword")
	flagSet.StringVar(&archAlias, "arch-alias", archAlias, "arch keyword alias")

}

func githubToken() string {
	if token != "" {
		return token
	}
	return os.Getenv(tokenEnv)
}

func makeAssetOptions() (*git.AssetOptions, error) {

	osAliasMap, err := parseAlias(osAlias)
	if err != nil {
		return nil, errors.New("parse alias error: see flags --os-alias")
	}

	archAliasMap, err := parseAlias(archAlias)
	if err != nil {
		return nil, errors.New("parse alias error: see flags --arch-alias")
	}

	return &git.AssetOptions{
		Owner:          owner,
		Repo:           repo,
		Tag:            tag,
		OS:             osName,
		OSAlias:        osAliasMap[osName],
		Arch:           arch,
		ArchAlias:      archAliasMap[arch],
		DestPath:       dest,
		Target:         target,
		VersionDirName: versionDirName,
	}, nil
}

func parseAlias(flagAlias string) (map[string][]string, error) {
	ret := map[string][]string{}
	aliases := strings.Split(flagAlias, ";")
	for _, alias := range aliases {
		kv := strings.Split(alias, ":")
		if len(kv) != 2 {
			return nil, fmt.Errorf("parse alias: %v", kv)
		}
		k, v := kv[0], strings.Split(kv[1], ",")
		ret[k] = v
	}
	return ret, nil
}
