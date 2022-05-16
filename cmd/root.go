package cmd

import (
	"context"
	"errors"
	"fmt"
	git "github.com/NubeIO/git/pkg/github"
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

func runRoot(cmd *cobra.Command, args []string) {
	ctx := context.Background()
	client := git.NewClient(githubToken(), verbose)

	opt, err := makeAssetOptions()
	if err != nil {
		color.Magenta(err.Error())
		fmt.Println(cmd.UsageString())
		os.Exit(1)
	}

	if verbose {
		color.Cyan("repository:\t%s", repo)
		color.Cyan("release tag:\t%s", tag)
	}

	ass, err := client.DownloadReleaseAsset(ctx, git.Repository(repo), opt)

	if err != nil {
		log.Errorln(err)
		return
	}
	log.Infoln("download completed", ass.GetName())

}

var (
	verbose  bool
	tokenEnv string
	token    string
	repo     string
)

var (
	asset     string
	tag       = "latest"
	osName    = runtime.GOOS
	osAlias   = "darwin:macos,osx;windows:win"
	arch      = runtime.GOARCH
	archAlias = "amd64:x86_64"
	dest, _   = os.Getwd()
	target    string
)

func init() {
	pFlagSet := rootCmd.PersistentFlags()
	pFlagSet.BoolVarP(&verbose, "verbose", "v", verbose, "verbose output")
	pFlagSet.StringVar(&tokenEnv, "token-env", "GITHUB_TOKEN", "github oauth2 token environment name")
	pFlagSet.StringVar(&token, "token", token, "github oauth2 token value (optional)")
	pFlagSet.StringVarP(&repo, "repo", "", "NubeIO/rubix-bios", "github repository (owner/name)")

	flagSet := rootCmd.Flags()
	flagSet.StringVar(&asset, "asset", asset, "asset name keyword")
	flagSet.StringVar(&tag, "tag", tag, "release tag")
	flagSet.StringVar(&osName, "os", osName, "os keyword")
	flagSet.StringVar(&osAlias, "os-alias", osAlias, "os keyword alias")
	flagSet.StringVar(&arch, "arch", arch, "arch keyword")
	flagSet.StringVar(&archAlias, "arch-alias", archAlias, "arch keyword alias")
	flagSet.StringVar(&dest, "dest", dest, "destination path")
	flagSet.StringVar(&target, "target", target, "rename destination file (optional)")

}

func githubToken() string {
	if token != "" {
		return token
	}
	return os.Getenv(tokenEnv)
}

func makeAssetOptions() (*git.AssetOptions, error) {
	if asset == "" {
		return nil, errors.New("require asset name: see flags --asset")
	}

	osAliasMap, err := parseAlias(osAlias)
	if err != nil {
		return nil, errors.New("parse alias error: see flags --os-alias")
	}

	archAliasMap, err := parseAlias(archAlias)
	if err != nil {
		return nil, errors.New("parse alias error: see flags --arch-alias")
	}

	return &git.AssetOptions{
		Name:      asset,
		Tag:       tag,
		OS:        osName,
		OSAlias:   osAliasMap[osName],
		Arch:      arch,
		ArchAlias: archAliasMap[arch],
		DestPath:  dest,
		Target:    target,
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
