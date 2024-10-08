package downloader

import (
	"io"
	"os"

	"github.com/ImuS663/bpd/pkg/file"
	"github.com/ImuS663/bpd/pkg/net"
	"github.com/ImuS663/bpd/pkg/parser"
	"github.com/ImuS663/bpd/pkg/pbar"
	"github.com/ImuS663/bpd/pkg/writer"
	"github.com/pterm/pterm"
)

func Download(filesUrls []string, headers map[string]string, outDir string, allConfirmed bool) {
	for _, url := range filesUrls {
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

		if _, err := io.Copy(writer, reader); err != nil {
			pterm.Error.Println(err)
			continue
		}

		pterm.Success.Printf("File '%s' downloaded successfully\n", fileName)
	}
}

func ValidateArgs(args []string, allConfirmed bool) []string {
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

func ParseFilesByUrlsAndXPath(urls []string, xpath string, allConfirmed bool) []string {
	filesUrls := make([]string, 0)
	p := parser.NewParser(xpath)

	for _, url := range urls {
		result, err := p.ParseFileURL(url)
		if err != nil {
			if allConfirmed || confirm(pterm.Warning.Sprintf("Error parsing '%s'. Do you want to continue without it?", url)) {
				continue
			} else {
				pterm.Error.Println(err)
				os.Exit(0)
			}
		}

		filesUrls = append(filesUrls, result)
	}
	return filesUrls
}

func ParseFilesByXPathsAndUrl(xpaths []string, url string, allConfirmed bool) []string {
	var p parser.Parser

	filesUrls := make([]string, 0)

	for _, xpath := range xpaths {
		p = *parser.NewParser(xpath)

		result, err := p.ParseFileURL(url)
		if err != nil {
			if allConfirmed || confirm(pterm.Warning.Sprintf("Error parsing '%s'. Do you want to continue without it?", url)) {
				continue
			} else {
				pterm.Error.Println(err)
				os.Exit(0)
			}
		}

		filesUrls = append(filesUrls, result)
	}

	return filesUrls
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
