package cmd

import (
	"context"
	"github.com/NubeIO/git/pkg/github"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Fetch github repository release list.",
	Long:  `Fetch github repository release list.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		client := github.NewClient(githubToken(), verbose)
		opt := &github.ListOptions{
			Page:    page,
			PerPage: perPage,
		}

		resp, err := client.ListReleases(ctx, github.Repository(repo), opt)
		if err != nil {
			return err
		}

		if !verbose {
			var results []*repoRelease
			for _, v := range resp {
				results = append(results, &repoRelease{
					ID:   *v.ID,
					Name: *v.Name,
					Tag:  *v.TagName,
					URL:  *v.HTMLURL,
				})
			}

			return printPrettyJSON(Cyan, results)
		}

		color.Cyan("repository:\t%s", repo)
		color.Cyan("page-num:\t%d", page)
		color.Cyan("per-page:\t%d", perPage)

		return printPrettyJSON(Cyan, resp)
	},
}

type repoRelease struct {
	ID   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Tag  string `json:"tag,omitempty"`
	URL  string `json:"url,omitempty"`
}

var (
	page    = 1
	perPage = 10
)

func init() {
	rootCmd.AddCommand(listCmd)

	flagSet := listCmd.Flags()
	flagSet.IntVar(&page, "page", page, "request page number")
	flagSet.IntVar(&perPage, "per-page", perPage, "request per page count")
}
