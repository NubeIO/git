package cmd

import (
	"context"
	"github.com/NubeIO/git/pkg/github"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Show github repository release info.",
	Long:  `Show github repository release info.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		client := github.NewClient(githubToken(), verbose)

		resp, err := client.GetRelease(ctx, github.Repository(repo), tag)
		if err != nil {
			return err
		}

		if verbose {
			color.Cyan("repository:\t%s", repo)
			color.Cyan("release tag:\t%s", tag)
		}

		return printPrettyJSON(Cyan, resp)
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)

	flagSet := infoCmd.Flags()
	flagSet.StringVar(&tag, "tag", tag, "release tag")
}
