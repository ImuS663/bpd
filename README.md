# BPD

BPD is a CLI tool for downloading files from websites by Xpath.

## Features

- Download files from websites using Xpath
- Support for multiple URLs
- Progress bar for large downloads
- Customizable request headers

## Build

To build BPD, run the following command:

```shell
go build -o bpd
```

## Usage

BPD can be used to download files from websites by specifying the URL and Xpath of the file. For example:

```shell
bpd 'https://example.com/p1' 'https://example.com/p2' -x '//*[@id="example"]/div[1]' -H 'header1=value1' -H 'header2=value2' -o '/absolute/path/to/output/directory'
```

This command will download the files from the specified URLs and save them to the current directory.

## Options

- `-x` or `--xpath`: Specify the Xpath of the file to download
- `-o` or `--out-dir`: Specify the output directory for the downloaded files
- `-H` or `--header`: Specify a custom request header

## License

BPD is released under the MIT License. See [LICENSE](https://github.com/ImuS663/bpd/blob/main/LICENSE) for details.
