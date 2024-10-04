package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var version string = "dev-build"

var rootCmd = &cobra.Command{
	Use:     "bpd",
	Short:   "BPD is a CLI tool for downloading files from websites by Xpath.",
	Version: version,
}

var (
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

	rootCmd.PersistentFlags().StringVarP(&outDir, "out-dir", "o", defaultOutDir, "output directory PATH")
	rootCmd.PersistentFlags().StringToStringVarP(&headers, "header", "H", defaultHeaders, "request header")
	rootCmd.PersistentFlags().BoolVarP(&allConfirmed, "yes", "y", false, "confirm all prompts")

	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
