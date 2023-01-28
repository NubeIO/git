package cmd

import (
	"context"
	"fmt"
	"github.com/NubeIO/git/pkg/git"
	"github.com/spf13/cobra"
	"os"
	"runtime"
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
}

var (
	tokenEnv string
	token    string
	repo     string
)

var (
	owner     string
	tag       = "latest"
	osName    = runtime.GOOS
	osAlias   = "darwin:macos,osx;windows:win"
	arch      = runtime.GOARCH
	archAlias = "amd64:x86_64"
	dest, _   = os.Getwd()
)

func init() {
	pFlagSet := rootCmd.PersistentFlags()
	pFlagSet.StringVar(&tokenEnv, "token-env", "GITHUB_TOKEN", "github oauth2 token environment name")
	pFlagSet.StringVar(&token, "token", token, "github oauth2 token value (optional)")
	pFlagSet.StringVarP(&owner, "owner", "", "NubeIO", "github repository (OWNER/name)")
	pFlagSet.StringVarP(&repo, "repo", "", "rubix-bios", "github repository (owner/NAME)")
	pFlagSet.StringVar(&dest, "dest", dest, "destination path")

	flagSet := rootCmd.Flags()

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
	return &git.AssetOptions{
		Owner: owner,
		Repo:  repo,
		Tag:   tag,
		Arch:  arch,
	}, nil
}
