package cmd

import (
	"os"

	"github.com/ImuS663/bpd/cmd/downloader"
	"github.com/ImuS663/bpd/pkg/net"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var msCmd = &cobra.Command{
	Use:     "ms xpath...",
	Short:   "Search and download files from `single` webpage by `multiple` Xpath",
	Args:    cobra.MatchAll(cobra.MinimumNArgs(1), cobra.OnlyValidArgs),
	Example: "bpd ms '//*[@id=\"example\"]/div[1]' '//*[@id=\"example\"]/div[2]' -u https://example.com",
	Run:     runMs,
}

var url string

func init() {
	rootCmd.AddCommand(msCmd)

	msCmd.Flags().StringVarP(&url, "url", "u", "", "url (required)")
	msCmd.MarkFlagRequired("url")
}

func runMs(cmd *cobra.Command, args []string) {
	if !net.ValidateURL(url) {
		pterm.Error.Println("Invalid URL")
		os.Exit(1)
	}

	filesUrls := downloader.ParseFilesByXPathsAndUrl(args, url, allConfirmed)

	if len(filesUrls) == 0 {
		pterm.Error.Println("No files found")
		os.Exit(1)
	}

	downloader.Download(filesUrls, headers, outDir, allConfirmed)
}
