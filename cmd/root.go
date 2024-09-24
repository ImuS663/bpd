package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "bpd [url...]",
	Short:   "BPD is a CLI tool for downloading files from websites by Xpath.",
	Args:    cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Example: "bpd https://example.com/p1 https://example.com/p2 -x '//*[@id=\"example\"]/div[1]' -H 'header1=value1' -H 'header2=value2'",
	Run:     func(cmd *cobra.Command, args []string) {},
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
