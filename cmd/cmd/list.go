package cmd

import (
	"github.com/NubeIO/git/pkg/git"
	pprint "github.com/NubeIO/git/pkg/helpers/print"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Fetch github repository release list.",
	Long:  `Fetch github repository release list.`,
	RunE:  runList,
}

func runList(cmd *cobra.Command, args []string) error {
	client, err := initClient()
	opt := &git.ListOptions{
		Page:    page,
		PerPage: perPage,
	}
	resp, err := client.ListReleases(opt)
	if err != nil {
		return err
	}

	var results []*repoRelease
	for _, v := range resp {
		results = append(results, &repoRelease{
			ID:   *v.ID,
			Name: *v.Name,
			Tag:  *v.TagName,
			URL:  *v.HTMLURL,
		})
	}
	return pprint.PrintPrettyJSON(pprint.Cyan, results)
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
