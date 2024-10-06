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
bpd sm 'https://example.com/p1' 'https://example.com/p2' -x '//*[@id="example"]/div[1]/a'
```

This command will download the files from the specified URLs and save them to the current directory.

or

```shell
bpd ms '//*[@id="example"]/div[1]/a' '//*[@id="example"]/div[2]/a' -u 'https://example.com'
```

This command will download the files from the specified URL and save them to the current directory.

## Global Options

- `-o` or `--out-dir`: Specify the output directory for the downloaded files
  - You can also set the `BPD_OUT_DIR` environment variable to override the default output directory. For example:
  
  ```shell
  export BPD_OUT_DIR=/path/to/output/directory
  ```

- `-H` or `--header`: Specify a custom request header
  - You can also set the `BPD_HEADERS` environment variable to override the default headers. For example:
  
  ```shell
  export BPD_HEADERS='Accept-Language=en-US|Referer=https://example.com/'
  ```

- `-y` or `--yes`: Confirm all prompts
- `-h` or `--help`: Show help message

## License

BPD is released under the MIT License. See [LICENSE](https://github.com/ImuS663/bpd/blob/main/LICENSE) for details.
