package cmd

import (
	"os"

	"github.com/ImuS663/bpd/cmd/downloader"
	"github.com/ImuS663/bpd/pkg/net"
	"github.com/ImuS663/bpd/pkg/parser"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var msCmd = &cobra.Command{
	Use:     "ms xpath...",
	Short:   "Serch and download files from `single` webpage by `multiple` Xpath",
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

	var fulesUrls = make([]string, 0)

	for _, xpath := range args {
		parser := parser.NewParser(xpath)

		result, err := parser.ParseFileURL(url)
		if err != nil {
			pterm.Error.Println(err)
			continue
		}

		fulesUrls = append(fulesUrls, result)
	}

	if len(fulesUrls) == 0 {
		pterm.Error.Println("No files found")
		os.Exit(1)
	}

	downloader.Download(fulesUrls, headers, outDir, allConfirmed)
}
