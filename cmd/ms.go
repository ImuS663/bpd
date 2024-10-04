package cmd

import (
	"io"
	"os"

	"github.com/ImuS663/bpd/pkg/file"
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

	for _, url := range fulesUrls {
		reader, count, err := net.InitReader(url, headers)
		if err != nil {
			pterm.Error.Println(err)
			continue
		}
		defer reader.Close()

		fileName, filePath := file.GetFilePathAndName(url, outDir)

		if file.Exists(filePath) {
			if !allConfirmed && !confirm(pterm.Warning.Sprintf("File '%s' already exists. Do you want to overwrite it?", fileName)) {
				continue
			}
		}

		writer, err := initWriter(fileName, count, filePath)
		if err != nil {
			pterm.Error.Println(err)
			continue
		}
		defer writer.Close()

		_, err = io.Copy(writer, reader)
		if err != nil {
			pterm.Error.Println(err)
			continue
		}

		pterm.Success.Printf("File '%s' downloaded successfully\n", fileName)
	}
}
