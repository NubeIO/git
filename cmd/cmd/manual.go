package cmd

import (
	git "github.com/NubeIO/git/pkg/github"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var unzipCmd = &cobra.Command{
	Use:   "manual",
	Short: "do a manual install",
	Long:  `when the user has no internet to download the build but has a copy of the zip on the PC`,
	Run:   runUnzip,
}

func runUnzip(cmd *cobra.Command, args []string) {

	client := git.NewClient("")
	opt := &git.AssetOptions{
		DestPath: dest,
		ManualInstall: git.ManualInstall{
			Path:        manualPath,
			Asset:       manualAsset,
			DeleteAsset: manualDeleteZip,
		},
	}
	err := client.DownloadAsset(&git.ReleaseAsset{}, opt)
	if err != nil {
		log.Errorln(err)
		return
	}
}

func init() {
	rootCmd.AddCommand(unzipCmd)
	flagSet := unzipCmd.Flags()
	flagSet.BoolVarP(&manualDeleteZip, "dont-delete", "", false, "delete the zip after the install")
	flagSet.StringVar(&manualPath, "manual-path", manualPath, "manual asset path (eg: /home/user)")
	flagSet.StringVar(&manualAsset, "manual-asset", manualAsset, "manual asset name (eg: rubix-service-1.19.0-eb71da61.amd64.zip)")

}
