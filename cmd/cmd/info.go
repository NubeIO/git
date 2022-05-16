package cmd

import (
	"context"
	"github.com/NubeIO/git/pkg/github"
	"github.com/NubeIO/git/pkg/helpers/print"
	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Show github repository release info.",
	Long:  `Show github repository release info.`,
	RunE:  runInfo,
}

func runInfo(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	client := git.NewClient(githubToken())

	resp, err := client.GetRelease(ctx, git.Repository(repo), tag)
	if err != nil {
		return err
	}

	var results repoRelease
	results.URL = resp.GetAssetsURL()
	results.Name = resp.GetName()

	return pprint.PrintPrettyJSON(pprint.Cyan, results)
}

func init() {
	rootCmd.AddCommand(infoCmd)

	flagSet := infoCmd.Flags()
	flagSet.StringVar(&tag, "tag", tag, "release tag")
}
