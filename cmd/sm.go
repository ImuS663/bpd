package cmd

import (
	"os"

	"github.com/ImuS663/bpd/cmd/downloader"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var smCmd = &cobra.Command{
	Use:     "sm url...",
	Short:   "Search and download files from `multiple` webpages by `single` Xpath",
	Args:    cobra.MatchAll(cobra.MinimumNArgs(1), cobra.OnlyValidArgs),
	Example: "bpd sm https://example.com/p1 https://example.com/p2 -x '//*[@id=\"example\"]/div[1]'",
	Run:     runSm,
}

var xpath string

func init() {
	rootCmd.AddCommand(smCmd)

	smCmd.Flags().StringVarP(&xpath, "xpath", "x", "", "xPath expression (required)")
	smCmd.MarkFlagRequired("xpath")
}

func runSm(cmd *cobra.Command, args []string) {
	urls := downloader.ValidateArgs(args, allConfirmed)

	if len(urls) == 0 {
		pterm.Error.Println("No valid URLs found")
		os.Exit(1)
	}

	filesUrls := downloader.ParseFilesByUrlsAndXPath(urls, xpath, allConfirmed)

	if len(filesUrls) == 0 {
		pterm.Error.Println("No files found")
		os.Exit(1)
	}

	downloader.Download(filesUrls, headers, outDir, allConfirmed)
}
