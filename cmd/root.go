package cmd

import (
	"io"
	"os"
	"strings"

	"github.com/ImuS663/bpd/pkg/file"
	"github.com/ImuS663/bpd/pkg/net"
	"github.com/ImuS663/bpd/pkg/parser"
	"github.com/ImuS663/bpd/pkg/pbar"
	"github.com/ImuS663/bpd/pkg/writer"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var version string = "dev-build"

var rootCmd = &cobra.Command{
	Use:     "bpd url...",
	Short:   "BPD is a CLI tool for downloading files from websites by Xpath.",
	Args:    cobra.MatchAll(cobra.MinimumNArgs(1), cobra.OnlyValidArgs),
	Example: "bpd https://example.com/p1 https://example.com/p2 -x '//*[@id=\"example\"]/div[1]' -H header1=value1 -H header2=value2 -o path/to/output/directory",
	Run:     run,
	Version: version,
}

var (
	xpath        string
	outDir       string
	headers      map[string]string
	allConfirmed bool
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	var defaultOutDir = "./"
	var defaultHeaders = make(map[string]string)

	if os.Getenv("BPD_OUT_DIR") != "" {
		defaultOutDir = os.Getenv("BPD_OUT_DIR")
	}

	if os.Getenv("BPD_HEADERS") != "" {
		headers := strings.Split(os.Getenv("BPD_HEADERS"), "|")

		for _, header := range headers {
			split := strings.Split(header, "=")
			if len(split) == 2 {
				defaultHeaders[split[0]] = strings.Join(split[1:], "")
			}
		}
	}

	rootCmd.Flags().StringVarP(&xpath, "xpath", "x", "", "xpath to the element (required)")
	rootCmd.Flags().StringVarP(&outDir, "out-dir", "o", defaultOutDir, "output directory PATH")
	rootCmd.Flags().StringToStringVarP(&headers, "header", "H", defaultHeaders, "request header")
	rootCmd.Flags().BoolVarP(&allConfirmed, "yes", "y", false, "confirm all prompts")

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

func validateArgs(args []string) []string {
	urls := make([]string, 0)

	for _, url := range args {
		if !net.ValidateURL(url) {
			if allConfirmed || confirm(pterm.Warning.Sprintf("'%s' is not a valid URL. Do you want to continue without it?", url)) {
				continue
			} else {
				os.Exit(0)
			}
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
			if allConfirmed || confirm(pterm.Warning.Sprintf("Error parsing '%s'. Do you want to continue without it?", url)) {
				continue
			} else {
				pterm.Error.Println(err)
				os.Exit(0)
			}
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

func confirm(msg string) bool {
	res, _ := pterm.DefaultInteractiveConfirm.Show(msg)

	// Move cursor up on one line
	pterm.Print("\x1b[A")
	// Clear line
	pterm.Print("\x1b[2K")

	return res
}
