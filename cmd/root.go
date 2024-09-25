package cmd

import (
	"io"
	"os"

	"github.com/ImuS663/bpd/pkg/file"
	"github.com/ImuS663/bpd/pkg/net"
	"github.com/ImuS663/bpd/pkg/parser"
	"github.com/ImuS663/bpd/pkg/pbar"
	"github.com/ImuS663/bpd/pkg/writer"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "bpd [url...]",
	Short:   "BPD is a CLI tool for downloading files from websites by Xpath.",
	Args:    cobra.MatchAll(cobra.MinimumNArgs(1), cobra.OnlyValidArgs),
	Example: "bpd https://example.com/p1 https://example.com/p2 -x '//*[@id=\"example\"]/div[1]' -H 'header1=value1' -H 'header2=value2'",
	Run:     run,
}

var (
	xpath   string
	outDir  string
	headers map[string]string
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&xpath, "xpath", "x", "", "Xpath to the element (required)")
	rootCmd.Flags().StringVarP(&outDir, "out-dir", "o", ".", "Output directory PATH")
	rootCmd.Flags().StringToStringVarP(&headers, "header", "H", nil, "Request header")

	rootCmd.MarkFlagRequired("xpath")
}

func run(cmd *cobra.Command, args []string) {
	urls := validateArgs(args)

	if len(urls) == 0 {
		pterm.Error.Println("No valid URLs found")
		os.Exit(1)
	}

	fulesUrls := parseFilesArgs(urls)

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

func validateArgs(args []string) []string {
	urls := make([]string, 0)

	for _, url := range args {
		if !net.ValidateURL(url) {
			pterm.Info.Printf("Invalid URL: %s\n", url)
		} else {
			urls = append(urls, url)
		}
	}
	return urls
}

func parseFilesArgs(urls []string) []string {
	fulesUrls := make([]string, 0)
	parser := parser.NewParser(xpath)

	for _, url := range urls {
		result, err := parser.ParseFileURL(url)
		if err != nil {
			pterm.Error.Println(err)
			continue
		}

		fulesUrls = append(fulesUrls, result)
	}
	return fulesUrls
}

func initWriter(fileName string, count int64, filePath string) (*writer.ProgressWriter, error) {
	pbar := pbar.NewPTermProgressBar(fileName, count)

	file, err := file.OpenFile(filePath)
	if err != nil {
		return nil, err
	}

	writer := writer.NewProgressWriter(file, pbar)
	return writer, nil
}
